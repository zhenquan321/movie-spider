## 电影爬虫

### [😀 点击访问电影 https://lmongo.com/view/movie](https://lmongo.com/view/movie)

#### Traefik + Docker Deploy

[golang](https://golang.org/) v1.12.x + [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) v1.x + [gin](https://github.com/gin-gonic/gin) v1.3.x + [mongodb](https://www.mongodb.com/) v4.0.6 + [JSONPlaceholder](http://jsonplaceholder.typicode.com/)

[使用 Docker](https://github.com/Kirk-Wang/Hello-Gopher/tree/master/mongo)

##### 运行 Dev
```sh
# api
go run .
# app
```

##### 发布 Pro
```sh

go build app.go

nohup ./app  </dev/null &>/dev/null & 

```

##### redis
- 启动：
  service redisd start
- 关闭：
  service redisd stop
- 查看
  ps -aux | grep redis

#### mac -redis 
- redis-server
