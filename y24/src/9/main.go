package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("res/day9.txt")
	s := bufio.NewScanner(f)
	files1 := [][3]int{}
	files2 := [][4]int{}
	idx := 0
	for s.Scan() {
		line := s.Text()
		file := 1
		for i := range line {
			x := int(line[i]) - '0'
			files1 = append(files1, [3]int{x, file, len(files1) / 2})
			files2 = append(files2, [4]int{x, file, len(files1) / 2, idx})
			idx += x
			file = 1 - file
		}
	}

	part1(files1)
	part2(files2)
}

func part1(files [][3]int) {
	idx := 0
	sum := 0
	fi := len(files) - 1
	for i := 0; i < len(files); i++ {
		if files[i][1] == 1 {
			for j := 0; j < files[i][0]; j++ {
				sum += files[i][2] * idx
				idx++
			}
		} else if files[i][1] == 0 {
			for files[i][0] > 0 && fi > i {
				for files[fi][1] == 0 || files[fi][0] == 0 {
					fi -= 2
				}
				for files[i][0] > 0 && files[fi][0] > 0 {
					sum += idx * files[fi][2]
					idx++
					files[fi][0] -= 1
					files[i][0] -= 1
				}
			}
		}
	}
	fmt.Println(sum)
}

func part2(files [][4]int) {
	sum := 0
	for fi := len(files) - 1; fi > 0; fi -= 2 {
		i := 1
		for ; i < fi && files[i][0] < files[fi][0]; i += 2 {
		}
		if i > fi {
			idx := files[fi][3]
			for j := 0; j < files[fi][0]; j++ {
				sum += files[fi][2] * idx
				idx++
			}
		} else {
			id := files[i][3]
			for files[fi][0] > 0 {
				sum += id * files[fi][2]
				id++
				files[fi][0] -= 1
				files[i][0] -= 1
				files[i][3] += 1
			}
		}
	}
	fmt.Println(sum)
}
