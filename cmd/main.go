package main

import (
	"flag"
	"fmt"
	"log"

	pkg "github.com/codescalersinternships/Datetime-client-RawanMostafa/pkg"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 0, "Specifies the port")

	var url string
	flag.StringVar(&url, "url", "", "Specifies the base url")

	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "", "Specifies the endpoint")

	flag.Parse()

	c := pkg.CreateClient()
	resp, err := pkg.SendRequest(c, url, endpoint, fmt.Sprint(port))
	if err != nil {
		log.Fatal(err)
	}
	body, err := pkg.ReadBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}
