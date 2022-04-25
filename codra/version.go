package codra

import (
	"fmt"

	"github.com/spf13/cobra"
)

//文件编译自动初始化
func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
