namespace go interaction

include "user.thrift"
include "video.thrift"

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Comment {
    1: i64 id,
    2: user.User user,
    3: string content,
    4: string publish_time,
}

struct LikeActionReq{
    1:i64 videoID,
    2:i64 action_type,
}

struct LikeActionResp{
    1:BaseResp base,
}

struct LikeListReq{
    1:i64 page_num,
}

struct LikeListResp{
    1:BaseResp base,
    2:list<video.Video> video_list,
}

service InteractionHandler{
    LikeActionResp LikeAction(1:LikeActionReq req)(api.post="/bibi/interaction/like/action"),
    LikeListResp LikeList(1:LikeListReq req)(api.get="/bibi/interaction/like/list"),

}