package ulog

const defaultLevel = "debug"

var defaultLogger Logger

func init() {
	defaultLogger = &ulogger{
		log: newProductionLogger(defaultLevel, 2),
	}
}
