package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const port = 8080

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to my Server!")

	})
	mux.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
		hasMsg := r.URL.Query().Has("msg")
		msg := r.URL.Query().Get("msg")

		io.WriteString(w, fmt.Sprintf("message availability:%t\n\n\nmessage content:%s\n", hasMsg, msg))

	})
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
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
