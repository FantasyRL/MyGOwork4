namespace go user

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct User {
    1: i64 id,
    2: string name,
    3: i64 follow_count,
    4: i64 follower_count,
    5: bool is_follow,
    6: string avatar,
    7: i64 video_count,
}

struct RegisterReq {
    1: string username,
    2: string password,
}

struct RegisterResp {
    1: BaseResp base,
    2: i64 user_id,
}

struct LoginReq {
    1: string username,
    2: string password,
}

struct LoginResp {
    1: BaseResp base,
    2: User user,
    3: string token,
}

struct InfoReq {
}

struct InfoResp {
    1: BaseResp base,
    2: User user,
}

struct AvatarReq{
    1:binary avatar_file,
}
struct AvatarResp{
    1: BaseResp base,
    2: User user,
}
service UserHandler {
    RegisterResp Register(1: RegisterReq req)(api.post="/bibi/user/register/"),
    LoginResp Login(1: LoginReq req)(api.post="/bibi/user/login/"),
    InfoResp Info(1: InfoReq req)(api.get="/bibi/user/"),
    AvatarResp Avatar(1:AvatarReq req)(api.put="/bibi/user/avatar/upload")
}