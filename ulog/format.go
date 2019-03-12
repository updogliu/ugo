package ulog

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	zc "go.uber.org/zap/zapcore"
)

func EncodeTime(t time.Time, encoder zc.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("060102-15:04:05.000 MST"))
}

func EncodeLevel(l zc.Level, encoder zc.PrimitiveArrayEncoder) {
	switch l {
	case zc.DebugLevel:
		encoder.AppendString("[D]")
	case zc.InfoLevel:
		encoder.AppendString("[I]")
	case zc.WarnLevel:
		encoder.AppendString("[W]")
	case zc.ErrorLevel:
		encoder.AppendString("[E]")
	case zc.DPanicLevel:
		encoder.AppendString("[A]")
	case zc.PanicLevel:
		encoder.AppendString("[P]")
	case zc.FatalLevel:
		encoder.AppendString("[F]")
	default:
		encoder.AppendString(fmt.Sprintf("[Level(%d)]", l))
	}
}

func EncodeCaller(caller zc.EntryCaller, encoder zc.PrimitiveArrayEncoder) {
	if !caller.Defined {
		encoder.AppendString("undefined")
		return
	}
	idx := strings.LastIndexByte(caller.File, '/')
	encoder.AppendString(fmt.Sprint(caller.File[idx+1:], ":", caller.Line))
}

// "filename:lineNum:funcName"
func EncodeCallerWithFuncName(caller zc.EntryCaller, enc zc.PrimitiveArrayEncoder) {
	s := fmt.Sprint(runtime.FuncForPC(caller.PC).Name(), ":", caller.TrimmedPath())
	enc.AppendString(s)
}
