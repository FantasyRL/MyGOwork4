BibiBibi
├─.hz
├─Dockerfile
├─README.md
├─bibi
├─build.sh
├─docker-compose.yml
├─go.mod
├─go.sum
├─main.go
├─router.go
├─router_gen.go
├─treer.md
├─script
|   └bootstrap.sh
├─pkg
|  ├─utils
|  |   ├─sender
|  |   |   └send.go
|  |   ├─pwd
|  |   |  └pwd.go
|  |   ├─otp2fa
|  |   |   └totp.go
|  |   ├─oss
|  |   |  └oss.go
|  ├─pack
|  |  ├─build_base.go
|  |  └pack.go
|  ├─errno
|  |   └errno.go
|  ├─conf
|  |  ├─config-example.yaml
|  |  ├─config.go
|  |  ├─config.yaml
|  |  ├─sql
|  |  |  └init.sql
├─idl
|  ├─chat.thrift
|  ├─follow.thrift
|  ├─interaction.thrift
|  ├─user.thrift
|  └video.thrift
├─docs
|  ├─docs.go
|  ├─swagger.json
|  └swagger.yaml
├─biz
|  ├─service
|  |    ├─video_service
|  |    |       ├─get_like_video.go
|  |    |       ├─hot_video.go
|  |    |       ├─list_video.go
|  |    |       ├─model.go
|  |    |       ├─search_video.go
|  |    |       └upload_video.go
|  |    ├─user_service
|  |    |      ├─2fa.go
|  |    |      ├─avatar.go
|  |    |      ├─basic.go
|  |    |      ├─get_user.go
|  |    |      └model.go
|  |    ├─interaction_service
|  |    |          ├─comment_action.go
|  |    |          ├─comment_list.go
|  |    |          ├─like_action.go
|  |    |          ├─like_count.go
|  |    |          ├─like_list.go
|  |    |          └model.go
|  |    ├─follow_service
|  |    |       ├─follow_action.go
|  |    |       ├─follow_count.go
|  |    |       ├─follow_list.go
|  |    |       └model.go
|  |    ├─chat_service
|  |    |      ├─model.go
|  |    |      ├─model_reply.go
|  |    |      ├─record.go
|  |    |      ├─reply_msgp.go
|  |    |      ├─monitor
|  |    |      |    ├─chat.go
|  |    |      |    ├─init.go
|  |    |      |    └model.go
|  ├─router
|  |   ├─register.go
|  |   ├─video
|  |   |   ├─middleware.go
|  |   |   └video.go
|  |   ├─user
|  |   |  ├─middleware.go
|  |   |  └user.go
|  |   ├─interaction
|  |   |      ├─interaction.go
|  |   |      └middleware.go
|  |   ├─follow
|  |   |   ├─follow.go
|  |   |   └middleware.go
|  |   ├─chat
|  |   |  ├─chat.go
|  |   |  └middleware.go
|  ├─mw
|  | ├─jwt
|  | |  └jwt.go
|  ├─model
|  |   ├─video
|  |   |   └video.go
|  |   ├─user
|  |   |  └user.go
|  |   ├─interaction
|  |   |      └interaction.go
|  |   ├─follow
|  |   |   └follow.go
|  |   ├─chat
|  |   |  └chat.go
|  ├─handler
|  |    ├─ping.go
|  |    ├─video
|  |    |   └video_handler.go
|  |    ├─user
|  |    |  └user_handler.go
|  |    ├─interaction
|  |    |      └interaction_handler.go
|  |    ├─follow
|  |    |   └follow_handler.go
|  |    ├─chat
|  |    |  └chat_handler.go
|  ├─dal
|  |  ├─db
|  |  | ├─chat_db.go
|  |  | ├─chat_msgp.go
|  |  | ├─comment_db.go
|  |  | ├─comment_msgp.go
|  |  | ├─follow_db.go
|  |  | ├─init.go
|  |  | ├─like_db.go
|  |  | ├─user_db.go
|  |  | └video_db.go
|  |  ├─cache
|  |  |   ├─chat_cache.go
|  |  |   ├─comment_cache.go
|  |  |   ├─follow_cache.go
|  |  |   ├─init.go
|  |  |   └like_cache.go