package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	// Бесконечный цикл для интерактивного сеанса
	for {
		// Выводим приглашение для ввода команды
		fmt.Print("MyShell> ")

		// Считываем ввод пользователя с консоли
		input := readInput()

		// Проверяем наличие команды выхода
		if input == "exit" {
			break
		}

		// Разбиваем ввод на команды по пайпу
		commands := strings.Split(input, "|")

		// Запускаем каждую команду в цепочке пайпов
		var output []byte
		var err error
		for _, cmdStr := range commands {
			// Убираем лишние пробелы
			cmdStr = strings.TrimSpace(cmdStr)

			// Разбиваем команду на аргументы
			args := strings.Fields(cmdStr)

			// Проверяем и выполняем встроенные команды
			switch args[0] {
			case "cd":
				err := changeDirectory(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				}
			case "pwd":
				dir, err := printWorkingDirectory()
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Printf("Текущая директория: %s\n", dir)
				}
			case "echo":
				result, err := echo(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Printf("%s\n", result)
				}
			case "kill":
				err := killProcess(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				}
			case "ps":
				processes, err := listProcesses()
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Printf("Процессы:\n%s\n", processes)
				}
			default:
				// Выполняем внешнюю команду
				output, err = executeExternalCommand(args)
			}

			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				break
			}
		}

		// Выводим результат выполнения цепочки команд
		if len(output) > 0 {
			fmt.Printf("%s\n", string(output))
		}
	}
}

// Функция для считывания ввода пользователя
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Функция для смены директории
func changeDirectory(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Использование: cd <директория>")
	}

	dir := args[1]
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	return nil
}

// Функция для вывода текущей директории
func printWorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Функция для команды echo
func echo(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("Использование: echo <текст>")
	}
	return strings.Join(args[1:], " "), nil
}

// Функция для команды kill
func killProcess(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Использование: kill <PID>")
	}

	pidStr := args[1]
	pid := strings.TrimSpace(pidStr)

	cmd := exec.Command("kill", "-9", pid)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Функция для команды ps
func listProcesses() (string, error) {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Функция для выполнения внешней команды
func executeExternalCommand(args []string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
