package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := make([]string, 0)
	links = append(links,
		"http://stackoverflow.com/",
		"http://www.google.com/",
		"http://pkg.go.dev/",
		"http://www.facebook.com/")

	c := make(chan string)

	for _, link := range links {
		go checkSiteStatus(link, c)
		fmt.Println(<-c)
	}
}

func checkSiteStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "webstie is down lol")
		c <- link + " webstie is down lol from go channel"
		return
	}
	// fmt.Println(link, "wow it works")
	c <- link + " wow it works - msg from channel"
}
