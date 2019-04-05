package etcd

import (
	"context"
	"errors"
	"go.etcd.io/etcd/clientv3"
)

type ValueWatcher interface {
	GetValue() ([]byte, int64)
	GetError() error
	GetPath() string
}

type valueEntity struct {
	Value   []byte
	version int64
	path    string
	Err     error
}

func (v *valueEntity) GetPath() string {
	return v.path
}
func (v *valueEntity) GetValue() ([]byte, int64) {
	return v.Value, v.version
}
func (v *valueEntity) GetError() error {
	return v.Err
}

// WatchValue 监控指定节点的值是否发生变化，变化时返回变化后的值
func (e *EtcdClient) WatchValue(path string) (data chan ValueWatcher, err error) {
	if !e.isConnect {
		return nil, errors.New("etcd 连接已经关闭")
	}

	data = make(chan ValueWatcher, 1)
	rch := e.conn.Watch(context.Background(), path)

	go func(data chan ValueWatcher) {
		for {
			select {
			case <-e.closeCh:
				data <- &valueEntity{path: path, Err: ErrClientConnClosing}
				return
			case ev := <-rch:
				if e.done {
					data <- &valueEntity{path: path, Err: ErrClientConnClosing}
					return
				}

				if ev.Err() != nil {
					data <- &valueEntity{path: path, Err: ev.Err()}
					return
				}

				for _, ev := range ev.Events {
					data <- &valueEntity{Value: ev.Kv.Value, version: ev.Kv.Version, path: path, Err: nil}
				}
			}
		}
	}(data)

	return
}


type ChildrenWatcher interface {
	GetValue() ([]byte, int64)
	GetPath() string
	GetError() error
}

type valuesEntity struct {
	Value   []byte
	version int64
	path    string
	Err     error
}


func (v *valuesEntity) GetValue() ([]byte, int64) {
	return v.Value, v.version
}
func (v *valuesEntity) GetError() error {
	return v.Err
}
func (v *valuesEntity) GetPath() string {
	return v.path
}
// WatchChildren 监控子节点变化
func (e *EtcdClient) WatchChildren(path string) (data chan ChildrenWatcher,err error) {
	if !e.isConnect {
		return nil, errors.New("etcd 连接已经关闭")
	}

	data = make(chan ChildrenWatcher, 1)
	rch := e.conn.Watch(context.Background(), path,clientv3.WithPrefix())

	go func(data chan ChildrenWatcher) {
		for {
			select {
			case <-e.closeCh:
				data <- &valuesEntity{path: path, Err: ErrClientConnClosing}
				return
			case ev := <-rch:
				if e.done {
					data <- &valuesEntity{path: path, Err: ErrClientConnClosing}
					return
				}

				if ev.Err() != nil {
					data <- &valuesEntity{path: path, Err: ev.Err()}
					return
				}

				for _, ev := range ev.Events {
					data <- &valuesEntity{Value: ev.Kv.Value, version: ev.Kv.Version, path: string(ev.Kv.Key), Err: nil}
				}
			}
		}
	}(data)

	return
}
