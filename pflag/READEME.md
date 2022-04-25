## pflag包的主要特点

pflag 包与 flag 包的工作原理甚至是代码实现都是类似的，下面是 pflag 相对 flag 的一些优势：

- 支持更加精细的参数类型：例如，flag 只支持 uint 和 uint64，而 pflag 额外支持 uint8、uint16、int32 等类型。
- 支持更多参数类型：ip、ip mask、ip net、count、以及所有类型的 slice 类型。
- 兼容标准 flag 库的 Flag 和 FlagSet：pflag 更像是对 flag 的扩展。
- 原生支持更丰富的功能：支持 shorthand、deprecated、hidden 等高级功能。

## 几种类型的用法
- bool类型
```shell
--flag
--flag = value
--flag value
```
- NoOptDefVal 用法（设置默认值）
```go
//即当age没有输入值时，默认输入的是25
var cliAge = flag.IntP("age", "a",22, "Input Your Age")
flag.Lookup("age").NoOptDefVal = "25"
```
- shorthand:`-`表示shorthand参数，`--`表示正常参数
- 标准化参数名称:通过重写commandLine的相关参数实现
如果我们创建了名称为 --des-detail 的参数，但是用户却在传参时写成了 --des_detail 或 --des.detail 会怎么样？默认情况下程序会报错退出，但是我们可以通过 pflag 提供的 SetNormalizeFunc 功能轻松的解决这个问题；
```go
//将--des_detail 或 --des—detail转化为--des.detail
func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
    from := []string{"-", "_"}
    to := "."
    for _, sep := range from {
        name = strings.Replace(name, sep, to, -1)
    }
    return flag.NormalizedName(name)
}
flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
```
- 把参数标记为即将废弃
```go
// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")
```

- 在帮助文档中隐藏参数
```go
// 在帮助文档中隐藏参数 badflag
flag.CommandLine.MarkHidden("badflag")
```

## Pflag 除了支持单个的 Flag 之外，还支持 FlagSet
- 获取并使用FlagSet的方法
  - 1.调用 NewFlagSet 创建一个 FlagSet。
  ```go
    var version bool
    flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
    flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
  ```
  - 2.使用 Pflag 包定义的全局 FlagSet：CommandLine。实际上 CommandLine 也是由 NewFlagSet 函数创建的。
  ```go
    import (
        "github.com/spf13/pflag"
    )
    pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit.") 
    func BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
        //CommandLine函数的定义：CommandLine is the default set of command-line flags, parsed from os.Args.
        //var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
        flag := CommandLine.VarPF(newBoolValue(value, p), name, shorthand, usage)
        flag.NoOptDefVal = "true"
    }
  ``` 