package main

import (
	"fmt"
	"sync"
)

// Exercise: Web Crawler

// In this exercise you'll use Go's concurrency features to parallelize a web crawler.

// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}
	fmt.Printf("found: %q %q\n", body, url)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	// Crawl("https://golang.org/", 4, fetcher)	
	cacheChan := make(chan string)
	depth := 4
	go ParallelCrawl("https://golang.org/", 4, fetcher, cacheChan)
	for i := 0; i < depth; i++ {
		fmt.Println(<-cacheChan)
	}
}

type Cache map[string]bool

type SafeCache struct {
	v Cache
	mu sync.Mutex
}

func (sf *SafeCache) Store(url string) {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	sf.v[url] = true
}

func (sf *SafeCache) Check(url string) bool {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	_, ok := sf.v[url]
	return ok
}

func ParallelCrawl(url string, depth int, fetcher Fetcher, ch chan string) {
	if depth <= 0 {
		return
	}
	
	if ok := cache.Check(url); !ok {
		cache.Store(url)
	} else {
		ch <- fmt.Sprintf("%q cached", url)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- fmt.Sprintf("%q not found", url)
		return
	}
	
	fmt.Printf("found: %q %q\n", body, url)
	
	for _, u := range urls {
		go ParallelCrawl(u, depth-1, fetcher, ch)
	}
	return
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var cache = SafeCache{
	v: make(Cache),
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
