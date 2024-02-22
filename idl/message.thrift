namespace go message

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Message{
    1:i64 id,
    2:i64 target_id,
    3:i64 from_id,
    4:string content,
    5:i64 create_time,
}

struct MessageChatReq{
    1:i64 target_id,
}

struct MessageChatResp{
    1:BaseResp base,
    2:list<Message> message_list,
}

struct MessageActionReq{
    1:i64 target_id,
    2:string content,
    3:i64 action_type,//todo:群聊
}

struct MessageActionResp{
    1:BaseResp base,
}

service MessageHandler{
    MessageChatResp MessageChat(1: MessageChatReq req) (api.get="/bibi/message/chat/"),
//    Get the chat message records
    MessageActionResp MessageAction(1: MessageActionReq req) (api.post="/bibi/message/action/"),
//    Send Message
}