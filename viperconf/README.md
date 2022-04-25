# Viper的常见使用说明

## Viper读取位置的优先级从高到低

1. 通过 viper.Set 函数显示设置的配置
2. 命令行参数
3. 环境变量
4. 配置文件
5. Key/Value存储
6. 默认值

   
**viper配置键不区分大小写**

1. 设置默认的配置文件名。
   ```go
    viper.SetDefault("ContentDir", "content")
    viper.SetDefault("LayoutDir", "layouts")
    viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
   ```

2. 读取配置文件。(在开发中常见的使用方式)
   支持 JSON、TOML、YAML、YML、Properties、Props、Prop、HCL、Dotenv、Env 格式的配置文件
   ```go
    package main
    import (
    "fmt"
    "github.com/spf13/pflag"
    "github.com/spf13/viper"
    )

    var (
        cfg  = pflag.StringP("config", "c", "", "Configuration file.")
        help = pflag.BoolP("help", "h", false, "Show this help message.")
    )

    func main() {
        pflag.Parse()
        if *help {
            pflag.Usage()
            return
        }

        // 从配置文件中读取配置
        if *cfg != "" {
            viper.SetConfigFile(*cfg)   // 指定配置文件名
            viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
        } else {
            viper.AddConfigPath(".")          // 把当前目录加入到配置文件的搜索路径中
            viper.AddConfigPath("$HOME/.iam") // 配置文件搜索路径，可以设置多个配置文件搜索路径
            viper.SetConfigName("config")     // 配置文件名称（没有文件扩展名）
        }

        if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
            panic(fmt.Errorf("Fatal error config file: %s \n", err))
        }

        fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
    }
   ```
   
3. 监听和重新读取配置文件。
   ```go
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
    // 配置文件发生变更之后会调用的回调函数
    fmt.Println("Config file changed:", e.Name)
    })
   ```
4. 从io.Reader 读取配置。`viper.Set("username","colin")`
5. 从环境变量读取。
   * AutomaticEnv
   * BindEnv(input …string) error
   * SetEnvPrefix(in string)
   * SetEnvKeyReplacer(r *strings.Replacer)
   * AllowEmptyEnv(allowEmptyEnv bool)
    ```go
    os.Setenv("VIPER_USER_SECRET_ID", "QLdywI2MrmDVjSSv6e95weNRvmteRjfKAuNV")
    os.Setenv("VIPER_USER_SECRET_KEY", "bVix2WBv0VPfrDrvlLWrhEdzjLpPCNYb")

    viper.AutomaticEnv()                                             // 读取环境变量
    viper.SetEnvPrefix("VIPER")                                      // 设置环境变量前缀：VIPER_，如果是viper，将自动转变为大写。
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将viper.Get(key) key字符串中'.'和'-'替换为'_'
    viper.BindEnv("user.secret-key")
    viper.BindEnv("user.secret-id", "USER_SECRET_ID") // 绑定环境变量名到key
    ```

6. 从命令行标志读取。
   ```go
    // 绑定单个标志
    viper.BindPFlag("token", pflag.Lookup("token")) 

    //绑定标志集
    viper.BindPFlags(pflag.CommandLine)
    ```
7. 从远程 Key/Value 存储读取

## 常见的读取配置文件

1. 访问嵌套的键`viper.GetString("datastore.metric.host") // 返回文件中对应的值，文件可以是json之类的`
2. 反序列化：将所有值解析到结构体或者map
    ```go
    type config struct {
    Port int
    Name string
    PathMap string `mapstructure:"path_map"`
    }
    var C config

    err := viper.Unmarshal(&C)//也可使用UnmarshalKey
    if err != nil {
    t.Fatalf("unable to decode into struct, %v", err)
    }
    ```
    如果键盘本身含有默认.分隔符号，则需要修改分割符号
    ```go
    v := viper.NewWithOptions(viper.KeyDelimiter("::"))

    v.SetDefault("chart::values", map[string]interface{}{
        "ingress": map[string]interface{}{
            "annotations": map[string]interface{}{
                "traefik.frontend.rule.type":                 "PathPrefix",
                "traefik.ingress.kubernetes.io/ssl-redirect": "true",
            },
        },
    })

    type config struct {
    Chart struct{
            Values map[string]interface{}
        }
    }

    var C config

    v.Unmarshal(&C)
    ```

3. 序列化成字符串
    ```go
    v := viper.NewWithOptions(viper.KeyDelimiter("::"))

    v.SetDefault("chart::values", map[string]interface{}{
        "ingress": map[string]interface{}{
            "annotations": map[string]interface{}{
                "traefik.frontend.rule.type":                 "PathPrefix",
                "traefik.ingress.kubernetes.io/ssl-redirect": "true",
            },
        },
    })

    type config struct {
    Chart struct{
            Values map[string]interface{}
        }
    }

    var C config

    v.Unmarshal(&C)
    ```