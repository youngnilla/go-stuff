// packages
package scrape

// imports
import (
	//"io"
	"fmt"
	"log"
	"net/http"
	//"os"
)

// main func
func main() {
	// Make HTTP GET req
	// (demo) response, err := http.Get("https://devdungeon.com")
	response, err := http.Get("https://www.pro-football-reference.com/years/2021/passing.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// print out result to stdout
	fmt.Println(response)
}
