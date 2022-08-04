package formatter_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	formatter "github.com/helmwave/logrus-emoji-formatter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type FormatterTestSuite struct {
	suite.Suite
}

func (s *FormatterTestSuite) createLogger() (*logrus.Logger, *bytes.Buffer) {
	logger := logrus.New()
	var buf bytes.Buffer
	logger.SetFormatter(&formatter.Config{Color: false})
	logger.SetOutput(&buf)
	logger.SetLevel(logrus.TraceLevel)

	return logger, &buf
}

func (s *FormatterTestSuite) TestDefaultFormatting() {
	logger, buf := s.createLogger()

	msg := "testblabla"

	s.Run("trace", func() {
		expected := fmt.Sprintf("[🤮 aka TRACE]: %s\n", msg)
		logger.Trace(msg)
		defer buf.Reset()

		s.Require().Equal(expected, buf.String())
	})
	s.Run("debug", func() {
		expected := fmt.Sprintf("[🤷 aka DEBUG]: %s\n", msg)
		logger.Debug(msg)
		defer buf.Reset()

		s.Require().Equal(expected, buf.String())
	})
	s.Run("info", func() {
		expected := fmt.Sprintf("[🙃 aka INFO]: %s\n", msg)
		logger.Info(msg)
		defer buf.Reset()

		s.Require().Equal(expected, buf.String())
	})
	s.Run("warn", func() {
		expected := fmt.Sprintf("[🙈 aka WARNING]: %s\n", msg)
		logger.Warn(msg)
		defer buf.Reset()

		s.Require().Equal(expected, buf.String())
	})
	s.Run("error", func() {
		expected := fmt.Sprintf("[💩 aka ERROR]: %s\n", msg)
		logger.Error(msg)
		defer buf.Reset()

		s.Require().Equal(expected, buf.String())
	})
}

func (s *FormatterTestSuite) TestCustomFormat() {
	logger, buf := s.createLogger()

	c, _ := logger.Formatter.(*formatter.Config)
	c.LogFormat = " "

	msg := "testblabla"
	logger.Info(msg)

	s.Require().Equal(" \n", buf.String())
}

func (s *FormatterTestSuite) TestTimeFormat() {
	logger, buf := s.createLogger()

	c, _ := logger.Formatter.(*formatter.Config)
	c.LogFormat = "%time%"

	msg := "testblabla"
	expected := time.Now()
	logger.Info(msg)

	t, err := time.Parse(time.RFC3339, strings.TrimSpace(buf.String()))
	s.Require().NoError(err)
	s.Require().WithinDuration(expected, t, time.Second)
}

func TestFormatterTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(FormatterTestSuite))
}
