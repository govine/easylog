package easylog_test

import (
	"encoding/json"

	"github.com/govine/easylog"
	"github.com/govine/easylog/handler"
)

func ExampleEasylog_SimpleStdout() {
	stdoutHandler := handler.NewStdoutHandler(nil) // 将输出定向到 stdout
	easylog.AddHandler(stdoutHandler)              // 添加到handlers
	easylog.SetLevel(easylog.DEBUG)                // 设置日志级别
	easylog.Debug().Msg("hello world")
	easylog.Fatal().Msg("error")
	// Output:
	// hello world
	// error
	easylog.RemoveHandler(stdoutHandler)
}

func format(record *easylog.Record) string {
	type output struct {
		Fields map[string]interface{} `json:"fields"`
		Msg    string                 `json:"msg"`
	}
	o := output{
		Fields: record.FieldMap,
		Msg:    record.Message,
	}
	b, err := json.Marshal(o)
	if err == nil {
		return string(b[:])
	}
	return ""
}

func ExampleEasylog_SimpleFormat() {
	stdoutHandler := handler.NewStdoutHandler(format)
	easylog.AddHandler(stdoutHandler)
	easylog.SetLevel(easylog.DEBUG)
	easylog.Debug().Fields(map[string]interface{}{"name": "dog"}).Msg("hello")
	// Output:
	// {"fields":{"name":"dog"},"msg":"hello"}
	easylog.RemoveHandler(stdoutHandler)
}
