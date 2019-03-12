package ulog

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ulogger struct {
	log   *zap.SugaredLogger
	level string
}

func unmarshalLevel(lvl string) zapcore.Level {
	// default level is info
	var level zapcore.Level
	if lvl != "" {
		if err := level.UnmarshalText([]byte(lvl)); err != nil {
			fmt.Fprintln(os.Stderr,
				fmt.Sprintf("Invalid log level %s. Fall back to use the default level.", lvl))
			return unmarshalLevel(defaultLevel)
		}
	}
	return level
}

func newProductionLogger(level string) *zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = EncodeLevel
	config.EncodeTime = EncodeTime
	config.EncodeCaller = EncodeCaller

	logLevel := unmarshalLevel(level)
	normalEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel
	})

	consoleOut := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewCore(consoleEncoder, consoleOut, normalEnabler)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}

// Debug logs a message at level Debug on the Logger.
func (l *ulogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *ulogger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *ulogger) Debugln(args ...interface{}) {
	l.log.Debug(fmt.Sprintln(args...))
}

func (l *ulogger) Debugw(msg string, args ...interface{}) {
	l.log.Debugw(msg, args...)
}

func (l *ulogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *ulogger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *ulogger) Infoln(args ...interface{}) {
	l.log.Info(fmt.Sprintln(args...))
}

func (l *ulogger) Infow(msg string, args ...interface{}) {
	l.log.Infow(msg, args...)
}

func (l *ulogger) Warning(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *ulogger) Warningf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *ulogger) Warningln(args ...interface{}) {
	l.log.Warn(fmt.Sprintln(args...))
}

func (l *ulogger) Warningw(msg string, args ...interface{}) {
	l.log.Warnw(msg, args...)
}

func (l *ulogger) Error(args ...interface{}) {
	l.log.Error(args...)
	//ReportError(args...)
}

func (l *ulogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
	//ReportErrorf(format, args...)
}

func (l *ulogger) Errorln(args ...interface{}) {
	l.log.Error(fmt.Sprintln(args...))
	//ReportErrorln(args...)
}

func (l *ulogger) Errorw(msg string, args ...interface{}) {
	l.log.Errorw(msg, args...)
	//ReportErrorw(msg, args...)
}

func (l *ulogger) Fatal(args ...interface{}) {
	//ReportFatal(args...)
	l.log.Fatal(args...)
}

func (l *ulogger) Fatalf(format string, args ...interface{}) {
	//ReportFatalf(format, args...)
	l.log.Fatalf(format, args...)
}

func (l *ulogger) Fatalln(args ...interface{}) {
	//ReportFatalln(args...)
	l.log.Fatal(fmt.Sprintln(args...))
}

func (l *ulogger) Fatalw(msg string, args ...interface{}) {
	//ReportFatalw(msg, args...)
	l.log.Fatalw(msg, args...)
}

func (l *ulogger) Panic(args ...interface{}) {
	//ReportPanic(args...)
	l.log.Panic(args...)
}

func (l *ulogger) Panicf(format string, args ...interface{}) {
	//ReportPanicf(format, args...)
	l.log.Panicf(format, args...)
}

func (l *ulogger) Panicln(args ...interface{}) {
	//ReportPanicln(args...)
	l.log.Panic(fmt.Sprintln(args...))
}

func (l *ulogger) Panicw(msg string, args ...interface{}) {
	//ReportPanicw(msg, args...)
	l.log.Panicw(msg, args...)
}

func (l *ulogger) With(key string, value interface{}) Logger {
	return &ulogger{l.log.With(zap.Any(key, value)), l.level}
}

func (l *ulogger) WithFields(fields map[string]interface{}) Logger {
	i := 0
	var clog Logger
	for k, v := range fields {
		if i == 0 {
			clog = l.With(k, v)
		} else {
			clog = clog.With(k, v)
		}
		i++
	}
	return clog
}
