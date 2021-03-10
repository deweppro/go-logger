/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

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
