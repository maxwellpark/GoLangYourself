package req

import (
	"io/ioutil" // IO
	"log"       // Logger
	"net/http"  // Http client
)

var url string = "https://reqres.in/api/users/1"

// GET request and read stream 
func GetUser() string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
	
	// Read bytes from body 
	bytes, errRead := ioutil.ReadAll(resp.Body)
	defer func() {
		// Close the stream 
		e := resp.Body.Close() 
		if e != nil {
			log.Fatal(e)
		}
	}()
	
	if errRead != nil {
		log.Fatal(errRead)
	}
	content := string(bytes)
	log.Print(content)
	return content
}
