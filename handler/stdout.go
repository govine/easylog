package handler

import (
	"fmt"
	"os"

	"git.qutoutiao.net/govine/easylog"
)

type StdoutHandler struct {
	format easylog.Formatter
}

func (s *StdoutHandler) Handle(record *easylog.Record) {
	var str string
	if s.format != nil {
		str = s.format(record)
	} else {
		str = fmt.Sprintf(record.Msg, record.Args)
	}

	os.Stdout.Write([]byte(str + "\n"))
}

func (f *StdoutHandler) Flush() {
}

func NewStdoutHandler(format easylog.Formatter) easylog.IEasyLogHandler {
	return easylog.NewEasyLogHandler(&StdoutHandler{
		format: format,
	})
}
