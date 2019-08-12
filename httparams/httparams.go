package httparams

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	st "github.com/golang/protobuf/ptypes/struct"
	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
	"google.golang.org/grpc"
)

type sessionValues map[string]*st.Value

// ParamsHandler make parampeeling
func ParamsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Params Init")

		parameters := filterParametersByURL(r)
		if len(parameters) > 0 {
			values := makeSessionValues(parameters)
			sessionID := session.SessionIDFromContext(r.Context())
			log.Println("SessionID:" + sessionID)
			saveToSession(sessionID, values)
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

func makeSessionValues(parameters url.Values) sessionValues {
	values := make(sessionValues)
	for key, value := range parameters {
		log.Println("i:" + key + " key:" + value[0])
		values[key] = &st.Value{Kind: &st.Value_StringValue{value[0]}} // TODO: Multiple value?
	}
	return values
}

func saveToSession(sessionID string, keys sessionValues) {
	var client session.DSessionServiceClient
	serverAddr := os.Getenv("SESSION_GRP_SERVER")
	if serverAddr == "" {
		serverAddr = "10.20.35.111:30645"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	client = session.NewDSessionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := &session.AddValuesToSessionMessage{Id: sessionID, Values: keys}
	response, err := client.AddValuesToSession(ctx, message)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Printf("Session response: %v", response)
}
