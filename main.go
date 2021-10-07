package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("status:%v Addr:%v\n", http.StatusOK, r.RemoteAddr)
	ver := runtime.Version()
	for k := range r.Header {
		value := r.Header.Get(k)
		w.Header().Set(k, value)
	}
	w.Header().Set("Version", ver)
}
