package handler

import (
	"os"

	"github.com/covine/easylog"
)

type StdoutHandler struct {
	format easylog.Formatter
}

func (s *StdoutHandler) Handle(record *easylog.Record) {
	var str string
	if s.format != nil {
		str = s.format(record)
	} else {
		str = record.Message
	}

	_, _ = os.Stdout.Write([]byte(str + "\n"))
}

func (s *StdoutHandler) Flush() {
}

func (s *StdoutHandler) Close() {
}

func NewStdoutHandler(format easylog.Formatter) easylog.IEasyLogHandler {
	return easylog.NewEasyLogHandler(&StdoutHandler{
		format: format,
	})
}
