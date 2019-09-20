package balancer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sereiner/library/balancer/pb"
	"github.com/sereiner/library/jsons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

type T struct {
	MerchantID int   `json:"merchant_id"`
	Spu        int64 `json:"spu"`
	Status     int   `json:"status"`
}

// {9632912394158080 14185116614529024}
// export ENDPOINTS=172.17.236.194:2379,172.17.236.197:2379,172.17.236.198:2379
func TestNewResolver(t *testing.T) {

	request, err := GetRPCRequest("receipt_server_debug", "flowserver")
	if err != nil {
		t.Error(err)
		return
	}

	type SpuSku struct {
		Spu int64 `json:"spu"`
		Sku int64 `json:"sku"`
	}
	input := &SpuSku{
		Spu: 9632912394158080,
		Sku: 14185116614529024,
	}

	status, err := request("/relation/record", input, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(status)

}

var rpcMap = sync.Map{}

type RequestFunc func(service string, input, result interface{}, header map[string]string, method ...string) (status int, err error)

func GetRPCRequest(platName, serverName string) (RequestFunc, error) {

	key := fmt.Sprintf("%s/%s", platName, serverName)

	v, ok := rpcMap.Load(key)
	if ok {
		return getRequestFunc(v.(pb.RPCClient)), nil
	}

	r := NewResolver("127.0.0.1:2379", platName, serverName)
	resolver.Register(r)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	conn, err := grpc.DialContext(
		ctx,
		r.Scheme()+"://authority/",
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithBlock())
	defer cancel()
	if err != nil {
		return nil, fmt.Errorf("创建 %s/%s grpc客户端失败,请确保服务端存在", platName, serverName)
	}

	client := pb.NewRPCClient(conn)

	rpcMap.Store(key, client)

	return getRequestFunc(client), nil
}

func getRequestFunc(cli pb.RPCClient) RequestFunc {
	return func(service string, input, result interface{}, header map[string]string, method ...string) (status int, err error) {
		h, err := jsons.Marshal(header)
		if err != nil {
			return
		}
		if len(h) == 0 {
			h = []byte("{}")
		}

		if result != nil {
			if r := reflect.TypeOf(result); r.Kind() != reflect.Ptr {
				return 0, fmt.Errorf("result Type must be ptr")
			}
		}

		t := reflect.TypeOf(input)
		if t.Kind() != reflect.Ptr {
			return 0, fmt.Errorf("input Type must be ptr")
		}

		f, err := jsons.Marshal(input)
		if err != nil {
			return
		}
		if len(f) == 0 {
			h = []byte("{}")
		}
		var methodName string
		if len(method) != 0 {
			methodName = strings.ToUpper(method[0])
		}
		fmt.Println(string(f))
		response, err := cli.Request(context.Background(),
			&pb.RequestContext{
				Method:  methodName,
				Service: service,
				Header:  string(h),
				Input:   string(f),
			},
			grpc.FailFast(true))
		if err != nil {
			status = 500
			return
		}

		status = int(response.Status)
		if result != nil {
			res := response.GetResult()
			err = json.Unmarshal([]byte(res), result)
			if err != nil {
				return
			}
		}

		return
	}

}


