// concurrent web crawler
// Adopted from the example from "The Go Programming Language"
package main

import (
    "log"
    "os"

    "gopl.io/ch5/links"
)

var done = make(chan bool)
//var rateLimiter = make(chan int, 4)

// visits given URL and returns all the links present in that page
func crawl(url string) []string {
    //rateLimiter <- 0
    log.Println(url)

    links, err := links.Extract(url)

    //<-rateLimiter

    if err != nil {
        log.Print(err) // print error and ignore
    }
    return links
}

func startCrawler(worklist chan string) {
    visited := make(map[string]bool)

    for link := range worklist {
        if visited[link] {
            continue // skip already visited pages
        }
        visited[link] = true // mark as visited

        // passing link as argument is very important here.
        // If we don't do that then the function will pick the value of link
        // at the time of running, not at the time of starting this
        // function. Since the loop variable `link` is changing in every
        // iteration, it is important to pass a copy of that to the
        // goroutine
        go func(link string) {
            newLinks := crawl(link)
            for _, newLink := range newLinks {
                worklist <- newLink
            }
        }(link)
    }
}

func main() {
    worklist := make(chan string)
    go startCrawler(worklist)

    worklist <- os.Args[1]

    <-done
}
