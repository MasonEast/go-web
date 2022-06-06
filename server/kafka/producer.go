package kafka

import (
	"errors"
	"fmt"
	"myapp/global"

	"github.com/Shopify/sarama"
)

var ProducerId = 1

type Producer struct {
	Producer sarama.SyncProducer
	Topic string
	ProducerId int
	MessageId int
}
func (p *Producer) InitProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll	// 发送完数据要等leader和follower确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner	// 新选出一个partition
	config.Producer.Return.Successes = true	// 成功交付的消息通过success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{global.SERVER_LIST}, config)
	if err != nil {
		fmt.Println(errors.New(err.Error()))
		return
	}

	p.Producer = client
	p.Topic = global.TOPIC
	p.ProducerId = ProducerId
	p.MessageId = 1

	ProducerId++
}

func (p *Producer) SendMessage(txt string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.Topic
	// txt := fmt.Sprintf("ProducerId: %d, this is a test log %d", p.ProducerId, p.MessageId)
	msg.Value = sarama.StringEncoder(txt)
	pid, offset, err := p.Producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:", err)
		return
	}
	fmt.Printf("ProducerId: %d, MessageId: %d, offset: %d, pid: %d\n", p.ProducerId, p.MessageId, offset, pid)

	p.MessageId++
}

func (p *Producer) Close() {
	p.Producer.Close()
}