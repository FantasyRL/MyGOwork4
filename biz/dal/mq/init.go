package mq

import (
	"bibi/pkg/conf"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
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
	ChatMQCli = new(ChatMQ)
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

	go ChatMQCli.Consumer()
	fmt.Println("RabbitMQ connect access")
}

func GenRabbitMQAddr(mq *conf.RabbitMQ) string {
	return strings.Join([]string{"amqp://", mq.Username, ":", mq.Password, "@localhost:", mq.Port, "/"}, "")
}
