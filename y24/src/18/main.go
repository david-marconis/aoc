package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("res/day18.txt")
	s := bufio.NewScanner(f)
	byteList := [][2]int{}
	for s.Scan() {
		line := s.Text()
		sp := strings.Split(line, ",")
		x, _ := strconv.Atoi(sp[0])
		y, _ := strconv.Atoi(sp[1])
		byteList = append(byteList, [2]int{x, y})
	}
	n := 71
	stop := 1024
	isSample := len(byteList) == 25
	if isSample {
		n = 7
		stop = 12
	}
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
	}
	for i := 0; i < stop; i++ {
		b := byteList[i]
		grid[b[1]][b[0]] = '#'
	}

	target := [2]int{n - 1, n - 1}
	length := bfs(target, n, grid)
	fmt.Println(length)
	for i := stop; i < len(byteList); i++ {
		b := byteList[i]
		grid[b[1]][b[0]] = '#'
		if bfs(target, n, grid) == 0 {
			fmt.Printf("%d,%d\n", b[0], b[1])
			break
		}
	}
}

func bfs(target [2]int, n int, grid [][]byte) int {
	seen := make(map[[2]int]struct{})
	seen[[2]int{0, 0}] = struct{}{}
	queue := [][3]int{[3]int{0, 0, 0}}
	length := 0
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		x, y, c := pos[0], pos[1], pos[2]
		if x == target[0] && y == target[1] {
			length = c
			break
		}
		dir := [4][4]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		for _, d := range dir {
			nx, ny := d[0]+x, d[1]+y
			if nx < 0 || nx >= n || ny < 0 || ny >= n || grid[ny][nx] == '#' {
				continue
			}
			key := [2]int{nx, ny}
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			queue = append(queue, [3]int{nx, ny, c + 1})
		}
	}
	return length
}
