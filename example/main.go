/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"os"
	"time"

	goLogger "github.com/deweppro/go-logger"
)

func main() {
	file, err := os.OpenFile("/tmp/demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	async := goLogger.NewAsync()
	async.SetOutput(file)
	async.Infof("async %d", 1)
	async.Warnf("async %d", 2)
	async.Errorf("async %d", 3)
	async.Debugf("async %d", 4)

	sync := goLogger.NewSync()
	sync.SetOutput(file)
	sync.Infof("sync %d", 1)
	sync.Warnf("sync %d", 2)
	sync.Errorf("sync %d", 3)
	sync.Debugf("sync %d", 4)

	<-time.After(time.Second)
}
