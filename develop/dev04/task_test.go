package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	expected := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	result := findAnagramSets(&words)

	// Проверяем соответствие ожидаемого результата и полученного результата
	if !reflect.DeepEqual(*result, expected) {
		t.Errorf("Expected %v, but got %v", expected, *result)
	}
}
