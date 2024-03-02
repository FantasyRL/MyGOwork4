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

Bonus:
    
对点赞引入redis缓存(评论也引入了，社交写一半懒了)

实现了WebSocket实时聊天(WebSocket接口好像没办法用swagger自动生成，所以接口文档里没有...)