package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	search := "Яндекс"
	urls := []string{
		"https://ya.ru",
		"https://google.com",
		"https://yandex.ru",
	}

	fmt.Println(findStringInUrls(search, urls))
}

func findStringInUrls(search string, urls []string) []string {
	var outUrls []string

	findInUrls := make(chan string)
	for _, url := range urls {
		go findStringInUrl(search, url, findInUrls)
	}

	for i := 0; i < len(urls); i++ {
		findInUrl := <-findInUrls
		if findInUrl != "" {
			outUrls = append(outUrls, findInUrl)
		}
	}

	return outUrls
}

func findStringInUrl(search string, url string, ch chan<- string) {
	var body = getBody(url)
	if strings.Contains(body, search) {
		ch <- url
	} else {
		ch <- ""
	}
}

func getBody(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}