# go-logger

[![Coverage Status](https://coveralls.io/repos/github/deweppro/go-logger/badge.svg?branch=master)](https://coveralls.io/github/deweppro/go-logger?branch=master)
[![Release](https://img.shields.io/github/release/deweppro/go-logger.svg?style=flat-square)](https://github.com/deweppro/go-logger/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-logger)](https://goreportcard.com/report/github.com/deweppro/go-logger)
[![Build Status](https://travis-ci.com/deweppro/go-logger.svg?branch=master)](https://travis-ci.com/deweppro/go-logger)


# How to use it

```go
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

	async := logger.NewAsync()
	async.SetOutput(file)
	async.Infof("async %d", 1)
	async.Warnf("async %d", 2)
	async.Errorf("async %d", 3)
	async.Debugf("async %d", 4)

	sync := logger.NewSync()
	sync.SetOutput(file)
	sync.Infof("sync %d", 1)
	sync.Warnf("sync %d", 2)
	sync.Errorf("sync %d", 3)
	sync.Debugf("sync %d", 4)

	<-time.After(time.Second)
}
```

Example log output:
```json
{"time":1602721013,"type":"INF","data":"sync 1"}
{"time":1602721013,"type":"WRN","data":"sync 2"}
{"time":1602721013,"type":"ERR","data":"sync 3"}
{"time":1602721013,"type":"DBG","data":"sync 4"}
...
```