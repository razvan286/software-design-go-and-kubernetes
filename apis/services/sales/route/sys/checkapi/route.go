package checkapi

import (
	"github.com/jmoiron/sqlx"
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/authclient"
	"github.com/razvan286/software-design-go-and-kubernetes/business/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(build string, app *web.App, log *logger.Logger, db *sqlx.DB, authClient *authclient.Client) {
	authen := mid.AuthenticateService(log, authClient)
	athAdminOnly := mid.AuthorizeService(log, authClient, auth.RuleAdminOnly)

	api := newAPI(build, log, db)
	app.HandleFuncNoMiddleware("GET /liveness", api.liveness)
	app.HandleFuncNoMiddleware("GET /readiness", api.readiness)
	app.HandleFunc("GET /testerror", api.testError)
	app.HandleFunc("GET /testpanic", api.testPanic)
	app.HandleFunc("GET /testauth", api.liveness, authen, athAdminOnly)
}
