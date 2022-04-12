package logger_test

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/deweppro/go-logger"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require.NotNil(t, logger.Default())

	filename, err := ioutil.TempFile(os.TempDir(), "test_new_default-*.log")
	require.NoError(t, err)

	logger.SetOutput(filename)
	logger.SetLevel(logger.LevelDebug)
	require.Equal(t, logger.LevelDebug, logger.GetLevel())

	go logger.Infof("async %d", 1)
	go logger.Warnf("async %d", 2)
	go logger.Errorf("async %d", 3)
	go logger.Debugf("async %d", 4)
	logger.Infof("sync %d", 1)
	logger.Warnf("sync %d", 2)
	logger.Errorf("sync %d", 3)
	logger.Debugf("sync %d", 4)
	logger.WithFields(logger.Fields{"ip": "0.0.0.0"}).Infof("context1")
	logger.WithFields(logger.Fields{"nil": nil}).Infof("context2")
	logger.WithFields(logger.Fields{"func": func() {}}).Infof("context3")

	<-time.After(time.Second * 1)
	logger.Close()

	require.NoError(t, filename.Close())
	data, err := ioutil.ReadFile(filename.Name())
	require.NoError(t, err)
	require.NoError(t, os.Remove(filename.Name()))

	sdata := string(data)
	require.Contains(t, sdata, `"lvl":"INF","msg":"async 1"`)
	require.Contains(t, sdata, `"lvl":"WRN","msg":"async 2"`)
	require.Contains(t, sdata, `"lvl":"ERR","msg":"async 3"`)
	require.Contains(t, sdata, `"lvl":"DBG","msg":"async 4"`)
	require.Contains(t, sdata, `"lvl":"INF","msg":"sync 1"`)
	require.Contains(t, sdata, `"lvl":"WRN","msg":"sync 2"`)
	require.Contains(t, sdata, `"lvl":"ERR","msg":"sync 3"`)
	require.Contains(t, sdata, `"msg":"context1","ctx":{"ip":"0.0.0.0"}`)
	require.Contains(t, sdata, `"msg":"context2","ctx":{"nil":null}`)
	require.Contains(t, sdata, `"msg":"context3","ctx":{"func":"unsupported field value: (func())`)
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	ll := logger.New()
	ll.SetOutput(ioutil.Discard)
	ll.SetLevel(logger.LevelDebug)
	wg := sync.WaitGroup{}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		wg.Add(1)
		for p.Next() {
			ll.WithFields(logger.Fields{"a": "b"}).Infof("hello")
		}
		wg.Done()
	})
	wg.Wait()
	ll.Close()
}
