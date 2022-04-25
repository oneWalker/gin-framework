package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	//"github.com/sirupsen/logrus"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

var (
	configFile string
)

func InitConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if configFile != "" {
		// Use config file from the flag.
		Viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "c" (without extension).
		// https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
		_, b, _, _ := runtime.Caller(0)
		Viper.AddConfigPath(filepath.Dir(b))

		Viper.AddConfigPath(home + "/viberconf") // 或者可以使用 viper.AddConfigPath(".")或者viper.AddConfigPath("b")直接读取当前的值
		Viper.SetConfigType("yaml")
		Viper.SetConfigName("config")

	}

	if err := Viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
