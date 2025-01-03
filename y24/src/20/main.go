package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("----------------NEW-------------")
	f, _ := os.Open("res/day20.txt")
	s := bufio.NewScanner(f)
	start := [2]int{0, 0}
	end := [2]int{0, 0}
	grid := [][]byte{}
	for s.Scan() {
		line := s.Text()
		for i, c := range line {
			if c == 'S' {
				start[0], start[1] = i, len(grid)
			} else if c == 'E' {
				end[0], end[1] = i, len(grid)
			}
		}
		grid = append(grid, []byte(line))
	}
	costs, maxCost := findCosts(grid, start, end)
	fmt.Println(findCheats(grid, costs, start, end, maxCost-100, 2))
	fmt.Println(findCheats(grid, costs, start, end, maxCost-100, 20))
}

func findCheats(grid [][]byte, costs [][]int, start [2]int, end [2]int, target, cheatSize int) int {
	seen := make(map[[2]int]struct{})
	queue := [][3]int{{start[0], start[1], 0}}
	seen[start] = struct{}{}
	cheats := 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		x, y, c := item[0], item[1], item[2]
		if c > target {
			break
		}
		for i := max(1, y-cheatSize); i <= min(len(grid)-1, y+cheatSize); i++ {
			dy := int(math.Abs(float64(i - y)))
			for j := max(1, x-cheatSize+dy); j <= min(len(grid[0])-1, x+cheatSize-dy); j++ {
				dx := int(math.Abs(float64(j - x)))
				nc := c + dx + dy + costs[i][j]
				if grid[i][j] == '#' || nc > target {
					continue
				}
				cheats += 1
			}
		}
		dirs := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		for _, dir := range dirs {
			nb := [2]int{x + dir[0], y + dir[1]}
			_, ok := seen[nb]
			if grid[nb[1]][nb[0]] == '#' || ok {
				continue
			}
			seen[nb] = struct{}{}
			queue = append(queue, [3]int{nb[0], nb[1], c + 1})
		}
	}
	return cheats
}

func findCosts(grid [][]byte, start, end [2]int) ([][]int, int) {
	seen := make(map[[2]int]struct{})
	queue := [][3]int{{end[0], end[1], 0}}
	seen[end] = struct{}{}
	costs := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		costs[i] = make([]int, len(grid[i]))
	}
	maxCost := 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		x, y, c := item[0], item[1], item[2]
		costs[y][x] = c
		maxCost = c
		dirs := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		for _, dir := range dirs {
			nb := [2]int{x + dir[0], y + dir[1]}
			_, ok := seen[nb]
			if grid[nb[1]][nb[0]] == '#' || ok {
				continue
			}
			seen[nb] = struct{}{}
			queue = append(queue, [3]int{nb[0], nb[1], c + 1})
		}
	}
	return costs, maxCost
}
