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

	for i := 0; i < 2000; i++ {
		for i, s := range nums {
			a := s * 64
			s = s ^ a
			s = s % 16777216
			b := s / 32
			s = s ^ b
			s = s % 16777216
			c := s * 2048
			s = s ^ c
			s = s % 16777216
			nums[i] = s
		}
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)
}
