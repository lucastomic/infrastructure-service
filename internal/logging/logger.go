package logging

import (
	"context"
	"net/http"
)

type Logger interface {
	Request(context.Context, *http.Request, int)
	Info(ctx context.Context, format string, a ...any)
	Error(ctx context.Context, format string, a ...any)
	Warn(args ...any)
	Fatal(args ...any)
}
