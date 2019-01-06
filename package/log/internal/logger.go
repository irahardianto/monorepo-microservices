package internal

import "github.com/irahardianto/monorepo-mocroservices/package/log/internal/zaplog"

// Debug add log entry with or without fields to debug level
func Debug(msg string, data interface{}) {
	zaplog.Debug(msg, zaplog.Any("data", data))
}

// Info add log entry with or without fields to info level
func Info(msg string, data interface{}) {
	zaplog.Info(msg, zaplog.Any("data", data))
}

// Warn add log entry with or without fields to warn level
func Warn(msg string, data interface{}) {
	zaplog.Warn(msg, zaplog.Any("data", data))
}

// Error add log entry with or without fields to error level
func Error(msg string, data interface{}) {
	zaplog.Error(msg, zaplog.Any("data", data))
}

// Fatal add log entry with or without fields to fatal level
func Fatal(msg string, data interface{}) {
	zaplog.Fatal(msg, zaplog.Any("data", data))
}

// Panic add log entry with or without fields to panic level
func Panic(msg string, data interface{}) {
	zaplog.Panic(msg, zaplog.Any("data", data))
}
