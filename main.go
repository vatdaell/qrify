package main

import (
	"net/http"

	render "gopkg.in/unrolled/render.v1"

	"github.com/gorilla/mux"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	m := mux.NewRouter()
	m.HandleFunc("/", IndexHandler)
	m.HandleFunc("/qrify", QRHandler)
	http.Handle("/", m)
	http.ListenAndServe(":80", nil)
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	r := render.New()
	r.HTML(w, http.StatusOK, "example", nil)
}

func QRHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(png)

}
