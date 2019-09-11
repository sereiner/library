package etcd_lib

import (
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/samuel/go-zookeeper/zk"
	logger "github.com/sereiner/library/log"
	"sync"
	"time"
)

// TIMEOUT 连接zk服务器操作的超时时间
var TIMEOUT = time.Second

var (
	ErrColientCouldNotConnect = errors.New("etcd: could not connect to the server")
	ErrClientConnClosing      = errors.New("etcd: the client connection is closing")
)

//ZookeeperClient zookeeper客户端
type ZookeeperClient struct {
	servers   []string
	timeout   time.Duration
	conn      *clientv3.Client
	Log       *logger.Logger
	useCount  int32
	isConnect bool
	once      sync.Once
	CloseCh   chan struct{}
	digest    bool
	userName  string
	password  string
	// 是否是手动关闭
	done bool
}
