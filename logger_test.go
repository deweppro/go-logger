/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package logger

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	log := Default()
	require.NotNil(t, log)

	filename, err := ioutil.TempFile(os.TempDir(), "test_new_default-*.log")
	require.NoError(t, err)

	SetOutput(filename)

	Infof("async %d", 1)
	Warnf("async %d", 2)
	Errorf("async %d", 3)
	Debugf("async %d", 4)
	Infof("sync %d", 1)
	Warnf("sync %d", 2)
	Errorf("sync %d", 3)
	Debugf("sync %d", 4)

	log.Close()

	require.NoError(t, filename.Close())
	data, err := ioutil.ReadFile(filename.Name())
	require.NoError(t, err)
	require.NoError(t, os.Remove(filename.Name()))

	sdata := string(data)
	require.Contains(t, sdata, "\"type\":\"INF\",\"data\":\"async 1\"")
	require.Contains(t, sdata, "\"type\":\"WRN\",\"data\":\"async 2\"")
	require.Contains(t, sdata, "\"type\":\"ERR\",\"data\":\"async 3\"")
	require.Contains(t, sdata, "\"type\":\"DBG\",\"data\":\"async 4\"")
	require.Contains(t, sdata, "\"type\":\"INF\",\"data\":\"sync 1\"")
	require.Contains(t, sdata, "\"type\":\"WRN\",\"data\":\"sync 2\"")
	require.Contains(t, sdata, "\"type\":\"ERR\",\"data\":\"sync 3\"")
	require.Contains(t, sdata, "\"type\":\"DBG\",\"data\":\"sync 4\"")
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	ll := New()
	ll.SetOutput(ioutil.Discard)
	wg := sync.WaitGroup{}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		wg.Add(1)
		for p.Next() {
			ll.Infof("hello")
		}
		wg.Done()
	})
	wg.Wait()
	ll.Close()
}
