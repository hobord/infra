package session

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
)

type sessionKey int

const sessionIDKey sessionKey = 0

// SessionHandler is a middleware handler
func SessionHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Session Handler")

		ctx := newContextWithSessionID(r.Context(), r)
		sessionID := SessionIDFromContext(ctx)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session", Value: sessionID, Expires: expiration}
		http.SetCookie(w, &cookie)

		log.Println("SessionID:" + sessionID)
		log.Println("Session NEXT")
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Println("Session NEXT END")
	}
}

// SessionIDFromContext return session id
func SessionIDFromContext(ctx context.Context) string {
	return ctx.Value(sessionIDKey).(string)
}

func newContextWithSessionID(ctx context.Context, r *http.Request) context.Context {
	var sessionID string
	cookie, err := r.Cookie(getSessionCookieKey(r))

	if err != nil || cookie == nil || cookie.Value == "" {
		sessionID = createSession()
	} else {
		sessionID = cookie.Value
	}

	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func createSession() string {
	var client DSessionServiceClient
	serverAddr := os.Getenv("SESSION_GRP_SERVER")
	if serverAddr == "" {
		serverAddr = "10.20.35.111:30645"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	client = NewDSessionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dsession, err := client.CreateSession(ctx, &CreateSessionMessage{Ttl: 0})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	return dsession.Id
}

func getSessionCookieKey(r *http.Request) string {
	ck := os.Getenv("SESSION_COOKIE_KEY")
	if ck == "" {
		ck = "session"
	}
	return ck
}
