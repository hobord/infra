package redirect

import (
	"log"
	"net/http"
	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
)
// RedirectHandler is a redirect midleware
func RedirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirect Init")
		sessionID := session.SessionIDFromContext(r.Context())
		log.Println("SessionID:" + sessionID)
		reqID := requestId.RequestIDFromContext(r.Context())
		log.Println("Hello request ID:" + reqID)
		log.Println("Redirect END")
	})
}