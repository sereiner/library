package balancer

import (
	"context"
	"github.com/sereiner/library/balancer/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"testing"
	"time"
)


func TestNewResolver(t *testing.T) {

	r := NewResolver("127.0.0.1:2379","receipt_server_debug", "flowserver")
	resolver.Register(r)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	conn, err := grpc.DialContext(
		ctx,
		r.Scheme()+"://authority/"+"receipt",
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithBlock())
	defer cancel()
	if err != nil {
		t.Errorf("创建grpc客户端失败,请确保服务端存在")
	}

	client := pb.NewRPCClient(conn)

	res,err := client.Request(context.Background(),&pb.RequestContext{
		Service:              "/sku/stock",
		Input:                "",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}