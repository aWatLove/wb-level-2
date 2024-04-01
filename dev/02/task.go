package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
Задача на распаковку

Создать Go-функцию, осуществляющую примитивную распаковку
строки, содержащую повторяющиеся символы/руны, например:
- "a4bc2d5e" => "aaaabccddddde"
- "abcd" => "abcd"
- "45" => "" (некорректная строка)
- "" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
- qwe\4\5 => qwe45 (*)
- qwe\45 => qwe44444 (*)
- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка, функция
должна возвращать ошибку. Написать unit-тесты.
*/

const escapeSymbol = 92

func unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", errors.New("incorrect string")
	}

	runes := []rune(s)
	var builder strings.Builder
	var printRune rune

	for i := 0; i < len(runes); i++ {

		if runes[i] == escapeSymbol { // если встретилась escape последовательность, то переходим к следующему символу
			i++
		}

		printRune = runes[i] // символ который нужно печатать

		if i+1 != len(runes) { // проверка на выход за преедлы слайса
			if unicode.IsDigit(runes[i+1]) {
				num, err := strconv.Atoi(string(runes[i+1]))
				if err != nil {
					return "", errors.New("cant convert int from string")
				}
				for j := 0; j < num; j++ { // печатает символ num раз
					builder.WriteRune(printRune)
				}
				i++
			} else { // иначе, просто добавляем символ
				builder.WriteRune(printRune)
			}
		} else { // запись последнего символа
			builder.WriteRune(printRune)
		}
	}

	return builder.String(), nil
}
