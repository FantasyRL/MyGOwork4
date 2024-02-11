namespace go video

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Video{
    1:i64 id,
    2:string title,
//    3:User author,
    3:i64 uid,
    4:string play_url,
    5:string cover_url,
    6:i64 like_count,
    7:i64 comment_count,
    8:bool is_like,
    9:string publish_time,
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

struct HotVideoReso{
    1:BaseResp base,
    2:list<Video> video_list,
}

service VideoHandler{
    PutVideoResp PutVideo(1:PutVideoReq req)(api.post="/bibi/video/upload"),
    ListUserVideoResp ListVideosByID(1:ListUserVideoReq req)(api.post="/bibi/video/myvideo"),
    SearchVideoResp SearchVideo(1:SearchVideoReq req)(api.post="/bibi/video/search"),
    HotVideoReq HotVideo(1:HotVideoReq req)(api.get="/bibi/video/hot"),
}