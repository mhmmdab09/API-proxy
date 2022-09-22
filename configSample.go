package configSample

type anyService struct {
	ID          string
	name        string
	baseURL     string
	requestPATH string
	secretKey   string
	secretValue string
}

var serviceName anyService = anyService{
	ID:          "service ID",
	name:        "service name",
	baseURL:     "https://api.base.url",
	requestPATH: "?request&path",
	secretKey:   "Api-Key",
}

type anyClient struct {
	ID          string
	name        string
	clientToken string
	pathPrefix  string
}

var clientName anyClient = anyClient{
	ID:          "client ID",
	name:        "client name",
	clientToken: "token",
	pathPrefix:  "/path/prefix",
}
