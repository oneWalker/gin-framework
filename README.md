# Gin-Practice实现相关功能或者Demo

没有上传vendor文件，所以初始化时需要使用以下命令
`go mod vendor`

* [x]  controller
  * [x]  Query/Uri解析
  * [x]  Form解析：FormData/x-www-x-www-formx-www-form-urlencoded
  * [x]  Json解析
  * [x]  文件上传的基本解析
  * [x]  各种数据格式的响应：json,struct,xml,yaml,protobuf
  * [x]  HTML模版渲染
  * [x]  重定向
  * [x]  异步
* [x]  router
  * [x]  定义基本路由规则
  * [x]  定义基本响应格式以及错误抛出code规则
* service
  * [x]  mysql操作的基本demo
    * [x]  mysql执行后如何对获取到的数组进行修改
    * 常见的orm framework
      * gorm:项目中使用的mysql的主要orm<https://gorm.io/index.html>
      * xorm<https://github.com/go-xorm/xorm>
    * 一些三方工具
      * 从数据库批量生成 models<https://github.com/Shelnutt2/db2struct>
      * 通过 model 生成 CRUD 及 docs<https://github.com/cweagans/apig/tree/dep-conversion>
  * [x]  mongodb操作的基本demo
  * [x]  redis操作的基本demo
    * 常见的库:go-redis(操作更方便，本教程使用<https://pkg.go.dev/github.com/go-redis/redis#section-documentation>)和redigo(官方推荐)
* Database初始化相关：
  * [x]  mysql
  * [x]  mogodb
  * [x]  redis
* [x]  定时任务
  * cron定时:第三方依赖包
    * robfig/cron<https://godoc.org/github.com/robfig/cron>
    * gocron<https://pkg.go.dev/github.com/jasonlvhit/gocron>
    * 定时任务管理系统（开箱即用）<https://github.com/ouqiang/gocron>
  * timer定时<https://golang.org/pkg/time/#pkg-examples>
    * NewTimer函数
    * NewTicker函数
* middleware
  * 数据统一处理中间件：go的参数校验需要根据相应的
  * [ ]  jwt或者token验证中间件
  * [ ]  数据校验插件
* [ ] RPC 统一路由设计
* [x]  单元测试
  * 参考demo<https://zhuanlan.zhihu.com/p/90632661>
* [ ]  消息中间件
* [x]  安全和跨域配置
* [x]  全局变量相关
* [x]  区分开发环境和正式环境
* [ ]  链路追踪
* 编译程序应用
  * [ ] Makefile相关
* [x]  优雅重启和停止
  * b4go1.8:[manners](https://github.com/braintree/manners),[graceful](https://github.com/tylerstillwater/graceful),[grace](https://github.com/facebookarchive/grace)
  * now:内置方法Shutdown()
* 请求第三方接口
  * [x]  HTTP方式请求
    * 参考网站:<https://www.cnblogs.com/Paul-watermelon/p/11386392.html>
  * [x]  RPC方式测试：postman已经支持
* [x] App网络传输序列化协议
  * [x] json json,struct,map之间的转换
    * json to struct
    * struct to json
    * json to Map
    * map to json
    * map to struct
    * struct to map
  * [x] protobuf - 一般提供给RPC
    * RPC例子：https://github.com/marmotedu/gopractise-demo/tree/main/apistyle/greeter
    * grpc的4种类型的服务方法：简单模式，服务端数据流模式，客户端数据流模式和双向数据流模式
    * 目录grpc下
  * [x] XML - xml,struct,json之间的转换
    * XMLtoStruct
    * XMLtoJson
      * XML到Json并不提供直接相关，可使用Struct作为中间转换
      * 第三方工具XML转换成json：<https://github.com/basgys/goxml2json>
    * XMLToMap
      * 通过重写XMl相关的生成和解析函数Marshal和Unmarshal

* Package相关,应用构建相关：https://time.geekbang.org/column/article/395705
  * [x] logrus 日志输送相关框架，与原有的go日志包相互兼容
  * [ ] Pflag：命令行参数解析工具
  * [ ] cobra<https://github.com/spf13/cobra> 现代化的命令行框架
  * [ ] viber<https://github.com/spf13/viber> 配置解析工具
* 其他脚手架参考链接
  * go-microservices-boilerplate<https://github.com/FeifeiyuM/go-microservices-boilerplate>
  * blog-service<https://github.com/go-programming-tour-book/blog-service>
  * 参考学习例子:<https://github.com/yeqown/playground/tree/master/gonic>
