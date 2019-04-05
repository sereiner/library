package etcd

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const requestTimeout = time.Second *2

var (
	ErrEtcdNodeIsExists = errors.New("etcd: node is exists")
	ErrEtcdNodeIsNotExists = errors.New("etcd: node is not exists")
	ErrClientConnClosing      = errors.New("etcd: the client connection is closing")
)

// EtcdClient .
type EtcdClient struct {
	servers   []string
	timeout   time.Duration
	conn      *clientv3.Client
	useCount  int32
	isConnect bool
	once      sync.Once
	closeCh   chan struct{}
	done      bool
}

// NewEtcdClient 创建etcd客户端
func NewEtcdClient(servers []string, timeout time.Duration) (*EtcdClient, error) {
	client := &EtcdClient{servers: servers, timeout: timeout, useCount: 0}
	client.closeCh = make(chan struct{})
	return client, nil
}

// Connect 连接etcd
func (e *EtcdClient) Connect() (err error) {
	if e.conn == nil {
		conn, err := clientv3.New(clientv3.Config{
			Endpoints:   e.servers,
			DialTimeout: e.timeout,
		})
		if err != nil {
			return fmt.Errorf("connect failed err:%v ", err)
		}
		e.conn = conn
	}
	atomic.AddInt32(&e.useCount, 1)
	time.Sleep(time.Second)
	e.isConnect = true
	return nil
}

// IsConnect 判断是否已经连接
func (e *EtcdClient) IsConnect() bool {
	if e.conn == nil {
		return false
	}
	return e.isConnect
}

// ReConnect 重连
func (e *EtcdClient) ReConnect() (err error) {
	e.isConnect = false
	if e.conn != nil {
		e.conn.Close()
	}
	e.done = false
	return e.Connect()
}

// Close 关闭连接
func (e *EtcdClient) Close() error {
	atomic.AddInt32(&e.useCount, -1)
	if e.useCount > 0 {
		return nil
	}
	if e.conn != nil {
		e.once.Do(func() {
			e.conn.Close()
		})
	}

	e.isConnect = false
	e.done = true
	e.once.Do(func() {
		close(e.closeCh)
	})
	return nil
}
