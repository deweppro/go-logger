/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package goLogger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	pool "github.com/deweppro/go-chan-pool"
)

var (
	_ Logger = (*asyncLog)(nil)
)

type (
	asyncLog struct {
		writer io.Writer
		pool   *pool.ChanPool
		async  chan LogMessage
	}
)

func NewAsync() Logger {
	log := &asyncLog{
		writer: os.Stdout,
		pool: &pool.ChanPool{
			Size: 64,
			New: func() interface{} {
				return LogMessage{}
			},
		},
		async: make(chan LogMessage, 1024),
	}

	go log.asyncWrite()

	return log
}

func (_async *asyncLog) asyncWrite() {
	for {
		select {
		case m := <-_async.async:
			b, err := json.Marshal(m)
			if err != nil {
				b = []byte(err.Error())
			}
			_, _ = _async.writer.Write(append(b, nl...))
			_async.pool.Put(m)
		default:
		}
	}
}

func (_async *asyncLog) SetOutput(out io.Writer) {
	_async.writer = out
}

func (_async *asyncLog) Infof(format string, args ...interface{}) {
	m := _async.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "INF", time.Now().Unix(), fmt.Sprintf(format, args...)
	_async.async <- m
}

func (_async *asyncLog) Warnf(format string, args ...interface{}) {
	m := _async.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "WRN", time.Now().Unix(), fmt.Sprintf(format, args...)
	_async.async <- m
}

func (_async *asyncLog) Errorf(format string, args ...interface{}) {
	m := _async.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "ERR", time.Now().Unix(), fmt.Sprintf(format, args...)
	_async.async <- m
}

func (_async *asyncLog) Debugf(format string, args ...interface{}) {
	m := _async.pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = "DBG", time.Now().Unix(), fmt.Sprintf(format, args...)
	_async.async <- m
}
