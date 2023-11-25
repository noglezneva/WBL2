package main

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://www.example.com"
	filename := "index.html"

	err := download(url, filename)
	if err != nil {
		t.Errorf("Ошибка при скачивании файла: %s", err.Error())
	}

	// Проверяем, что файл был создан
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("Файл не был создан")
	}

	// Удаляем файл после выполнения теста
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Ошибка при удалении файла: %s", err.Error())
	}
}
