package monitor

import (
	"bibi/biz/dal/cache"
	"bibi/biz/dal/db"
	"bibi/biz/service/chat_service"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
	"encoding/json"
	"github.com/hertz-contrib/websocket"
	"log"
)

func (manager *ClientManager) Listen() {
	for {
		log.Println("监听管道通信")

		select {

		case client := <-Manager.Register:
			log.Printf("%v:online\n", client.ID)
			Manager.Clients[client.ID] = client //把连接放到用户管理上

			baseResp := pack.BuildChatBaseResp(errno.WebSocketSuccess)
			resp, _ := json.Marshal(baseResp)
			_ = client.Socket.WriteMessage(websocket.TextMessage, resp)

		case client := <-Manager.Unregister:
			log.Printf("%v:offline\n", client.ID)
			baseResp := pack.BuildChatBaseResp(errno.WebSocketLogoutSuccess)
			resp, _ := json.Marshal(baseResp)
			_ = client.Socket.WriteMessage(websocket.TextMessage, resp)
			close(client.Send)                 //close chan
			delete(Manager.Clients, client.ID) //delete map

		case broadcast := <-Manager.Broadcast:
			if broadcast.Type == 1 {
				marshalMsg := broadcast.Message
				targetId := broadcast.Client.TargetId

				flag := false
				for id, client := range Manager.Clients {
					if id != targetId {
						continue
					}
					select {
					case client.Send <- marshalMsg: //Write()
						flag = true
					default:
						close(client.Send)
						delete(Manager.Clients, client.ID)
					}
				}
				var replyMsg chat_service.ReplyMsg
				_, _ = replyMsg.UnmarshalMsg(marshalMsg)

				if flag {
					baseResp := pack.BuildChatBaseResp(errno.WebSocketTargetOnlineSuccess)
					resp, _ := json.Marshal(baseResp)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, resp)

					if _, err := db.CreateMessage(&db.Message{
						Uid:      replyMsg.From,
						TargetId: targetId,
						Content:  replyMsg.Content,
					}); err != nil {
						log.Println("database error")
						baseResp = pack.BuildChatBaseResp(errno.ServiceError)
						resp, _ = json.Marshal(baseResp)
						_ = broadcast.Client.Socket.WriteMessage(websocket.CloseMessage, resp)
					}

				} else { //flag==false对方不在线
					baseResp := pack.BuildChatBaseResp(errno.WebSocketTargetOfflineError)
					resp, _ := json.Marshal(baseResp)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, resp)

					_, err := db.CreateMessage(&db.Message{
						Uid:      replyMsg.From,
						TargetId: targetId,
						Content:  replyMsg.Content,
					})
					if err != nil {
						log.Println("database error")
						baseResp = pack.BuildChatBaseResp(errno.ServiceError)
						resp, _ = json.Marshal(baseResp)
						_ = broadcast.Client.Socket.WriteMessage(websocket.CloseMessage, resp)
					}

					if err = cache.SetMessage(context.TODO(), targetId, marshalMsg); err != nil {
						log.Println(err)
						baseResp = pack.BuildChatBaseResp(errno.ServiceError)
						resp, _ = json.Marshal(baseResp)
						_ = broadcast.Client.Socket.WriteMessage(websocket.CloseMessage, resp)
					}
				}
			}
		}
	}
}
