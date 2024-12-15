package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("res/day15.txt")
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	start := [2]int{0, 0}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		for i := 0; i < len(line); i++ {
			if line[i] == '@' {
				start = [2]int{i, len(grid)}
			}
		}
		grid = append(grid, []byte(line))
	}
	movements := []byte{}
	for s.Scan() {
		for _, c := range s.Text() {
			movements = append(movements, byte(c))
		}
	}
	grid2 := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		width := len(grid[0]) * 2
		grid2[i] = make([]byte, width)
		for j := 0; j < width; j += 2 {
			switch grid[i][j/2] {
			case 'O':
				grid2[i][j] = '['
				grid2[i][j+1] = ']'
			case '@':
				grid2[i][j] = '@'
				grid2[i][j+1] = '.'
			case '#':
				grid2[i][j] = '#'
				grid2[i][j+1] = '#'
			case '.':
				grid2[i][j] = '.'
				grid2[i][j+1] = '.'
			}
		}
	}
	start2 := [2]int{start[0] * 2, start[1]}

	part1(start, movements, grid)
	part2(start2, movements, grid2)

}

func part1(start [2]int, movements []byte, grid [][]byte) {
	x, y := start[0], start[1]
outer:
	for _, m := range movements {
		dx, dy := 0, 0
		switch m {
		case '<':
			dx = -1
		case '^':
			dy -= 1
		case '>':
			dx += 1
		case 'v':
			dy += 1
		}
		nx, ny := x+dx, y+dy
		if grid[ny][nx] == '.' {
			grid[ny][nx] = '@'
			grid[y][x] = '.'
			x = nx
			y = ny
			continue
		}
		if grid[ny][nx] == '#' {
			continue
		}
		mx, my := nx, ny
		for grid[my][mx] == 'O' {
			mx += dx
			my += dy
			if mx < 0 || mx >= len(grid[0]) || my < 0 || my >= len(grid) {
				continue outer
			}
		}
		if grid[my][mx] == '#' {
			continue
		}

		for ox, oy := mx, my; ox != nx || oy != ny; {
			grid[oy][ox] = 'O'
			ox -= dx
			oy -= dy
		}
		grid[ny][nx] = '@'
		grid[y][x] = '.'
		x = nx
		y = ny
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)
}

func part2(start [2]int, movements []byte, grid [][]byte) {
	x, y := start[0], start[1]
outer:
	for _, m := range movements {
		dx, dy := 0, 0
		switch m {
		case '<':
			dx = -1
		case '^':
			dy -= 1
		case '>':
			dx += 1
		case 'v':
			dy += 1
		}
		nx, ny := x+dx, y+dy
		if grid[ny][nx] == '.' {
			grid[ny][nx] = '@'
			grid[y][x] = '.'
			x = nx
			y = ny
			continue
		}
		if grid[ny][nx] == '#' {
			continue
		}
		mx, my := nx, ny
		for grid[my][mx] == '[' || grid[my][mx] == ']' {
			mx += dx
			my += dy
			if mx < 0 || mx >= len(grid[0]) || my < 0 || my >= len(grid) {
				continue outer
			}
		}
		if grid[my][mx] == '#' {
			continue
		}

		if !canMove(nx, ny, dx, dy, grid) {
			continue
		}
		move(nx, ny, dx, dy, grid)
		grid[ny][nx] = '@'
		grid[y][x] = '.'
		x = nx
		y = ny
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '[' {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)
}

func canMove(x, y, dx, dy int, grid [][]byte) bool {
	var ox int
	if grid[y][x] == '[' {
		ox = x + 1
	} else {
		ox = x - 1
	}

	if dy != 0 && grid[y+dy][x] == '.' && grid[y+dy][ox] == '.' {
		return true
	}
	if dx != 0 && grid[y][x+dx+dx] == '.' {
		return true
	}
	if dy != 0 && (grid[y+dy][x] == '#' || grid[y+dy][ox] == '#') {
		return false
	}
	if dx != 0 && grid[y][x+dx+dx] == '#' {
		return false
	}
	if dx != 0 {
		return canMove(x+dx+dx, y, dx, dy, grid)
	}
	result := true
	if grid[y+dy][x] != '.' {
		result = canMove(x, y+dy, dx, dy, grid)
	}
	if grid[y+dy][ox] != '.' && result {
		result = canMove(ox, y+dy, dx, dy, grid)
	}
	return result
}

func move(x, y, dx, dy int, grid [][]byte) {
	var ox int
	if grid[y][x] == '[' {
		ox = x + 1
	} else {
		ox = x - 1
	}

	if dx != 0 && grid[y][x+dx+dx] != '.' {
		move(x+dx+dx, y, dx, dy, grid)
	}
	if dy != 0 && grid[y+dy][x] != '.' {
		move(x, y+dy, dx, dy, grid)
	}
	if dy != 0 && grid[y+dy][ox] != '.' {
		move(ox, y+dy, dx, dy, grid)
	}
	if dy != 0 && grid[y+dy][x] == '.' && grid[y+dy][ox] == '.' {
		grid[y+dy][x] = grid[y][x]
		grid[y+dy][ox] = grid[y][ox]
		grid[y][x] = '.'
		grid[y][ox] = '.'
	}
	if dx != 0 && grid[y][x+dx+dx] == '.' {
		grid[y][x+dx+dx] = grid[y][x+dx]
		grid[y][x+dx] = grid[y][x]
		grid[y][x] = '.'
	}
}
