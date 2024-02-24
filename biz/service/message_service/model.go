package message_service

import (
	"context"
	"github.com/hertz-contrib/websocket"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService { return &MessageService{ctx: ctx} }

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID           int64
	TargetId     int64
	Socket       *websocket.Conn
	MessageQueue chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

// ClientManager Manager client user
type ClientManager struct {
	Clients    map[int64]*Client //manager
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client //login
	Unregister chan *Client //exit
}

var Manager = ClientManager{
	Clients:    make(map[int64]*Client),
	Broadcast:  make(chan *Broadcast),
	Reply:      make(chan *Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}
