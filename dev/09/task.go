package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

/*
Утилита wget

Реализовать утилиту wget с возможностью скачивать сайты
целиком.
*/

const outputFile = "output.txt"

func writeHTMLToFile(file string, response http.Response) error {
	data, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func(data *os.File) {
		err = data.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(data)

	writer := bufio.NewWriter(data)
	_, err = io.Copy(writer, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func wget(file, url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	err = writeHTMLToFile(file, *r)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("there are no args")
	}

	url := flag.Arg(0)

	err := wget(outputFile, url)
	if err != nil {
		log.Fatal(err.Error())
	}
}
