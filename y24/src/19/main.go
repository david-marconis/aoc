package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("res/day19.txt")
	s := bufio.NewScanner(f)
	s.Scan()
	towles := strings.Split(s.Text(), ", ")
	designs := []string{}
	s.Scan()
	for s.Scan() {
		designs = append(designs, s.Text())
	}
	towleMap := make(map[byte]*[]string)
	for _, t := range towles {
		l, ok := towleMap[t[0]]
		if ok {
			*l = append(*l, t)
		} else {
			l := []string{t}
			towleMap[t[0]] = &l
		}
	}
	sum := 0
	sum2 := 0
	for _, d := range designs {
		cache := make(map[int]int)
		count := dfs(d, 0, towleMap, cache)
		if count > 0 {
			sum += 1
		}
		sum2 += count
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func dfs(design string, i int, towleMap map[byte]*[]string, cache map[int]int) int {
	if i == len(design) {
		return 1
	}
	if i > len(design) {
		return 0
	}
	towles, ok := towleMap[design[i]]
	if !ok {
		return 0
	}
	if c, ok := cache[i]; ok {
		return c
	}
	sum := 0
	for _, t := range *towles {
		if !strings.HasPrefix(design[i:], t) {
			continue
		}
		sum += dfs(design, i+len(t), towleMap, cache)
	}
	cache[i] = sum
	return sum
}
