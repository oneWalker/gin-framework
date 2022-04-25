package tools

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// CleanHandler define
type CleanHandler func()

// GracefulShutdownAndCleanup method
func GracefulShutdownAndCleanup(cleanHdl CleanHandler) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	logrus.Info("start successfully.")

	<-done
	cleanHdl()
	logrus.Info("stop successfully")
}
