package main

import (
    "log"
    "os"

    "gopl.io/ch5/links"
)

var done = make(chan bool)

// visits given URL and returns all the links present in that page
func crawl(url string) []string {
    log.Println(url)
    links, err := links.Extract(url)
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

        // WARNING: there is a subtle bug here!
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

    <-done // blocks indefinitely
}
