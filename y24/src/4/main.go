package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("res/day4.txt")
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	for s.Scan() {
		line := s.Text()
		grid = append(grid, []byte(line))
	}

	sum := 0
	n := len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j <= n-4; j++ {
			sum += countXmas(grid[i][j : j+4])
			v := make([]byte, 4)
			for k := 0; k < 4; k++ {
				v[k] = grid[j+k][i]
			}
			sum += countXmas(v)
		}
	}

	for i := 0; i <= n-4; i++ {
		for j := 0; j <= n-4; j++ {
			d := make([]byte, 4)
			r := make([]byte, 4)
			for k := 0; k < 4; k++ {
				d[k] = grid[i+k][j+k]
				r[k] = grid[i+k][n-1-j-k]
			}
			sum += countXmas(d)
			sum += countXmas(r)
		}
	}

	sum2 := 0
	for i := 0; i <= n-3; i++ {
		for j := 0; j <= n-3; j++ {
			d := make([]byte, 3)
			r := make([]byte, 3)
			for k := 0; k < 3; k++ {
				d[k] = grid[i+k][j+k]
				r[k] = grid[i+k][j+2-k]
			}
			ds := string(d)
			rs := string(r)
			if (ds == "MAS" || ds == "SAM") && (rs == "MAS" || rs == "SAM") {
				sum2 += 1
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func countXmas(bytes []byte) int {
	sum := 0
	s := string(bytes)
	if s == "XMAS" {
		sum += 1
	}
	if s == "SAMX" {
		sum += 1
	}
	return sum
}
