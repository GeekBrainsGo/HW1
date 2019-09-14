package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}

}

func findLinks(findText string, findLinks []string) (findedLinks []string) {

	for _, link := range findLinks {

		resp, err := http.Get(link)
		failOnError(err, "Get link error:")
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if strings.Contains(string(content), findText) {
			findedLinks = append(findedLinks, link)
		}

	}

	return
}

func main() {

	findText := "Поиск"

	arrLinks := []string{"https://www.yandex.ru/", "https://google.com/", "https://rbc.ru"}

	findIt := findLinks(findText, arrLinks)

	fmt.Println(findIt)

}
