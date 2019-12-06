package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getOrbits(orbiting map[string]string, start string) []string {
	var orbits []string
	for next, existing := orbiting[start]; existing; {
		orbits = append(orbits, next)
		next, existing = orbiting[next]
	}

	return orbits
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	orbiting := make(map[string]string)
	for scanner.Scan() {
		s := scanner.Text()
		els := strings.Split(s, ")")
		orbiting[els[1]] = els[0]
	}

	youOrbits := getOrbits(orbiting, "YOU")
	sanOrbits := getOrbits(orbiting, "SAN")
	dist := 0
orbTrace:
	for youIdx, youOrb := range youOrbits {
		for sanIdx, sanOrb := range sanOrbits {
			if youOrb == sanOrb {
				dist = youIdx + sanIdx
				break orbTrace
			}
		}
	}
	fmt.Println(dist)
}
