package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	utils "ashish.com/m/internal/utils"
	"github.com/sirupsen/logrus"
)

// SeverityLevel represents severity level enum.
type SeverityLevel uint

const (
	// SeverityLevelTrace represents 'trace' severity level.
	SeverityLevelTrace SeverityLevel = 1

	// SeverityLevelDebug represents 'debug' severity level.
	SeverityLevelDebug SeverityLevel = 5

	// SeverityLevelInfo represents 'info' severity level.
	SeverityLevelInfo SeverityLevel = 9

	// SeverityLevelWarn represents 'warn' severity level.
	SeverityLevelWarn SeverityLevel = 13

	// SeverityLevelError represents 'error' severity level.
	SeverityLevelError SeverityLevel = 17

	// SeverityLevelFatal represents 'fatal' severity level.
	SeverityLevelFatal SeverityLevel = 21
)

// String converts the severity level to a string representation.
func (s SeverityLevel) String() string {
	switch s {
	case SeverityLevelTrace:
		return "trace"
	case SeverityLevelDebug:
		return "debug"
	case SeverityLevelInfo:
		return "info"
	case SeverityLevelWarn:
		return "warning"
	case SeverityLevelError:
		return "error"
	case SeverityLevelFatal:
		return "fatal"
	default:
		return "unknown"
	}
}

// Log represents a standard data model for log record. This model is
// recommended by opentelemetry standard.
type Log struct {
	Timestamp      uint64         `json:"timestamp,omitempty"`
	SeverityText   string         `json:"severityText,omitempty"`
	SeverityNumber uint           `json:"severityNumber,omitempty"`
	Resource       map[string]any `json:"resource,omitempty"`
	Attributes     map[string]any `json:"attributes,omitempty"`
	Body           any            `json:"body,omitempty"`
}

// JSONFormatter represents a json formatter for log entry.
type JSONFormatter struct {
	resource        map[string]any
	fileStartString string
}

// NewJSONFormatter creates a new instance of json formatter.
func NewJSONFormatter(resource map[string]any, fileStartString string) *JSONFormatter {
	return &JSONFormatter{
		resource:        resource,
		fileStartString: fileStartString,
	}
}

// Format renders the single log entry in json format.
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// collect all fields
	fields := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			fields[k] = v.Error()
		default:
			fields[k] = v
		}
	}

	// add caller info
	if entry.HasCaller() && entry.Context != nil {
		substr := func(s, sub string) string {
			if index := strings.Index(s, sub); index != -1 {
				return s[index:]
			}
			return s
		}
		fields["context.file"] = fmt.Sprintf("%s:%d", substr(entry.Caller.File, f.fileStartString), entry.Caller.Line)
		fields["context.function"] = substr(entry.Caller.Function, f.fileStartString)
	}

	// populate log entry
	log := Log{
		Timestamp:      uint64(entry.Time.UnixNano()),
		SeverityNumber: uint(toSeverityLevel(entry.Level)),
		SeverityText:   toSeverityLevel(entry.Level).String(),
		Resource:       f.resource,
		Attributes:     fields,
		Body:           entry.Message,
	}

	// marshal log entry into json
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(true)
	encoder.SetIndent(utils.Empty, strings.Repeat(utils.Space, 2))
	if err := encoder.Encode(log); err != nil {
		return nil, fmt.Errorf("error while marshalling log entry to json: %v", err)
	}

	// return json formatted log entry
	return b.Bytes(), nil
}

func toSeverityLevel(lvl logrus.Level) SeverityLevel {
	switch lvl {
	case logrus.TraceLevel:
		return SeverityLevelTrace
	case logrus.DebugLevel:
		return SeverityLevelDebug
	case logrus.InfoLevel:
		return SeverityLevelInfo
	case logrus.WarnLevel:
		return SeverityLevelWarn
	case logrus.ErrorLevel, logrus.PanicLevel:
		return SeverityLevelError
	case logrus.FatalLevel:
		return SeverityLevelFatal
	default:
		return 0
	}
}
