package rabbitMq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
	"video_service/config"
)

type RMProducer struct {
	Q      amqp.Queue // rm生产者队列
	Conn   *amqp.Connection
	Ch     *amqp.Channel // rm生产者ch管道
	Cancel context.CancelFunc
	Ctx    context.Context
}

func NewRMQProducer(name string) *RMProducer {
	// 初始化RMQ
	// 连接到RabbitMQ服务器
	url := config.InitRabbitMQUrl()
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	// 创建队列
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	return &RMProducer{Q: q, Ch: ch}
}

func (p *RMProducer) RMWithTimeout(time time.Duration) {
	// 5s 后，如果没有消费则自动删除
	// 过期时间和上下文
	ctx, cancel := context.WithTimeout(context.Background(), time)
	p.Ctx, p.Cancel = ctx, cancel
}

func (p *RMProducer) PublishMessage(reqString []byte, contentType string) (err error) {
	err = p.Ch.PublishWithContext(
		p.Ctx,
		"",       // exchange
		p.Q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{ // 发送的消息，固定有消息体和一些额外的消息头，包中提供了封装对象
			ContentType: contentType, // 是一个通用的 MIME 类型，表示二进制流，常用于传输不具备明确定义类型的二进制数据。
			Body:        reqString,   // 请求信息重新封装为json，加入消息队列
		})
	return
}

func (p *RMProducer) Close() error {
	p.Cancel()
	err := p.Ch.Close()
	if err != nil {
		return err
	}
	// 确保关闭连接（如果你在 InitMQ 中初始化了 conn）
	if p.Conn != nil {
		return p.Conn.Close()
	}
	return err
}
