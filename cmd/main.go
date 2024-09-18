package main

import (
	"log"

	pkg "github.com/codescalersinternships/Datetime-client-RawanMostafa/pkg"
)

func main() {
	c := pkg.CreateClient()
	resp, err := pkg.SendRequest(c, "8080")
	if err != nil {
		log.Fatal(err)
	}
	body, err := pkg.ReadBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}
