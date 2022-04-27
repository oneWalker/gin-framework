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
  * [x]  jwt或者token验证中间件
  * [x]  数据校验插件
  * [x]  返回error代码规则化
* [x] RPC
* [ ]  单元测试
  * 参考demo<https://zhuanlan.zhihu.com/p/90632661>
  * 利用Go gets可自动生成相关的单元测试的代码
  * [ ] xmlDemo中的单元测试
  * [ ] jsonDemo中的单元测试
* [X]  压力测试： 压力测试文件参考：`service/xmlDemobech_test.go`
* [x]  消息中间件
* [x]  安全和跨域配置
* [x]  全局变量相关
* [x]  区分开发环境和正式环境
* [ ]  链路追踪
* 编译程序应用
  * [x] Makefile相关
    * 相关学习说明文档：https://time.geekbang.org/column/article/388920
    ```makefile
    target ...: prerequisites ...
        command
        ...
        ...
    ```

* [x]  优雅重启和停止
  * b4go1.8:[manners](https://github.com/braintree/manners),[graceful](https://github.com/tylerstillwater/graceful),[grace](https://github.com/facebookarchive/grace)
  * now:内置方法Shutdown()
* [ ] 常用的调试工具，比如Delve
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
  * [x] viper<https://github.com/spf13/viper> 解析配置文件参数解析工具，也具有5个函数支持环境变量：在文件夹viperconf
  * [x] Pflag<https://github.com/govenue/pflag>：命令行参数解析工具
  * [x] cobra<https://github.com/spf13/cobra> 现代化的命令行框架：
    * cobra+viperconf实现的功能和main.go里面的作用是一样的
    * Pflag + Cobra替代方案urfave/cli<https://github.com/urfave/cli>
* [x] SDK Demo
  * 注意初始化项目的时候，项目名字采用：`/github.com/${organization}/${projectName}/${version}`
  * SDK在导出的时候就是导出的当前的主要目录，不需要在其他地方再进行导出
  * 区分：Node.js相关的包的导出函数都是放置在`/${package}/src`下的`.js`文件
* 其他参考资料
  * go career path:`/resources/go_career_path`
  * 保姆级学习教程：<https://geektutu.com/post/geerpc-day7.html>
  * colin在极客时间上的课程和项目：<https://github.com/marmotedu>
  * GoNotes：<https://github.com/xuesongbj/Go-Notes>
  * go-microservices-boilerplate<https://github.com/FeifeiyuM/go-microservices-boilerplate>
  * blog-service<https://github.com/go-programming-tour-book/blog-service>
  * 参考三方包demo:<https://github.com/yeqown/playground/tree/master/gonic>

* 注意，db相关的数据连接放置在pkg/db