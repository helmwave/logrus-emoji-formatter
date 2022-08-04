package formatter

import (
	"time"
)

const (
	defaultLogFormat       = "[%emoji% aka %lvl%]: %msg%"
	defaultTimestampFormat = time.RFC3339
)

// Config is a logrus formatter.
type Config struct {
	// Timestamp format.
	TimestampFormat string
	// LogFormat is a format string for log.
	// Available standard keys: time, msg, lvl.
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside `%%` i.e `%time% %msg%`.
	LogFormat string
	// Color enables colors output.
	Color bool
}
