package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	awake := make(map[string][]int)
	dutySchedule := make(map[int][]string)

	for {
		shiftNote := readLine(reader)
		if shiftNote == "" {
			break
		}

		shiftNoteArr := strings.Split(shiftNote, " ")

		layout := "2006-01-02"
		date, err := time.Parse(layout, shiftNoteArr[0][1:])
		checkError(err)

		timeArr := strings.Split(shiftNoteArr[1], ":")
		hTemp, err := strconv.ParseInt(timeArr[0], 10, 64)
		checkError(err)
		h := int(hTemp)
		mTemp, err := strconv.ParseInt(timeArr[1][:len(timeArr)], 10, 64)
		checkError(err)
		m := int(mTemp)

		if h > 0 {
			date = date.AddDate(0, 0, 1)
		}

		dateString := date.Format(layout)

		status := shiftNoteArr[2]
		if status == "Guard" {
			gTemp, err := strconv.ParseInt(shiftNoteArr[3][1:], 10, 64)
			checkError(err)
			gaurdID := int(gTemp)

			dutySchedule[gaurdID] = append(dutySchedule[gaurdID], dateString)
		} else {
			if awake[dateString] == nil {
				awake[dateString] = make([]int, 60)
			}
			if status == "wakes" {
				awake[dateString][m] = 1
			} else {
				awake[dateString][m] = -1
			}
		}
	}

	mostAsleepID := -1
	mostSingleMin := -1
	mostMinNum := -1
	for id, arr := range dutySchedule {
		timesAsleep := make([]int, 60)

		for _, date := range arr {
			if awake[date] == nil {
				awake[date] = make([]int, 60)
			}
			current := 1
			for i := 0; i < 60; i++ {
				if awake[date][i] != 0 {
					current = awake[date][i]
				}

				if current == -1 {
					awake[date][i] = 0
					timesAsleep[i]++
					if timesAsleep[i] > mostSingleMin {
						mostSingleMin = timesAsleep[i]
						mostAsleepID = id
						mostMinNum = i
					}
				} else {
					awake[date][i] = 1
				}
			}
		}
	}

	fmt.Println(mostMinNum * mostAsleepID)
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
