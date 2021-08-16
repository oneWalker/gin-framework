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

* [x]  router
  * [x]  定义基本路由规则
  * [x]  定义基本响应格式以及错误抛出code规则
* [x]  service
  * [ ]  mysql操作的基本demo
  * [ ]  mongodb操作的基本demo
  * [ ]  redis操作的基本demo
* [ ]  Database初始化相关：
  * [ ]  mysql
  * [ ]  mogodb
  * [ ]  redis
* [ ]  定时任务
* [ ]  middleware
  * [ ]  数据统一处理中间件
  * [ ]  jwt或者token验证中间件
  * [ ]  数据校验插件
* [ ]  RPC API相关
* [ ]  单元测试
* [ ]  消息中间件
* [x]  安全和跨域配置
* [ ]  全局变量相关
* [ ]  链路追踪
* [ ]  编译程序应用
* [ ]  优雅重启和停止
