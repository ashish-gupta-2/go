package app

import (
	"io"

	logpkg "ashish.com/m/internal/log"
	"ashish.com/m/internal/version"
	"github.com/sirupsen/logrus"
)

// LogLevel represents log level enum.
type LogLevel uint16

const (
	// LogLevelTrace represents 'trace' log level.
	LogLevelTrace LogLevel = iota

	// LogLevelDebug represents 'debug' log level.
	LogLevelDebug

	// LogLevelInfo represents 'info' log level.
	LogLevelInfo

	// LogLevelWarn represents 'warn' log level.
	LogLevelWarn

	// LogLevelError represents 'error' log level.
	LogLevelError

	// LogLevelFatal represents 'fatal' log level.
	LogLevelFatal
)

// InitLog initializes the global logger.
func InitLog(level LogLevel, output io.Writer, jsonFormat bool) {
	logrus.SetLevel(toLogrusLevel(level))
	logrus.SetOutput(output)
	if jsonFormat {
		logrus.SetReportCaller(true)
		logrus.AddHook(&logpkg.ContextHook{})
		resource := map[string]any{
			"app.name":    version.GetName(),
			"app.version": version.GetVersion(),
		}
		logrus.SetFormatter(logpkg.NewJSONFormatter(resource, "app"))
	}
}

func toLogrusLevel(lvl LogLevel) logrus.Level {
	switch lvl {
	case LogLevelTrace:
		return logrus.TraceLevel
	case LogLevelDebug:
		return logrus.DebugLevel
	case LogLevelInfo:
		return logrus.InfoLevel
	case LogLevelWarn:
		return logrus.WarnLevel
	case LogLevelError:
		return logrus.ErrorLevel
	case LogLevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.TraceLevel
	}
}
