# MyGOwork4

**bibi-demo** is a small video website backend using hertz(Thrift)+gorm(mysql)+redis(+msgp)+oss(aliyun)+gorilla

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

