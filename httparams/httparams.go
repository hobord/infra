package httparams

import (
	"log"
	"net/http"
	"net/url"

	st "github.com/golang/protobuf/ptypes/struct"
	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
)

// ParamsHandler make parampeeling
func ParamsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Params Init")

		parameters := filterParametersByURL(r)
		if len(parameters) > 0 {
			values := makeSessionValues(parameters)
			sessionID := session.SessionIDFromContext(r.Context())
			log.Println("SessionID:" + sessionID)
			session.AddValuesToSession(sessionID, values)
		}

		reqID := requestId.RequestIDFromContext(r.Context())
		log.Println("Hello request ID:" + reqID)
		next.ServeHTTP(w, r)
		log.Println("Params END")
	})
}

func filterParametersByURL(r *http.Request) url.Values {
	// TODO: make filter logic, for saving different parameters by domain
	parameters := r.URL.Query()
	return parameters
}

func makeSessionValues(parameters url.Values) session.Values {
	values := make(session.Values)
	for key, value := range parameters {
		log.Println("i:" + key + " key:" + value[0])
		values[key] = &st.Value{Kind: &st.Value_StringValue{value[0]}} // TODO: Multiple value?
	}
	return values
}
