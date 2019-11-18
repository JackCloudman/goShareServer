package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getInfo/", searchByHashHandler).Methods("POST")
	r.HandleFunc("/uploadFile/", uploadFileHandler).Methods("POST")
	r.HandleFunc("/searchFile/", searchFileHandler).Methods("POST", "OPTIONS")
	startSql()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Escuchando...")
	log.Fatal(srv.ListenAndServe())
}
