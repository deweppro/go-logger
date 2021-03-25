package main

import (
	"os"

	"github.com/deweppro/go-logger"
)

func main() {
	logger.SetOutput(os.Stdout)

	logger.Infof("async %d", 1)
	logger.Warnf("async %d", 2)
	logger.Errorf("async %d", 3)
	logger.Debugf("async %d", 4)
	logger.Infof("sync %d", 1)
	logger.Warnf("sync %d", 2)
	logger.Errorf("sync %d", 3)
	logger.Debugf("sync %d", 4)

	logger.Close()
}
