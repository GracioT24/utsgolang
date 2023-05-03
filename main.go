package main

import (
	"PopularProgrammingUTSno3/mahasiswa"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port = ":5050"

	router := mux.NewRouter()

	router.HandleFunc("/nama", mahasiswa.BuilData).Methods("POST")
	router.HandleFunc("/semuadata", mahasiswa.InspectData).Methods("GET")

	server := &http.Server{Addr: port, Handler: router}

	server.ListenAndServe()
}
