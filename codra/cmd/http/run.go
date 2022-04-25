package http

import (
	"context"
	"gin-practice/initialize"
	"gin-practice/pkg/db/mongodb"
	"gin-practice/pkg/db/mysql"
	"gin-practice/pkg/tools"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCommand .
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		DisableFlagParsing: true, // 不解析参数
		Use:                "http",
		Short:              "Run HTTP server",
		Run: func(_ *cobra.Command, _ []string) {

			mongodb.Init()
			mysql.Init()

			router := initialize.Routers()
			srv := &http.Server{
				Addr:    viper.GetString("listen"),
				Handler: router,
			}

			logrus.Info("Server Starting...")
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					logrus.Fatalf("Server start failed: %s", err)
				}
			}()
			logrus.Infof("Server started at http://%s", viper.GetString("listen"))

			// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
			tools.GracefulShutdownAndCleanup(func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := srv.Shutdown(ctx); err != nil {
					logrus.Fatalf("Server Shutdown: %v", err)
				}
				mongodb.Close()
				mysql.Close()
			})
		},
	}

	return cmd

}
