# EasyFrame

以gin为基础，结合领域驱动概念实现的MVC框架。

## 框架原则
	1.代码是给人看的，所以易读性应该是第一位。
	2.代码要尽量内聚，降低耦合，因此推荐DDD风格的模块目录

## 目录结构

```go

├── app
│   ├── controller
│   ├── middleware
│   └── router.go
├── bin
│   └── main
├── build.sh
├── cmd
│   └── main.go
├── config
│   └── develop.yml
├── dockerfile
│   └── main
├── internal
│   ├── database
│   ├── mlog
│   ├── redis
│   ├── sys
│   ├── utils
│   └── viper
├── services
│   └── user
└── vendor
    ├── github.com
    └── modules.txt
```

## 框架逻辑分层
![](http://123.206.89.80:4999/server/index.php?s=/api/attachment/visitFile/sign/0dc61f6d263969744d2e2b0f6015e7b3&showdoc=.jpg)


app目录为应用层，是一个请求命令真正开始处理的入口。
该层主要有三个功能:
1.路由  将请求转发给指定的控制器
2.路由中间件，可在路由转发前，进行一系列的中间判断。例如：登录校验、签名校验、访问权限控制、访问日志打印、请求耗时统计等。开发者可以根据业务需求自由定制业务功能。
3.控制器层: 主要是进行业务逻辑编排。该层通过调用领域层暴露的功能函数，获得业务数据。
但是该层不处理具体的业务逻辑。

services 为领域层，主要负责业务逻辑的处理。
该层一般由不同的领域模块组成，以订单系统为例:
该层会有 负责用户逻辑的 user领域模块，负责订单逻辑的 order模块。负责交易的 payment领域模块。
每个领域模块的结构基本上相同以user领域模块举例:

└── user
    ├── cache_repository.go
    ├── db_repository.go
    ├── repository.go
    ├── service.go
    ├── user_model.go
    └── user_test.go

主要有:
负责对外提供服务的 service.go
负责定义用户数据模型的 user_model.go
负责用户数据来源的*_repository.go  注意repository可以根据需要设置一个或多个，例如cache_repository.go 可以实现和redis的交互，用做数据缓存仓库。
db_repository.go 可以实现和mysql的交互，用于持久化数据仓库。

一般情况下，service.go中的逻辑想获取用户信息，可以调用user_model提供的方法，而不需要关心该数据是缓存提供的还是持久化仓库提供的。

而整个领域模块对外提供服务只能通过service.go暴露的函数。模块内部的有修改，不会影响到其他模块。
模块之间若有交互逻辑，可以在应用层 即controller中进行处理。以次为规则，我们就可以做到高内聚，低耦合。

甚至开发时，可以一部分开发人员开发底层数据仓库repository，主要关注数据安全存储的问题。一部分开发service.go 主要关注业务逻辑实现的问题。
以下就是一个请求命令途径模块的顺序。
![](http://123.206.89.80:4999/server/index.php?s=/api/attachment/visitFile/sign/a16601f31ed39a675c11d8731a2b1d46&showdoc=.jpg)



## 基础设置说明
本框架对配置、日志、数据库、redis等常用组件进行了封装。采用的多是开源项目

|  组件名称 | git地址 | stars  |  说明 |
| ------------ | ------------ | ------------ | ------------ |
| logrus  |  github.com/sirupsen/logrus |  15.9K |   支持多种日志格式，通过和rotatelog的结合实现了日志切割功能|
|go-redis   |  github.com/go-redis/redis | 9.6k  |  官方推荐 |
| grom  |  github.com/jinzhu/gorm |  20.5k |文档健全、功能完善   |
| viper  |  github.com/spf13/viper |  13.3k |支持多种格式、来源配置   |











