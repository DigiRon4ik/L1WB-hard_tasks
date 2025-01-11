package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

// TimeFetcher - interface for getting time.
type TimeFetcher interface {
	FetchTime(server string) (time.Time, error)
}

// NTPTimeFetcher - TimeFetcher implementation for working with a real NTP server.
type NTPTimeFetcher struct{}

// FetchTime gets the exact time from a real NTP server.
func (n NTPTimeFetcher) FetchTime(server string) (time.Time, error) {
	return ntp.Time(server)
}

// GetExactTime gets the exact time via TimeFetcher.
func GetExactTime(fetcher TimeFetcher, server string) (string, error) {
	exactTime, err := fetcher.FetchTime(server)
	if err != nil {
		return "", err
	}
	return exactTime.Format("2006-01-02 15:04:05 MST"), nil
}

func main() {
	const ntpServer = "time.google.com"

	// Using NTPTimeFetcher to get time
	fetcher := NTPTimeFetcher{}

	// Getting the exact time
	timeStr, err := GetExactTime(fetcher, ntpServer)
	if err != nil {
		log.Printf("Error getting exact time: %v\n", err)
		os.Exit(1)
	}

	// Time stamp
	fmt.Printf("Exact time: %s\n", timeStr)
}

/*
 - Output: -
Exact time: 2025-01-11 04:45:12 MSK
*/
