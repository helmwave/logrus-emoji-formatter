package formatter

import (
	"time"
)

const (
	defaultLogFormat       = "[%emoji% aka %lvl%]: %msg%"
	defaultTimestampFormat = time.RFC3339
)

// Formatter implements logrus. Formatter interface.
type Config struct {
	// Timestamp format
	TimestampFormat string

	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string

	// Color enabled
	Color bool

	// Emoji Formatter
	Emoji bool
}
