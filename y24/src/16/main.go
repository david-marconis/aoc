package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.Open("res/day16.txt")
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	start := [2]int{0, 0}
	end := [2]int{0, 0}
	for s.Scan() {
		line := s.Text()
		grid = append(grid, []byte(line))
		for i, c := range line {
			if c == 'E' {
				end = [2]int{i, len(grid) - 1}
			} else if c == 'S' {
				start = [2]int{i, len(grid) - 1}
			}
		}
	}
	queue := [][5]int{[5]int{start[0], start[1], 1, 0, 0}}
	scores := make(map[[3]int]int)
	sum := 0
	target := math.MaxInt
	visitedList := []map[[2]int]struct{}{{[2]int{start[0], start[1]}: struct{}{}}}
	ends := []int{}
	for len(queue) > 0 {
		minI := 0
		minS := math.MaxInt
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i][3] < minS {
				minS = queue[i][3]
				minI = i
			}
		}
		queue[minI], queue[n-1] = queue[n-1], queue[minI]
		e := queue[n-1]
		queue = queue[:n-1]
		x, y, d, c, p := e[0], e[1], e[2], e[3], e[4]
		visited := visitedList[p]
		if s, ok := scores[[3]int{x, y, d}]; ok && s < c {
			continue
		}
		if c > target {
			break
		}
		if x == end[0] && y == end[1] {
			sum = c
			target = c
			ends = append(ends, p)
		}
		dx, dy := 0, 0
		if d == 0 {
			dy = -1
		} else if d == 1 {
			dx = 1
		} else if d == 2 {
			dy = 1
		} else if d == 3 {
			dx = -1
		}
		nbs := [][5]int{}
		if grid[y+dy][x+dx] != '#' {
			nVisited := make(map[[2]int]struct{}, len(visited)+1)
			for k, v := range visited {
				nVisited[k] = v
			}
			nVisited[[2]int{x + dx, y + dy}] = struct{}{}
			visitedList = append(visitedList, nVisited)
			nbs = append(nbs, [5]int{x + dx, y + dy, d, c + 1, len(visitedList) - 1})
		}
		nbs = append(nbs, [5]int{x, y, (d + 1) % 4, c + 1000, p})
		nbs = append(nbs, [5]int{x, y, (d + 4 - 1) % 4, c + 1000, p})
		for _, nb := range nbs {
			nx, ny, nd, nc := nb[0], nb[1], nb[2], nb[3]
			key := [3]int{nx, ny, nd}
			if s, ok := scores[key]; ok && s < nc {
				continue
			}
			scores[key] = nc
			queue = append(queue, nb)
		}
	}
	fmt.Println(sum)
	allVisited := make(map[[2]int]struct{})
	for _, e := range ends {
		for k, v := range visitedList[e] {
			allVisited[k] = v
			grid[k[1]][k[0]] = 'O'
		}
	}
	fmt.Println(len(allVisited))
}
