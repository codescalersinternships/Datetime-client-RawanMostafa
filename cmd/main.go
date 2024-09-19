package main

import (
	"flag"
	"log"
	"os"

	pkg "github.com/codescalersinternships/Datetime-client-RawanMostafa/pkg"
)

func setEnv(url string, endpoint string, port string) {
	err := os.Setenv("DATETIME_URL", url)
	if err != nil {
		log.Fatalf("Error in setting url env var: %v", err)
	}
	err = os.Setenv("DATETIME_ENDPOINT", endpoint)
	if err != nil {
		log.Fatalf("Error in setting endpoint env var: %v", err)
	}
	err = os.Setenv("DATETIME_PORT", port)
	if err != nil {
		log.Fatalf("Error in setting port env var: %v", err)
	}
}

func getFlags() (string, string, string) {
	var port string
	flag.StringVar(&port, "port", "", "Specifies the port")

	var url string
	flag.StringVar(&url, "url", "", "Specifies the base url")

	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "", "Specifies the endpoint")

	flag.Parse()
	return url, endpoint, port
}
func main() {

	url, endpoint, port := getFlags()

	setEnv(url, endpoint, port)

	c := pkg.CreateClient()
	resp, err := pkg.SendRequest(c)
	if err != nil {
		log.Fatal(err)
	}
	body, err := pkg.ReadBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}
