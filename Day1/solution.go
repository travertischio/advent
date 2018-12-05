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

	for {
		nRead := readLine(reader)
		if nRead == "" {
			break
		}

		n, err := strconv.ParseInt(nRead, 10, 64)
		checkError(err)
		frequency += n
	}

	fmt.Println(frequency)
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
