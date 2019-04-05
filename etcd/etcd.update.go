package etcd

import (
	"context"
	"errors"
	"github.com/sereiner/log"
)

// Update 更新节点的值,如果节点不存在则报错
func(e *EtcdClient) Update(path string,data string) error {
	if !e.isConnect {
		return  errors.New("etcd 连接已经关闭")
	}

	ctx,cancel:= context.WithTimeout(context.Background(),requestTimeout)
	resp,err := e.conn.Get(ctx,path)
	cancel()
	if err != nil {
		return err
	}
	log.Info(resp.Count)
	if resp.Count == 0 {
		return  ErrEtcdNodeIsNotExists
	}

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	_, err = e.conn.Put(ctx, path, data)
	cancel()
	if err != nil {
		return err
	}

	return nil
}