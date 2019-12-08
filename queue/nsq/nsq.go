package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/sereiner/library/queue"
)

type nsqClient struct {
	servers []string
	client  *nsq.Producer
}

// New 根据配置文件创建一个redis连接
func New(addrs []string, conf string) (m *nsqClient, err error) {
	m = &nsqClient{servers: addrs}
	m.client, err = nsq.NewProducer(addrs[0], nsq.NewConfig())
	if err != nil {
		return
	}
	return
}

// Push 向存于 key 的列表的尾部插入所有指定的值
func (c *nsqClient) Push(key string, value string) error {
	return c.client.Publish(key, []byte(value))
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *nsqClient) Pop(key string) (string, error) {
	return "", nil
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *nsqClient) Count(key string) (int64, error) {
	return 0, nil
}

// Close 释放资源
func (c *nsqClient) Close() error {
	c.client.Stop()
	return nil
}

type nsqResolver struct {
}

func (s *nsqResolver) Resolve(address []string, conf string) (queue.IQueue, error) {
	return New(address, conf)
}
func init() {
	queue.Register("nsq", &nsqResolver{})
}
