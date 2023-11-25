package main

import (
	"testing"
	"time"

	"github.com/beevik/ntp"
	"github.com/stretchr/testify/assert"
)

func TestNTPTime(t *testing.T) {
	expectedLocation, _ := time.LoadLocation("UTC")                               // Ожидаемое местоположение времени
	expectedTime := time.Date(2023, time.March, 25, 8, 0, 0, 0, expectedLocation) // Ожидаемое точное время

	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		t.Fatalf("Ошибка получения времени NTP: %v", err)
	}

	// Проверка, что текущее время близко к точному времени (с погрешностью около 1 секунды)
	assert.InDelta(t, expectedTime.Unix(), ntpTime.Unix(), 1, "Точное время должно быть близко к текущему времени")
}
