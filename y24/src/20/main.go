package main

import (
	"bufio"
	"container/heap"
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
	ends := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		ends[i] = make([]int, len(grid[i]))
	}
	cheats := make(map[[4]int]struct{})
	target := sp(grid, start, end, math.MaxInt, -1, ends, cheats)
	score := 0
	count := -1
	for ; score < target; count++ {
		score = sp(grid, start, end, target-100, 0, ends, cheats)
	}
	fmt.Println(count)
}

func sp(grid [][]byte, start, end [2]int, target, startPhase int, ends [][]int, cheats map[[4]int]struct{}) int {
	scores := make(map[[4]int]int)
	queue := make(PriorityQueue, 0)
	startItem := Item{start, 0, startPhase, startPhase, nil}
	queue.Push(&startItem)
	heap.Init(&queue)
	lowest := math.MaxInt
	for len(queue) > 0 {
		item := heap.Pop(&queue).(*Item)
		e := item.value
		x, y, c, cheatStart, cheatEnd := e[0], e[1], item.priority, item.cheatStart, item.cheatEnd
		costToEnd := ends[y][x]
		if c > target {
			break
		}
		if item.cheatStart != 0 {
			c += costToEnd
		}
		if c > target {
			continue
		}
		if cheatStart > 0 && item.parent.cheatStart == 0 {
			pv := item.parent.value
			cheats[[4]int{pv[0], pv[1], x, y}] = struct{}{}
		}
		if (x == end[0] && y == end[1]) || c != item.priority {
			parentCostToEnd := costToEnd + 1
			parent := item.parent
			for parent != nil && parent.cheatStart != 0 {
				ends[parent.value[1]][parent.value[0]] = parentCostToEnd
				parent = parent.parent
				parentCostToEnd++
			}
			return c
		}
		if s, ok := scores[[4]int{x, y, cheatStart, cheatEnd}]; ok && s < c {
			continue
		}
		dirs := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		nbs := []Item{}
		for _, dir := range dirs {
			nb := [2]int{x + dir[0], y + dir[1]}
			if grid[nb[1]][nb[0]] == '#' {
				continue
			}
			nbs = append(nbs, Item{nb, c + 1, cheatStart, cheatEnd, item})
		}
		for _, dir := range dirs {
			nb := [2]int{x + dir[0], y + dir[1]}
			if grid[nb[1]][nb[0]] != '#' || cheatStart != 0 {
				continue
			}
			for _, dir := range dirs {
				nb2 := [2]int{nb[0] + dir[0], nb[1] + dir[1]}
				if nb2[0] < 1 || nb2[0] >= len(grid[0])-1 || nb2[1] < 1 || nb2[1] >= len(grid) || grid[nb2[1]][nb2[0]] == '#' || (nb2[0] == x && nb2[1] == y) {
					continue
				}
				if _, ok := cheats[[4]int{x, y, nb2[0], nb2[1]}]; ok {
					continue
				}
				newCheatStart := len(grid[0])*y + x
				newCheatEnd := len(grid[0])*nb2[1] + nb2[0]
				nbs = append(nbs, Item{nb2, c + 2, newCheatStart, newCheatEnd, item})
			}
		}

		for _, nb := range nbs {
			key := [4]int{nb.value[0], nb.value[1], nb.cheatStart, nb.cheatEnd}
			if s, ok := scores[key]; ok && s <= nb.priority {
				continue
			}
			scores[key] = nb.priority
			heap.Push(&queue, &nb)
		}
	}
	return lowest
}

type Item struct {
	value      [2]int
	priority   int
	cheatStart int
	cheatEnd   int
	parent     *Item
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
