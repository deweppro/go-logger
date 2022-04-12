package main

import (
	"os"

	"github.com/deweppro/go-logger"
)

func main() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.LevelDebug)

	logger.Infof("log %s", "info")
	logger.Warnf("log %s", "warn")
	logger.Errorf("log %s", "error")
	logger.Debugf("log %s", "debug")
	logger.WithFields(logger.Fields{"a": "b"}).Infof("with context")

	logger.Close()
}
