# MyGOwork4

**bibi-demo** is a small video website backend using hertz(Thrift)+gorm(mysql)+redis(+msgp)+oss(aliyun)

## deploy by docker(net=host)
`bash`
```
docker-compose up -d # 启动相关容器
docker build -t bibi-demo . # 构建镜像
docker run -d --net=host bibi-demo go run bibi # 运行程序
```

使用：将docs/swagger.* 丢到apifox/postman，然后就能用了(**Authorization格式**:Bearer {token})


