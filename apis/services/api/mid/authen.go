package mid

import (
	"context"
	"net/http"

	"github.com/razvan286/software-design-go-and-kubernetes/app/api/authclient"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/business/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// AuthenticateService validates authentication via the auth service.
func AuthenticateService(log *logger.Logger, client *authclient.Client) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.AuthenticateService(ctx, log, client, r.Header.Get("authorization"), hdl)
		}

		return h
	}

	return m
}

// AuthenticateLocal processes the authentication requirements locally.
func AuthenticateLocal(auth *auth.Auth) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.AuthenticateLocal(ctx, auth, r.Header.Get("authorization"), hdl)
		}

		return h
	}

	return m
}
