package utils

import (
	errpkg "ashish.com/m/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MapToGRPCStatus translates error into grpc status. It contains status code and message.
func MapToGRPCStatus(err error) *status.Status {
	var s *status.Status
	switch e := err.(type) {
	case *errpkg.ResourceNotFoundError:
		s = status.New(codes.NotFound, e.Error())
	case *errpkg.RecordNotFoundError:
		s = status.New(codes.NotFound, e.Error())
	case *errpkg.EmptyError:
		s = status.New(codes.InvalidArgument, e.Error())
	case *errpkg.FormatError:
		s = status.New(codes.InvalidArgument, e.Error())
	default:
		msg := "Internal error occurred."
		s = status.New(codes.Internal, msg)
	}
	return s
}
