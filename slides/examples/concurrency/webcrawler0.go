package main

// version 0: crawls just two levels, sequentially

import (
    "log"
    "os"

    "gopl.io/ch5/links"
)

// visits given URL and returns all the links present in that page
func crawl(url string) []string {
    log.Println(url)
    links, err := links.Extract(url)
    if err != nil {
        log.Print(err) // print error and ignore
    }
    return links
}

func main() {
    url := os.Args[1]
    visited := make(map[string]bool)

    links := crawl(url)
    for _, link := range links {
        if visited[link] {
            continue
        }
        visited[link] = true
        crawl(link)
    }
}
