package main

import (
	"context"
	"glooko/brazeservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	lg := log.New(os.Stdout, "braze-in", log.LstdFlags)
	bh := handlers.NewBraze(lg)
	ih := handlers.NewIndex(lg)
	//gz := handlers.NewGzip(lg)
	sm := mux.NewRouter()

	getSubRtr := sm.Methods("GET").Subrouter()
	getSubRtr.HandleFunc("/braze", bh.ListBrazeCalls)
	getSubRtr.HandleFunc("/", ih.EchoBody)
	postSubRtr := sm.Methods("POST").Subrouter()
	postSubRtr.HandleFunc("/braze", bh.AddBrazeCall)

	postSubRtr.Use(bh.BrazeCallValidate)
	//getSubRtr.Use(gz.Compress)

	ch := gHandlers.CORS(gHandlers.AllowedOrigins([]string{"*"}))
	sv := &http.Server{
		Addr:         ":9090",
		Handler:      gHandlers.CompressHandler(ch(sm)),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := sv.ListenAndServe()
		if err != nil {
			lg.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, os.Interrupt)
	sig := <-sigCh
	lg.Println("Shutting down on signal ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	sv.Shutdown(ctx)
}
