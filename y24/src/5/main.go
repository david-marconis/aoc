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
	f, _ := os.Open("res/day5.txt")
	s := bufio.NewScanner(f)
	rules := [][]int{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		sp := strings.Split(line, "|")
		x, _ := strconv.Atoi(sp[0])
		y, _ := strconv.Atoi(sp[1])
		rules = append(rules, []int{x, y})
	}

	manuals := [][]int{}
	for s.Scan() {
		sp := strings.Split(s.Text(), ",")
		manual := make([]int, len(sp))
		for i, v := range sp {
			page, _ := strconv.Atoi(v)
			manual[i] = page
		}
		manuals = append(manuals, manual)
	}

	badMans := [][]int{}
	sum := 0
outer:
	for _, m := range manuals {
		indices := make(map[int]int)
		for i, v := range m {
			indices[v] = i
		}
		for _, r := range rules {
			ix, okX := indices[r[0]]
			iy, okY := indices[r[1]]
			if okX && okY && ix >= iy {
				badMans = append(badMans, m)
				continue outer
			}
		}
		sum += m[len(m)/2]
	}
	fmt.Println(sum)

	sum2 := 0
	for _, m := range badMans {
		slices.SortFunc(m, func(a, b int) int {
			for _, r := range rules {
				if a == r[0] && b == r[1] {
					return -1
				} else if a == r[1] && b == r[0] {
					return 1
				}
			}
			return 0
		})
		sum2 += m[len(m)/2]
	}
	fmt.Println(sum2)
}
