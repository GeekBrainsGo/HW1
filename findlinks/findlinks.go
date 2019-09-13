// Package findlinks implement function that search query in given links.
package findlinks

/*
	Basics Go.
	Rishat Ishbulatov, dated Sep 11, 2019.

	Write a function that will receive an input string with a search
	query (string) and an array of links to pages that should be searched
	([]string). The result of the function should be an array of strings
	with links to the pages on which the search query was found. The function
	should look for an exact match to the phrase in the response text from the
	server for each of the links.
*/

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

// FindLinks return links to pages on which the search query was found.
func FindLinks(s string, in []string) []string {
	buffer, out := new(bytes.Buffer), make([]string, 0, len(in))
	for _, v := range in {
		resp, err := http.Get(v)
		if err != nil {
			log.Fatalf("fetch: %v\n", err)
		}
		buffer.ReadFrom(resp.Body)
		resp.Body.Close()
		if strings.Contains(buffer.String(), s) {
			out = append(out, v)
		}
		buffer.Reset()
	}
	return out
}
