package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type anyService struct {
	ID          string
	name        string
	baseURL     string
	secretKey   string
	secretValue string
}

var pointToAddressService anyService
var distanceService anyService

type anyClient struct {
	ID          string
	name        string
	clientToken string
	pathPrefix  string
}

var myClient anyClient = anyClient{
	ID:          "01",
	name:        "client one",
	clientToken: "",
	pathPrefix:  "api1/v1/",
}

func readConfig() {

	//todo : yaml file has not been read correctly
	confContent, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	conf := make(map[string]anyService)
	err2 := yaml.Unmarshal(confContent, &conf)

	if err2 != nil {
		log.Fatal(err2)
	}

	pointToAddressService = conf["pointToAddressCon"]
	distanceService = conf["distanceCon"]
	fmt.Println(pointToAddressService)
	fmt.Println(distanceService)

	/*
	   pointToAddressService.ID = "01"
	   pointToAddressService.name = "Address API"
	   pointToAddressService.baseURL = "https://api.neshan.org/v5/reverse"
	   pointToAddressService.secretKey = "Api-Key"

	   distanceService.ID = "02"
	   distanceService.name = "Distance API"
	   distanceService.baseURL = "https://api.neshan.org/v1/distance-matrix"
	   distanceService.secretKey = "Api-Key"
	*/
}

func callService(authKey string, authValue string, serviceID string, baseU string, params string) (out string) {
	url := baseU + "?" + params
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add(authKey, authValue)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
	out = string(body)
	return
}

func getSecret(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/secret" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "secret.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		pointToAddressService.secretValue = r.FormValue("address_secret")
		distanceService.secretValue = r.FormValue("distance_secret")
		fmt.Fprintf(w, "Address API secret = %s\n", pointToAddressService.secretValue)
		fmt.Fprintf(w, "Distance API secret = %s\n", distanceService.secretValue)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func clientHandler(w http.ResponseWriter, r *http.Request) {

	var realPath string = strings.ReplaceAll(r.URL.Path, myClient.pathPrefix, "")

	switch realPath {
	case "/distance/":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(callService(distanceService.secretKey,
				distanceService.secretValue,
				distanceService.ID,
				distanceService.baseURL,
				r.URL.RawQuery)))
		}
	case "/address/":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(callService(pointToAddressService.secretKey,
				pointToAddressService.secretValue,
				pointToAddressService.ID,
				pointToAddressService.baseURL,
				r.URL.RawQuery)))
		}
	default:
		{
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Asset not found\n"))
		}
	}

}

func main() {
	http.HandleFunc("/secret", getSecret)
	http.HandleFunc("/api1/v1/address/", clientHandler)
	http.HandleFunc("/api1/v1/distance/", clientHandler)

	readConfig()

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":11111", nil); err != nil {
		log.Fatal(err)
	}

}
