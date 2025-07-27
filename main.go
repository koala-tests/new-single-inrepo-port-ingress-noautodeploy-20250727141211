package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// defineFlags contains all the flag definitions, which are called before flag.Parse().
func parseFlags() {
	flag.Parse()
}

func main() {
	log.Printf("Server started.")
	parseFlags()

	router := http.NewServeMux()
	router.HandleFunc("/koala", HelloKoala)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

func HelloKoala(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	VERSION, err := ioutil.ReadFile("VERSION")
	if err != nil {
		fmt.Fprintln(w, "Error reading VERSION file (new-single-inrepo-port-ingress-noautodeploy-20250727141211): "+err.Error())
		return
	}
	fmt.Fprintln(w, "Hello world! service:new-single-inrepo-port-ingress-noautodeploy-20250727141211 version: "+string(VERSION))
}
