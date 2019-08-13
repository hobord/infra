package main

import (
	"log"
	"net/http"
	"os"

	httparams "github.com/hobord/infra/httparams"
	redirect "github.com/hobord/infra/redirect"
	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
)

type config struct {
	port        string
	sessionkey  string
	log         bool
	metrics     bool
	metricsPort string
}

func main() {
	// cfg := config{
	// 	port:        "8100",
	// 	metrics:     false,
	// 	metricsPort: "9090"}

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8100"
	}

	// Proxy
	// router
	rdh := redirect.RedirectHandler()
	pmh := httparams.ParamsHandler(rdh)
	sh := session.SessionHandler(pmh)
	ridh := requestId.RequestIDHandler(sh)

	logger := logger(ridh)

	log.Fatal(http.ListenAndServe(":"+httpPort, logger))
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before Logger")
		next.ServeHTTP(w, r) // call original
		log.Println("After Logger")
	})
}
