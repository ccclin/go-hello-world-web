package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

// Hello is hello struct
type Hello struct {
	IP string
}

func main() {
	http.HandleFunc("/", RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RootHandler is path '/'
func RootHandler(w http.ResponseWriter, r *http.Request) {
	hello := Hello{readUserIP(r)}

	fp := path.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(fp)

	if err := tmpl.Execute(w, hello); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
