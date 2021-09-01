package global

// 定义一个全局的包，大写变量可导出，实现全局变量的目的；小写可用包.参数
// 可以考虑将所有需要全局和初始化的都放入该目录中，比如：数据库，校验统一返回等等
// 将初始化还是定义在原来的地方，只是将其引用到该目录下
// 使用 global.DemoString
var (
	DemoString string
)