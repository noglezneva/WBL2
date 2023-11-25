package main

import (
	"bufio"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	// Запуск программы в отдельном процессе
	cmd := exec.Command("go", "run", "main.go", "--timeout=5s", "localhost", "12345")
	// Перехват вывода программы
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		t.Fatalf("Не удалось запустить программу: %s", err.Error())
	}
	defer cmd.Process.Kill()

	// Ждем, пока программа полностью запустится
	time.Sleep(1 * time.Second)

	// Проверка подключения к серверу
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		t.Fatalf("Ошибка подключения к серверу: %s", err.Error())
	}
	defer conn.Close()

	// Отправка данных в сокет
	data := "Hello, server!\n"
	stdin.Write([]byte(data))

	// Получение данных из сокета
	reader := bufio.NewReader(stdout)
	response, _ := reader.ReadString('\n')

	// Проверка полученного ответа
	expectedResponse := "Received: Hello, server!\n"
	if response != expectedResponse {
		t.Errorf("Некорректный ответ от сервера. Ожидаемый: %s, полученный: %s", expectedResponse, response)
	}

	// Отправка сигнала завершения программы (Ctrl+D)
	stdin.Close()

	// Ожидание завершения программы
	cmd.Wait()
}

func TestMain(m *testing.M) {
	// Запуск тестов и получение кода завершения
	exitCode := m.Run()
	os.Exit(exitCode)
}
