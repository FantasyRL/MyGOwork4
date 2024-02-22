package mq

import (
	"bibi/pkg/conf"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type ChatMQ struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queueName string
	////交换机名称
	//Exchange string
	//连接信息
	MqUrl string
}

var (
	ChatMQCli *ChatMQ
)

func Init() {
	ChatMQCli.MqUrl = GenRabbitMQAddr(conf.MQ)
	conn, err := amqp.Dial(ChatMQCli.MqUrl)
	if err != nil {
		log.Fatalf("Init rabbitMQ error:%s", err)
		return
	}
	ChatMQCli.conn = conn

	ch, err := ChatMQCli.conn.Channel()
	if err != nil {
		log.Fatalf("Init rabbitMQ error:%s", err)
		return
	}

	ChatMQCli.ch = ch

	ChatMQCli.queueName = "chatQueue"

	go fmt.Println("RabbitMQ connect access")
}

func GenRabbitMQAddr(mq *conf.RabbitMQ) string {
	return fmt.Sprintf("amqp://%s:%s@localhost:%s/", mq.Username, mq.Password, mq.Port)
}
