package main

import (
	"cmp"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

func serve() {
	context, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	port := cmp.Or(os.Getenv("PORT"), "8080")

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: router,
	}

	go func() {
		server.ListenAndServe()
	}()

	slog.Info("Serving", "port", port)

	<-context.Done()
	server.Shutdown(context)
}

func main() {
	serve()
}
