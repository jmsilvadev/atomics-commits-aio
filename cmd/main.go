package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	server "github.com/jmsilvadev/atomics-commits-aio/http"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	log.Println("starting API server")
	router := server.NewRouter()
	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		cancel()
	}()

	go func() {
		log.Println("server listening at 8080")
		srv.ListenAndServe()
	}()
	<-ctx.Done()

	srv.Shutdown(ctx)
	log.Println("Shutting down service")
}
