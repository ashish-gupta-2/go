package http

import (
	"context"
	"net/http"
	"strings"

	ctxpkg "ashish.com/m/internal/context"
	httppkg "ashish.com/m/internal/http"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
)

// HandlerFunc represents the handler function for incoming http api request.
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string)

// InterceptHTTPServer intercepts http api from server side and writes appropriate log messages.
func InterceptHTTPServer(handler HandlerFunc) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		// build new context with context values
		newContext := context.WithValue(context.Background(), ctxpkg.Key{}, ctxpkg.Value{
			Namespace: ctxpkg.NamespaceRoot,
			TransID:   uuid.NewString(),
			Endpoint:  r.URL.String(),
		})

		// log http headers from request
		fields := log.Fields{}
		for key, value := range r.Header {
			fields[strings.ToLower(key)] = strings.Join(value, ", ")
		}
		log.WithContext(newContext).WithFields(fields).Info("started http api invocation")

		// call the original handler
		internalWriter := httppkg.NewInternalResponseWriter(w)
		handler(newContext, internalWriter, r, params)

		// log handler invocation status
		code := internalWriter.GetCode()
		message := internalWriter.GetMessage()
		var statusFields log.Fields
		if len(message) > 0 {
			statusFields = log.Fields{"status.code": code, "status.message": message}
		} else {
			statusFields = log.Fields{"status.code": code}
		}
		log.WithContext(newContext).WithFields(statusFields).Info("finished http api invocation")
	}
}
