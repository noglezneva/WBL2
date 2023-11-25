package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Флаги командной строки
	filePath := flag.String("file", "", "путь к файлу для сортировки")
	column := flag.Int("k", 0, "номер колонки для сортировки (по умолчанию 0, разделитель - пробел)")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	monthSort := flag.Bool("M", false, "сортировать по названию месяца")
	ignoreTrailingSpace := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "проверять отсортированы ли данные")
	numericSuffixSort := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Не указан путь к файлу")
	}

	fileContent, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fileContent), "\n")

	sort.Slice(lines, func(i, j int) bool {
		lineA := lines[i]
		lineB := lines[j]

		if *ignoreTrailingSpace {
			lineA = strings.TrimRight(lineA, " ")
			lineB = strings.TrimRight(lineB, " ")
		}

		if *column > 0 && *column <= len(strings.Fields(lineA)) && *column <= len(strings.Fields(lineB)) {
			// Обрабатываем случай с указанием колонки
			fieldA := strings.Fields(lineA)[*column-1]
			fieldB := strings.Fields(lineB)[*column-1]

			if *numericSuffixSort {
				// Извлекаем числовое значение и суффикс из поля
				numA, suffA := extractNumericSuffix(fieldA)
				numB, suffB := extractNumericSuffix(fieldB)

				// Сортируем по числовому значению и сравниваем суффиксы
				if numA != numB {
					return numA < numB
				}
				return suffA < suffB
			}

			if *numeric {
				// Сортировка по числовому значению
				numA, errA := strconv.Atoi(fieldA)
				numB, errB := strconv.Atoi(fieldB)

				if errA == nil && errB == nil {
					return numA < numB
				}
			}

			// Сортировка по строковому значению колонки
			return fieldA < fieldB
		}

		if *monthSort {
			// Обрабатываем сортировку по названию месяца
			dateA, errA := time.Parse("Jan", lineA)
			dateB, errB := time.Parse("Jan", lineB)

			if errA == nil && errB == nil {
				return dateA.Before(dateB)
			}
		}

		// Сортировка по всей строке
		return lineA < lineB
	})

	if *reverse {
		reverseLines(lines)
	}

	if *unique {
		lines = removeDuplicates(lines)
	}

	output := strings.Join(lines, "\n")

	// Проверка отсортированности данных
	if *checkSorted && isSorted(lines) {
		fmt.Println("Данные отсортированы")
	}

	// Вывод отсортированных данных в файл
	outputFilePath := *filePath + ".sorted"
	err = ioutil.WriteFile(outputFilePath, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Отсортированные данные сохранены в файл: %s\n", outputFilePath)
}

// Извлекает числовое значение и суффикс из строки
func extractNumericSuffix(s string) (int, string) {
	numericPart := ""
	suffixPart := ""

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			numericPart = string(s[i]) + numericPart
		} else {
			suffixPart = string(s[i]) + suffixPart
		}
	}

	num, _ := strconv.Atoi(numericPart)

	return num, suffixPart
}

// Разворачивает порядок строк
func reverseLines(lines []string) {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// Удаляет повторяющиеся строки из среза
func removeDuplicates(lines []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, line := range lines {
		if !encountered[line] {
			encountered[line] = true
			result = append(result, line)
		}
	}

	return result
}

// Проверяет, отсортирован ли срез строк
func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if lines[i-1] > lines[i] {
			return false
		}
	}
	return true
}
