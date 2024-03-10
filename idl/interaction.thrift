namespace go interaction

include "user.thrift"
include "video.thrift"

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Comment {
    1: i64 id,
    2: i64 video_id,
    3: user.User user,
    4: string content,
    5: string publish_time,
}

struct LikeActionReq{
    1:i64 video_id,
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
    2:i64 video_count,
    3:list<video.Video> video_list,
}

struct CommentCreateReq{
    1:i64 video_id,
    2:string content,
}

struct CommentCreateResp{
    1:BaseResp base,
    2:Comment comment,
}

struct CommentDeleteReq{
    1:i64 video_id,
    2:i64 comment_id,
}

struct CommentDeleteResp{
    1:BaseResp base,
}

struct CommentListReq{
    1:i64 video_id,
    2:i64 page_num,
}

struct CommentListResp{
    1:BaseResp base,
    2:i64 comment_count,
    3:list<Comment> comment_list,
}

service InteractionHandler{
    LikeActionResp LikeAction(1:LikeActionReq req)(api.post="/bibi/interaction/like/action"),
    LikeListResp LikeList(1:LikeListReq req)(api.get="/bibi/interaction/like/list"),
    CommentCreateResp CommentCreate(1:CommentCreateReq req)(api.post="/bibi/interaction/comment/create"),
    CommentDeleteResp CommentDelete(1:CommentDeleteReq req)(api.post="/bibi/interaction/comment/delete"),
    CommentListResp CommentList(1:CommentListReq req)(api.post="/bibi/interaction/comment/list"),
}