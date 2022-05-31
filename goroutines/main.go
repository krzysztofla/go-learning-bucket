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

	for _, link := range links {
		checkSiteStatus(link)
	}
}

func checkSiteStatus(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "webstie is down lol")
		return
	}
	fmt.Println(link, "wow it works")
}
