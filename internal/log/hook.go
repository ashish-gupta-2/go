package log

import (
	ctxpkg "ashish.com/m/internal/context"
	"github.com/sirupsen/logrus"
)

// ContextHook represents the hook implementation for logrus.
type ContextHook struct{}

// Levels returns all log levels which needs interception.
func (h *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire intercepts every log entry and inserts additional fields.
func (h *ContextHook) Fire(entry *logrus.Entry) error {
	if entry.Context != nil {
		value := entry.Context.Value(ctxpkg.Key{})
		if ctxValue, ok := value.(ctxpkg.Value); ok {
			entry.Data["context.namespace"] = ctxValue.Namespace
			if len(ctxValue.TransID) > 0 {
				entry.Data["context.transid"] = ctxValue.TransID
			}
			if len(ctxValue.Endpoint) > 0 {
				entry.Data["context.endpoint"] = ctxValue.Endpoint
			}
			return nil
		}
	}
	return nil
}
