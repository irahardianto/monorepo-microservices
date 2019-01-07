package log

import log "github.com/irahardianto/monorepo-microservices/package/log/internal"

// Debug add log entry with debug level
func Debug(msg string, data interface{}) {
	log.Debug(msg, data)
}

// Info add log entry with info level
func Info(msg string, data interface{}) {
	log.Info(msg, data)
}

// Warn add log entry with warn level
func Warn(msg string, data interface{}) {
	log.Warn(msg, data)
}

// Error add log entry with error level
func Error(msg string, data interface{}) {
	log.Error(msg, data)
}

// Fatal add log entry with fatal level
func Fatal(msg string, data interface{}) {
	log.Fatal(msg, data)
}

// Panic add log entry with panic level
func Panic(msg string, data interface{}) {
	log.Panic(msg, data)
}
