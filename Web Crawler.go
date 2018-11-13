// make_http_request_with_timeout.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request
	resp, err := client.Get("http://tauerperfumes.com/l-air-du-desert-marocain.html")
	bytes, err := ioutil.ReadAll(resp.Body)

	fmt.Print("HTML:\n\n", string(bytes))

	resp.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}
