package implementation

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// WithTokenContextFromHeader returns a new context with the token from the header
func WithTokenContextFromHeader(ctx context.Context, header http.Header) context.Context {
	token := header.Get("Authorization")
	bearerToken := strings.TrimPrefix(token, "Bearer ")
	return WithTokenContext(ctx, bearerToken)
}

// WithTokenContext returns a new context with the token
func WithTokenContext(ctx context.Context, bearerToken string) context.Context {
	if bearerToken == "" {
		return ctx
	}
	return context.WithValue(ctx, tokenCtxKey, bearerToken)
}

// GetTokenFromContext returns the token from the context, otherwise an empty string
func GetTokenFromContext(ctx context.Context) string {
	if token, ok := ctx.Value(tokenCtxKey).(string); ok {
		return token
	}
	return ""
}

type contextKey struct {
	name string
}

var tokenCtxKey = &contextKey{"token"}

// ErrAccessDenied is returned when the user is not authenticated
var ErrAccessDenied = errors.New("access denied")
