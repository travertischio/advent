package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	frequency := int64(0)
	list := make([]int64, 0)
	frequencies := make(map[int64]bool)
	frequencies[0] = true

	for {
		nRead := readLine(reader)
		if nRead == "" {
			break
		}

		n, err := strconv.ParseInt(nRead, 10, 64)
		checkError(err)
		list = append(list, n)
		frequency += n
		if frequencies[frequency] {
			fmt.Println(frequency)
			return
		}
		frequencies[frequency] = true
	}

	for {
		for _, listN := range list {
			frequency += int64(listN)
			if frequencies[frequency] {
				fmt.Println(frequency)
				return
			}
			frequencies[frequency] = true
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
