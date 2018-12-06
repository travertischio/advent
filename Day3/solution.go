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
	fabric := make(map[int]map[int]int)
	overlaps := 0

	for {
		claim := readLine(reader)
		if claim == "" {
			break
		}

		claimArr := strings.Split(claim, " ")
		startArr := strings.Split(claimArr[2], ",")
		startY := startArr[1][:len(startArr[1])-1]
		yTemp, err := strconv.ParseInt(startY, 10, 64)
		checkError(err)
		y := int(yTemp)
		startX := startArr[0]
		xTemp, err := strconv.ParseInt(startX, 10, 64)
		checkError(err)
		x := int(xTemp)

		dimentionsArr := strings.Split(claimArr[3], "x")
		wTemp, err := strconv.ParseInt(dimentionsArr[0], 10, 64)
		checkError(err)
		w := int(wTemp)
		hTemp, err := strconv.ParseInt(dimentionsArr[1], 10, 64)
		checkError(err)
		h := int(hTemp)

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				if fabric[x+i] == nil {
					fabric[x+i] = make(map[int]int)
				}
				fabric[x+i][y+j]++
				if fabric[x+i][y+j] == 2 {
					overlaps++
				}
			}
		}
	}
	fmt.Println(overlaps)
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
