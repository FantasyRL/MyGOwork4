package chat_service

import (
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"encoding/json"
	"github.com/hertz-contrib/websocket"
	"log"
)

func (c *Client) Read() {
	defer func() { //闭包
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler() //心跳

		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(sendMsg) // 接收消息
		if err != nil {
			log.Printf(errno.ParamErrMsg)
			break
		}

		if sendMsg.Type == 1 { //发消息
			//r1, _ := cache.RedisClient.Get(c.ID).Result()//ID是否在缓存里
			//r2, _ := cache.RedisClient.Get(c.SendID).Result()
			r1 := "2" //test
			r2 := ""
			if r1 > "3" && r2 == "" { //1给2 发了三条 2没有回复，就停止1发送
				baseResp := pack.BuildChatBaseResp(errno.WebSocketTargetOfflineError)
				msg, _ := json.Marshal(baseResp)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				//cache.RedisClient.Incr(c.ID)
				//_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result()
				//3个月的持久化
			}

			Manager.Broadcast <- &Broadcast{ //传到broadcast来发给target
				Client:  c,
				Message: []byte(sendMsg.Content),
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
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replymsg := ReplyMsg{
				Code:    errno.SuccessCode, //随便写的
				Content: string(message),
			}
			msg, _ := json.Marshal(replymsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)

		}
	}
}
