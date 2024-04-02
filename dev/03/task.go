package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFileName  = "input.txt"
	outputFileName = "output.txt"
)

var (
	columnSortPtr  = flag.Int("k", -1, "указание колонки для сортировки ")
	numbersSortPtr = flag.Bool("n", false, "Сортировка по числовому значению")
	isReversePtr   = flag.Bool("r", false, "сортировать в обратном порядке")
	isUniquePtr    = flag.Bool("u", false, "не выводить повторяющиеся строки")
	isHelpPtr      = flag.Bool("help", false, "Помощь")
)

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

func writeLinesToFile(file string, strings []string) error {
	data, err := os.Create(file)
	if err != nil {
		return err
	}
	defer data.Close()

	if *isReversePtr {
		strings = reverse(strings)
	}
	if *isUniquePtr {
		strings = excludeRepeatedStrings(strings)
	}

	for _, v := range strings {
		_, err := data.WriteString(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func sortByColumn(lines []string, columnIdx int) ([]string, error) {
	var strByColumns [][]string
	var res []string

	for _, v := range lines {
		values := strings.Split(v, " ")
		if len(values)-1 < columnIdx {
			return nil, fmt.Errorf("there is no such column = %d", columnIdx)
		}
		strByColumns = append(strByColumns, values)
	}

	sort.Slice(strByColumns, func(i, j int) bool {
		if len(strByColumns[i]) == 0 || len(strByColumns[j]) == 0 {
			return len(strByColumns[i]) == 0
		}
		return strByColumns[i][columnIdx] < strByColumns[j][columnIdx]
	})

	for _, v := range strByColumns {
		res = append(res, strings.Join(v, " "))
	}

	return res, nil
}

func sortByNumbers(lines []string) ([]string, error) {
	var res []string
	var tempNums []int
	var tempStr string

	for _, v := range lines {
		tempStr = strings.TrimSuffix(v, "\n")
		tempStr = strings.Replace(tempStr, "\r", "", 1)
		n, err := strconv.Atoi(tempStr)

		if err != nil {
			return nil, err
		}
		tempNums = append(tempNums, n)
	}

	sort.Ints(tempNums)
	for _, v := range tempNums {
		res = append(res, fmt.Sprintf("%s\n", strconv.Itoa(v)))
	}
	return res, nil
}

func reverse(lines []string) []string {
	res := make([]string, len(lines))
	l := len(lines) - 1
	for i := range lines {
		res = append(res, lines[l-i])
	}
	return res
}

func excludeRepeatedStrings(lines []string) []string {
	set := make(map[string]struct{})
	var res []string

	for _, v := range lines {
		if _, exist := set[v]; !exist {
			res = append(res, v)
			set[v] = struct{}{}
		}
	}
	return res
}

func main() {
	flag.Parse()
	if *isHelpPtr {
		flag.PrintDefaults()
		os.Exit(0)
	}

	r, err := readLinesFromFile(inputFileName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if *numbersSortPtr {
		r, err = sortByNumbers(r)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else if *columnSortPtr >= 0 {
		r, err = sortByColumn(r, *columnSortPtr)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	err = writeLinesToFile(outputFileName, r)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
