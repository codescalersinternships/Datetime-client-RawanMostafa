package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	c := http.Client{}
	resp, err := c.Get("http://localhost:8080/datetime")
	if err != nil {
		log.Fatalf("Error in sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error in reading request body: %v", err)
		return
	}
	fmt.Println(string(body))
}
