package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	str    string
	counts []int
}

func main() {
	fmt.Println("-------------NEW----------")
	f, _ := os.Open("res/day12.txt")
	s := bufio.NewScanner(f)
	records := []record{}
	for s.Scan() {
		l := s.Text()
		sp := strings.Fields(l)
		str := sp[0]
		nums := strings.Split(sp[1], ",")
		counts := make([]int, len(nums))
		for i, v := range nums {
			x, _ := strconv.Atoi(v)
			counts[i] = x
		}
		records = append(records, record{str, counts})
	}

	records2 := []record{}
	for _, rec := range records {
		nc := len(rec.counts)
		counts := make([]int, nc*5)
		str2 := rec.str
		copy(counts, rec.counts)
		for z := nc; z < 5*nc; z += nc {
			copy(counts[z:z+nc], rec.counts)
			str2 += "?" + rec.str
		}
		records2 = append(records2, record{str2, counts})
	}

	for _, recs := range [2][]record{records, records2} {
		sum := 0
		for _, rec := range recs {
			cache := make([][]int, len(rec.str)*2)
			for i := 0; i < len(rec.str)*2; i++ {
				cache[i] = make([]int, len(rec.counts))
				for j := 0; j < len(rec.counts); j++ {
					cache[i][j] = -1
				}
			}
			sum += dfs(rec, 0, 0, cache)
		}
		fmt.Println(sum)
	}
}

func dfs(rec record, i, j int, cache [][]int) int {
	for i < len(rec.str) && rec.str[i] == '.' {
		i++
	}
	if j == len(rec.counts) {
		for n := i; n < len(rec.str); n++ {
			if rec.str[n] == '#' {
				return 0
			}
		}
		return 1
	} else if i >= len(rec.str) {
		return 0
	}

	hit := cache[i][j]
	if hit != -1 {
		return hit
	}

	c := rec.counts[j]
	match := i+c >= len(rec.str) || rec.str[i+c] != '#'
	for k := i; k < c+i && match; k++ {
		if k >= len(rec.str) || rec.str[k] == '.' {
			match = false
			break
		}
	}
	score := 0
	if match {
		score += dfs(rec, i+c+1, j+1, cache)
	}
	if rec.str[i] != '#' {
		score += dfs(rec, i+1, j, cache)
	}
	cache[i][j] = score
	return score
}
