/*
 * HomeWork-1: Yandex file
 * Created on 12.09.19 19:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const reqURL = "https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key="

// YandexFile is the base struct for downloading files by public URL.
type yandexFile struct {
	Href      string
	Method    string
	Templated bool
}

func main() {

	// get user's url or use examples
	//fileURL := "https://yadi.sk/i/pBfU5WBqFWO0FA" // docx
	//fileURL := "https://yadi.sk/d/0JhGPmrfvgSHEw" // jpg

	fileURL := ""
	for {
		fmt.Printf("Enter Yandex share URL (Ctrl-C for exit): ")
		_, err := fmt.Scanln(&fileURL)
		if err != nil {
			log.Println("error parse URL", err)
			continue
		}

		fileName, err := getFileFromURL(fileURL)
		if err != nil {
			log.Printf("Error while downloading file from URL:\n%s \n%v\n", fileURL, err)
			continue
		}

		fmt.Println("File", fileName, "saved successfully.")
	}
}

func getFileFromURL(URL string) (fileName string, err error) {

	// get file metadata
	resp, err := http.Get(reqURL + URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("file not found")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// unmarshal metadata - get direct link
	yf := &yandexFile{}

	err = json.Unmarshal([]byte(body), yf)
	if err != nil {
		return
	}

	// get file body
	respFile, err := http.Get(yf.Href)
	if err != nil {
		return
	}
	defer respFile.Body.Close()

	bodyFile, err := ioutil.ReadAll(respFile.Body)
	if err != nil {
		return
	}

	// parse name from url
	u, err := url.Parse(yf.Href)
	if err != nil {
		return
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}
	fileName = m["filename"][0]

	// save file
	err = ioutil.WriteFile(fileName, bodyFile, 0644)

	return
}
