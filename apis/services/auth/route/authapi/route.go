package authapi

import (
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/business/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, a *auth.Auth) {
	authen := mid.Authorization(a)

	api := newAPI(a)
	app.HandleFunc("GET /auth/token/{kid}", api.token, authen)
	app.HandleFunc("GET /auth/authenticate", api.authenticate, authen)
	app.HandleFunc("POST /auth/authorize", api.authorize)
}
