package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

type region struct {
	area       int
	perimiter  int
	perimiter2 int
}

func main() {
	f, _ := os.Open("res/day12.txt")
	s := bufio.NewScanner(f)
	grid := [][]byte{}
	var width int
	for s.Scan() {
		line := "." + s.Text() + "."
		if len(grid) == 0 {
			width = len(line)
			grid = append(grid, []byte(strings.Repeat(".", width)))
		}
		grid = append(grid, []byte(line))
	}
	grid = append(grid, []byte(strings.Repeat(".", width)))
	regions := findRegions(grid)
	sum2 := 0
	sum := 0
	for _, r := range regions {
		sum += r.area * r.perimiter
		sum2 += r.area * r.perimiter2
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

var adj = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func findRegions(grid [][]byte) []region {
	regions := []region{}
	allVisited := make(map[[2]int]int)
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if _, ok := allVisited[[2]int{j, i}]; ok {
				continue
			}
			visited := make(map[[2]int]struct{})
			edge := make(map[[4]int]struct{})
			exploreRegion(i, j, grid, visited, edge)
			keys := slices.Collect(maps.Keys(edge))
			slices.SortFunc(keys, func(a, b [4]int) int { return a[1] - b[1] + a[0] - b[0] })
			for _, n := range keys {
				mx, my := n[0]+1-(n[2]*n[2]), n[1]+1-(n[3]*n[3])
				if _, ok := edge[[4]int{mx, my, n[2], n[3]}]; ok {
					delete(edge, n)
				}
			}
			regions = append(regions, region{len(visited), len(keys), len(edge)})
			for k := range visited {
				allVisited[k] = 1
			}
		}
	}
	return regions
}

func exploreRegion(i, j int, grid [][]byte, visited map[[2]int]struct{}, edge map[[4]int]struct{}) {
	pos := [2]int{j, i}
	if _, ok := visited[pos]; ok {
		return
	}
	visited[pos] = struct{}{}
	for _, n := range adj {
		nx, ny := j+n[0], i+n[1]
		if grid[i][j] != grid[ny][nx] {
			edge[[4]int{nx, ny, n[0], n[1]}] = struct{}{}
		} else {
			exploreRegion(ny, nx, grid, visited, edge)
		}
	}
}
