package checkapi

import (
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/app/api/authclient"
	"github.com/razvan286/software-design-go-and-kubernetes/business/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, log *logger.Logger, authClient *authclient.Client) {
	authen := mid.AuthenticateService(log, authClient)
	athAdminOnly := mid.AuthorizeService(log, authClient, auth.RuleAdminOnly)

	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
	app.HandleFunc("GET /testauth", liveness, authen, athAdminOnly)
}
