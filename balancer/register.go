package balancer

import (
	"context"
	"fmt"
	"github.com/sereiner/library/envs"
	"net"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Prefix should start and end with no slash
var Deregister = make(chan struct{})

// Register
func Register(target, platName, svrName, host, port string, interval time.Duration, ttl int) error {
	serviceValue := net.JoinHostPort(host, port)
	serviceKey := fmt.Sprintf("/%s/%s/%s", platName, svrName, serviceValue)
	var targetArr []string
	if len(target) == 0 {
		endpoints := envs.GetString("ENDPOINTS", "127.0.0.1:2379")
		targetArr = strings.Split(endpoints, ",")
	} else {
		targetArr = strings.Split(target, ",")
	}

	// get endpoints for register dial address
	var err error
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: targetArr,
	})
	if err != nil {
		return fmt.Errorf("grpc: create clientv3 client failed: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), interval)
	resp, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		return fmt.Errorf("grpc: create clientv3 lease failed: %v", err)
	}

	if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		return fmt.Errorf("grpc: set platName:%s service '%s' with ttl to clientv3 failed: %s", platName, svrName, err.Error())
	}

	if _, err := cli.KeepAlive(context.TODO(), resp.ID); err != nil {
		return fmt.Errorf("grpc: refresh platName:%s service '%s' with ttl to clientv3 failed: %s", platName, svrName, err.Error())
	}

	// wait deregister then delete
	go func() {
		<-Deregister
		cli.Delete(context.Background(), serviceKey)
		Deregister <- struct{}{}
	}()

	return nil
}

// UnRegister delete registered service from etcd
func UnRegister() {
	Deregister <- struct{}{}
	<-Deregister
}
