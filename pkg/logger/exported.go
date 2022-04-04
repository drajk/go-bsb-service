package logger

import "context"

var (
	std = New()
)

// GetLogger returns the default logger
func GetLogger() Logger {
	return std
}

// With - Add key and value to logger context
func With(key string, value interface{}) Logger {
	return std.With(context.Background(), key, value)
}

// WithError - logs error
func WithError(err error) Logger {
	std.Error(err)
	return std
}
