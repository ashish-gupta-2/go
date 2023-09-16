package http

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// Constant declaration for http headers.
const (
	HeaderContentType        = "Content-Type"
	HeaderContentDisposition = "Content-Disposition"
	ContentTypeJSON          = "application/json"
	ContentTypeYAML          = "application/yaml"
	ContentTypeOctetStream   = "application/octet-stream"
)

// GetStatusCode returns http status code which is equivalent to grpc code.
func GetStatusCode(code codes.Code) int {
	status := http.StatusOK
	switch code {
	case codes.OK:
		status = http.StatusOK
	case codes.Canceled:
		status = 499
	case codes.Unknown:
		status = http.StatusInternalServerError
	case codes.InvalidArgument:
		status = http.StatusBadRequest
	case codes.DeadlineExceeded:
		status = http.StatusGatewayTimeout
	case codes.NotFound:
		status = http.StatusNotFound
	case codes.AlreadyExists:
		status = http.StatusConflict
	case codes.PermissionDenied:
		status = http.StatusForbidden
	case codes.Unauthenticated:
		status = http.StatusUnauthorized
	case codes.ResourceExhausted:
		status = http.StatusTooManyRequests
	case codes.FailedPrecondition:
		status = http.StatusPreconditionFailed
	case codes.Aborted:
		status = http.StatusConflict
	case codes.OutOfRange:
		status = http.StatusBadRequest
	case codes.Unimplemented:
		status = http.StatusNotImplemented
	case codes.Internal:
		status = http.StatusInternalServerError
	case codes.Unavailable:
		status = http.StatusServiceUnavailable
	case codes.DataLoss:
		status = http.StatusInternalServerError
	}

	return status
}
