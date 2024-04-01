package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
Базовая задача

Создать программу печатающую точное время с использованием
NTP -библиотеки. Инициализировать как go module. Использовать
библиотеку github.com/beevik/ntp. Написать программу
печатающую текущее время / точное время с использованием этой
библиотеки.

Требования:
1. Программа должна быть оформлена как go module
2. Программа должна корректно обрабатывать ошибки
библиотеки: выводить их в STDERR и возвращать ненулевой
код выхода в OS
*/

const hostURL = "0.beevik-ntp.pool.ntp.org"

func getHostTime(host string) (time.Time, error) {
	hostTime, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}
	return hostTime, nil
}

func main() {
	curTime, err := getHostTime(hostURL)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}

	fmt.Print(curTime)
}
