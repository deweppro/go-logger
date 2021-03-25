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
	"sync/atomic"
	"time"

	p "github.com/deweppro/go-chan-pool"
)

//go:generate easyjson

//nolint: golint
const (
	LevelError uint32 = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

var (
	nl   = []byte("\n")
	std  = New()
	pool = &p.ChanPool{
		Size: 64,
		New: func() interface{} {
			return Message{}
		},
	}
)

var _ Logger = (*Log)(nil)

type (
	//Logger base interface
	Logger interface {
		SetOutput(out io.Writer)
		SetLevel(v uint32)
		Fatalf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Debugf(format string, args ...interface{})
		Close()
	}

	//Message model
	//easyjson:json
	Message struct {
		Time int64  `json:"time"`
		Type string `json:"type"`
		Data string `json:"data"`
	}

	//Log base model
	Log struct {
		status uint32
		writer io.Writer
		cmsg   chan []byte
		close  chan struct{}
	}
)

//Default logger
func Default() *Log {
	return std
}

//New init new logger
func New() *Log {
	l := &Log{
		status: LevelError,
		writer: os.Stdout,
		cmsg:   make(chan []byte, runtime.GOMAXPROCS(0)*1024),
		close:  make(chan struct{}),
	}
	go l.queue()
	return l
}

func (l *Log) send(level, format string, args ...interface{}) {
	m := pool.Get().(Message)
	m.Type, m.Time, m.Data = level, time.Now().Unix(), fmt.Sprintf(format, args...)

	b, err := json.Marshal(m)
	if err != nil {
		b = []byte(err.Error())
	}

	select {
	case l.cmsg <- b:
	default:
	}

	pool.Put(m)
}

func (l *Log) write(b []byte) {
	_, _ = l.writer.Write(append(b, nl...))
}

func (l *Log) queue() {
	for {
		select {
		case m, ok := <-l.cmsg:
			if ok {
				l.write(m)
			} else {
				close(l.close)
				return
			}
		}
	}
}

//Close waiting for all messages to finish recording
func (l *Log) Close() {
	close(l.cmsg)
	<-l.close
}

//SetOutput change writer
func (l *Log) SetOutput(out io.Writer) {
	l.writer = out
}

//SetLevel change log level
func (l *Log) SetLevel(v uint32) {
	atomic.StoreUint32(&l.status, v)
}

//GetLevel getting log level
func (l *Log) GetLevel() uint32 {
	return atomic.LoadUint32(&l.status)
}

//Infof info message
func (l *Log) Infof(format string, args ...interface{}) {
	if atomic.LoadUint32(&l.status) >= LevelInfo {
		l.send("INF", format, args...)
	}
}

//Warnf warning message
func (l *Log) Warnf(format string, args ...interface{}) {
	if atomic.LoadUint32(&l.status) >= LevelWarn {
		l.send("WRN", format, args...)
	}
}

//Errorf error message
func (l *Log) Errorf(format string, args ...interface{}) {
	l.send("ERR", format, args...)
}

//Debugf debug message
func (l *Log) Debugf(format string, args ...interface{}) {
	if atomic.LoadUint32(&l.status) >= LevelDebug {
		l.send("DBG", format, args...)
	}
}

//Fatalf fatal message and exit
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.send("FAT", format, args...)
	l.Close()
	os.Exit(1)
}

//SetOutput change writer
func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

//SetLevel change log level
func SetLevel(v uint32) {
	std.SetLevel(v)
}

//GetLevel getting log level
func GetLevel() uint32 {
	return std.GetLevel()
}

//Close waiting for all messages to finish recording
func Close() {
	std.Close()
}

//Infof info message
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

//Warnf warning message
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

//Errorf error message
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

//Debugf debug message
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

//Fatalf fatal message and exit
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}
