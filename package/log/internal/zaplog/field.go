package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field is the type for koollog field
type Field = zapcore.Field

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
//
// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
// them. To minimize surprises, []byte values are treated as binary blobs, byte
// values are treated as uint8, and runes are always treated as integers.
func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}
