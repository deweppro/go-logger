/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"

	p "github.com/deweppro/go-chan-pool"
)

//go:generate easyjson

var (
	nl   = []byte("\n")
	std  = New()
	pool = &p.ChanPool{
		Size: 64,
		New: func() interface{} {
			return LogMessage{}
		},
	}
)

type (
	Logger interface {
		SetOutput(out io.Writer)
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Debugf(format string, args ...interface{})
	}

	//easyjson:json
	LogMessage struct {
		Time int64  `json:"time"`
		Type string `json:"type"`
		Data string `json:"data"`
	}

	log struct {
		writer io.Writer
		async  bool
		c      chan LogMessage
		lock   sync.RWMutex
	}
)

func Default() *log {
	return std
}

func New() *log {
	_log := &log{
		writer: os.Stdout,
		async:  false,
		c:      make(chan LogMessage, runtime.GOMAXPROCS(0)*512),
	}
	go _log.routine()
	return _log
}

func (_log *log) Sync() {
	_log.lock.Lock()
	defer _log.lock.Unlock()

	_log.async = false
}

func (_log *log) Async() {
	_log.lock.Lock()
	defer _log.lock.Unlock()

	_log.async = true
}

func (_log *log) is() bool {
	_log.lock.RLock()
	defer _log.lock.RUnlock()

	return _log.async
}

func (_log *log) send(level, format string, args ...interface{}) {
	m := pool.Get().(LogMessage)
	m.Type, m.Time, m.Data = level, time.Now().Unix(), fmt.Sprintf(format, args...)

	if _log.is() {
		select {
		case _log.c <- m:
		default:
			pool.Put(m)
		}
	} else {
		_log.write(m)
	}
}

func (_log *log) write(m LogMessage) {
	_log.lock.RLock()
	defer _log.lock.RUnlock()

	b, err := json.Marshal(m)
	if err != nil {
		b = []byte(err.Error())
	}
	_, _ = _log.writer.Write(append(b, nl...))
	pool.Put(m)
}

func (_log *log) routine() {
	for m := range _log.c {
		_log.write(m)
	}
}

func (_log *log) SetOutput(out io.Writer) {
	_log.lock.Lock()
	defer _log.lock.Unlock()

	_log.writer = out
}

func (_log *log) Infof(format string, args ...interface{}) {
	_log.send("INF", format, args...)
}

func (_log *log) Warnf(format string, args ...interface{}) {
	_log.send("WRN", format, args...)
}

func (_log *log) Errorf(format string, args ...interface{}) {
	_log.send("ERR", format, args...)
}

func (_log *log) Debugf(format string, args ...interface{}) {
	_log.send("DBG", format, args...)
}

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
