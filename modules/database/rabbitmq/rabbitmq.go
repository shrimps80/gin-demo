package rabbitmq

import (
	"log"
	
	"github.com/streadway/amqp"
	"fmt"
	"gin-demo/config"
)

type MqType struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var (
	Client MqType
)

func init() {
	mqUrl := fmt.Sprintf("amqp://%s:%s@%s",
		config.GetEnv().MQName,
		config.GetEnv().MQPassword,
		config.GetEnv().MqServers)
	
	var err error
	Client.Conn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		panic(err)
	}
	
	Client.Channel, err = Client.Conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		panic(err)
	}
}

func (client *MqType) Destory() {
	_ = client.Channel.Close()
	_ = client.Conn.Close()
}

func (client *MqType) PublishSimple(queueName, message string) bool {
	var err error
	// 保证队列存在，消息能发送到队列中
	_, err = client.Channel.QueueDeclare(
		queueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		log.Fatalf("Failed QueueDeclare: %s", err)
		return false
	}
	
	//发送消息到队列中
	err = client.Channel.Publish(
		"", //交换机
		queueName,
		// 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return false
	}
	
	return true
}
