package kafka

import (
	"fmt"
	"myapp/global"

	"github.com/Shopify/sarama"
)

var ConsumerId = 100

type Consumer struct {
	Consumer   sarama.Consumer
	Topic      string
	ConsumerId int //消费者Id
}

func (c *Consumer) InitConsumer() error {
	consumer, err := sarama.NewConsumer([]string{global.SERVER_LIST}, nil)
	if err != nil {
		return err
	}
	c.Consumer = consumer
	c.Topic = global.TOPIC
	c.ConsumerId = ConsumerId
	ConsumerId++
	return nil
}

//指定partition
//offset 可以指定，传-1为获取最新offest
func (c *Consumer) GetMessage(partitionId int32, offset int64) {
	if offset == -1 {
		offset = sarama.OffsetNewest
	}
	pc, err := c.Consumer.ConsumePartition(c.Topic, partitionId, offset)
	if err != nil {
		fmt.Printf("failed to start consumer for partition %d,err:%v\n", partitionId, err)
		//That topic/partition is already being consumed
		return
	}

	// 异步从每个分区消费信息
	go func(sarama.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Printf("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v\n", c.ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		}
	}(pc)
}

//遍历所有分区
func (c *Consumer) GetMessageToAll(offset int64) {

	partitionList, err := c.Consumer.Partitions(c.Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v", err)
		return
	}
	fmt.Println("所有partition:", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		c.GetMessage(int32(partition), offset)
	}
}

func (c *Consumer) Start() {

	consumer := new(Consumer)
	err := consumer.InitConsumer()
	if err != nil {
		fmt.Printf("fail to init consumer, err:%v", err)
		return
	}
	go consumer.GetMessageToAll(1)

}
