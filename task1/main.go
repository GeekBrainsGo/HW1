package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func scanLinks(query string, links []string) (resultUrls []string, err error) {
	for _, link := range links {
		// Get запрос по адресам в массиве ссылок
		resp, err := http.Get(link)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		// Читаем полученный body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(body))
		defer resp.Body.Close()

		// Проверяем содержится ли наш запрос в полученном body
		if strings.Contains(string(body), query) {
			resultUrls = append(resultUrls, link)
		}
	}
	return resultUrls, err
}

func main() {
	links := []string{
		"https://yandex.ru/",
		"https://meduza.io/",
		"https://golang.org/",
		"https://rambler.ru/",
		"https://rbc.ru",
	}
	query := "РБК"

	result, err := scanLinks(query, links)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(strings.Join(result, "\n"))
}
