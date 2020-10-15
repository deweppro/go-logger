/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package goLogger

//go:generate easyjson

var nl = []byte("\n")

type (
	//easyjson:json
	LogMessage struct {
		Time int64  `json:"time"`
		Type string `json:"type"`
		Data string `json:"data"`
	}
)
