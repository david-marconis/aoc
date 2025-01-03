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
	var output = run(a, b, c, program, false)
	// Jesus chris Go...
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(output), " ", ",", -1), "[]"))
	l := 4
	d := 1
	for i := 0; i < int(math.Pow(8, float64(len(program)))); i += d {
		output = run(i, 0, 0, program, true)
		if len(output) > l {
			l++
			d *= 8
		}
		if len(output) != len(program) {
			continue
		}
		fmt.Println(i)
		break
	}
}

func run(a, b, c int, program []int, selfValidate bool) []int {
	output := []int{}
	pc := 0
	z, o, t, e := 0, 1, 2, 3
	r := [7]*int{&z, &o, &t, &e, &a, &b, &c}
	for pc < len(program) {
		in, op := program[pc], program[pc+1]
		switch in {
		case 0:
			a = a / int(math.Pow(2, float64(*r[op])))
		case 1:
			b = b ^ op
		case 2:
			b = *r[op] % 8
		case 3:
			if a != 0 {
				pc = op
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			output = append(output, *r[op]%8)
			if selfValidate && output[len(output)-1] != program[len(output)-1] {
				return output[:len(output)-1]
			}
		case 6:
			b = a / int(math.Pow(2, float64(*r[op])))
		case 7:
			c = a / int(math.Pow(2, float64(*r[op])))
		}
		pc += 2
	}
	return output
}
