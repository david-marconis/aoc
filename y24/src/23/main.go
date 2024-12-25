package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("----------NEW-----------")
	f, _ := os.Open("res/day23.txt")
	s := bufio.NewScanner(f)
	m := make(map[string]map[string]struct{})
	for s.Scan() {
		line := s.Text()
		sp := strings.Split(line, "-")
		add(m, sp[0], sp[1])
		add(m, sp[1], sp[0])
	}
	nodes := slices.Collect(maps.Keys(m))
	slices.Sort(nodes)
	nets := [][]string{}
	newNets := [][]string{}
	for _, n := range nodes {
		newNets = append(newNets, []string{n})
	}
	for len(newNets) > 0 {
		nets = newNets
		if len(nets[0]) == 3 {
			sum := 0
		fornet:
			for _, net := range nets {
				for _, n := range net {
					if n[0] == 't' {
						sum += 1
						continue fornet
					}
				}
			}
			fmt.Println(sum)
		}
		newNets = [][]string{}
		for _, node := range nodes {
		fornets:
			for _, net := range nets {
				if net[len(net)-1] >= node {
					continue
				}
				for _, o := range net {
					_, ok1 := m[o][node]
					_, ok2 := m[node][o]
					if !ok1 || !ok2 {
						continue fornets
					}
				}
				newNet := make([]string, len(net)+1)
				copy(newNet, net)
				newNet[len(net)] = node
				newNets = append(newNets, newNet)
			}
		}
	}
	fmt.Println(strings.Join(nets[0], ","))
}

func add(m map[string]map[string]struct{}, s1, s2 string) {
	var s map[string]struct{}
	if l, ok := m[s1]; ok {
		s = l
	} else {
		s = make(map[string]struct{}, 0)
	}
	s[s2] = struct{}{}
	m[s1] = s
}
