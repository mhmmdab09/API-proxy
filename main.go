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

	var pointToAddressService anyService
	pointToAddressService.ID = "01"
	pointToAddressService.name = "Neshan API"
	pointToAddressService.baseURL = "https://api.neshan.org/v5/reverse"
	pointToAddressService.requestPATH = "?lat=32.654012&lng=51.666944"
	pointToAddressService.secretKey = "Api-Key"
	pointToAddressService.secretValue = "service.f586da437b9147999e42808212e4b573"

	var distanceService anyService
	distanceService.ID = "02"
	distanceService.name = "Weather API"
	distanceService.baseURL = "https://api.neshan.org/v1/distance-matrix"
	distanceService.requestPATH = "?type=car&origins=36.3177579,59.5323219&destinations=36.35067,59.5451965"
	distanceService.secretKey = "Api-Key"
	distanceService.secretValue = "service.21db1d0baa3c42838f0cdf68a8c8073a"

	callService(distanceService.secretKey,
		distanceService.secretValue,
		distanceService.ID,
		distanceService.baseURL,
		distanceService.requestPATH)

	callService(pointToAddressService.secretKey,
		pointToAddressService.secretValue,
		pointToAddressService.ID,
		pointToAddressService.baseURL,
		pointToAddressService.requestPATH)

	/*
		http.HandleFunc("/", rootHandler)
		err := http.ListenAndServe("localhost:11111", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/
}
