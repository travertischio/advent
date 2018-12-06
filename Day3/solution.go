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
	overlaps := make(map[int]bool)
	claimTotal := 0

	for {
		claim := readLine(reader)
		if claim == "" {
			break
		}
		claimTotal++

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

		cTemp, err := strconv.ParseInt(claimArr[0][1:], 10, 64)
		checkError(err)
		claimID := int(cTemp)

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				if fabric[x+i] == nil {
					fabric[x+i] = make(map[int]int)
				}

				if fabric[x+i][y+j] > 0 {
					overlaps[claimID] = true
					overlaps[fabric[x+i][y+j]] = true
				}
				fabric[x+i][y+j] = claimID
			}
		}
	}

	for i := 1; i <= claimTotal; i++ {
		if overlaps[i] == false {
			fmt.Println(i)
			break
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
