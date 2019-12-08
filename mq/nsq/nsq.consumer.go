package nsq

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/nsqio/go-nsq"

	"github.com/sereiner/library/concurrent/cmap"
	"github.com/sereiner/library/mq"
)

type NsqConsumer struct {
	address   string
	consumers cmap.ConcurrentMap
	quitChan  chan struct{}
	*mq.OptionConf
}

type nsqConsumer struct {
	consumer *nsq.Consumer
	msgQueue chan *nsq.Message
}

func (n *nsqConsumer) HandleMessage(message *nsq.Message) error {

	n.msgQueue <- message
	return nil
}

func (n *nsqConsumer) Messages() chan *nsq.Message {
	return n.msgQueue
}

// NewNsqConsumer 初始化nsq Consumer
func NewNsqConsumer(address string, opts ...mq.Option) (nsq *NsqConsumer, err error) {
	nsq = &NsqConsumer{address: address, quitChan: make(chan struct{}, 0)}
	nsq.OptionConf = &mq.OptionConf{}
	nsq.consumers = cmap.New(2)
	for _, opt := range opts {
		opt(nsq.OptionConf)
	}
	return
}

//Connect 连接到服务器
func (n *NsqConsumer) Connect() error {
	return nil
}

//Consume 订阅消息
func (n *NsqConsumer) Consume(queue string, concurrency int, call func(mq.IMessage)) (err error) {
	_, cname, _ := n.consumers.SetIfAbsentCb(queue, func(i ...interface{}) (interface{}, error) {

		c := &nsqConsumer{}
		queueArr := strings.Split(queue, "#")
		if len(queueArr) != 2 {
			panic("nsq 消息队列格式错误 topic#channel")
		}

		config := nsq.NewConfig()
		c.consumer, err = nsq.NewConsumer(queueArr[0], queueArr[1], config)
		c.msgQueue = make(chan *nsq.Message, 10000)
		c.consumer.AddHandler(c)
		err := c.consumer.ConnectToNSQD(n.address)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return c, nil
	})

	consumer := cname.(*nsqConsumer)

	var (
		chanmsg = make(chan *nsq.Message, 10000)
		closing = make(chan struct{})
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Kill, os.Interrupt)
		<-signals
		fmt.Println("Initiating shutdown of consumer...")
		close(closing)
	}()

	go func(consumer *nsqConsumer) {
		for {
			select {
			case message, ok := <-consumer.Messages():
				if ok {
					chanmsg <- message
				}
			}
		}
	}(consumer)

	go func() {
	LOOP:
		for {
			select {
			case msg, ok := <-chanmsg:
				if ok {
					go call(NewNsqMessage(msg))
				} else {
					break LOOP
				}
			}
		}
	}()

	//close(chanmsg)

	//if err := consumer.consumer.Close(); err != nil {
	//fmt.Println("Failed to close consumer: ", err)
	//}
	return nil
}

//UnConsume 取消注册消费
func (n *NsqConsumer) UnConsume(queue string) {
	if c, ok := n.consumers.Get(queue); ok {
		consumer := c.(*nsqConsumer)
		close(consumer.msgQueue)
	}
}

//Close 关闭当前连接
func (n *NsqConsumer) Close() {
	close(n.quitChan)
	n.consumers.IterCb(func(key string, value interface{}) bool {
		consumer := value.(*nsqConsumer)
		close(consumer.msgQueue)
		return true
	})
}

type nsqConsumerResolver struct {
}

func (s *nsqConsumerResolver) Resolve(address string, opts ...mq.Option) (mq.MQConsumer, error) {
	return NewNsqConsumer(address, opts...)
}
func init() {
	mq.RegisterCosnumer("nsq", &nsqConsumerResolver{})
}
