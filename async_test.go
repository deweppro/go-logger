/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package goLogger

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAsync(t *testing.T) {
	log := NewAsync()
	require.NotNil(t, log)

	filename, err := ioutil.TempFile(os.TempDir(), "test_new_default-*.log")
	require.NoError(t, err)

	log.SetOutput(filename)
	log.Infof("test %d", 1)
	log.Warnf("test %d", 2)
	log.Errorf("test %d", 3)
	log.Debugf("test %d", 4)

	<-time.After(time.Second)

	require.NoError(t, filename.Close())
	data, err := ioutil.ReadFile(filename.Name())
	require.NoError(t, err)
	require.NoError(t, os.Remove(filename.Name()))

	sdata := string(data)
	require.Contains(t, sdata, "\"type\":\"INF\",\"data\":\"test 1\"")
	require.Contains(t, sdata, "\"type\":\"WRN\",\"data\":\"test 2\"")
	require.Contains(t, sdata, "\"type\":\"ERR\",\"data\":\"test 3\"")
	require.Contains(t, sdata, "\"type\":\"DBG\",\"data\":\"test 4\"")
}
