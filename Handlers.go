package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string  `json:"message"`
	Files   []*File `json:"files,omitempty"`
	Error   string  `json:"error,omitempty"`
	Users   []*Peer `json:"users,omitempty"`
}

func writeResponse(w http.ResponseWriter, message Response, status int) {
	w.WriteHeader(status)
	response, _ := json.Marshal(message)
	w.Write(response)
}
func searchByHashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()              // Parses the request body
	hash := r.Form.Get("hash") // x will be "" if parameter is not set
	m := Response{}
	if hash == "" {
		m.Message = "Falta hash para hacer la busqueda"
		writeResponse(w, m, http.StatusBadRequest)
		return
	}
	f := searchByHash(hash)
	if f == nil {
		m.Message = "Archivo no encontrado"
		writeResponse(w, m, http.StatusNotFound)
		return
	}
	m.Files = append(m.Files, f)
	writeResponse(w, m, http.StatusOK)
}
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := &Response{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		response := &Response{}
		response.Message = "Error al procesar la solicitud"
		response.Error = err.Error()
		writeResponse(w, *response, http.StatusBadRequest)
		log.Print(err.Error())
		return
	}
	if request.Files == nil || request.Users == nil {
		response := &Response{}
		response.Message = "Error al procesar la solicitud"
		writeResponse(w, *response, http.StatusBadRequest)
		return
	}
	uploadfile(request.Files[0], request.Users[0])
	fmt.Printf("%+v\n", request.Files[0])
	fmt.Printf("%+v\n", request.Users[0])
}
func searchFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	request := &Response{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		response := &Response{}
		response.Message = "Error al procesar la solicitud"
		response.Error = err.Error()
		writeResponse(w, *response, http.StatusBadRequest)
		log.Print(err.Error())
		log.Println(r.Body)
		return
	}
	if request.Message == "" {
		response := &Response{}
		response.Message = "No he encontrado ningun archivo"
		writeResponse(w, *response, http.StatusOK)
		return
	}
	fmt.Printf("%s\n", request.Message)
	files := searchByName(request.Message)
	response := &Response{}
	response.Files = files
	writeResponse(w, *response, http.StatusOK)
}
