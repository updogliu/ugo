package ulog

const defaultLevel = "debug"

var defaultLogger Logger

func init() {
	defaultLogger = NewLogger()
}
