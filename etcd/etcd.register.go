package etcd

import (
	"context"
	"errors"
	"go.etcd.io/etcd/clientv3"
)

// Register 注册一个服务
func (e *EtcdClient) Register(path ,data string) (ch <-chan *clientv3.LeaseKeepAliveResponse,err error) {
	if !e.isConnect {
		return nil, errors.New("etcd 连接已经关闭")
	}

	resp, err := e.conn.Grant(context.TODO(), 10)
	if err != nil {
		return nil, err
	}

	_, err = e.conn.Put(context.TODO(), path, data, clientv3.WithLease(resp.ID))
	if err != nil {
		return nil,err
	}


	ch, err = e.conn.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		return nil,err
	}

	return
}

func (e *EtcdClient) Revoke(leaseID clientv3.LeaseID) error {
	if !e.isConnect {
		return errors.New("etcd 连接已经关闭")
	}

	// revoking lease expires the key attached to its lease ID
	_, err := e.conn.Revoke(context.TODO(), leaseID)
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) RevokeByPath(path string) error {
	if !e.isConnect {
		return errors.New("etcd 连接已经关闭")
	}

	ctx,cancel:= context.WithTimeout(context.Background(),requestTimeout)

	resp,err := e.conn.Get(ctx,path)
	cancel()
	if err != nil {
		return err
	}

	if resp.Count == 0 {
		return  ErrEtcdNodeIsNotExists
	}

	id := resp.Kvs[0].Lease

	_, err = e.conn.Revoke(context.TODO(), clientv3.LeaseID(id))
	if err != nil {
		return err
	}

	return nil
}
