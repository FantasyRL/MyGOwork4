package monitor

import (
	"bibi/biz/dal/cache"
	"bibi/biz/service/chat_service"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hertz-contrib/websocket"
	"log"
)

func (c *Client) Read() {
	defer func() { //闭包
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		sendMsg := new(chat_service.SendMsg)
		err := c.Socket.ReadJSON(sendMsg) // 接收消息
		if err != nil {
			log.Printf(errno.ParamErrMsg)
			break
		}
		if len(sendMsg.Content) == 0 || len(sendMsg.Content) > 2000 {
			baseResp := pack.BuildChatBaseResp(errno.CharacterBeyondLimitError)
			resp, _ := json.Marshal(baseResp)
			_ = c.Socket.WriteMessage(websocket.TextMessage, resp)
			break
		}

		if sendMsg.Type == 1 {

			marshalMsg, _ := (chat_service.ReplyMsg{
				Code:    errno.WebSocketSuccessCode,
				From:    c.ID,
				Content: sendMsg.Content,
			}).MarshalMsg(nil)
			Manager.Broadcast <- &Broadcast{ //传到broadcast来发给target
				Client:  c,
				Message: marshalMsg,
				Type:    sendMsg.Type,
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case marshalMsg, ok := <-c.Send:
			if !ok {
				baseResp := pack.BuildChatBaseResp(errno.WebSocketError)
				resp, _ := json.Marshal(baseResp)
				_ = c.Socket.WriteMessage(websocket.CloseMessage, resp)
				return
			}
			var replyMsg chat_service.ReplyMsg
			_, _ = replyMsg.UnmarshalMsg(marshalMsg)
			resp, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, resp)

		}
	}
}

func (c *Client) IfNotReadMessage(uid int64) error {
	ok, err := cache.IsUserChattedByOthers(context.TODO(), uid)
	if err != nil {
		return err
	}
	if ok {
		count, replyMsgs, err := cache.GetMessages(context.TODO(), uid)
		if err != nil {
			return err
		}
		countMsg := fmt.Sprintf("你有%v条未读消息", count)
		_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(countMsg))
		for _, replyMsg := range replyMsgs {
			resp, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, resp)
		}
	}
	return nil
}
