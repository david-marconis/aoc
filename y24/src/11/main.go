package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\n-----------NEW----------")
	bytes, _ := os.ReadFile("res/day11.txt")
	ns := strings.Fields(string(bytes))
	nums := make([]int, len(ns))
	for i, v := range ns {
		x, _ := strconv.Atoi(v)
		nums[i] = x
	}
	runSim(nums, 25)
	runSim(nums, 75)
}

func runSim(nums []int, target int) {
	cache := make(map[[2]int]int)
	sum := 0
	for _, num := range nums {
		sum += dfs(num, 0, target, cache)
	}
	fmt.Println(sum)
}

func dfs(num, i, target int, cache map[[2]int]int) int {
	key := [2]int{num, i}
	if v, ok := cache[key]; ok {
		return v
	}
	if i == target {
		return 1
	}
	score := 0
	str := strconv.Itoa(num)
	if num == 0 {
		score += dfs(1, i+1, target, cache)
	} else if len(str)%2 == 0 {
		half := len(str) / 2
		x, _ := strconv.Atoi(str[0:half])
		y, _ := strconv.Atoi(str[half:])
		score += dfs(x, i+1, target, cache) + dfs(y, i+1, target, cache)
	} else {
		score += dfs(num*2024, i+1, target, cache)
	}
	cache[key] = score
	return score
}
