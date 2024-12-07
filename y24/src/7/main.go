package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("res/day7.txt")
	s := bufio.NewScanner(f)
	lines := [][]uint64{}
	for s.Scan() {
		line := s.Text()
		sp := strings.Split(line, ":")
		ops := strings.Fields(sp[1])
		nums := []uint64{}
		a, _ := strconv.ParseUint(sp[0], 10, 64)
		nums = append(nums, a)
		for _, v := range ops {
			a, _ := strconv.ParseUint(v, 10, 64)
			nums = append(nums, a)
		}
		lines = append(lines, nums)
	}

	sum := uint64(0)
	badlines := [][]uint64{}
	for _, ns := range lines {
		target := ns[0]
		ops := ns[1:]
		if operate(ops, 0, target) {
			sum += target
		} else {
			badlines = append(badlines, ns)
		}
	}
	fmt.Println(sum)

	sum2 := sum
	for _, ns := range badlines {
		target := ns[0]
		ops := ns[1:]
		if operate2(ops, 0, target) {
			sum2 += target
		}
	}
	fmt.Println(sum2)
}

func operate(slc []uint64, total uint64, target uint64) bool {
	if len(slc) == 0 {
		return total == target
	} else if total > target {
		return false
	}
	return operate(slc[1:], total+slc[0], target) ||
		operate(slc[1:], total*slc[0], target)
}

func operate2(slc []uint64, total uint64, target uint64) bool {
	if len(slc) == 0 {
		return total == target
	} else if total > target {
		return false
	}
	if operate2(slc[1:], total+slc[0], target) ||
		operate2(slc[1:], total*slc[0], target) {
		return true
	}
	t3s := strconv.FormatUint(total, 10) + strconv.FormatUint(slc[0], 10)
	t3, _ := strconv.ParseUint(t3s, 10, 64)
	return operate2(slc[1:], t3, target)
}
