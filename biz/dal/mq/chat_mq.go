package mq

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"context"
	"github.com/streadway/amqp"
	"log"
	"time"
)

//go:generate msgp -tests=false -o=chat_msgp.go -io=false
type MiddleMessage struct {
	Uid       int64     `msg:"uid"`
	TargetId  int64     `msg:"target"`
	Content   string    `msg:"content"`
	CreatedAt time.Time `msg:"publish_time"`
}

//todo:log
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (m *ChatMQ) Publisher(marshalMsg []byte) error {
	_, err := m.ch.QueueDeclare(
		m.queueName, // name
		true,        // 声明为持久队列
		false,       // 是否自动删除
		false,       // 是否具有排他性
		false,       // 是否阻塞处理
		nil,         // 额外的属性
	)
	failOnError(err, "Failed to declare a queue")

	err = m.ch.Publish(
		"",
		m.queueName, // routing key
		false,       // mandatory
		false,
		amqp.Publishing{
			/*消息持久化
			确保消息不会丢失，需要做两件事：我们需要将队列和消息都标记为持久的。
			将消息标记为持久的——通过使用amqp.Publishing中的持久性选项amqp.Persistent。
			*/
			DeliveryMode: amqp.Persistent, // 持久（交付模式：瞬态/持久）
			ContentType:  "text/plain",
			Body:         marshalMsg,
		})
	return err
}

func (m *ChatMQ) Consumer() {
	defer m.conn.Close()
	defer m.ch.Close()
	_, err := m.ch.QueueDeclare(
		m.queueName, // name
		true,        // 声明为持久队列
		false,       // 是否自动删除
		false,       // 是否具有排他性
		false,       // 是否阻塞处理
		nil,         // 额外的属性
	)
	failOnError(err, "Failed to declare a queue")

	//公平分发(就一个Consumer那就没得分发了)
	//err = m.ch.Qos(
	//	1,     // prefetch count
	//	0,     // prefetch size
	//	false, // global
	//)
	//failOnError(err, "ch.Qos() failed, err:%v\n")

	msgs, err := m.ch.Consume(
		m.queueName, // queue
		"",          // 区分多个消费者
		true,        // 是否自动应答
		false,       // 是否独有
		false,       // 设置为true，表示不能将同一个Connection中生产者发送的消息传递给这个Connection中的消费者
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go m.DeliverMessageToUser(msgs)
	<-forever
}

func (m *ChatMQ) DeliverMessageToUser(msgs <-chan amqp.Delivery) {
	for req := range msgs { //不断接收channel的消息
		var midMessage MiddleMessage
		_, err := midMessage.UnmarshalMsg(req.Body)
		if err != nil {
			log.Printf("unmarshal message error:%v", err)
			continue
		}

		message, err := db.CreateMessage(&db.Message{
			Uid:       midMessage.Uid,
			TargetId:  midMessage.TargetId,
			Content:   midMessage.Content,
			CreatedAt: midMessage.CreatedAt,
		})
		if err != nil {
			log.Printf("database error:%v", err)
			continue
		}

		//context.TODO():when it's unclear which Context to use or it is not
		err = cache.SetMessage(context.TODO(), message)
		if err != nil {
			log.Printf("database error:%v", err)
			continue
		}
	}
}
