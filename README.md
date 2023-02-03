# go-logger

## Deprecated. Use https://github.com/deweppro/go-sdk

# How to use it

```go
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
```

Example log output:
```json
{"time":1649896276,"lvl":"INF","msg":"log info"}
{"time":1649896276,"lvl":"WRN","msg":"log warn"}
{"time":1649896276,"lvl":"ERR","msg":"log error"}
{"time":1649896276,"lvl":"DBG","msg":"log debug"}
{"time":1649896276,"lvl":"INF","msg":"with context","ctx":{"a":"b"}}

...
```
