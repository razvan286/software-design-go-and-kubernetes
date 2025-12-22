package mid

import (
	"context"
	"net/http"

	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
			// Logging

			err := handler(ctx, w, r)

			// Logging
			log.Info(ctx, "request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
			
			return err
		}

		return h
	}
	return m

}
