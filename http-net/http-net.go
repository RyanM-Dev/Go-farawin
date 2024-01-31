package main

// test for commit
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const port = 8080

// func authMiddleware(w http.ResponseWriter, r *http.Request) bool {
// 	key := r.Header.Get("Authorization")
// 	if key == "test" {
// 		return true

// 	} else {
// 		fmt.Printf("failed to authorize %d", http.StatusUnauthorized)
// 		io.WriteString(w, fmt.Sprintf("failed to authorize %d", http.StatusUnauthorized))

// 		return false
// 	}
// }

//list of available APIs

var apis = map[string]string{
	"GET /":        "apis list",
	"POST /print":  "print Query or body(body has priority )",
	"GET /hello":   "Greeting",
	"DELETE /stop": "exit API server",
}

// this function lists the apis by / url

func apiList(w http.ResponseWriter, r *http.Request) {

	jsonAPI, err := json.MarshalIndent(apis, "\n", " ")
	if err != nil {
		fmt.Printf("failed to marshal: %s", err)
		return
	}
	io.WriteString(w, string(jsonAPI))

}

// this function greets the client on /hello url

func greet(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Welcome to my Server!")

}

// this method prints the data received through body and query

func queryAndBodyPrinter(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if string(body) != "" {

		if err != nil {
			fmt.Printf("could not read body: %s\n", err)
			return
		}
		fmt.Printf("POST / request body:\n%s\n", body)
		io.WriteString(w, string(body))
	} else if string(body) == "" {

		if hasMsg := r.URL.Query().Has("msg"); hasMsg {

			msg := r.URL.Query().Get("msg")
			fmt.Printf("message is:%s\n", msg)

			io.WriteString(w, fmt.Sprintf("message availability:%t\n\n\nmessage content:%s\n", hasMsg, msg))
		}
	}

}

//this method stops the server on /stop url

func stopServer(w http.ResponseWriter, r *http.Request) {

	os.Exit(1)
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform authentication check here
		authToken := r.Header.Get("Authorization")
		if authToken != "test2" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication passes, call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", authenticationMiddleware(http.HandlerFunc(apiList)))
	mux.Handle("/hello", authenticationMiddleware(http.HandlerFunc(greet)))
	mux.Handle("/print", authenticationMiddleware(http.HandlerFunc(queryAndBodyPrinter)))
	mux.Handle("/stop", authenticationMiddleware(http.HandlerFunc(stopServer)))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("Server on localhost:%d is starting...\n", port)

	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("http server is closed")

		} else {
			fmt.Printf("filed to start server: %s", err)

		}
		return
	}

}
