package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

// findAnagramSets finds all sets of anagrams.
func findAnagramSets(words *[]string) *map[string][]string {
	// Converts all words to lowercase.
	normalizedWords := make([]string, len(*words))
	for i, word := range *words {
		normalizedWords[i] = strings.ToLower(word)
	}

	// Storage for grouping anagrams.
	anagramGroups := make(map[string][]string)

	for _, word := range normalizedWords {
		// Transform the word into a sorted string of characters (anagram group key).
		sortedKey := sortString(word)
		anagramGroups[sortedKey] = append(anagramGroups[sortedKey], word)
	}

	// Formulate the result.
	result := make(map[string][]string)
	for _, group := range anagramGroups {
		if len(group) > 1 { // Exclude groups with one element.
			// Remove duplicates and sort the group alphabetically.
			sGroup := uniqueAndSort(group)
			// The first element in the group is used as the key.
			result[group[0]] = sGroup
		}
	}

	return &result
}

// sortString sorts the characters in a string alphabetically.
func sortString(s string) string {
	runes := []rune(s)
	sort.SliceStable(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// uniqueAndSort removes duplicates and sorts the array of strings.
func uniqueAndSort(words []string) []string {
	wordSet := make(map[string]struct{})
	uniqueWords := make([]string, 0, 2)

	for _, word := range words {
		if _, ok := wordSet[word]; !ok {
			wordSet[word] = struct{}{}
			uniqueWords = append(uniqueWords, word)
		}
	}

	sort.Strings(uniqueWords)
	return uniqueWords
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Слиток", "ПЯТАК", "салфетка", "графин", "нифгра", "нифгра", "ниФгра", "Кслито"}
	anagramSets := findAnagramSets(&words)

	for key, group := range *anagramSets {
		fmt.Printf("%s: %v\n", key, group)
	}
}

/*
 - Output: -
пятак: [пятак пятка тяпка]
листок: [кслито листок слиток столик]
графин: [графин нифгра]
*/
