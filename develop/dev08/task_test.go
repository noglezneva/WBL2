package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestShell(t *testing.T) {
	// Создаем пайпы для ввода и вывода
	reader, writer := io.Pipe()
	oldStdin := os.Stdin
	os.Stdin = reader

	var output bytes.Buffer
	oldStdout := os.Stdout
	os.Stdout = &output

	// Восстанавливаем входной и вывод после выполнения теста
	defer func() {
		writer.Close()
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// Устанавливаем входной поток с помощью строкового ридера
	input := "echo Hello, World!\n"
	io.WriteString(writer, input)
	writer.Close()

	// Запускаем шелл
	main()

	// Проверяем ожидаемый вывод
	expectedOutput := "Hello, World!"
	actualOutput := strings.TrimSpace(output.String())
	if actualOutput != expectedOutput {
		t.Errorf("Ожидаемый вывод: %s, полученный вывод: %s", expectedOutput, actualOutput)
	}
}
