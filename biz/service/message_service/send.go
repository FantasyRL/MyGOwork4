package message_service

import (
	"bibi/biz/dal/mq"
	"bibi/biz/model/message"
	"bibi/pkg/errno"
	"time"
)

func (s *MessageService) SendMessage(req *message.MessageActionReq, uid int64) error {
	if len(req.Content) == 0 || len(req.Content) > 2000 {
		return errno.CharacterBeyondLimitError
	}
	msg := &mq.MiddleMessage{
		Uid:       uid,
		TargetId:  req.TargetID,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	marshalMsg, err := msg.MarshalMsg(nil)
	if err != nil {
		return err
	}
	err = mq.ChatMQCli.Publisher(marshalMsg)
	return err
}
