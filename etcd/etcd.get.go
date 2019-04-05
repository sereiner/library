package etcd

import (
	"context"
	"errors"
	"go.etcd.io/etcd/clientv3"
)

// GetValue 获取节点当前的值
func(e *EtcdClient) GetValue(path string) (value []byte, version int64, err error){
	if !e.isConnect {
		return nil,-1, errors.New("etcd 连接已经关闭")
	}

	ctx,cancel:= context.WithTimeout(context.Background(),requestTimeout)

	resp,err := e.conn.Get(ctx,path)
	cancel()
	if err != nil {
		return nil,-1,err
	}

	if resp.Count == 0 {
		return nil,-1, ErrEtcdNodeIsNotExists
	}

	env := resp.Kvs[0]

	return env.Value, env.Version,nil
}

// GetChildren 获取节点下的子节点
func (e *EtcdClient) GetChildren(path string) (paths []string,  err error) {
	if !e.isConnect {
		return nil, errors.New("etcd 连接已经关闭")
	}

	ctx ,cancel := context.WithTimeout(context.Background(),requestTimeout)
	resp,err := e.conn.Get(ctx,path,clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil,err
	}

	for _,ev := range resp.Kvs{
		paths = append(paths,string(ev.Key))
	}

	return paths, nil
}

// GetRevValue 获取节点指定version的值
func (e *EtcdClient) GetRevValue(path string,revision int64) (value []byte, version int64, err error) {
	if !e.isConnect {
		return nil,-1, errors.New("etcd 连接已经关闭")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := e.conn.Get(ctx, path, clientv3.WithRev(revision))
	cancel()
	if err != nil {
		return nil,-1,err
	}

	if resp.Count == 0 {
		return nil,-1, ErrEtcdNodeIsNotExists
	}

	env := resp.Kvs[0]

	return env.Value, env.Version,nil
}