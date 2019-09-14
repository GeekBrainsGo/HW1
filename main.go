package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)


type yaFile struct {
	Href      string
	Method    string
	Templated bool
}

func main() {
	var urls []string
	var search string
	fmt.Println("Введите массив ссылок, для окончания введите \"все\"")
	for i := 1; ; i++ {
		var url string
		fmt.Print(strconv.Itoa(i) + ": ")
		_, err := fmt.Scanln(&url)
		if err != nil {
			log.Panic(err)
		}
		if strings.Compare(url, "все") == 0 {
			break
		} else {
			urls = append(urls, url)
		}
	}

	fmt.Print("Введите строку для поиска: ")
	_, err := fmt.Scanln(&search)
	if err != nil {
		log.Panic(err)
	}

	search = strings.ToLower(search)

	var matchedUrls []string
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Panic(err)
		}
		cont, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Panic(err)
		}
		if strings.Contains(strings.ToLower(string(cont)), search) {
			matchedUrls = append(matchedUrls, url)
		}
	}
	if len(matchedUrls) < 1 {
		fmt.Println("Искомое слово не нашлось ни по одной из ссылок")
	} else {
		fmt.Println("Искомое слово нашлось на:")
		for i, url := range matchedUrls {
			fmt.Println(strconv.Itoa(i+1) + ":" + url)
		}
	}

}

func yaDownload(url string) {
	var yf yaFile
	resp, err := http.Get("https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key=" + url)
	if err != nil {
		log.Panic(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
		return
	}
	err = json.Unmarshal(body, &yf)
	if err != nil {
		log.Panic(err)
		return
	}

	linkFile, err := http.Get(yf.Href)
	if err != nil {
		log.Panic(err)
		return
	}
	defer linkFile.Body.Close()

	dataFile, err := ioutil.ReadAll(linkFile.Body)
	if err != nil {
		log.Panic(err)
		return
	}

	parsedName, err := url.Parse(yf.Href)
	if err != nil {
		log.Panic(err)
		return
	}

	parsedKeys, err := url.ParseQuery(parsedName.RawQuery)
	if err != nil {
		log.Panic(err)
		return
	}

	name := parsedKeys["filename"][0]

	err = ioutil.WriteFile(name, dataFile, 0777)

}
}