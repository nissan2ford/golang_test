package main

import (
	"fmt"
	"log"
	"net/http"
)

func getStatus(urls []string)(<-chan string, <-chan string) {
	// make channel in fucntion
	reqUrl := make(chan string)
	statusChan := make(chan string)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			reqUrl <- url
			statusChan <- res.Status
		}(url)
	}
	return reqUrl,statusChan // return url and channel
}

func main() {
	urls := []string{
		"http://192.168.11.1",
		"http://www.yahoo.co.jp",
	}

	url,statusChan := getStatus(urls)
	
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-url,<-statusChan)
	}
}
