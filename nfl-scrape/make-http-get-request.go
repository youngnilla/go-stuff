// packages
package main

// imports
import (
	"io"
	"log"
	"net/http"
	"os"
)

// main func
func main() {
	// Make HTTP GET req
	response, err := http.Get("https://devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}
