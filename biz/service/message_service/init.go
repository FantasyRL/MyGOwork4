package message_service

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

			baseResp := pack.BuildMessageBaseResp(errno.WebSocketSuccess)
			resp, _ := json.Marshal(baseResp)
			_ = client.Socket.WriteMessage(websocket.TextMessage, resp)

		case client := <-Manager.Unregister:
			log.Printf("%v:offline\n", client.ID)

			if _, ok := Manager.Clients[client.ID]; ok { //一般应该不会有不ok的情况吧...
				baseResp := pack.BuildMessageBaseResp(errno.WebSocketError)
				resp, _ := json.Marshal(baseResp)
				_ = client.Socket.WriteMessage(websocket.TextMessage, resp)
				close(client.Message)              //close chan
				delete(Manager.Clients, client.ID) //delete map
			}

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
					case client.Message <- message:
						flag = true
					default:
						close(client.Message)
						delete(Manager.Clients, client.ID)
					}
				}

				if flag {
					replyMsg := &ReplyMsg{
						Code:    errno.SuccessCode,
						From:    broadcast.Client.TargetId,
						Content: "对方在线",
					}
					resp, _ := json.Marshal(replyMsg)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, resp)
					//err := InsertMsg(conf.MongoDBName, broadcast.Client.ID, string(message), 1, int64(3*month))
					//if err != nil {
					//	fmt.Println("插入消息失败", err)
					//}
				} else { //flag==false对方不在线
					log.Println("对方不在线")
					baseResp := pack.BuildMessageBaseResp(errno.WebSocketTargetOfflineError)
					resp, _ := json.Marshal(baseResp)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, resp)
					//err := InsertMsg(conf.MongoDBName, broadcast.Client.ID, string(message), 0, int64(3*month))
					//if err != nil {
					//	fmt.Println("插入消息失败", err)
					//}
				}
			}
			//else if broadcast.Type == 4 { //在群聊发消息
			//	message := broadcast.Message
			//	gid := broadcast.Client.GroupID
			//	for _, conn := range Manager.Clients {
			//		if conn.GroupID != gid {
			//			continue
			//		}
			//		select {
			//		case conn.Send <- message:
			//		default:
			//			close(conn.Send)
			//			delete(Manager.Clients, conn.ID)
			//		}
			//	}
			//	err := InsertMsg(conf.MongoDBName, gid, string(message), 1, int64(3*month))
			//	if err != nil {
			//		fmt.Println("插入消息失败", err)
			//	}
			//}

		}
	}
}
