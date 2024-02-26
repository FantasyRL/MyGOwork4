package chat_service

import (
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"encoding/json"
	"github.com/hertz-contrib/websocket"
	"log"
)

func (manager *ClientManager) Start() {
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
				message := broadcast.Message
				targetId := broadcast.Client.TargetId

				flag := false
				for id, client := range Manager.Clients {
					if id != targetId {
						continue
					}
					select {
					case client.Send <- message: //Write()
						flag = true
					default:
						close(client.Send)
						delete(Manager.Clients, client.ID)
					}
				}

				if flag {
					baseResp := pack.BuildChatBaseResp(errno.WebSocketTargetOnlineSuccess)
					resp, _ := json.Marshal(baseResp)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, resp)
					//err := InsertMsg(conf.MongoDBName, broadcast.Client.ID, string(message), 1, int64(3*month))
					//if err != nil {
					//	fmt.Println("插入消息失败", err)
					//}
				} else { //flag==false对方不在线
					baseResp := pack.BuildChatBaseResp(errno.WebSocketTargetOfflineError)
					resp, _ := json.Marshal(baseResp)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, resp)
					//err := InsertMsg(conf.MongoDBName, broadcast.Client.ID, string(message), 0, int64(3*month))
					//if err != nil {
					//	fmt.Println("插入消息失败", err)
					//}
				}
			}
		}
	}
}
