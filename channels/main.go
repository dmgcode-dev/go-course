package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	//for i := 0; i < len(urls); i++ {
	//fmt.Println(<-c)
	// for {
	//    go checkUrl(<-c, c)
	for l := range c {
		go func(u string) {
			time.Sleep(5 * time.Second)
			checkUrl(u, c)
		}(l)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		//c <- "Might be down I think"
		c <- url
		return
	}

	fmt.Println(url, "is up")
	//c <- "yep it's up"
	c <- url
}
