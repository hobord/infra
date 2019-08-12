package httparams

import (
	"log"
	"net/http"

	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
)

// ParamsHandler make parampeeling
func ParamsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Params Init")
		next.ServeHTTP(w, r)
		sessionID := session.SessionIDFromContext(r.Context())
		log.Println("SessionID:" + sessionID)
		reqID := requestId.RequestIDFromContext(r.Context())
		log.Println("Hello request ID:" + reqID)
		log.Println("Params END")
	})
}
