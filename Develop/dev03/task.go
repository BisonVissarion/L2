package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов
	column := flag.Int("k", 0, "Колонка для сортировки (по умолчанию 0 - разделение по пробелу)")
	numeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	// Открытие входного файла
	inputFileName := flag.Arg(0)
	if inputFileName == "" {
		fmt.Println("Укажите входной файл")
		return
	}
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer inputFile.Close()

	// Считывание строк из файла
	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Функция сравнения для сортировки
	compare := func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		// Разделение на колонки
		aColumns := strings.Fields(a)
		bColumns := strings.Fields(b)

		// Выбор колонки для сравнения
		aValue := ""
		bValue := ""

		if *column < len(aColumns) {
			aValue = aColumns[*column]
		}

		if *column < len(bColumns) {
			bValue = bColumns[*column]
		}

		// Преобразование к числу, если нужно
		if *numeric {
			aInt, errA := strconv.Atoi(aValue)
			bInt, errB := strconv.Atoi(bValue)

			if errA == nil && errB == nil {
				aValue = strconv.Itoa(aInt)
				bValue = strconv.Itoa(bInt)
			}

		}

		// Сравнение
		if *reverse {
			return aValue > bValue
		}
		return aValue < bValue
	}

	// Сортировка
	sort.SliceStable(lines, func(i, j int) bool {
		return compare(i, j)
	})

	// Удаление повторяющихся строк
	if *unique {
		var uniqueLines []string
		seen := make(map[string]struct{})
		for _, line := range lines {
			if _, ok := seen[line]; !ok {
				seen[line] = struct{}{}
				uniqueLines = append(uniqueLines, line)
			}
		}
		lines = uniqueLines
	}

	// Вывод отсортированных данных
	for _, line := range lines {
		fmt.Println(line)
	}
}
