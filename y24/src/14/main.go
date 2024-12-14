package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("------------NEW---------------")
	f, _ := os.Open("res/day14.txt")
	s := bufio.NewScanner(f)
	robots := [][4]int{}
	for s.Scan() {
		line := s.Text()
		sp := strings.Fields(line)
		robot := [4]int{}
		for i := 0; i < 2; i++ {
			p := strings.Split(sp[i], ",")
			robot[i*2], _ = strconv.Atoi(p[0][2:])
			robot[i*2+1], _ = strconv.Atoi(p[1])
		}
		robots = append(robots, robot)
	}
	width := 11
	height := 7
	qs := runSim(robots, width, height, 100)
	fmt.Println(((8 % 4) + 4) % 4)

	sum := 1
	for _, q := range qs {
		fmt.Println(q)
		sum *= q
	}
	fmt.Println("Part1: ", sum)
	fmt.Println("Press enter to start part 2 scanning. Enter to exit")
	bufio.NewScanner(os.Stdin).Scan()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	var b []byte = make([]byte, 1)
	i := 1
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		runSim(robots, width, height, i)
		fmt.Println("right: +1, left: -1, up: +width, down -width")
		os.Stdin.Read(b)
		if b[0] == 68 {
			i--
		} else if b[0] == 67 {
			i++
		} else if b[0] == 65 {
			i += width
		} else if b[0] == 66 {
			i -= width
		}
		if b[0] == 10 {
			break
		}
	}
}

func runSim(robots [][4]int, width int, height int, n int) [4]int {
	grid := [][]rune{}
	for i := 0; i < height; i++ {
		grid = append(grid, make([]rune, width))
		for j := 0; j < width; j++ {
			grid[i][j] = ' '
		}
	}
	qs := [4]int{}
	for _, r := range robots {
		x, y, dx, dy := r[0], r[1], r[2], r[3]
		x = (((x + dx*n) % width) + width) % width
		y = (((y + dy*n) % height) + height) % height
		if x < width/2 && y < height/2 {
			qs[0]++
		} else if x > width/2 && y < height/2 {
			qs[1]++
		} else if x < width/2 && y > height/2 {
			qs[2]++
		} else if x > width/2 && y > height/2 {
			qs[3]++
		}
		grid[y][x] = '■'
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
	fmt.Println(n)
	fmt.Println()
	return qs
}
