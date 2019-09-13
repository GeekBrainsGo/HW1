/*
 * HomeWork-1: Search string
 * Created on 11.09.19 22:41
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

const sitesFile = "sites.txt"

func main() {

	// get URLs from file
	file, err := os.Open(sitesFile)
	if err != nil {
		log.Fatalln("Can't open file with sites:", sitesFile, err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading file body:", file, err)
	}

	urls := strings.Split(string(b), "\n")

	//examples: "Бим", "Книга", "1973", "2033", "bug"
	search := ""
	for {
		fmt.Printf("Enter search URL (Ctrl-C for exit): ")
		_, err := fmt.Scanln(&search)
		if err != nil {
			log.Println("error parse search string", err)
			continue
		}

		fmt.Printf("Found string '%s' in sites:\n", search)
		found := searchStringURL(search, urls)

		for _, f := range found {
			fmt.Println(f)
		}
	}
}

func searchStringURL(search string, urls []string) (res []string) {

	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}

	for _, url := range urls {
		if len(url) < 3 { // no fake strings
			continue
		}

		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error getting url: %v", err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading body: %v", err)
				return
			}

			if strings.Contains(string(body), search) {
				mux.Lock()
				res = append(res, url)
				mux.Unlock()
			}
		}(url)
	}

	wg.Wait()
	return
}
