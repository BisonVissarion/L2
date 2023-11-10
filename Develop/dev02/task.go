package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func unpackString(s string) (string, error) {
	var result strings.Builder // Используем strings.Builder для эффективной конкатенации строк
	var escape bool

	for _, r := range s {
		if escape {
			// Если включен режим escape, добавляем символ независимо от его значения
			result.WriteRune(r)
			escape = false
		} else {
			if r >= '0' && r <= '9' {
				// Если символ - цифра, проверяем следующий символ
				nextIdx := result.Len()
				nextChar := ' '
				if nextIdx < len(s) {
					nextChar = rune(s[nextIdx])
				}
				if nextChar == '\\' {
					// Если следующий символ - обратный слеш, выключаем escape
					escape = true
				} else if nextChar >= '0' && nextChar <= '9' {
					// Если следующий символ тоже цифра, это некорректная строка
					return "", errors.New("некорректная строка")
				} else {
					// Иначе, распаковываем символ заданное количество раз
					count, err := strconv.Atoi(string(r))
					if err != nil {
						return "", errors.New("некорректная строка")
					}
					result.WriteString(strings.Repeat(string(nextChar), count))
				}
			} else {
				// Если символ не цифра, добавляем его к результату
				result.WriteRune(r)
			}
		}
	}

	return result.String(), nil
}

func main() {
	testCases := []struct {
		input  string
		output string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"qwe\\4\\5", "qwe45"},
		{"qwe", "qwerty"},
		{"qwe\\\\5", "qwe\\\\\\\\"},
		{"45", ""},
	}

	for _, testCase := range testCases {
		result, err := unpackString(testCase.input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else if result == testCase.output {
			fmt.Printf("Вход: %s => Выход: %s (Ожидаемый результат)\n", testCase.input, result)
		} else {
			fmt.Printf("Вход: %s => Выход: %s (Ожидался %s)\n", testCase.input, result, testCase.output)
		}
	}
}
