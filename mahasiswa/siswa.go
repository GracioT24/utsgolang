package mahasiswa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Mahasiswa struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"Nim"`
	Alamat string `json:"Alamat"`
}

var data []Mahasiswa

func BuilData(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person Mahasiswa

	person.Nama = request.FormValue("Nama")
	person.NIM = request.FormValue("NIM")
	person.Alamat = request.FormValue("Alamat")

	if person.Nama == "" || person.NIM == "" || person.Alamat == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Status":"Error", "Note":"Nama,NIM, dan Alamat wajib untuk diisi"}`)
		return
	} else {
		data = append(data, person)

		ImporData, err := json.Marshal(person)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `{"Status":"Error", "Note":"terdapat error dalam server"}`)
			return
		} else {
			fmt.Fprint(w, string(ImporData))
		}
	}
}

func InspectData(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"Status":"Error", "Note":"Tambahkan data terlebih dahulu"}`)
		return
	}

	ImporData, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"Status":"Error", "Note":"terdapat error dalam server"}`)
		return

	} else {
		fmt.Fprint(w, string(ImporData))
	}
}
