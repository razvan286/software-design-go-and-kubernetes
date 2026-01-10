// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/sales/route/sys/checkapi"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/logger"
	"github.com/razvan286/software-design-go-and-kubernetes/foundation/web"
	"github.com/razvan286/software-design-go-and-kubernetes/apis/services/api/mid"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(mux)

	return mux
}
