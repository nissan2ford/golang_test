package main

import (
	"fmt"
	"httpfunc"
	"time"
)

func main() {
	urls := []string{
		"http://192.168.11.1",
		"http://www.yahoo.co.jp",
	}

	starttime := time.Now()

	reqUrl,statusChan := httpfunc.GetHttpStatus(urls)

	endtime := time.Now()

	duration := endtime.Sub(starttime)

//	reqUrl,statusChan,duration := httpfunc.GetHttpStatusDur(urls)

	for i := 0; i < len(urls); i++ {
//		fmt.Println(<-reqUrl,<-statusChan)
		fmt.Println(<-reqUrl,<-statusChan,"TAT=",duration)
	}
}
