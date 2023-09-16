package http

import (
	"encoding/json"
	"net/http"
)

// InternalResponseWriter represents a wrapper for http response writer.
type InternalResponseWriter struct {
	code     int
	message  string
	original http.ResponseWriter
}

// NewInternalResponseWriter creates a new instance of internal response writer.
func NewInternalResponseWriter(writer http.ResponseWriter) *InternalResponseWriter {
	return &InternalResponseWriter{
		code:     0,
		original: writer,
	}
}

// Write captures the incoming message on local field and then forwards the data
// to original http response stream.
func (w *InternalResponseWriter) Write(data []byte) (int, error) {
	var s map[string]any
	if err := json.Unmarshal(data, &s); err == nil {
		if m, ok := s["message"]; ok {
			if msg, ok := m.(string); ok {
				w.message = msg
			}
		}
	}
	return w.original.Write(data)
}

// Header returns the original http header.
func (w *InternalResponseWriter) Header() http.Header {
	return w.original.Header()
}

// WriteHeader captures the incoming status code on local field and then forwards the
// status code to original http response.
func (w *InternalResponseWriter) WriteHeader(code int) {
	w.code = code
	w.original.WriteHeader(code)
}

// GetCode returns the http status code.
func (w *InternalResponseWriter) GetCode() int {
	return w.code
}

// GetMessage returns the previously captured message.
func (w *InternalResponseWriter) GetMessage() string {
	return w.message
}
