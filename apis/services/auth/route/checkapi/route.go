package checkapi

import (
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Routes adds specific routes for this group.
func Routes(build string, app *web.App, log *logger.Logger, db *sqlx.DB) {
	api := newAPI(build, log, db)
	app.HandleFuncNoMiddleware("GET /liveness", api.liveness)
	app.HandleFuncNoMiddleware("GET /readiness", api.readiness)
}
