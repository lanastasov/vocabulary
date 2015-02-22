package main

import (
	"flag"
	"net/http"
)

var (
	port = flag.String("http", ":8083", "http address")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", Index)
	http.HandleFunc("/lanastasov", Lanastasov)
	http.ListenAndServe(*port, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func Lanastasov(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./lanastasov/index.html")
}
