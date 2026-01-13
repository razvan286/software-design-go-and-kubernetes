package checkapi

import (
	"github.com/razvan286/software-design-go-and-kubernetes/business/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, a *auth.Auth) {
	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
}