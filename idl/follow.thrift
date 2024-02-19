namespace go follow

include "user.thrift"

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct FollowActionReq{
    1:i64 object_uid,
    2:i64 action_type,
}

struct FollowActionResp{
    1:BaseResp base,
}

struct FollowingListReq{
    1:i64 page_num,
}

struct FollowingListResp{
    1:BaseResp base,
    2:i64 count,
    3:list<user.User> following_list,
}

struct FollowerListReq{
    1:i64 page_num,
}

struct FollowerListResp{
    1:BaseResp base,
    2:i64 count,
    3:list<user.User> follower_list,
}

struct FriendListReq{
    1:i64 page_num,
}

struct FriendListResp{
    1:BaseResp base,
    2:i64 count,
    3:list<user.User> friend_list,
}

service FollowHandler{
    FollowActionResp FollowAction(1:FollowActionReq req)(api.post="/bibi/follow/action"),
    FollowerListResp FollowerList(1:FollowerListReq req)(api.get="/bibi/follow/follower"),
    FollowingListResp FollowingList(1:FollowingListReq req)(api.get="/bibi/follow/following"),
    FriendListResp FriendList(1:FriendListReq req)(api.get="/bibi/follow/friend"),
}
