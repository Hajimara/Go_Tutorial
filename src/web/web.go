package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	myChannel := make(chan Page)
	urls := []string{"https://naver.com", "https://google.co.kr", "https://daum.net"}
	for _, url := range urls {
		go responseSize(url, myChannel)
	}
	for i := 0; i < len(urls); i++ {
		page := <-myChannel
		fmt.Println(page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting... ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
