package mid

import (
	"context"
	"net/http"

	"github.com/razvan286/software-design-go-and-kubernetes/app/api/authclient"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Authorize executes the authorize middleware functionality.
func Authorize(log *logger.Logger, client *authclient.Client, rule string) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Authorize(ctx, log, client, rule, hdl)
		}

		return h
	}

	return m
}
