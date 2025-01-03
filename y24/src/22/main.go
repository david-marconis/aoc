package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("res/day22.txt")
	s := bufio.NewScanner(f)
	nums := []int{}
	for s.Scan() {
		line := s.Text()
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}
	fmt.Println(secretNumber(nums))

	best := findBest(nums)
	fmt.Println(sell(nums, best))
}

func secretNumber(oldNums []int) int {
	nums := make([]int, len(oldNums))
	copy(nums, oldNums)
	sum := 0
	for i, _ := range nums {
		for j := 0; j < 2000; j++ {
			nums[i] = compute(nums[i])
		}
		sum += nums[i]
	}
	return sum
}

func findBest(oldNums []int) [4]int {
	nums := make([]int, len(oldNums))
	copy(nums, oldNums)
	seqCounts := make(map[[4]int]map[int]int)
	for i, _ := range nums {
		changes := [4]int{10, 10, 10, 10}
		for j := 0; j < 2000; j++ {
			oldPrice := nums[i] % 10
			nums[i] = compute(nums[i])
			price := nums[i] % 10
			copy(changes[0:3], changes[1:4])
			changes[3] = price - oldPrice
			if j < 3 {
				continue
			}
			m, ok := seqCounts[changes]
			if ok {
				if _, ok2 := m[i]; !ok2 {
					m[i] = price
				}
			} else {
				m = make(map[int]int)
				m[i] = price
				seqCounts[changes] = m
			}
		}
	}

	bestPrice := 0
	bestSeq := [4]int{10, 10, 10, 10}
	for seq, priceMap := range seqCounts {
		price := 0
		for _, p := range priceMap {
			price += p
		}
		if price > bestPrice {
			bestPrice = price
			bestSeq = seq
		}
	}
	return bestSeq
}

func sell(oldNums []int, seq [4]int) int {
	nums := make([]int, len(oldNums))
	copy(nums, oldNums)
	sum := 0
	for i, _ := range nums {
		changes := [4]int{10, 10, 10, 10}
		for j := 0; j < 2000; j++ {
			oldPrice := nums[i] % 10
			nums[i] = compute(nums[i])
			price := nums[i] % 10
			copy(changes[0:3], changes[1:4])
			changes[3] = price - oldPrice
			if changes == seq {
				sum += price
				break
			}
		}
	}
	return sum
}

func compute(s int) int {
	a := s * 64
	s = s ^ a
	s = s % 16777216
	b := s / 32
	s = s ^ b
	s = s % 16777216
	c := s * 2048
	s = s ^ c
	s = s % 16777216
	return s
}
