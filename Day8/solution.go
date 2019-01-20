package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers []int
var metadataTotal int

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nums := readLine(reader)
	numArr := strings.Split(nums, " ")

	for _, e := range numArr {
		num, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, int(num))
	}

	addMetadata(0)

	fmt.Println(metadataTotal)
}

func addMetadata(start int) int {
	children := numbers[start]
	metas := numbers[start+1]

	childStart := start + 2
	for i := 0; i < children; i++ {
		childStart = addMetadata(childStart)
	}

	metaBegin := childStart
	for j := 0; j < metas; j++ {
		metadataTotal += numbers[metaBegin+j]
	}

	return metaBegin + metas
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
