# MyGOwork4

**bibi-demo** is a small video website backend using hertz(hz-gen、jwt、websocket)+gorm(mysql)+redis(+msgp)+oss(aliyun)

## deploy by docker(net=host)
`bash`
```
#oss的配置是什么都没有的(上次传github上瞬间就被警告了)
mv pkg/conf/config-example.yaml pkg/conf/config.yaml
docker-compose up -d # 启动相关容器
docker build -t bibi-demo . # 构建镜像
docker run -d --net=host bibi-demo go run bibi # 运行程序
```

使用：将docs/swagger.* 丢到apifox/postman，然后就能用了(**Header:Authorization格式**:Bearer {token})

本项目的构建历程：抄项目架构、抄结构体、看demo遇到不会的学一下然后继续抄...

真是一场酣畅淋漓的ctrl+c.jpg

## 完成情况：
分页管理：做了一部分，后面社交之类懒得做了

未遵循接口文档：hertz-jwt默认是header:Authorization，当我发现修改方法时项目已经基本完成了，肥肠爆芡

Bonus:
    
对点赞引入redis缓存(评论也引入了，社交写一半懒了)

实现WebSocket(WebSocket接口好像没办法用swagger自动生成，所以接口文档里没有...)