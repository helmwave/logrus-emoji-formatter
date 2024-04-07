package formatter

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

//nolint:gochecknoglobals // used as constants
var (
	emojisLevel   = [7]string{"💀", "🤬", "💩", "🙈", "🙃", "🤷", "🤮"}
	colors        = [7]*color.Color{color.New(color.BgMagenta), color.New(color.FgRed), color.New(color.FgRed), color.New(color.FgYellow), color.New(color.FgCyan), color.New(color.FgWhite), color.New(color.FgMagenta)}
	logFieldColor = color.New(color.FgMagenta).SprintFunc()
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

	color.NoColor = !f.Color
	color := colors[i].SprintFunc()

	l := color(level)
	m := color(entry.Message)

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
			output += fmt.Sprintf("\n\t%s: \n\t  - %v", logFieldColor(k), v)
		default:
			output += fmt.Sprintf("\n\t%s: %v", logFieldColor(k), val)
		}
	}
	output += "\n"

	return []byte(output), nil
}
