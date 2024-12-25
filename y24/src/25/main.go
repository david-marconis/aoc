package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-----------NEW-----------")
	f, _ := os.Open("res/day25.txt")
	s := bufio.NewScanner(f)
	locks := [][5]int{}
	keys := [][5]int{}
	grid := [][]byte{}
	for s.Scan() {
		line := s.Text()

		if line == "" {
			thing := [5]int{-1, -1, -1, -1, -1}
			for i := 0; i < len(grid[0]); i++ {
				for j := 0; j < len(grid); j++ {
					if grid[j][i] == '#' {
						thing[i]++
					}
				}
			}
			if grid[0][0] == '#' {
				locks = append(locks, thing)
			} else {
				keys = append(keys, thing)
			}
			grid = grid[0:0]
		} else {
			grid = append(grid, []byte(line))
		}
	}

	sum := 0
	for _, l := range locks {
	outer:
		for _, k := range keys {
			for i := 0; i < 5; i++ {
				if l[i]+k[i] >= 6 {
					continue outer
				}
			}
			sum += 1
		}
	}

	fmt.Println(sum)
}
