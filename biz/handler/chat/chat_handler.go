// Code generated by hertz generator.

package chat

import (
	chat "bibi/biz/model/chat"
	"bibi/biz/service/chat_service"
	"bibi/biz/service/chat_service/monitor"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/websocket"
	"log"
)

// Chat .
// @Summary chat(websocket)
// @Description chat online
// @Accept json/form
// @Produce json
// @Param target_id query int true "目标id"
// @Param Authorization header string true "token"
// @router /bibi/message/ws [GET]
func Chat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req chat.MessageChatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(chat.MessageChatResp)

	v, _ := c.Get("current_user_id")
	id := v.(int64)

	if id == req.TargetID {
		resp.Base = pack.BuildChatBaseResp(errno.ParamError)
		c.JSON(consts.StatusOK, resp)
		return
	}

	var upGrader = websocket.HertzUpgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(c *app.RequestContext) bool {
			return true
		},
	}
	//hertz wcnm
	//HertzHandler receives a websocket connection after the handshake has been completed.
	//在匿名函数里面开启send和recv
	client := new(monitor.Client)

	err = upGrader.Upgrade(c, func(conn *websocket.Conn) {
		client = &monitor.Client{
			ID:       id,
			TargetId: req.TargetID,
			Socket:   conn,
			Send:     make(chan []byte),
		}
		//将用户注册到用户管理上
		monitor.Manager.Register <- client
		err = client.IfNotReadMessage(id)
		if err != nil {
			log.Println(err)
			return
		}
		go client.Read()
		go client.Write()
		forever := make(chan bool)
		<-forever //直到conn被关闭才会退出
	})
	if err != nil {
		log.Println("upgrade:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
}

// MessageRecord .
// @Summary message_record
// @Description get message record
// @Accept json/form
// @Produce json
// @Param target_id query int true "目标id"
// @Param from_time query string true "2024-02-29"
// @Param to_time query string true "2024-03-01"
// @Param Authorization header string true "token"
// @Param action_type query int true "1"
// @router /bibi/message/record [GET]
func MessageRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req chat.MessageRecordReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(chat.MessageRecordResp)

	v, _ := c.Get("current_user_id")
	id := v.(int64)

	switch req.ActionType {
	case 1:
		msgList, count, err := chat_service.NewMessageService(ctx).MessageRecord(&req, id)
		if err != nil {
			resp.Base = pack.BuildChatBaseResp(err)
			break
		}
		resp.Base = pack.BuildChatBaseResp(nil)
		resp.MessageCount = count
		resp.Record = chat_service.BuildMessageResp(msgList)
	default:
		resp.Base = pack.BuildChatBaseResp(errno.ParamError)
	}

	c.JSON(consts.StatusOK, resp)
}
