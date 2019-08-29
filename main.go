package main

import (
	"net/http"
	"os"

	ctxrouter "github.com/hobord/infra/ctxrouter"
	httparams "github.com/hobord/infra/httparams"
	log "github.com/hobord/infra/log"
	redirect "github.com/hobord/infra/redirect"
	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
)

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8100"
	}

	rh := ctxrouter.RouterHandler()
	rdh := redirect.RedirectHandler(rh)
	pmh := httparams.ParamsHandler(rdh)
	sh := session.SessionHandler(pmh)
	ridh := requestId.RequestIDHandler(sh)

	httplogger := log.HttpLoggerHandler(ridh)

	log.Logger.Fatal(http.ListenAndServe(":"+httpPort, httplogger))
}
