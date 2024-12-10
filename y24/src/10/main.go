package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("res/day10.txt")
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	zeroes := [][2]int{}
	for s.Scan() {
		line := s.Text()
		row := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			x := line[i] - '0'
			row[i] = x
			if x == 0 {
				zeroes = append(zeroes, [2]int{i, len(grid)})
			}
		}
		grid = append(grid, row)
	}

	visited := make(map[[2]int]map[[2]int]int)
	visited2 := make(map[[2]int]int)
	sum := 0
	sum2 := 0
	for _, z := range zeroes {
		sum += len(part1(grid, visited, z))
		sum2 += part2(grid, visited2, z)
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func part1(grid [][]byte, visited map[[2]int]map[[2]int]int, pos [2]int) map[[2]int]int {
	c, ok := visited[pos]
	if ok {
		return c
	}
	if grid[pos[1]][pos[0]] == 9 {
		m := make(map[[2]int]int)
		m[pos] = 1
		return m
	}
	ns := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	score := make(map[[2]int]int)
	for _, n := range ns {
		nx, ny := pos[0]+n[0], pos[1]+n[1]
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) || grid[ny][nx] != grid[pos[1]][pos[0]]+1 {
			continue
		}
		for k, v := range part1(grid, visited, [2]int{nx, ny}) {
			score[k] = v
		}
	}
	visited[pos] = score
	return score
}

func part2(grid [][]byte, visited map[[2]int]int, pos [2]int) int {
	c, ok := visited[pos]
	if ok {
		return c
	}
	if grid[pos[1]][pos[0]] == 9 {
		return 1
	}
	ns := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	score := 0
	for _, n := range ns {
		nx, ny := pos[0]+n[0], pos[1]+n[1]
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) || grid[ny][nx] != grid[pos[1]][pos[0]]+1 {
			continue
		}
		score += part2(grid, visited, [2]int{nx, ny})
	}
	visited[pos] = score
	return score
}
