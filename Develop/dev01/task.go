package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем текущее точное время с NTP-сервера
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка при получении времени: %s\n", err)
		os.Exit(1)
	}

	// Выводим текущее точное время
	fmt.Printf("Текущее точное время: %s\n", ntpTime.Format(time.RFC3339))
}
