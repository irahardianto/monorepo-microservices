package zaplog

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	encoder = zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	)

	options = []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)}

	output   zapcore.WriteSyncer
	loglevel zapcore.Level
	log      *zap.Logger
	logMu    sync.Mutex
)

func init() {
	defaultLog()
}

func createLogger(output zapcore.WriteSyncer, level zapcore.Level) {
	core := zapcore.NewCore(encoder, output, level)
	log = zap.New(core).WithOptions(options...)
}

func defaultLog() {
	logMu.Lock()
	defer logMu.Unlock()

	output = os.Stderr
	loglevel = zap.InfoLevel
	createLogger(output, loglevel)
}

// DebugMode sets the log level to debug
func DebugMode() {
	logMu.Lock()
	defer logMu.Unlock()

	loglevel = zap.DebugLevel
	createLogger(output, loglevel)
}

// Reset will reset the log to the original setup
func Reset() {
	defaultLog()
}

// Debug add log entry with debug level
func Debug(msg string, fields ...Field) {
	log.Debug(msg, fields...)
}

// Info add log entry with info level
func Info(msg string, fields ...Field) {
	log.Info(msg, fields...)
}

// Warn add log entry with warn level
func Warn(msg string, fields ...Field) {
	log.Warn(msg, fields...)
}

// Error add log entry with error level
func Error(msg string, fields ...Field) {
	log.Error(msg, fields...)
}

// Fatal add log entry with fatal level
func Fatal(msg string, fields ...Field) {
	log.Fatal(msg, fields...)
}

// Panic add log entry with panic level
func Panic(msg string, fields ...Field) {
	log.Panic(msg, fields...)
}
