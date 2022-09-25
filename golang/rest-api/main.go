// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strconv"
// )

// type Note struct {
// 	Index   int    `json:"index"`
// 	Name    string `json:"name`
// 	Content string `json:"content`
// }

// type Response struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// var (
// 	notes []Note
// )

// func sendResponse(w http.ResponseWriter, code int, message string, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	var response Response
// 	response.Code = code
// 	response.Message = message
// 	response.Data = data

// 	dByte, _ := json.Marshal(response)

// 	w.Write(dByte)
// }

// func remove(slice []Note, s int) []Note {
// 	return append(slice[:s], slice[s+1:]...)
// }

// func main() {
// 	// tambah note
// 	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodPost {
// 			defer r.Body.Close()
// 			dByte, err := io.ReadAll(r.Body)
// 			if err != nil {
// 				sendResponse(w, http.StatusInternalServerError, "Terjadi Masalah", nil)
// 				return
// 			}

// 			var note Note
// 			index := len(notes)
// 			json.Unmarshal(dByte, &note)

// 			note.Index = index

// 			notes = append(notes, note)
// 			sendResponse(w, http.StatusCreated, "Data Note Berhasil diTambahkan", "")
// 			return
// 		}

// 		if r.Method == http.MethodGet {
// 			id := r.URL.Query().Get("id")
// 			if id != "" {
// 				idT, _ := strconv.Atoi(id)
// 				sendResponse(w, http.StatusOK, "Data Berhasil didapatkan", notes[idT])
// 				return
// 			}
// 			sendResponse(w, http.StatusOK, "Data Berhasil didapatkan", notes)
// 			return
// 		}

// 		if r.Method == http.MethodDelete {
// 			id := r.URL.Query().Get("id")

// 			idT, _ := strconv.Atoi(id)
// 			remove(notes, idT)
// 			sendResponse(w, http.StatusOK, "Data Berhasil dihapus", "")
// 			return

// 		}
// 	})
// 	// lihat semua note
// 	port := 5000
// 	server := http.Server{
// 		Addr: fmt.Sprintf(":%d", port),
// 	}
// 	fmt.Println("server running on port : ", port)
// 	server.ListenAndServe()
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "strconv"
)

type Note struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Request struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	notes []Note
)

// respon
func sendResponse(w http.ResponseWriter, code int, status string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	var response Response
	response.Code = code
	response.Status = status
	response.Message = message

	dByte, _ := json.Marshal(response)

	w.Write(dByte)
}

// GetNote
func NOTES(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// id := r.URL.Query().Get("id")
		// if id != "" {
		// 	idT, _ := strconv.Atoi(id)
		// 	notes[idT]
		// }
		DNote, err := json.Marshal(notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(DNote)
		return
	}

	if r.Method == "POST" {
		var nt Note
		decodeJSON := json.NewDecoder(r.Body)
		if err := decodeJSON.Decode(&nt); err != nil {
			log.Fatal(err)
		}
		nt.Id = len(notes) + 1
		notes = append(notes, nt)

		DNote, _ := json.Marshal(nt) // ke byte
		w.Write(DNote)
		return
	}

	http.Error(w, "hayo  capek ya wkwkw?", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/notes", NOTES)
	fmt.Println("on proses")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
