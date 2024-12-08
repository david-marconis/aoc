package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("res/day8.txt")
	s := bufio.NewScanner(f)
	antennas := [][]byte{}
	n := byte(0)
	i := byte(0)
	for s.Scan() {
		line := s.Text()
		n = byte(len(line))
		for j := range line {
			if line[j] != '.' {
				antennas = append(antennas, []byte{byte(j), i, line[j]})
			}
		}
		i++
	}
	part1(antennas, n)
	part2(antennas, n)
}

func part1(antennas [][]byte, n byte) {
	locas := make(map[[2]byte]int)
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			if antennas[i][2] != antennas[j][2] {
				continue
			}
			ax := antennas[i][0]
			ay := antennas[i][1]
			bx := antennas[j][0]
			by := antennas[j][1]
			dx := bx - ax
			dy := by - ay
			n1 := [2]byte{ax - dx, ay - dy}
			if n1[0] >= 0 && n1[0] < n && n1[1] >= 0 && n1[1] < n {
				locas[n1] = 1
			}
			n1 = [2]byte{bx + dx, by + dy}
			if n1[0] >= 0 && n1[0] < n && n1[1] >= 0 && n1[1] < n {
				locas[n1] = 1
			}
		}
	}
	fmt.Println(len(locas))
}

func part2(antennas [][]byte, n byte) {
	locas := make(map[[2]byte]int)
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			if antennas[i][2] != antennas[j][2] {
				continue
			}
			ax := antennas[i][0]
			ay := antennas[i][1]
			bx := antennas[j][0]
			by := antennas[j][1]
			dx := bx - ax
			dy := by - ay
			n1 := [2]byte{ax, ay}
			for n1[0] >= 0 && n1[0] < n && n1[1] >= 0 && n1[1] < n {
				locas[n1] = 1
				n1[0] -= dx
				n1[1] -= dy
			}
			n1 = [2]byte{bx, by}
			for n1[0] >= 0 && n1[0] < n && n1[1] >= 0 && n1[1] < n {
				locas[n1] = 1
				n1[0] += dx
				n1[1] += dy
			}
		}
	}
	fmt.Println(len(locas))
}
