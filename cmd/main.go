package main

import (
	"flag"
	"log"
	"os"
	"time"

	pkg "github.com/codescalersinternships/Datetime-client-RawanMostafa/pkg"
)

const defaltBaseUrl = "http://localhost"
const defaultEndpoint = "/datetime"
const defaultPort = "8083"

func getFlags() (string, string, string) {
	var port string
	flag.StringVar(&port, "port", "", "Specifies the port")

	var baseUrl string
	flag.StringVar(&baseUrl, "baseUrl", "", "Specifies the base url")

	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "", "Specifies the endpoint")

	flag.Parse()
	return baseUrl, endpoint, port
}

func loadConfigs() (string, string, string) {

	baseUrl, endpoint, port := getFlags()
	var found bool

	if baseUrl == "" {
		baseUrl, found = os.LookupEnv("DATETIME_BASEURL")
		if !found {
			baseUrl = defaltBaseUrl
		}
	}
	if endpoint == "" {
		endpoint, found = os.LookupEnv("DATETIME_ENDPOINT")
		if !found {
			endpoint = defaultEndpoint
		}
	}
	if port == "" {
		port, found = os.LookupEnv("DATETIME_PORT")
		if !found {
			port = defaultPort
		}
	}
	return baseUrl, endpoint, port

}

func main() {
	baseUrl, endpoint, port := loadConfigs()

	c := pkg.NewClient(baseUrl, endpoint, port, "text/plain", time.Second)
	timeNow, err := c.GetTime()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(timeNow)
}
