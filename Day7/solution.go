package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type requirement struct {
	name     string
	parents  []*requirement
	children []*requirement
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	reqList := make(map[string]*requirement)
	totalReqs := 0

	for {
		dep := readLine(reader)
		if dep == "" {
			break
		}

		taskArr := strings.Split(dep, " ")
		parent := taskArr[1]
		child := taskArr[7]

		c, cExists := reqList[child]
		if !cExists {
			cp := make([]*requirement, 0)
			cc := make([]*requirement, 0)
			c = &requirement{
				name:     child,
				parents:  cp,
				children: cc,
			}
			totalReqs++
			reqList[child] = c
		}

		p, pExists := reqList[parent]
		if !pExists {
			pp := make([]*requirement, 0)
			pc := make([]*requirement, 0)
			p = &requirement{
				name:     parent,
				parents:  pp,
				children: pc,
			}
			totalReqs++
			reqList[parent] = p
		}

		p.children = append(p.children, c)
		c.parents = append(c.parents, p)
	}

	visited := make(map[string]bool)
	out := ""
	for totalReqs > 0 {
		for i := 0; i < 26; i++ {
			name := string(rune('A' + i))
			if _, ok := visited[name]; !ok {
				avaliable := true
				for _, p := range reqList[name].parents {
					if _, ok := visited[p.name]; !ok {
						avaliable = false
						break
					}
				}

				if avaliable {
					visited[name] = true
					totalReqs--
					out = fmt.Sprintf("%s%s", out, name)
					break
				}
			}
		}
	}

	fmt.Println(out)
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
