package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"generative-web/internal/config"
	"generative-web/pkg/database"
	"generative-web/pkg/handlers"
)

func Start(wg *sync.WaitGroup) {

	var conf, err = config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var conn = database.DBConnection{}

	err = conn.InitConnection(conf.Database.ConnStr)

	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)

	// setup router
	router := mux.NewRouter()

	// setup routes
	// basic routes
	router.HandleFunc("/", handlers.Welcome).Methods("GET")
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")
	router.HandleFunc("/status", handlers.Status(conn.Conn)).Methods("GET")
	// route for loading template using query parameter
	router.HandleFunc("/template/load", handlers.LoadTemplate).Methods("GET")
	// make channel for graceful shutdown
	c := make(chan os.Signal, 1)

	go func() {
		defer wg.Done()
		log.Println("Starting server...")
		srv := &http.Server{
			Handler:      router,
			Addr:         fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		c <- os.Interrupt
		log.Println("Server stopped")
	}()
	log.Println("Server started on port " + os.Getenv("PORT"))
	wg.Wait()
	log.Println("Server stopped")
}

func main() {
	var wg sync.WaitGroup
	Start(&wg)
}
