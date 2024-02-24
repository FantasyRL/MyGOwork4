package message_service

import (
	"bibi/pkg/errno"
	"encoding/json"
	"github.com/hertz-contrib/websocket"
	"log"
)

func (manager *ClientManager) Start() {
	for {
		log.Println("监听管道通信")
		select {
		case conn := <-Manager.Register:
			log.Printf("有新连接 %v\n", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放到用户管理上
			replymsg := ReplyMsg{
				Code:    errno.SuccessCode,
				Content: "连接到服务器",
			}
			msg, _ := json.Marshal(replymsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			log.Printf("断开连接 %v\n", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replymsg := &ReplyMsg{
					Code:    errno.ParamErrCode,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replymsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.MessageQueue)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-Manager.Broadcast:
			if broadcast.Type == 1 {
				message := broadcast.Message
				targetId := broadcast.Client.TargetId
				flag := false
				for id, conn := range Manager.Clients {
					if id != targetId {
						continue
					}
					select {
					case conn.MessageQueue <- message:
						flag = true
					default:
						close(conn.MessageQueue)
						delete(Manager.Clients, conn.ID)
					}
				}
				if flag {
					replymsg := &ReplyMsg{
						Code:    errno.SuccessCode,
						Content: "对方在线应答",
					}
					msg, _ := json.Marshal(replymsg)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
					//err := InsertMsg(conf.MongoDBName, broadcast.Client.ID, string(message), 1, int64(3*month))
					//if err != nil {
					//	fmt.Println("插入消息失败", err)
					//}
				} else {
					log.Println("对方不在线")
					replyMsg := ReplyMsg{
						Code:    errno.ParamErrCode,
						Content: "对方不在线应答",
					}
					msg, _ := json.Marshal(replyMsg)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
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
