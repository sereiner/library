package nsq

import "github.com/go-redis/redis"

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
func NewNsqMessage(cmd *redis.StringSliceCmd) *NsqMessage {
	msg, err := cmd.Result()
	hasData := err == nil && len(msg) > 0
	ndata := ""
	if hasData {
		ndata = msg[len(msg)-1]
	}
	return &NsqMessage{Message: ndata, HasData: hasData}
}
