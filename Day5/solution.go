package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	polymer := []byte(readLine(reader))

	for i := 0; i < len(polymer)-1; i++ {
		if abs(int(polymer[i])-int(polymer[i+1])) == 32 {
			polymer = append(polymer[:i], polymer[i+2:]...)
			if i == 0 {
				i = -1
			} else {
				i -= 2
			}

		}
	}
	fmt.Println(len(polymer))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
