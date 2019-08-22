package redirect

import (
	context "context"
	"log"
	"net/http"
	"os"
	"time"

	requestId "github.com/hobord/infra/requestId"
	session "github.com/hobord/infra/session"
	"google.golang.org/grpc"
)

var redirectConn *grpc.ClientConn

func init() {
	serverAddr := os.Getenv("REDIRECT_GRPC_SERVER")
	if serverAddr == "" {
		serverAddr = "localhost:50052"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	redirectConn = conn
}

// RedirectHandler is a redirect midleware
func RedirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := session.SessionIDFromContext(r.Context())
		reqID := requestId.RequestIDFromContext(r.Context())
		log.Println("SessionID:" + sessionID)

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		u := scheme + "://" + r.Host + r.URL.Path
		redirection := requestRedirection(sessionID, reqID, u, r.Method, r.Header)
		if redirection.HttpStatusCode != 200 {
			http.Redirect(w, r, redirection.Location, int(redirection.HttpStatusCode))
		}

	})
}

func requestRedirection(sessionID, requestID, url, httpMethod string, httpHeaders http.Header) *GetRedirectionResponse {
	log.Printf("call redirect service with: %v", url)
	var client RedirectServiceClient
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client = NewRedirectServiceClient(redirectConn)
	defer cancel()

	headers := make(map[string]*HttpHeader)
	for headerKey, h := range httpHeaders {
		values := []string{}
		for _, val := range h {
			values = append(values, val)
		}
		headers[headerKey] = &HttpHeader{
			Header: values,
		}
	}

	message := &GetRedirectionMessage{
		SessionID:  sessionID,
		RequestID:  requestID,
		Url:        url,
		HttpMethod: httpMethod,
		Headers:    headers,
	}
	redirection, err := client.GetRedirection(ctx, message)

	if err != nil {
		// TODO: logs
		redirection = &GetRedirectionResponse{
			Location:       url,
			HttpStatusCode: 200,
		}
	}

	log.Println(redirection)
	return redirection
}
