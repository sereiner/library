package etcd

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
)

// Delete 删除一个节点
func (e *EtcdClient) Delete(path string) error {
	if !e.isConnect {
		return errors.New("etcd 连接已经关闭")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	// count keys about to be deleted
	resp, err := e.conn.Get(ctx, path)
	if err != nil {
		return err
	}

	if resp.Count != 1 {
		return errors.New("节点不存在,删除操作无效")
	}

	// delete the keys
	dresp, err := e.conn.Delete(ctx, path)
	if err != nil {
		return err
	}

	if int64(len(resp.Kvs)) == dresp.Deleted {
		return nil
	}

	return errors.New("删除节点失败")
}


// DeleteWithPrefix 根据前缀删除节点
func(e *EtcdClient) DeleteWithPrefix(path string) error {
	if !e.isConnect {
		return errors.New("etcd 连接已经关闭")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	// count keys about to be deleted
	resp, err := e.conn.Get(ctx, path, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if resp.Count == 0 {
		return fmt.Errorf("没有前缀名称为:%s 的节点,删除操作无效",path)
	}

	// delete the keys
	dresp, err := e.conn.Delete(ctx, path, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	if int64(len(resp.Kvs)) == dresp.Deleted {
		return nil
	}

	return errors.New("删除节点失败")
}
