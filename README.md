# Ten-Minutes-App


### [😀 haha~ demo](https://ten-minutes.lotteryjs.com/)

#### Traefik + Docker Deploy

<img src="https://github.com/go-training/drone-golang-example/raw/master/screenshots/traefik+docker+golang.png" width="400">


[golang](https://golang.org/) v1.12.x + [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) v1.x + [gin](https://github.com/gin-gonic/gin) v1.3.x + [mongodb](https://www.mongodb.com/) v4.0.6 + [JSONPlaceholder](http://jsonplaceholder.typicode.com/)

[使用 Docker](https://github.com/Kirk-Wang/Hello-Gopher/tree/master/mongo)

# Dev
```sh
# api
go run .
# app
```
# Pro
```sh
#守护进程启动
supervisord -c /etc/supervisor/golang.conf 

#查看与停止、重启

supervisorctl
rixingyike  RUNNING  pid 5976, uptime 0:27:47
> shutdown
Really shut the remote supervisord process down y/N? y
Shut down


```


## redis
- 启动：
  service redisd start
- 关闭：
  service redisd stop
- 查看
  ps -aux | grep redis