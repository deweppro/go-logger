/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"os"
	"time"

	"github.com/deweppro/go-logger"
)

func main() {
	file, err := os.OpenFile("/tmp/demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	logger.SetOutput(file)

	/*
		Synchronous recording
	*/
	logger.Default().Sync()
	logger.Infof("async %d", 1)
	logger.Warnf("async %d", 2)
	logger.Errorf("async %d", 3)
	logger.Debugf("async %d", 4)

	/*
		Asynchronous recording
	*/
	logger.Default().Async()
	logger.Infof("sync %d", 1)
	logger.Warnf("sync %d", 2)
	logger.Errorf("sync %d", 3)
	logger.Debugf("sync %d", 4)

	<-time.After(time.Second)
}
