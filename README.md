# MyGOwork4

**bibi-demo** is a small video website backend using hertz(hz-gen、jwt、websocket)+gorm(mysql)+redis+oss(aliyun)

## deploy by docker(net=host)

(使用前请先关闭本机的mysql与redis服务)

`快速启动`
```bash
#oss与email的配置需自行填写
mv pkg/conf/config-example.yaml pkg/conf/config.yaml
docker-compose up -d # 启动相关容器
docker build -t bibi-demo . # 构建镜像
docker run -d --net=host bibi-demo go run bibi # 运行程序
```

使用：将docs/swagger.* 丢到apifox/postman，然后就能用了(**Header:Authorization格式**:Bearer {token})

本项目的构建历程：抄项目架构、抄结构体、看demo遇到不会的学一下然后继续抄...

真是一场酣畅淋漓的ctrl+c.jpg

对于结构体加密存储redis使用了msgp(就一个地方偷懒直接用了JSON) 

(commit 都是瞎写的不要在意...)

## 完成情况：
分页管理：做了一部分，后面社交之类懒得做了

未遵循接口文档：hertz-jwt默认是header:Authorization，当我发现修改方法时项目已经基本完成了，肥肠爆芡

完整项目目录：/treer.md

`docker`
```
├─Dockerfile
├─docker-compose.yml
```

`主要业务`
```
├─biz
|  ├─service    #服务层，负责添加缓存和处理错误
|  ├─router     #hz-gen 路由与注入中间件
|  ├─mw
|  | ├─jwt
|  | |  └jwt.go #hz-jwt
|  ├─model      #hz-gen
|  ├─handler    #hz-gen 负责收发数据与逻辑
|  ├─dal
|  |  ├─db      #mysql
|  |  ├─cache   #redis
├─main.go       #hz-gen 加了一些init
├─pkg           #放了一些utils、错误处理和config
```

`pkg`
```
├─pkg
|  ├─utils
|  |   ├─sender         #发邮件
|  |   ├─pwd            #pwd verify
|  |   ├─otp2fa         #totp
|  |   ├─oss            #store videos in aliyun
|  ├─pack
|  |  ├─build_base.go   #这是一个没有在thrift使用通用base的遗留问题，后续重构时会修改
|  |  └pack.go          #乱放东西，应该都是属于video模块里的
|  ├─errno
|  |   └errno.go        #bytedance提供的errno
|  ├─conf
|  |  ├─config-example.yaml
|  |  ├─config.go       #conf配置，并没有const.go因为我又忘了
|  |  ├─config.yaml     #未传入github
|  |  ├─sql
|  |  |  └init.sql      #docker-compose sql初始化
```

## Bonus:

使用了pquerna/totp进行login时的2FA认证,搭配Google Authenticator扫码食用,同时还对邮件引入了静态html页([sender](./pkg/utils/sender/send.go)里有深夜3点因base64包空值破防时刻)
    
对点赞引入redis缓存(评论也引入了，社交写一半懒了)

实现了WebSocket实时聊天(WebSocket接口好像没办法用swagger自动生成，所以接口文档里没有...)

## Future...

在idl中添加optional以优化response

将会改进comment缓存的逻辑

将会进行重构rpc以改进混沌的handler层

将会更加贴合接口文档需求

将会添加双token(为什么还不添加，是不想吗)