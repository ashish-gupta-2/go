package app

import (
	"context"

	ctxpkg "ashish.com/m/internal/context"
	"ashish.com/m/internal/utils"
	"github.com/google/uuid"
)

// NewRootContext returns a new log context created using root namespace.
func NewRootContext() context.Context {
	ctxValue := ctxpkg.Value{
		Namespace: ctxpkg.NamespaceRoot,
		TransID:   uuid.NewString(),
		Endpoint:  utils.Empty,
	}
	return context.WithValue(context.Background(), ctxpkg.Key{}, ctxValue)
}
