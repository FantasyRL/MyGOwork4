namespace go user

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct User {
    1: i64 id,
    2: string name,
    3: optional string email,
    4: i64 follow_count,
    5: i64 follower_count,
    6: bool is_follow,
    7: string avatar,
    8: i64 video_count,
}

struct RegisterReq {
    1: string username,
    2: string email,
    3: string password,
}

struct RegisterResp {
    1: BaseResp base,
    2: i64 user_id,
}

//struct OTP2FAReq{
//    1:i64 uid,
//}

//struct OTP2FAResp{
//    1:BaseResp base,
//}

struct Switch2FAReq{
    1:i64 action_type,
    2:optional string totp,
}

struct Switch2FAResp{
    1:BaseResp base,
}

struct LoginReq {
    1: string username,
    2: string password,
    3: optional string otp,
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
    AvatarResp Avatar(1:AvatarReq req)(api.put="/bibi/user/avatar/upload"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bibi/user/2fa"),
    Switch2FAResp Switch2FA(1:Switch2FAReq req)(api.get="/bibi/user/switch2fa"),
}