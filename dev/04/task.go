package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
)

/*
Поиск анаграмм по словарю

Написать функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
1. Входные данные для функции: ссылка на массив, каждый
элемент которого - слово на русском языке в кодировке
utf8
2. Выходные данные: ссылка на мапу множеств анаграмм
3. Ключ - первое встретившееся в словаре слово из
множества. Значение - ссылка на массив, каждый элемент
которого,
слово из множества.
4. Массив должен быть отсортирован по возрастанию.
5. Множества из одного элемента не должны попасть в
результат.
6. Все слова должны быть приведены к нижнему регистру.
7. В результате каждое слово должно встречаться только один
раз.
*/

func getByteSliceHash(bytes []byte) uint32 {
	hash := fnv.New32a()
	hash.Write(bytes)
	return hash.Sum32()
}

// parse and sort string to bytes
func strToBytes(word string) []byte {
	bytes := []byte(word)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return bytes
}

func getStringsLowercase(words []string) []string {
	for i, v := range words {
		words[i] = strings.ToLower(v)
	}
	return words
}

func getMapHashStrings(words []string) map[uint32][]string {
	m := make(map[uint32][]string)
	for _, v := range words {
		hash := getByteSliceHash(strToBytes(v))
		_, ok := m[hash]
		if !ok {
			m[hash] = make([]string, 0)
		}
		m[hash] = append(m[hash], v)
	}
	return m
}

func getConvertedMap(hashMap map[uint32][]string) map[string][]string {
	res := make(map[string][]string, len(hashMap))
	for _, v := range hashMap {
		if len(v) <= 1 {
			continue
		}
		t := getStringsLowercase(v)
		h := t[0] // ключ
		t = t[1:] // значение
		sort.Strings(t)
		res[h] = t
	}
	return res
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "дым"}
	fmt.Println(getConvertedMap(getMapHashStrings(words)))
}
