package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("res/day1.txt")
	s := bufio.NewScanner(f)
	l1 := []int{}
	l2 := []int{}
	for s.Scan() {
		line := strings.Fields(s.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		l1 = append(l1, x)
		l2 = append(l2, y)
	}
	slices.Sort(l1)
	slices.Sort(l2)
	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += Abs(l1[i] - l2[i])
	}
	fmt.Println(sum)
	sum2 := 0
	for i := 0; i < len(l1); i++ {
		x := 0
		for j := 0; j < len(l1); j++ {
			if l2[j] == l1[i] {
				x++
			}
		}
		sum2 += l1[i] * x
	}
	fmt.Println(sum2)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
