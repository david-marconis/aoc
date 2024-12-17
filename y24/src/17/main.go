package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("res/day17.txt")
	s := bufio.NewScanner(f)
	a, b, c := 0, 0, 0
	for _, r := range []*int{&a, &b, &c} {
		s.Scan()
		sp := strings.Fields(s.Text())
		*r, _ = strconv.Atoi(sp[2])
	}
	s.Scan()
	s.Scan()
	pline := strings.Fields(s.Text())
	sp := strings.Split(pline[1], ",")
	program := make([]int, len(sp))
	for i, n := range sp {
		program[i], _ = strconv.Atoi(n)
	}
	var output = run(a, b, c, program)
	// Jesus chris Go...
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(output), " ", ",", -1), "[]"))
}

func run(a, b, c int, program []int) []int {
	output := []int{}
	pc := 0
	for pc < len(program) {
		in, op := program[pc], program[pc+1]
		switch in {
		case 0:
			a = a / int(math.Pow(2, float64(combo(op, a, b, c))))
		case 1:
			b = b ^ op
		case 2:
			b = combo(op, a, b, c) % 8
		case 3:
			if a != 0 {
				pc = op
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			output = append(output, combo(op, a, b, c)%8)
		case 6:
			b = a / int(math.Pow(2, float64(combo(op, a, b, c))))
		case 7:
			c = a / int(math.Pow(2, float64(combo(op, a, b, c))))
		}
		pc += 2
	}
	return output
}

func combo(op, a, b, c int) int {
	r := [7]int{0, 1, 2, 3, a, b, c}
	return r[op]
}
