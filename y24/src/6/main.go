package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	f, _ := os.Open("res/day6.txt")
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	var start [2]int
	for s.Scan() {
		line := s.Text()
		row := make([]byte, len(line))
		for i := range line {
			row[i] = line[i]
			if line[i] == '^' {
				start = [2]int{i, len(grid)}
			}
		}
		grid = append(grid, row)
	}

	dir := 0 // 0123 = NESW
	m := make(map[[2]int]uint8)
	pos := [2]int{start[0], start[1]}

	for pos[0] >= 0 && pos[0] < len(grid[0]) && pos[1] >= 0 && pos[1] < len(grid) {
		m[pos] = 1
		dx, dy := 0, 0
		switch dir {
		case 0:
			dy = -1
		case 1:
			dx = 1
		case 2:
			dy = 1
		case 3:
			dx = -1
		}
		nx, ny := pos[0]+dx, pos[1]+dy
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
			break
		}
		if grid[ny][nx] == '#' {
			dir += 1
			dir %= 4
		} else {
			pos[0] = nx
			pos[1] = ny
		}
	}
	fmt.Println(len(m))

	var wg sync.WaitGroup
	results := make(chan int, len(grid)*len(grid[0]))

	for p := range m {
		i, j := p[1], p[0]
		if grid[i][j] == '^' || grid[i][j] == '#' {
			continue
		}
		x, y := j, i
		wg.Add(1)
		go func(x, y int) {
			defer wg.Done()
			visited := make(map[[3]int]uint8)
			pos := [2]int{start[0], start[1]}
			dir := 0 // 0123 = NESW
			for pos[0] >= 0 && pos[0] < len(grid[0]) && pos[1] >= 0 && pos[1] < len(grid) {
				if visited[[3]int{pos[0], pos[1], dir}] == 1 {
					results <- 1
					return
				}
				visited[[3]int{pos[0], pos[1], dir}] = 1
				dx, dy := 0, 0
				switch dir {
				case 0:
					dy = -1
				case 1:
					dx = 1
				case 2:
					dy = 1
				case 3:
					dx = -1
				}
				nx, ny := pos[0]+dx, pos[1]+dy
				if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
					results <- 0
					return
				}
				if grid[ny][nx] == '#' || nx == x && ny == y {
					dir += 1
					dir %= 4
				} else {
					pos[0] = nx
					pos[1] = ny
				}
			}
		}(x, y)
	}
	wg.Wait()
	close(results)
	sum2 := 0
	for v := range results {
		sum2 += v
	}
	fmt.Println(sum2)
}
