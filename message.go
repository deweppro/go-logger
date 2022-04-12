package logger

//go:generate easyjson

import p "github.com/deweppro/go-chan-pool"

var poolMessage = &p.ChanPool{
	Size: 1024,
	New: func() interface{} {
		return message{
			Ctx: make(map[string]interface{}),
		}
	},
}

//easyjson:json
type message struct {
	Time    int64                  `json:"time"`
	Level   string                 `json:"lvl"`
	Message string                 `json:"msg"`
	Ctx     map[string]interface{} `json:"ctx,omitempty"`
}

func (v message) Reset() {
	v.Time = 0
	v.Level = ""
	v.Message = ""
	for s := range v.Ctx {
		delete(v.Ctx, s)
	}
}
