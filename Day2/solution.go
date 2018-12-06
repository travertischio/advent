package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func findClosest(ids []string) {
	finalf := ""
	finals := ""
Outerloop:
	for _, first := range ids {
		for _, second := range ids {
			if first == second {
				continue
			}
			misses := 0
			for ci := range first {
				if first[ci] != second[ci] {
					misses++
				}
				if misses > 1 {
					break
				}
				if ci == len(first)-1 {
					fmt.Printf("Exiting with %s and %s\n", first, second)
					finalf = first
					finals = second
					break Outerloop
				}
			}
		}
	}
	finalString := ""
	for fi := range finalf {
		if finalf[fi] == finals[fi] {
			finalString += string([]byte{finalf[fi]})
		}
	}

	fmt.Printf("Final string: %s\n", finalString)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	ids := make([]string, 0)

	for {
		id := readLine(reader)
		if id == "" {
			break
		}

		ids = append(ids, id)
	}

	findClosest(ids)
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
