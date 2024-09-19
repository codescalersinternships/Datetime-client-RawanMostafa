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

func decideConfigs() (string, string, string) {

	baseUrl, endpoint, port := getFlags()

	if baseUrl == "" {
		envBaseUrl, found := os.LookupEnv("DATETIME_BASEURL")

		if found {
			baseUrl = envBaseUrl
		} else {
			baseUrl = defaltBaseUrl
		}
	}
	if endpoint == "" {
		envEndpoint, found := os.LookupEnv("DATETIME_ENDPOINT")

		if found {
			endpoint = envEndpoint
		} else {
			endpoint = defaultEndpoint
		}
	}
	if port == "" {
		envPort, found := os.LookupEnv("DATETIME_PORT")

		if found {
			port = envPort
		} else {
			port = defaultPort
		}
	}
	return baseUrl, endpoint, port

}

func main() {
	baseUrl, endpoint, port := decideConfigs()

	c := pkg.NewClient(baseUrl, endpoint, port, time.Second)
	resp, err := c.SendRequest()
	if err != nil {
		log.Fatal(err)
	}
	body, err := c.ReadBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}
