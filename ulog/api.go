package ulog

// Logger interface is compatible with grpc-go grpclog.Loggerv2 and etcd raft.Logger interfaces
type Logger interface {
	// Debug logs are typically voluminous, and are usually disabled in production.
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Debugw(msg string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Infow(msg string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Warningw(msg string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Errorw(msg string, args ...interface{})

	// Panic logs a message, then panics.
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Panicw(msg string, args ...interface{})

	// Fatal logs a message, then calls os.Exit(1).
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Fatalw(msg string, args ...interface{})

	With(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

// NewLogger creates a logger with "info" level.
func NewLogger() Logger {
	return NewLoggerWithLevel(defaultLevel)
}

// NewLoggerWithLevel creates a ulog logger instance with custom log level. An "info" level logger
// will be created if `level` is "".
//
// Log levels available: "debug", "info", "warn", "error", "panic", "fatal".
func NewLoggerWithLevel(level string) Logger {
	return &ulogger{
		log:   newProductionLogger(level),
		level: level,
	}
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

func Debugw(msg string, args ...interface{}) {
	defaultLogger.Debugw(msg, args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

func Infow(msg string, args ...interface{}) {
	defaultLogger.Infow(msg, args...)
}

func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	defaultLogger.Warningf(format, args...)
}

func Warningln(args ...interface{}) {
	defaultLogger.Warningln(args...)
}

func Warningw(msg string, args ...interface{}) {
	defaultLogger.Warningw(msg, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

func Errorw(msg string, args ...interface{}) {
	defaultLogger.Errorw(msg, args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

func Fatalw(msg string, args ...interface{}) {
	defaultLogger.Fatalw(msg, args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}

func Panicw(msg string, args ...interface{}) {
	defaultLogger.Panicw(msg, args...)
}
