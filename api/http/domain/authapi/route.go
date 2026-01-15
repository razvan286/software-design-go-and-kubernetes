package authapi

import (
	"github.com/razvan286/software-design-go-and-kubernetes/api/http/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Auth *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	bearer := mid.Bearer(cfg.Auth)
	basic := mid.Basic(cfg.Auth)

	api := newAPI(cfg.Auth)

	app.HandleFunc("GET /auth/token/{kid}", api.token, basic)
	app.HandleFunc("GET /auth/authenticate", api.authenticate, bearer)
	app.HandleFunc("POST /auth/authorize", api.authorize)
}
