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

	polymer := []rune(readLine(reader))

	shortest := len(polymer)
	for j := 'A'; j <= 'Z'; j++ {
		thisPoly := polymer

		thisPoly = []rune(strings.Replace(string(thisPoly), string(j), "", -1))
		thisPoly = []rune(strings.Replace(string(thisPoly), string(rune(j+32)), "", -1))

		for i := 0; i < len(thisPoly)-1; i++ {
			if abs(int(thisPoly[i])-int(thisPoly[i+1])) == 32 {
				thisPoly = append(thisPoly[:i], thisPoly[i+2:]...)
				if i == 0 {
					i = -1
				} else {
					i -= 2
				}

			}
		}
		if len(thisPoly) < shortest {
			shortest = len(thisPoly)
		}
	}
	fmt.Println(shortest)
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
