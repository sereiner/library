package main

import (
	"context"
	"github.com/sereiner/log"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	resp, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "/foo", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key 'foo' will be kept forever
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}



	//go func() {
	//	time.Sleep(time.Second*12)
	//	_, err = cli.Revoke(context.TODO(), resp.ID)
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}()
	for {
		select {
		case ka, ok := <-ch:
			if !ok {
				return
			} else {
				log.Infof("ttl:%+v", ka)
			}
		}
	}

}
