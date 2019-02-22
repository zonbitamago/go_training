package main

import "fmt"

type Fetcher interface {
	// fetch結果のbody部とURLの組み合わせを返却する
	// Fetch(url string) (body string, urls []string, err error)
	Fetch(url string, body chan string, urls chan []string, err chan error)
}

// Crawl 引数にcrawl開始のURLと、depthを指定し、再帰的にdepth分のfetchを行う。
func Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation dosen't do either.
	if depth <= 0 {
		return
	}
	cUrl := make(chan string)
	cUrls := make(chan []string)
	cErr := make(chan error)
	go fetcher.Fetch(url, cUrl, cUrls, cErr)
	body := <-cUrl
	urls := <-cUrls
	err := <-cErr
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// func (f fakeFetcher) Fetch(url string) (string, []string, error) {
func (f fakeFetcher) Fetch(url string, body chan string, urls chan []string, err chan error) {
	if res, ok := f[url]; ok {
		body <- res.body
		urls <- res.urls
		err <- nil
		// return res.body, res.urls, nil
	}

	body <- ""
	urls <- nil
	err <- fmt.Errorf("not found: %s", url)

	// return "", nil, fmt.Errorf("not found: %s", url)
}

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
