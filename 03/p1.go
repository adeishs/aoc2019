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

func process(grids *[][]int, wiring string, wireId int) int {
	mvts := strings.Split(wiring, ",")
	currX := 10000
	currY := 10000
	minDist := 2 * (currX + currY)
	for m := range mvts {
		mvtDist, _ := strconv.Atoi(mvts[m][1:])
		dir := mvts[m][0:1]

		for d := 0; d < mvtDist; d++ {
			if dir == "U" {
				currY--
			} else if dir == "D" {
				currY++
			} else if dir == "L" {
				currX--
			} else if dir == "R" {
				currX++
			}
			if (*grids)[currY][currX] == 0 ||
				(*grids)[currY][currX] == wireId {
				(*grids)[currY][currX] = wireId
			} else {
				crossDist := abs(currX-10000) +
					abs(currY-10000)
				if minDist > crossDist {
					minDist = crossDist
				}
			}
		}
	}

	return minDist
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
