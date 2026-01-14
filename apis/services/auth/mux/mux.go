// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/api/mid"
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/auth/route/authapi"
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/auth/route/checkapi"
	"github.com/razvan286/software-design-go-and-kubernetes/business/api/auth"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(build string, log *logger.Logger, db *sqlx.DB, auth *auth.Auth, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(build, app, log, db)
	authapi.Routes(app, auth)

	return app
}
