package nsq

import (
	"github.com/nsqio/go-nsq"
)

//NsqMessage reids消息
type NsqMessage struct {
	Message string
	HasData bool
}

//Ack 确定消息
func (m *NsqMessage) Ack() error {
	return nil
}

//Nack 取消消息
func (m *NsqMessage) Nack() error {
	return nil
}

//GetMessage 获取消息
func (m *NsqMessage) GetMessage() string {
	return m.Message
}

//Has 是否有数据
func (m *NsqMessage) Has() bool {
	return m.HasData
}

//NewNsqMessage 创建消息
func NewNsqMessage(msg *nsq.Message) *NsqMessage {
	ndata := string(msg.Body)
	hasData := msg.HasResponded()
	return &NsqMessage{Message: ndata, HasData: hasData}
}
