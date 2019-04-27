package httpfunc

import (
//	"fmt"
	"log"
	"net/http"
	"time"
)

func GetHttpStatus(urls []string)(<-chan string, <-chan string) {
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
	return reqUrl,statusChan // return url,status
}

func GetHttpStatusDur(urls []string)(<-chan string, <-chan string, <-chan time.Duration) {
	// make channel in fucntion
	reqUrl := make(chan string)
	statusChan := make(chan string)
	duration := make(chan time.Duration)

	for _, url := range urls {
		go func(url string) {
			// starttime
			starttime := time.Now()

			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			// duration
			endtime := time.Now()
			duration <- endtime.Sub(starttime)

			reqUrl <- url
			statusChan <- res.Status
		}(url)
	}
	return reqUrl,statusChan,duration // return url,status,duration
}

func ConnHttp(url string)(<-chan string) {
	// make channel in function
	statusChan := make(chan string)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	statusChan <- res.Status
	
	return statusChan
}
