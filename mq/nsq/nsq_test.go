package nsq

import (
	"fmt"
	"testing"
	"time"

	"github.com/sereiner/library/mq"
)

func Test_NewNsqProducer(t *testing.T) {
	p, err := NewNsqProducer("127.0.0.1:4150")
	if err != nil {
		t.Fatal(err)
	}
	p.Connect()
	for {
		time.Sleep(time.Microsecond * 500)
		p.Send("hello", `{"name":"jack","age":22}`, time.Second)
	}

}

func Test_NewNsqConsumer(t *testing.T) {
	c, err := NewNsqConsumer("127.0.0.1:4150")
	if err != nil {
		t.Fatal(err)
	}
	err = c.Consume("hello#ch", 1, M)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	c.Close()
}

func M(message mq.IMessage) {
	fmt.Println(message.GetMessage())
}
