# go-logger

[![Coverage Status](https://coveralls.io/repos/github/deweppro/go-logger/badge.svg?branch=main)](https://coveralls.io/github/deweppro/go-logger?branch=main)
[![Release](https://img.shields.io/github/release/deweppro/go-logger.svg?style=flat-square)](https://github.com/deweppro/go-logger/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-logger)](https://goreportcard.com/report/github.com/deweppro/go-logger)
[![Build Status](https://travis-ci.com/deweppro/go-logger.svg?branch=main)](https://travis-ci.com/deweppro/go-logger)


# How to use it

```go
package main

import (
	"os"
	"time"

	"github.com/deweppro/go-logger"
)

func main() {
	logger.SetOutput(os.Stdout)

	logger.Infof("sync %d", 1)
	logger.Warnf("sync %d", 2)
	logger.Errorf("sync %d", 3)
	logger.Debugf("sync %d", 4)

	logger.Close()
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