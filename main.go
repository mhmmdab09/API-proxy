package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type anyService struct {
	ID          string
	name        string
	baseURL     string
	requestPATH string
	secretKey   string
	secretValue string
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		//return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}

func getSecret() {
	//todo
}

func callService(authKey string, authValue string, serviceID string, baseU string, pathURL string) {
	url := baseU + pathURL
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add(authKey, authValue)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

func main() {
	/*
		http.HandleFunc("/", rootHandler)
		err := http.ListenAndServe("localhost:11111", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/
	var pointToAddressService anyService
	pointToAddressService.ID = "01"
	pointToAddressService.name = "Neshan API"
	pointToAddressService.baseURL = "https://api.neshan.org/v5/reverse"
	pointToAddressService.requestPATH = "?lat=32.654012&lng=51.666944"
	pointToAddressService.secretKey = "Api-Key"
	pointToAddressService.secretValue = "service.76e24479dd3f4e3b83d686a2b179d08d"

	var weatherService anyService
	weatherService.ID = "02"
	weatherService.name = "Weather API"
	weatherService.baseURL = ""
	weatherService.secretKey = ""
	weatherService.secretValue = ""

	callService(pointToAddressService.secretKey,
		pointToAddressService.secretValue,
		pointToAddressService.ID,
		pointToAddressService.baseURL,
		pointToAddressService.requestPATH)
}
