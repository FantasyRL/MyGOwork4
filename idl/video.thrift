namespace go video

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct User {
    1: i64 id,
    2: string name,
    3: i64 follow_count,
    4: i64 follower_count,
    5: bool is_follow
    6: string avatar,
    7: i64 total_favorited,
    8: i64 video_count,
    9: i64 favorite_count,
}

struct Video{
    1:i64 id,
    2:string title,
    3:User author,
    4:string play_url,
    5:string cover_url,
    6:i64 like_count,
    7:i64 comment_count,
    8:bool is_like,
}

struct PutVideoReq{
    1:binary video_file,
    2:string title,
    3:string cover,
    4:string token,
}

struct PutVideoResp{
    1:BaseResp base,
}

struct ListUserVideoReq{
    1:string token,
}

struct ListUserVideoResp{
    1:BaseResp base,
    2:list<Video> video_list,
}

struct SearchVideoReq{
    1:string param,
}

struct SearchVideoResp{
    1:BaseResp base,
    2:list<Video> video_list,
}

struct HotVideoReq{
}

struct HotVideoReso{
    1:BaseResp base,
    2:list<Video> video_list,
}

service VideoHandler{
    PutVideoResp PutVideo(1:PutVideoReq req)(api.post="/bibi/video/upload"),
    ListUserVideoResp ListVideo(1:ListUserVideoReq req)(api.post="/bibi/video/myvideo"),
    SearchVideoResp SearchVideo(1:SearchVideoReq req)(api.post="/bibi/video/search"),
    HotVideoReq HotVideo(1:HotVideoReq req)(api.get="/bibi/video/hot"),
}