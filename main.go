package main

import (
	"fmt"
	"net/http"
	"os"
)

type anyService struct {
	ID      string
	name    string
	baseURL string
	auth    string
}

type anyClient struct {
	ID            string
	name          string
	clientToken   string
	auth          string
	calledService string
	pathPrefix    string
}

func readConfig() {
	//todo
}

func callService() {
	//todo
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}

func getSecret() {
	//todo
}

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
