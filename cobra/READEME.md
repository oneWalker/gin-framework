## Codra的主要目的

主要使用main.go或者类似的文件中，用以进行启动相关的命令行方式进行启动。
在此文件中，运行codra和运行最外层的main.go文件完全一样，但配置的来源方式不一样。主文件main.go的环境变量来源于.env

1. 创建rootCmd，也可以在init（）函数中定义标志和处理配置
2. 创建main.go，调用rootcmd中的相关东西,原则上main.go是放在项目的home下面的
3. 添加其他命令，如version.go
4. 编译与运行

使用标志
1. 使用持久化的标志
   ```go
    rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
   ```
2. 使用本地标志:本地标志只能在它所绑定的命令上使用，--source只能在rootCmd上引用，而不能在rootCmd的子命令上进行使用
   ```go
    rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
   ```
3. 将标志绑定到viper,这样就可以使用viper.Get()获取标志的值
   ```go
    var author string

    func init() {
    rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
    viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
    }
   ```
4. 设置标志为必选：当标志为必选，但未提供标志的时，cobra会报错
   ```go
    rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
    rootCmd.MarkFlagRequired("region")
   ```

## Cobra 也内置了一些验证函数：
* NoArgs：如果存在任何非选项参数，该命令将报错。
* ArbitraryArgs：该命令将接受任何非选项参数。
* OnlyValidArgs：如果有任何非选项参数不在 Command 的 ValidArgs 字段中，该命令将报错。
* MinimumNArgs(int)：如果没有至少 N 个非选项参数，该命令将报错。MaximumNArgs(int)：如果有多于 N 个非选项参数，该命令将报错。
* ExactArgs(int)：如果非选项参数个数不为 N，该命令将报错。
* ExactValidArgs(int)：如果非选项参数的个数不为 N，或者非选项参数不在 Command 的 ValidArgs 字段中，该命令将报错。
* RangeArgs(min, max)：如果非选项参数的个数不在 min 和 max 之间，该命令将报错。

* 使用与定义验证函数
```go
var cmd = &cobra.Command{
  Short: "hello",
  Args: cobra.MinimumNArgs(1), // 使用内置的验证函数
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```

* 类似于RUN函数的钩子函数
  * PresistentPreRun
  * PreRun
  * Run
  * PostRun
  * PersistentPostRun
