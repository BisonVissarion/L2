package main

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	// Пройдемся по каждому слову в словаре
	for _, word := range words {
		// Приведем слово к нижнему регистру
		word = strings.ToLower(word)

		// Преобразуем слово в сортированный набор символов (анаграмму)
		sortedWord := sortString(word)

		// Добавим слово в множество анаграмм для соответствующей анаграммы
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	// Удалим множества из одного элемента
	for key, value := range anagramMap {
		if len(value) == 1 {
			delete(anagramMap, key)
		}
	}

	return anagramMap
}

// Функция для сортировки символов в слове
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "рыба", "бары", "абры"}

	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		// Сортируем множество анаграмм по возрастанию
		sort.Strings(value)
		// Выводим ключ и множество анаграмм
		println(key, value)
	}
}
