package log

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type StdKratosLog struct {
}

func (k *StdKratosLog) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if (len(keyvals) & 1) == 1 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}
	logStr := "kratos log: "

	for i := 0; i < len(keyvals); i += 2 {
		logStr += fmt.Sprintf(" %s=%v", keyvals[i], keyvals[i+1])
	}

	l := DefaultLogger

	switch level {
	case log.LevelDebug:
		{
			l.Debug(logStr)
		}
	case log.LevelInfo:
		{
			l.Info(logStr)
		}
	case log.LevelWarn:
		{
			l.Warn(logStr)
		}
	case log.LevelError:
		{
			l.Error(logStr)
		}
	case log.LevelFatal:
		{
			l.Fatal(logStr)
		}
	}

	return nil
}
