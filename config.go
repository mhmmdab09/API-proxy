package config

type anyService struct {
	ID          string
	name        string
	baseURL     string
	requestPATH string
	secretKey   string
	secretValue string
}

var pointToAddressService anyService = anyService{
	ID:        "01",
	name:      "Address API",
	baseURL:   "https://api.neshan.org/v5/reverse",
	secretKey: "Api-Key",
}

var distanceService anyService = anyService{
	ID:        "02",
	name:      "Distance API",
	baseURL:   "https://api.neshan.org/v1/distance-matrix",
	secretKey: "Api-Key",
}

type anyClient struct {
	ID           string
	name         string
	clientToken  string
	pathPrefix   string
	requestParam string
}

var myClient anyClient = anyClient{
	ID:          "01",
	name:        "client one",
	clientToken: "",
	pathPrefix:  "api1/v1/",
}
