package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
)

func main() {
	var strSearch string
	arr := make([]string, 0, 1)
	arr = append(arr, "http://www.mail.ru")
	arr = append(arr, "http://www.google.com")
	arr = append(arr, "http://www.yandex.ru")
	arr = append(arr, "http://www.ibm.ru")

	strSearch = "Поиск"
	ret := multiSearch(arr, strSearch)

	fmt.Println("Ищем:", strSearch)
	fmt.Println(ret)
}

// multiSearch - функция поиска заданой строки на ресурсах представленных в массиве URL
func multiSearch(arrLink []string, pattern string) []string {
	arrReturn := make([]string, 0, 1) // массив для результат работы функции
	wg := &sync.WaitGroup{}           // используем для контроля окончания всех горутин
	mu := &sync.Mutex{}               // используем для корректного обновления map-ы
	arrRet := map[string]bool{}       // map-а для сохранения результатов поиска

	for _, item := range arrLink {
		wg.Add(1) // добавим запуск новой горутины в группу ожидания

		go func(hRef string, pattern string, arr map[string]bool, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done() // перед выходом ОБЯЗАТЕЛЬНО укажем, что горутина закончилась

			resp, err := http.Get(hRef) // получим страницу
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close() // не забыть закрыть тело ответа на запрос

			if resp.StatusCode != http.StatusOK {
				log.Println(resp.Status)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}
			matched, err := regexp.MatchString(pattern, string(body))
			if err != nil {
				log.Println(err)
				return
			}

			mu.Lock()
			arr[hRef] = matched
			mu.Unlock()
		}(item, pattern, arrRet, wg, mu)
	}
	wg.Wait() // ожидаем окончания работы всех горутин

	// сформируем выходной массив
	for key, item := range arrRet {
		if item {
			arrReturn = append(arrReturn, key)
		}
	}
	return arrReturn
}
