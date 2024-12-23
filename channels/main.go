package main

import (
	"fmt"
	"net/http"
	"time"
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

	// note keyword: concurrency, parralelism
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	greeting := "Hi There!"

// 	var w sync.WaitGroup

// 	w.Add(1)

// 	go (func() {
// 		fmt.Println(greeting)
// 		defer w.Done()

// 	})()
// 	w.Wait()
// }
