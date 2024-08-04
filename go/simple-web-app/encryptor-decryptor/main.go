package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/encrypt", EncryptionHandler)
	http.HandleFunc("/decrypt", DecryptHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("body")
	_, err := w.Write([]byte(caesarCipher(input, shift)))
	if err != nil {
		log.Fatalln(err)
	}

}
func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	cipherText := r.FormValue("body")
	password := caesarCipher(cipherText, 26-shift)
	_, err := w.Write([]byte(password))
	if err != nil {
		log.Fatalln(err)
	}
}
