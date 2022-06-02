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
	}

	// loop variable link captured by func literal
	// - because go routine is poiting to reference of variable from main scope
	// this can cause incorrect programm behaviour
	// for link := range c {
	// 	go func() {
	// 		fmt.Println(link)
	// 	}()
	// }

	// let's fix this by taking adventage that go is copy by value language
	// and instead of passing reference let's just pass it as parameter to
	// anonymous function ( function literal in go nomenclature)
	for link := range c {
		go func(lin string) {
			fmt.Println(lin)
		}(link)
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
