package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Утилита cut

Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по
разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

// ArrayFlags - тип для реализации интерфейса Value из пакета flag
type ArrayFlags []string

// String - реализация метода String интерфейса Value из пакета flag
func (f *ArrayFlags) String() string {
	return ""
}

// Set - реализация метода Set интерфейса Value из пакета flag
func (f *ArrayFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}

var (
	arrayFlags    ArrayFlags
	delimFlag     = flag.String("d", "\t", "delimiter - использовать другой разделитель")
	separatedFlag = flag.Bool("s", true, "separated - только строки с разделителем")
)

func cut(lines []string) ([]string, error) {
	var res []string
	var builder strings.Builder

	for _, v := range lines {
		tempStrings := strings.Split(v, *delimFlag)
		if (*separatedFlag && len(tempStrings) > 1) || (!*separatedFlag) {
			for _, column := range arrayFlags {
				val, err := strconv.Atoi(column)
				if err != nil {
					return nil, err
				}

				if val <= len(tempStrings) {
					builder.WriteString(tempStrings[val-1])
					builder.WriteString(*delimFlag)
				}
			}
			builder.WriteString("\n")
			res = append(res, builder.String())
			builder.Reset()
		}
	}

	return res, nil
}

func main() {
	arrayFlags = []string{"1", "3"}
	flag.Var(&arrayFlags, "f", "fields - выбрать поля (колонки)\n")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	lines := make([]string, n)
	i := 0

	for range lines {
		str, err := bufio.NewReader(os.Stdin).ReadString('\r')
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		lines[i] = str
		i++
	}

	r, err := cut(lines)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println()
	for _, v := range r {
		fmt.Print(v)
	}

}
