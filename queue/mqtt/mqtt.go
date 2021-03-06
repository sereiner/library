package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	logger "github.com/sereiner/library/log"
	"github.com/sereiner/library/net"
	"github.com/sereiner/library/utility"

	"github.com/sereiner/library/queue"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

// MQTTClient memcache配置文件
type MQTTClient struct {
	servers []string
	client  *client.Client
	once    sync.Once
	Logger  *logger.Logger
	lk      sync.Mutex
	connCh  chan int
	done    bool
	conf    *queue.Config
}

//New 根据配置文件创建一个redis连接
func New(addrs []string, raw string) (m *MQTTClient, err error) {
	m = &MQTTClient{servers: addrs, Logger: logger.GetSession("mqtt.publisher", logger.CreateSession())}
	m.connCh = make(chan int, 1)
	m.conf = &queue.Config{}
	if err := json.Unmarshal([]byte(raw), &m.conf); err != nil {
		return nil, err
	}
	cc, _, err := m.connect()
	if err != nil {
		return nil, err
	}
	m.client = cc
	go m.reconnect()
	return m, nil
}

func (c *MQTTClient) reconnect() {
	for {
		select {
		case <-time.After(time.Second * 3): //延迟重连
			select {
			case <-c.connCh:
				c.Logger.Debug("publisher与服务器断开连接，准备重连")
				func() {
					defer recover()
					c.client.Disconnect()
					c.client.Terminate()
				}()
				client, b, err := c.connect()
				if err != nil {
					c.Logger.Error("连接失败:", err)
				}
				if b {
					c.Logger.Info("publisher成功连接到服务器")
					c.client = client
				}
			default:

			}
		}
	}
}

func (c *MQTTClient) connect() (*client.Client, bool, error) {
	c.lk.Lock()
	defer c.lk.Unlock()
	cert, err := c.getCert(c.conf)
	if err != nil {
		return nil, false, err
	}
	cc := client.New(&client.Options{
		ErrorHandler: func(err error) {
			select {
			case c.connCh <- 1: //发送重连消息
			default:
			}
		},
	})
	if err := cc.Connect(&client.ConnectOptions{
		Network:         "tcp",
		Address:         c.conf.Addr,
		UserName:        []byte(c.conf.UserName),
		Password:        []byte(c.conf.Password),
		ClientID:        []byte(fmt.Sprintf("%s-%s", net.GetLocalIPAddress(), utility.GetGUID()[0:6])),
		TLSConfig:       cert,
		PINGRESPTimeout: time.Second,
		CleanSession:    true,
		KeepAlive:       3,
	}); err != nil {
		return nil, false, fmt.Errorf("连接失败:%v(%s-%s/%s)", err, c.conf.Addr, c.conf.UserName, c.conf.Password)
	}
	return cc, true, nil
}
func (c *MQTTClient) getCert(conf *queue.Config) (*tls.Config, error) {
	if conf.CertPath == "" {
		return nil, nil
	}
	b, err := ioutil.ReadFile(conf.CertPath)
	if err != nil {
		return nil, fmt.Errorf("读取证书失败:%s(%v)", conf.CertPath, err)
	}
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(b); !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	return &tls.Config{
		RootCAs: roots,
	}, nil
}

// Push 向存于 key 的列表的尾部插入所有指定的值
func (c *MQTTClient) Push(key string, value string) error {
	if c.done {
		return fmt.Errorf("队列已关闭")
	}
	if key == "" {
		return fmt.Errorf("队列名称不能为空")
	}
	if value == "" {
		return fmt.Errorf("放入队列的数据不能为空")
	}
	return c.client.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(key),
		Message:   []byte(value),
	})
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *MQTTClient) Pop(key string) (string, error) {
	return "", fmt.Errorf("mqtt不支持pop方法")
}

// Count 移除并且返回 key 对应的 list 的第一个元素。
func (c *MQTTClient) Count(key string) (int64, error) {
	return 0, fmt.Errorf("mqtt不支持pop方法")
}

// Close 释放资源
func (c *MQTTClient) Close() error {
	c.done = true
	c.once.Do(func() {
		c.client.Disconnect()
		c.client.Terminate()
	})
	return nil
}

type redisResolver struct {
}

func (s *redisResolver) Resolve(address []string, conf string) (queue.IQueue, error) {
	return New(address, conf)
}
func init() {
	queue.Register("mqtt", &redisResolver{})
}
