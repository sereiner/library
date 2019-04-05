package etcd

import (
	"context"
	"errors"
)

// CreateNode 创建节点
func (e *EtcdClient) CreateNode(path string, data string) (err error) {

	if !e.isConnect {
		return errors.New("etcd 连接已经关闭")
	}

	isExists, err := e.Exists(path)
	if err != nil {
		return err
	}
	if isExists {
		return ErrEtcdNodeIsExists
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = e.conn.Put(ctx, path, data)
	cancel()
	if err != nil {
		return err
	}

	return nil
}

// Exists 判断节点是否存在
func (e *EtcdClient) Exists(path string) (bool, error) {
	if !e.isConnect {
		return false, errors.New("etcd 连接已经关闭")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := e.conn.Get(ctx, path)
	cancel()
	if err != nil {
		return false, err
	}

	if resp.Count != 0 {
		return true, nil
	}
	return false, nil
}
