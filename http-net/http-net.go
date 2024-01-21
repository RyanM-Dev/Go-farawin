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

func authMiddleware(w http.ResponseWriter, r *http.Request) bool {
	key := r.Header.Get("Authorization")
	if key == "test" {
		return true

	} else {
		fmt.Printf("failed to authorize %d", http.StatusUnauthorized)
		io.WriteString(w, fmt.Sprintf("failed to authorize %d", http.StatusUnauthorized))

		return false
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apis := map[string]string{
			"GET /":        "apis list",
			"POST /print":  "print Query or body(body has priority )",
			"GET /hello":   "Greeting",
			"DELETE /stop": "exit API server",
		}
		jsonAPI, err := json.MarshalIndent(apis, "\n", " ")
		if err != nil {
			fmt.Printf("failed to marshal: %s", err)
		}
		io.WriteString(w, string(jsonAPI))
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if authMiddleware(w, r) {

			io.WriteString(w, "Welcome to my Server!")

		}
	})
	mux.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
		if authMiddleware(w, r) {

			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("could not read body: %s\n", err)
			}
			fmt.Printf("POST / request body:\n%s\n", body)
			io.WriteString(w, string(body))
			if string(body) == "" {

				if hasMsg := r.URL.Query().Has("msg"); hasMsg {

					msg := r.URL.Query().Get("msg")

					io.WriteString(w, fmt.Sprintf("message availability:%t\n\n\nmessage content:%s\n", hasMsg, msg))
				}
			}

		}
	})
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		if authMiddleware(w, r) {

			os.Exit(1)
		}
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("http server is closed")
		} else {
			fmt.Printf("filed to start server: %s", err)
		}
	}

}
