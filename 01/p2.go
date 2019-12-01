package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcFuel(m int) int {
	return m/3 - 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum = 0
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		f := calcFuel(m)
		for f > 0 {
			sum += f
			f = calcFuel(f)
		}
	}
	fmt.Printf("%d\n", sum)
}
