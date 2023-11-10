package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Выводим приглашение
	fmt.Print("MyShell> ")

	// Создаем сканер для чтения пользовательского ввода
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Читаем введенную команду
		command := scanner.Text()

		// Разбиваем введенную команду на аргументы
		args := strings.Fields(command)

		// Обработка команды
		switch args[0] {
		case "cd":
			if len(args) > 1 {
				// Смена директории
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case "pwd":
			// Показать текущую директорию
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			// Вывод аргумента
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			}
		case "kill":
			// Убить процесс
			if len(args) > 1 {
				cmd := exec.Command("kill", args[1])
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case "ps":
			// Вывод информации о запущенных процессах
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "exit", "quit":
			// Завершение работы
			os.Exit(0)
		default:
			// Запуск произвольной команды
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		// Вывод нового приглашения
		fmt.Print("MyShell> ")
	}
}
