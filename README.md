# Ten-Minutes-App


### [😀 haha~ demo](https://ten-minutes.lotteryjs.com/)

#### Traefik + Docker Deploy

<img src="https://github.com/go-training/drone-golang-example/raw/master/screenshots/traefik+docker+golang.png" width="400">


[golang](https://golang.org/) v1.12.x + [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) v1.x + [gin](https://github.com/gin-gonic/gin) v1.3.x + [mongodb](https://www.mongodb.com/) v4.0.6 + [JSONPlaceholder](http://jsonplaceholder.typicode.com/), [react-admin](https://github.com/marmelab/react-admin)

[使用 Docker](https://github.com/Kirk-Wang/Hello-Gopher/tree/master/mongo)

# Dev
```sh
# api
go run .
# app
cd app
yarn & yarn start
```

# Refs

* [mongo-go-dirver offical examples](https://github.com/mongodb/mongo-go-driver/blob/master/examples/documentation_examples/examples.go)
* [Go by Example](https://gobyexample.com/)
* [gotify/server](https://github.com/gotify/server)
* [gin-jwt](https://github.com/appleboy/gin-jwt)


## redis
- 启动：
  service redisd start
- 关闭：
  service redisd stop
- 查看
  ps -aux | grep redis