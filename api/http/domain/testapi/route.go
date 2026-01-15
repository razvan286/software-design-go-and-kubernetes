package testapi

import (
	"github.com/razvan286/software-design-go-and-kubernetes/api/http/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/authclient"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	athAdminOnly := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI()

	app.HandleFunc("GET /testerror", api.testError)
	app.HandleFunc("GET /testpanic", api.testPanic)
	app.HandleFunc("GET /testauth", api.testAuth, authen, athAdminOnly)
}
