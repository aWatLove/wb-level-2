package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
	Утилита grep

Реализовать утилиту фильтрации по аналогии с консольной
утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

const (
	inputFile = "input.txt"
)

var (
	afterFlag      = flag.Int("A", 0, "after. Печатать +N строк после совпадения")
	beforeFlag     = flag.Int("B", 0, "before. Печатать +N строк до совпадения")
	contextFlag    = flag.Int("C", 0, "context. (A+B) печатать +-N строк вокруг совпадения")
	countFlag      = flag.Bool("c", false, "count. Количество строк")
	ignoreCaseFlag = flag.Bool("i", false, "ignore-case. Игнорировать регистр")
	invertFlag     = flag.Bool("v", false, "invert. Вместо совпадения, исключать")
	fixedFlag      = flag.Bool("F", false, "fixed. Точное совпадение со строкой, не паттерн")
	lineNumFlag    = flag.Bool("n", false, "line num. Напечатать номер строки")
)

// Pair - ключ - номер строки, значение - сама строка
type Pair struct {
	Key   int
	Value string
}

// PairList - тип, срез структур Pair, который реализует интерфейс из библиотеки "sort"
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func readLinesFromFile(file string) ([]string, error) {
	data, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var res []string
	r := bufio.NewReader(data)
	const delim = '\n'
	for {
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}
			res = append(res, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return res, nil
}

func trimEndSymbols(s string) string {
	s = strings.TrimSuffix(s, "\n")
	s = strings.Replace(s, "\r", "", 1)
	return s
}

func addAfterStrings(n int, idx int, lines []string, m *map[int]string) {
	var tmpStr string
	endBorder := len(lines) - 1
	var rightBorder int
	idx++
	if idx+n > endBorder {
		rightBorder = endBorder
	} else {
		rightBorder = idx + n
	}

	for i := idx; i < rightBorder; i++ {
		tmpStr = trimEndSymbols(lines[i])
		(*m)[i+1] = tmpStr
	}
}

func addBeforeStrings(n int, idx int, lines []string, m *map[int]string) {
	var tmpStr string

	var leftBorder int
	if n < idx {
		leftBorder = idx - n
	}

	for i := leftBorder; i < idx; i++ {
		tmpStr = trimEndSymbols(lines[i])
		(*m)[i+1] = tmpStr
	}
}

func convertMapToStrings(m *map[int]string) []string {
	res := make([]string, len(*m))
	pairs := make(PairList, len(*m))

	var tmpStr string

	i := 0
	for k, v := range *m {
		pairs[i] = Pair{k, v}
		i++
	}
	i = 0
	sort.Sort(pairs)
	for _, v := range pairs {
		if *lineNumFlag {
			tmpStr = fmt.Sprintf("%d %s", v.Key, v.Value)
		} else {
			tmpStr = v.Value
		}
		res[i] = tmpStr
		i++
	}

	return res
}

func grep(lines []string, pattern string) (int, []string, error) {
	var res []string
	var m = make(map[int]string, len(lines))
	var tmpStr string

	if *fixedFlag {
		for i, v := range lines {
			tmpStr = trimEndSymbols(v)

			if *ignoreCaseFlag {
				tmpStr = strings.ToLower(v)
				pattern = strings.ToLower(pattern)
			}

			if (tmpStr == pattern && !*invertFlag) || (tmpStr != pattern && *invertFlag) {
				m[i+1] = tmpStr
				if *contextFlag > 0 {
					addAfterStrings(*contextFlag, i, lines, &m)
					addBeforeStrings(*contextFlag, i, lines, &m)
				} else {
					if *afterFlag > 0 {
						addAfterStrings(*afterFlag, i, lines, &m)
					}
					if *beforeFlag > 0 {
						addBeforeStrings(*beforeFlag, i, lines, &m)
					}
				}
			}
		}
		res = convertMapToStrings(&m)
		return len(res), res, nil
	}

	if *ignoreCaseFlag {
		pattern = fmt.Sprintf("(?i)%s", pattern)
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return 0, nil, err
	}

	for i, v := range lines {
		tmpStr = trimEndSymbols(v)
		if (r.MatchString(tmpStr) && !*invertFlag) || (!r.MatchString(tmpStr) && *invertFlag) {
			m[i+1] = tmpStr
			if *contextFlag > 0 {
				addAfterStrings(*contextFlag, i, lines, &m)
				addBeforeStrings(*contextFlag, i, lines, &m)
			} else {
				if *afterFlag > 0 {
					addAfterStrings(*afterFlag, i, lines, &m)
				}
				if *beforeFlag > 0 {
					addBeforeStrings(*beforeFlag, i, lines, &m)
				}
			}
		}
	}
	if *countFlag {
		return len(m), nil, nil
	}
	res = convertMapToStrings(&m)
	return len(res), res, nil
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("there are no arguments")
		flag.PrintDefaults()
		os.Exit(1)
	}

	strs, err := readLinesFromFile(inputFile)
	if err != nil {
		fmt.Printf("error while reading file: %s", err)
		os.Exit(1)
	}

	pattern := flag.Arg(0)

	count, res, err := grep(strs, pattern)
	if err != nil {
		fmt.Printf("error during grep file: %s", err)
		os.Exit(1)
	}
	if len(res) != 0 {
		for _, v := range res {
			fmt.Println(v)
		}
	} else {
		fmt.Println(count)
	}
}
