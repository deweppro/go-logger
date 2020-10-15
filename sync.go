/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package goLogger

//go:generate easyjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	pool "github.com/deweppro/go-chan-pool"
)

var (
	_ Logger = (*syncLog)(nil)
)

type (
	syncLog struct {
		writer io.Writer
		pool   *pool.ChanPool
	}
)

func NewSync() Logger {
	return &syncLog{
		writer: os.Stdout,
		pool: &pool.ChanPool{
			Size: 64,
			New: func() interface{} {
				return LogMessage{}
			},
		},
	}
}

func (_sync *syncLog) write(m LogMessage) {
	b, err := json.Marshal(m)
	if err != nil {
		b = []byte(err.Error())
	}
	_, _ = _sync.writer.Write(append(b, nl...))
	_sync.pool.Put(m)
}

func (_sync *syncLog) SetOutput(out io.Writer) {
	_sync.writer = out
}

func (_sync *syncLog) Infof(format string, args ...interface{}) {
	m := _sync.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "INF", time.Now().Unix(), fmt.Sprintf(format, args...)
	_sync.write(m)
}

func (_sync *syncLog) Warnf(format string, args ...interface{}) {
	m := _sync.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "WRN", time.Now().Unix(), fmt.Sprintf(format, args...)
	_sync.write(m)
}

func (_sync *syncLog) Errorf(format string, args ...interface{}) {
	m := _sync.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "ERR", time.Now().Unix(), fmt.Sprintf(format, args...)
	_sync.write(m)
}

func (_sync *syncLog) Debugf(format string, args ...interface{}) {
	m := _sync.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "DBG", time.Now().Unix(), fmt.Sprintf(format, args...)
	_sync.write(m)
}
