package nsq

import (
	"errors"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/sereiner/library/mq"
	logger "github.com/sereiner/library/log"
)

//NsqProducer Producer
type NsqProducer struct {
	address   string
	client    *nsq.Producer
	backupMsg chan *mq.ProcuderMessage
	closeCh   chan struct{}
	done      bool
	*mq.OptionConf
}

//NewRedisProducer 创建新的producer
func NewRedisProducer(address string, opts ...mq.Option) (producer *NsqProducer, err error) {
	producer = &NsqProducer{address: address}
	producer.OptionConf = &mq.OptionConf{
		Logger: logger.GetSession("mq.nsq", logger.CreateSession()),
	}
	producer.closeCh = make(chan struct{})
	for _, opt := range opts {
		opt(producer.OptionConf)
	}
	return
}

//Connect  循环连接服务器
func (producer *NsqProducer) Connect() (err error) {
	producer.client, err = nsq.NewProducer(producer.address, nsq.NewConfig())
	if err != nil {
		return err
	}

	return producer.client.Ping()
}

//GetBackupMessage 获取备份数据
func (producer *NsqProducer) GetBackupMessage() chan *mq.ProcuderMessage {
	return producer.backupMsg
}

//Send 发送消息
func (producer *NsqProducer) Send(queue string, msg string, timeout time.Duration) (err error) {
	if producer.done {
		return errors.New("mq producer 已关闭")
	}

	return producer.client.Publish(queue, []byte(msg))
}

//Close 关闭当前连接
func (producer *NsqProducer) Close() {
	producer.done = true
	close(producer.closeCh)
	close(producer.backupMsg)
	producer.client.Stop()
}

type nsqProducerResolver struct {
}

func (s *nsqProducerResolver) Resolve(address string, opts ...mq.Option) (mq.MQProducer, error) {
	return NewRedisProducer(address, opts...)
}

func init() {
	mq.RegisterProducer("nsq", &nsqProducerResolver{})
}
