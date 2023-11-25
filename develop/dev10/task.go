package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
func main() {
	// Парсинг аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут на подключение к серверу")
	flag.Parse()

	// Получение аргументов - адрес хоста и порт
	host := flag.Arg(0)
	port := flag.Arg(1)

	// Установка обработчика сигнала для завершения программы при нажатии Ctrl+C или Ctrl+D
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh
		fmt.Println("Программа завершена")
		os.Exit(0)
	}()

	// Подключение к серверу
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения к серверу: %s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключено к", conn.RemoteAddr())

	// Чтение данных из STDIN и запись в сокет, чтение данных из сокета и вывод в STDOUT
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Bytes()
			_, err := conn.Write(data)
			if err != nil {
				fmt.Printf("Ошибка записи в сокет: %s\n", err.Error())
				break
			}

			response := make([]byte, 1024)
			_, err = conn.Read(response)
			if err != nil {
				fmt.Printf("Ошибка чтения из сокета: %s\n", err.Error())
				break
			}

			fmt.Println(string(response))
		}
	}()

	// Ожидание завершения программы при закрытии со стороны сервера
	<-sigCh
	fmt.Println("Соединение разорвано")
}
