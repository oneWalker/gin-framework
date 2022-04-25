package codra

import (
	"fmt"
	"gin-practice/codra/cmd/http"
	"gin-practice/codra/cmd/task"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	viperconf "gin-practice/viperconf"
)

// service configcmd
const (
	HTTP = "http"
	TASK = "task"
)

var (
	cfgFile     string
	projectBase string
	userLicense string
	verbose     bool
)

//1st rootCmd Method way
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//启动的时候自动初始化
func init() {
	cobra.OnInitialize(viperconf.InitConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viperconf.Viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viperconf.Viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viperconf.Viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viperconf.Viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viperconf.Viper.SetDefault("license", "apache")
}

//Different way to init
//初始化函数，如果要将其他的启动方式也加入到初始化中，可以在这里进行添加
//NewCommand()等于init+rootCmd
//如果需要使用此函数，需要将Execute()函数中的rootCmd改为当前值

func NewCommand() *cobra.Command {

	//cobra.OnInitialize(viperconf.InitConfig)

	cmd := &cobra.Command{
		Use:   "ebs",
		Short: "Examination Bill Server",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.SetReportCaller(true)
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: func(_ *cobra.Command, _ []string) {
			//通过获取的参数进行相关的加载，不在init初始化进行加载
			if viperconf.Viper.GetString("service") == TASK {
				task.NewCommand().Execute()
			} else {
				http.NewCommand().Execute()
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.yaml", "config file")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")

	// cmd.AddCommand(http.NewCommand())
	// cmd.AddCommand(task.NewCommand())

	return cmd
}
