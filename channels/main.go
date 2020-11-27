package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Channels")
	links := []string{
		"https://www.google.com",
		"https://stackoverflow.com",
		"https://9gag.com",
		"https://www.reddit.com",
		"https://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		// fmt.Println(link)
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		// time.Sleep(5 * time.Second)
		// Use function literals
		// go checkLink(l, c)
		go func(link string, channel chan string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l, c)
	}
}
func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	// fmt.Println("Reached here")
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down")
		// c <- " Might be down"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// c <- " is up!"
	c <- link
	// return
}
