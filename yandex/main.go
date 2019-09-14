package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type YandexResoureDownload struct {
	Href       string `json:"href"`
	Method     string `json:"method"`
	Tempalated bool   `json:"templated"`
}

const apiURL = "https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key="

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}

}

func main() {

	downloadURL := "https://yadi.sk/d/mp9UiYd7eG-JZg" //default link

	flag.StringVar(&downloadURL, "d", downloadURL, "URL to download")
	flag.Parse()

	fmt.Println("Download from: ", downloadURL)

	resp, err := http.Get(apiURL + downloadURL)
	failOnError(err, "Get download href fail")
	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)
	failOnError(err, "Read error")

	var yaResoureData YandexResoureDownload
	err = json.Unmarshal(jsonData, &yaResoureData)
	failOnError(err, "YandexResoureDownload Unmarshal error")

	hrefParams, err := url.ParseQuery(yaResoureData.Href)
	failOnError(err, "Href parse error")

	filename := hrefParams.Get("filename")
	fileLength := hrefParams.Get("fsize")

	download, err := http.Get(yaResoureData.Href)
	failOnError(err, "File download error")
	defer download.Body.Close()

	file, err := os.Create(filename)
	failOnError(err, "File open error")
	defer file.Close()

	bytesCopied, err := io.Copy(file, download.Body)
	failOnError(err, "File save error")

	if strconv.FormatInt(bytesCopied, 10) != fileLength {
		fmt.Println("File length error")
	}

	fmt.Println("Saved on: ", filename)

}
