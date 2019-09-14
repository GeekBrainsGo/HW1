package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Link struct {
	Href      string `json:"href"`      // URL. Может быть шаблонизирован, см. ключ templated.
	Method    string `json:"method"`    // HTTP-метод для запроса URL из ключа href.
	Templated bool   `json:"templated"` // Признак URL, который был шаблонизирован согласно RFC 6570. Возможные значения:
	// «true» — URL шаблонизирован: прежде чем отправлять запрос на этот адрес, следует указать нужные значения параметров вместо значений в фигурных скобках.
	// «false» — URL может быть запрошен без изменений.
}

func yandexDownload(linkURL, filename string) {

	// https://yandex.ru/dev/disk/api/reference/content-docpage/

	resp, err := http.Get("https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key=" + linkURL)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)

	link := &Link{}
	err = json.Unmarshal(body, link) // заполним структуру для получения файла
	if err != nil {
		log.Println(err)
		return
	}

	resp, err = http.Get(link.Href) // получим содержимое ссылки
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
	}

	// сохраним содержимое ссылки в файл или выведем в стандартый вывод
	outFile := os.Stdout
	if filename != "" {
		if outFile, err = os.Create(filename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}
	writer := bufio.NewWriter(outFile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()
	io.Copy(writer, resp.Body)

	return
}

func main() {
	yandexDownload("https://yadi.sk/d/BR5U66ex9rHqXg", "./mongoDB.go")
}
