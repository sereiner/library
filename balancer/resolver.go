package balancer

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	logger "github.com/sereiner/library/log"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Resolver struct {
	target  string
	schema  string
	service string
	cli     *clientv3.Client
	cc      resolver.ClientConn
	logger  *logger.Logger
}

// NewResolver return resolver builder
// target example: "http://127.0.0.1:2379,http://127.0.0.1:12379,http://127.0.0.1:22379"
// service is service name
func NewResolver(target string, platName, service string) resolver.Builder {
	return &Resolver{
		target:  target,
		schema:  platName,
		service: service,
		logger:  logger.GetSession("grpc", logger.CreateSession()),
	}
}

// Scheme return etcdv3 schema
func (r *Resolver) Scheme() string {
	return r.schema
}

// ResolveNow
func (r *Resolver) ResolveNow(rn resolver.ResolveNowOption) {
}

// Close
func (r *Resolver) Close() {
}

// Build to resolver.Resolver
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	var err error

	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints: strings.Split(r.target, ","),
	})
	if err != nil {
		return nil, fmt.Errorf("grpc: create clientv3 client failed: %v", err)
	}

	r.cc = cc

	go r.watch(fmt.Sprintf("/%s/%s/", r.schema, r.service))

	return r, nil
}

func (r *Resolver) watch(prefix string) {
	addrDict := make(map[string]resolver.Address)

	update := func() {
		addrList := make([]resolver.Address, 0, len(addrDict))
		for _, v := range addrDict {
			addrList = append(addrList, v)
		}
		r.cc.NewAddress(addrList)
	}

	resp, err := r.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err == nil {
		for i := range resp.Kvs {
			addrDict[string(resp.Kvs[i].Value)] = resolver.Address{Addr: string(resp.Kvs[i].Value)}
		}
	}
	update()
	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range rch {
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				addrDict[string(ev.Kv.Key)] = resolver.Address{Addr: string(ev.Kv.Value)}
				r.logger.Info(" grpc节点更新 ", string(ev.Kv.Key))
			case mvccpb.DELETE:
				delete(addrDict, string(ev.PrevKv.Key))
				r.logger.Info(" grpc节点更新下线 ", string(ev.Kv.Key))
			}
		}
		update()
	}
}
