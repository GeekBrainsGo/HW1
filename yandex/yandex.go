// Package yandex implement Yandex.Disk api interactions.
package yandex

/*
	Basics Go.
	Rishat Ishbulatov, dated Sep 13, 2019.

	Write a function that receives a public link to the file from
	Yandex.Disk as an input and saves the received file to the user's disk.
*/

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var filename string

const yandex = "https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key="

// Answer stands for yandex api link object.
type Answer struct {
	Link      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run yandex.go <yandex pablic link>")
	}
	if err := Disk(os.Args[1]); err != nil {
		log.Fatalf("disk: %v\n", err)
	}
}

// Disk downloads file from Yandex.Disk with public link to user's disk.
func Disk(s string) error {
	resp, err := http.Get(yandex + url.QueryEscape(s))
	buffer := new(bytes.Buffer)
	if err != nil {
		log.Fatalf("query fetch: %v\n", err)
	}
	buffer.ReadFrom(resp.Body)
	resp.Body.Close()

	answer := Answer{}
	if err := json.Unmarshal(buffer.Bytes(), &answer); err != nil {
		log.Fatalf("json parse: %v\n", err)
	}

	v, err := url.ParseQuery(answer.Link)
	if err != nil {
		log.Fatalf("parse query: %v\n", err)
	}
	filename = v.Get("filename")

	resp, err = http.Get(answer.Link)
	if err != nil {
		log.Fatalf("file fetch: %v\n", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os create: %v\n", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
