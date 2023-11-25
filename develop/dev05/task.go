package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	// Определение флагов командной строки
	after := flag.Int("A", 0, "Print N lines after the match")
	before := flag.Int("B", 0, "Print N lines before the match")
	context := flag.Int("C", 0, "Print N lines before and after the match")
	count := flag.Bool("c", false, "Count the number of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the matching results")
	fixed := flag.Bool("F", false, "Fixed string matching")
	lineNum := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	// Чтение аргументов командной строки
	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	// Вызов функции фильтрации для каждого файла
	for _, file := range files {
		matchLines := filterFile(file, pattern, *after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)
		printMatchLines(matchLines)
	}
}

// Функция фильтрации файла
func filterFile(file, pattern string, after, before, context int, count, ignoreCase, invert, fixed, lineNum bool) []string {
	matchLines := make([]string, 0)

	// Открытие файла для чтения
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Cannot open file: %s\n", err)
		return matchLines
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	prevLines := make([]string, 0)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		match := false

		if fixed {
			match = strings.Contains(line, pattern)
		} else {
			if ignoreCase {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}

		if match && !invert {
			if count {
				// Если указан флаг -c, увеличиваем счетчик
				matchLines = append(matchLines, line)
			} else {
				// Сохраняем строки до совпадения
				if before > 0 {
					for _, prevLine := range prevLines {
						matchLines = append(matchLines, prevLine)
					}
				}

				// Сохраняем текущую строку
				if lineNum {
					line = fmt.Sprintf("%d:%s", lineNumber, line)
				}
				matchLines = append(matchLines, line)

				// Сохраняем строки после совпадения
				if after > 0 {
					afterLines := make([]string, 0)
					for i := 0; i < after; i++ {
						if scanner.Scan() {
							afterLine := scanner.Text()
							matchLines = append(matchLines, afterLine)
							afterLines = append(afterLines, afterLine)
						} else {
							break
						}
					}
					prevLines = afterLines
				} else {
					prevLines = make([]string, 0)
				}

				// Сохраняем контекстные строки
				if context > 0 {
					contextLines := make([]string, 0)
					for i := 0; i < context; i++ {
						if scanner.Scan() {
							contextLine := scanner.Text()
							matchLines = append(matchLines, contextLine)
							contextLines = append(contextLines, contextLine)
						} else {
							break
						}
					}
					prevLines = contextLines
				}
			}

			found = true
		} else if !match && invert {
			if count {
				// Если указан флаг -c, увеличиваем счетчик
				matchLines = append(matchLines, line)
			} else {
				// Сохраняем строки до и после совпадения
				if (before > 0 || after > 0) && (prevLines != nil || after > 0) {
					if lineNum {
						line = fmt.Sprintf("%d:%s", lineNumber, line)
					}
					matchLines = append(matchLines, line)
				}

				// Сохраняем контекстные строки
				if context > 0 {
					contextLines := make([]string, 0)
					for i := 0; i < context; i++ {
						if scanner.Scan() {
							contextLine := scanner.Text()
							matchLines = append(matchLines, contextLine)
							contextLines = append(contextLines, contextLine)
						} else {
							break
						}
					}
					prevLines = contextLines
				} else {
					prevLines = make([]string, 0)
				}
			}

			found = true
		} else {
			// Сохраняем строки до и после совпадения
			if (before > 0 || after > 0) && (prevLines != nil || after > 0) {
				prevLines = append(prevLines, line)

				if len(prevLines) > before+after {
					prevLines = prevLines[1:]
				}
			} else if context > 0 {
				prevLines = append(prevLines, line)

				if len(prevLines) > context {
					prevLines = prevLines[1:]
				}
			}
		}
	}

	if count && !found {
		matchLines = []string{"0"}
	}

	return matchLines
}

// Функция вывода совпадающих строк
func printMatchLines(matchLines []string) {
	for _, line := range matchLines {
		fmt.Println(line)
	}
}
