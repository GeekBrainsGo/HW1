package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	filekey := flag.String("filekey", "https://yadi.sk/i/wLk_CIJUGNZcqw", "path or filekey on yadisk")
	filename := flag.String("filename", "./maincat.jpg", "path where file must be stored")
	flag.Parse()

	if err := SaveYaFile(*filekey, *filename); err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Println("Everything is awesome!")
}

// YaFileInfo - Информация о файле
type YaFileInfo struct {
	Name         string    `json:"name"`
	FileLink     string    `json:"file"`
	Size         int64     `json:"size"`
	MimeType     string    `json:"mime_type"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Error        *string   `json:"error"` //< Если указан указатель, при отсутствии поля в json будет nil
	ErrorMessage *string   `json:"message"`
}

// GetYaFileInfo - Получает информацию о файле
func GetYaFileInfo(filekey string) (*YaFileInfo, error) {
	const PUBLIC_FILE_LINK = "https://cloud-api.yandex.net/v1/disk/public/resources?public_key="

	filelink := fmt.Sprintf("%s%s", PUBLIC_FILE_LINK, filekey)

	// Получаем ответ от сервера
	finfoResp, err := http.Get(filelink)
	if err != nil {
		return nil, err
	}

	// Читаем информацию о файле
	finfoData, err := ioutil.ReadAll(finfoResp.Body)
	if err != nil {
		return nil, err
	}

	// Парсим в структуру
	finfo := &YaFileInfo{}
	if err := json.Unmarshal(finfoData, finfo); err != nil {
		return nil, err
	}

	return finfo, nil
}

// GetYaFile - Возвращает стрим файла
func GetYaFile(finfo *YaFileInfo) (io.ReadCloser, error) {
	// Проверяем, что Яндекс действительно нашёл файл
	if finfo.Error != nil {
		return nil, fmt.Errorf("yadisk error: %s: %s", finfo.Error, finfo.ErrorMessage)
	}

	// Получаем стрим на файл
	fresp, err := http.Get(finfo.FileLink)
	if err != nil {
		return nil, err
	}

	return fresp.Body, nil
}

// SaveYaFile - Скачивает файл
func SaveYaFile(filekey, filepath string) error {
	// Получаем информацию о файле
	finfo, err := GetYaFileInfo(filekey)
	if err != nil {
		return err
	}

	// Получаем стрим на файл
	fstream, err := GetYaFile(finfo)
	if err != nil {
		return err
	}
	defer fstream.Close()

	// Создаём локальный файл и получаем его стрим
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	// Копируем сетевой стрим в локальный
	if _, err := io.Copy(file, fstream); err != nil {
		return err
	}

	return nil
}
