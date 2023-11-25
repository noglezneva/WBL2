package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCutUtility(t *testing.T) {
	// Создание временного файла для тестовых данных
	tempFile, err := ioutil.TempFile("", "test_data.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Запись тестовых данных в файл
	data := "1\tabc\n2\txyz\n3\tdef"
	err = ioutil.WriteFile(tempFile.Name(), []byte(data), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Запуск программы с тестовыми аргументами
	os.Args = []string{"cmd", "-f", "2", "-d", "\t"}
	outputBuffer := &bytes.Buffer{}
	errorBuffer := &bytes.Buffer{}
	log.SetOutput(errorBuffer)
	oldStdout := os.Stdout
	os.Stdout = outputBuffer
	main()
	os.Stdout = oldStdout

	// Проверка вывода программы
	expectedOutput := "abc\nxyz\ndef\n"
	assert.Equal(t, expectedOutput, outputBuffer.String())

	// Проверка отсутствия ошибок
	assert.Empty(t, errorBuffer.String())

	// Проверка работы флага -s
	os.Args = []string{"cmd", "-f", "2", "-d", "\t", "-s=false"}
	outputBuffer.Reset()
	errorBuffer.Reset()
	log.SetOutput(errorBuffer)
	os.Stdout = outputBuffer
	main()
	os.Stdout = oldStdout

	// Проверка вывода программы с флагом -s=false
	expectedOutput = "abc\nxyz\ndef\n"
	assert.Equal(t, expectedOutput, outputBuffer.String())

	// Проверка отсутствия ошибок при флаге -s=false
	assert.Empty(t, errorBuffer.String())

	// Проверка обработки ошибки при неверном списке полей
	os.Args = []string{"cmd", "-f", "2,a", "-d", "\t"}
	outputBuffer.Reset()
	errorBuffer.Reset()
	log.SetOutput(errorBuffer)
	os.Stdout = outputBuffer
	main()
	os.Stdout = oldStdout

	// Проверка вывода ошибки при неверном списке полей
	expectedError := "Bad field list\n"
	assert.Equal(t, expectedError, errorBuffer.String())
}
