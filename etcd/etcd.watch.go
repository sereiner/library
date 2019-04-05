package etcd

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
func(e *EtcdClient) WatchValue(path string) (data chan ValueWatcher,err error) {

}

// WatchChildren 监控子节点变化
func (e *EtcdClient) WatchChildren(path string) (err error) {

}