namespace go video

include "user.thrift"

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Video{
    1:i64 id,
    2:string title,
    3:user.User author,
    4:i64 uid,
    5:string play_url,
    6:string cover_url,
    7:i64 like_count,
    8:i64 comment_count,
    9:i64 is_like,
    10:string publish_time,
}

struct PutVideoReq{
    1:binary video_file,
    2:string title,
    3:binary cover,
}

struct PutVideoResp{
    1:BaseResp base,
}

struct ListUserVideoReq{
    1:i64 page_num,
}

struct ListUserVideoResp{
    1:BaseResp base,
    2:i64 count,
    3:list<Video> video_list,
}

struct SearchVideoReq{
    1:string param,
    2:i64 page_num,
}

struct SearchVideoResp{
    1:BaseResp base,
    2:i64 count,
    3:list<Video> video_list,
}

struct HotVideoReq{
}

struct HotVideoResp{
    1:BaseResp base,
    2:list<Video> video_list,
}

service VideoHandler{
    PutVideoResp PutVideo(1:PutVideoReq req)(api.post="/bibi/video/upload"),
    ListUserVideoResp ListVideo(1:ListUserVideoReq req)(api.post="/bibi/video/published"),
    SearchVideoResp SearchVideo(1:SearchVideoReq req)(api.post="/bibi/video/search"),
    HotVideoResp HotVideo(1:HotVideoReq req)(api.get="/bibi/video/hot"),
}