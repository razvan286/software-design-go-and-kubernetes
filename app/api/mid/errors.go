package mid

import (
	"context"

	"github.com/razvan286/software-design-go-and-kubernetes/app/api/errs"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
)

// Errors handles errors coming out of the call chain. It detects normal application
// erors which are used to respond to the client in a uniform way.
// Unexpected error (status >= 500) are logged.
func Errors(ctx context.Context, log *logger.Logger, handler Handler) error {
	err := handler(ctx)
	if err == nil {
		return nil
	}

	log.Error(ctx, "message", "ERROR", err.Error())

	if errs.IsError(err) {
		return errs.GetError(err)
	}

	return errs.Newf(errs.Unknown, "%s", errs.Unknown.String())
}
