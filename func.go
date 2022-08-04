package formatter

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

//nolint:gochecknoglobals // used as constants
var (
	emojisLevel = [7]string{"💀", "🤬", "💩", "🙈", "🙃", "🤷", "🤮"}
	colors      = [7]string{"[44;1m", "[31;1m", "[31;1m", "[33m", "[36m", "[37;1m", "[35;1m"}
)

const (
	logFieldColor = "[35;1m"
	start         = "\033"
	end           = start + "[0m"
)

// Format building log message.
func (f *Config) Format(entry *logrus.Entry) ([]byte, error) {
	format := f.LogFormat
	if format == "" {
		format = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	level := strings.ToUpper(entry.Level.String())

	i, _ := logrus.ParseLevel(level)
	emoji := emojisLevel[i]
	l := level
	m := entry.Message
	fieldPattern := "%s"
	if f.Color {
		color := colors[i]
		l = start + color + level + end
		m = start + color + entry.Message + end
		fieldPattern = start + logFieldColor + fieldPattern + end
	}

	replacer := strings.NewReplacer(
		"%time%", entry.Time.Format(timestampFormat),
		"%msg%", m,
		"%lvl%", l,
		"%emoji%", emoji,
	)

	output := replacer.Replace(format)

	for k, val := range entry.Data {
		switch val := val.(type) {
		case []string:
			v := strings.Join(val, "\n\t  - ")
			output += fmt.Sprintf("\n\t"+fieldPattern+": \n\t  - %v", k, v)
		default:
			output += fmt.Sprintf("\n\t"+fieldPattern+": %v", k, val)
		}
	}
	output += "\n"

	return []byte(output), nil
}
