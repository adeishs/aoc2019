package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func process(grids *[][]int, wiring string, wire_id int) int {
	mvts := strings.Split(wiring, ",")
	curr_x := 10000
	curr_y := 10000
	min_dist := 2 * (curr_x + curr_y)
	for m := range mvts {
		mvt_dist, _ := strconv.Atoi(mvts[m][1:])
		dir := mvts[m][0:1]

		for d := 0; d < mvt_dist; d++ {
			if dir == "U" {
				curr_y--
			} else if dir == "D" {
				curr_y++
			} else if dir == "L" {
				curr_x--
			} else if dir == "R" {
				curr_x++
			}
			if (*grids)[curr_y][curr_x] == 0 ||
				(*grids)[curr_y][curr_x] == wire_id {
				(*grids)[curr_y][curr_x] = wire_id
			} else {
				cross_dist := abs(curr_x-10000) +
					abs(curr_y-10000)
				if min_dist > cross_dist {
					min_dist = cross_dist
				}
			}
		}
	}

	return min_dist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grids := make([][]int, 20000)
	for y := range grids {
		grids[y] = make([]int, 20000)
	}
	dist := 0
	for i := 1; i <= 2; i++ {
		scanner.Scan()
		s := scanner.Text()
		dist = process(&grids, s, i)
	}
	fmt.Println(dist)
}
