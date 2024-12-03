package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("res/day3.txt")
	s := string(bytes)
	r, _ := regexp.Compile("(mul\\(\\d+,\\d+\\))|(do\\(\\))|(don't\\(\\))")
	res := r.FindAllString(s, -1)
	var sum, sum2 uint64 = 0, 0
	do := true
	for _, v := range res {
		switch v {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			sp := strings.Split(v[4 : len(v)-1], ",")
			x, _ := strconv.ParseUint(sp[0], 10, 64)
			y, _ := strconv.ParseUint(sp[1], 10, 64)
			if do {
				sum2 += x * y
			}
			sum += x * y
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
