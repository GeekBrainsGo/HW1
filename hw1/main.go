package main


import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

func main(){

	query := "автор"
	links := []string{"https://book24.ru", "https://avidreaders.ru/books/", "https://www.iherb.com"}

	foundByLinks := findQuery(query, links)

	fmt.Println(links)
	fmt.Println(foundByLinks)
}


func findQuery(q string, m1 []string) (m2 []string){
	for x := range m1{
		//fmt.Println(m1[x])
		if checkContainsQuery(q,m1[x]){
			m2 = append(m2, m1[x])
		}
	}
	return
}

func checkContainsQuery(query, url string) bool {
	res := false

	body := getBodyFromUrl(url)
	if strings.Contains(body, query){
		res = true
	}

	return res
}

func getBodyFromUrl(url string) string{

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n",err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
		os.Exit(1)
	}
	//fmt.Printf("%s", body)
	return string(body)

}