package rabbitMq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"video_service/config"
	"video_service/logger"
)

type RMConsumer struct {
	Q    amqp.Queue
	Conn *amqp.Connection
	Ch   *amqp.Channel
	Done chan error // 关闭管道
}

// NewConsumer 获取RM消费者
func NewConsumer(name string) *RMConsumer {
	// 初始化RMQ
	// 连接到RabbitMQ服务器
	url := config.InitRabbitMQUrl()
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Print(err)
	}

	// 声明队列
	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err)
	}
	consumer := RMConsumer{
		Q:    q,
		Ch:   ch,
		Done: make(chan error),
	}
	return &consumer
}

// Consume 消费
func (c *RMConsumer) Consume(consumer string) <-chan amqp.Delivery {
	// 消费者
	message, err := c.Ch.Consume(
		c.Q.Name,
		consumer,
		false, //手动确认
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err)
	}
	return message
}

func (c *RMConsumer) Close() error {
	defer close(c.Done) // 确保关闭 Done 通道
	if err := c.Ch.Close(); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}
	// 确保关闭连接
	if err := c.Conn.Close(); err != nil {
		return fmt.Errorf("RMQ connection close error: %s", err)
	}
	defer logger.Log.Printf("RMQ Close OK")
	// wait for handle() to exit
	return <-c.Done
}
