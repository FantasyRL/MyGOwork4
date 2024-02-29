namespace go chat

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Message{
    1:i64 id,
    2:i64 target_id,
    3:i64 from_id,
    4:string content,
    5:string create_time,
}

struct MessageChatReq{
    1:i64 target_id,
}

struct MessageChatResp{
    1:BaseResp base,
}

struct MessageRecordReq{
    1:i64 target_id,
    2:string from_time,
    3:string to_time,
    4:i64 action_type,//todo:群聊
}

struct MessageRecordResp{
    1:BaseResp base,
    2:i64 message_count,
    3:list<Message> record,
}


service ChatHandler{
    MessageChatResp Chat(1: MessageChatReq req) (api.get="/bibi/message/ws"),
    MessageRecordResp MessageRecord(1: MessageRecordReq req) (api.get="/bibi/message/record"),

}