package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type macinhe struct{ a, b, p [2]int }

func main() {
	f, _ := os.Open("res/day13.txt")
	s := bufio.NewScanner(f)
	macs := []*macinhe{}
	for s.Scan() {
		p := [3][2]int{}
		for i := 0; i < 3; i++ {
			sp := strings.Fields(s.Text())
			x := sp[2-i/2]
			p[i][0], _ = strconv.Atoi(x[2 : len(x)-1])
			p[i][1], _ = strconv.Atoi(sp[3-i/2][2:])
			s.Scan()
		}
		macs = append(macs, &macinhe{p[0], p[1], p[2]})
	}
	calculate(macs)
	for _, m := range macs {
		m.p[0] += 10000000000000
		m.p[1] += 10000000000000
	}
	calculate(macs)
}

func calculate(macs []*macinhe) {
	sum := 0
	for _, m := range macs {
		A := [2][2]float64{
			{float64(m.a[0]), float64(m.b[0])},
			{float64(m.a[1]), float64(m.b[1])},
		}
		AI := [2][2]float64{
			{A[1][1], -A[0][1]},
			{-A[1][0], A[0][0]},
		}
		c := (A[0][0]*A[1][1] - A[0][1]*A[1][0])
		X := [2]float64{
			(AI[0][0]*float64(m.p[0]) + AI[0][1]*float64(m.p[1])) / c,
			(AI[1][0]*float64(m.p[0]) + AI[1][1]*float64(m.p[1])) / c,
		}
		ca, cb := int(math.Round(X[0])), int(math.Round(X[1]))
		if m.a[0]*ca+m.b[0]*cb == m.p[0] && m.a[1]*ca+m.b[1]*cb == m.p[1] {
			sum += 3*ca + cb
		}
	}
	fmt.Println(sum)
}
