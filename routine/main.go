package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checker(link, c)
	}

	//time.Sleep(time.Second * 10)

	for range c {
		fmt.Println(<-c)
	}
}

func checker(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	c <- link + "is up"
}
