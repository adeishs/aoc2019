package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countOrbits(orbiting map[string]string) int {
	numOfOrbits := 0
	for orb := range orbiting {
		for next, existing := orbiting[orb]; existing; {
			next, existing = orbiting[next]
			numOfOrbits++
		}
	}

	return numOfOrbits
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	orbiting := make(map[string]string)
	for scanner.Scan() {
		s := scanner.Text()
		els := strings.Split(s, ")")
		orbiting[els[1]] = els[0]
	}

	numOfOrbits := countOrbits(orbiting)
	fmt.Println(numOfOrbits)
}
