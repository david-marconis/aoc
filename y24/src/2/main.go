package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("res/day2.txt")
	s := bufio.NewScanner(f)
	sum := 0
	sum2 := 0
	for s.Scan() {
		line := strings.Fields(s.Text())
		sum += safe(line)
		for i := range line {
			nLine := make([]string, len(line)-1)
			copy(nLine, line[:i])
			copy(nLine[i:], line[i+1:])
			if safe(nLine) == 1 {
				sum2 += 1
				break
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func safe(line []string) int {
	inc := 0
	for i := 1; i < len(line); i++ {
		x, _ := strconv.Atoi(line[i-1])
		y, _ := strconv.Atoi(line[i])
		if inc == 0 {
			if x < y {
				inc = 1
			} else {
				inc = -1
			}
		}
		if (y-x)*inc < 1 || (y-x)*inc > 3 {
			return 0
		}
	}
	return 1
}
