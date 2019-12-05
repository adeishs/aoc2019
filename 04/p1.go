package main

import (
	"fmt"
)

func main() {
	min := 372304
	max := 847060

	cnt := 0
	for p := min; p <= max; p++ {
		password := p
		digits := make([]int, 6)
		d := 5
		for password > 0 {
			digits[d] = password % 10
			password /= 10
			d--
		}

		isValid := true
		for d = 0; d < len(digits)-1; d++ {
			if digits[d] > digits[d+1] {
				isValid = false
				break
			}
		}

		if !isValid {
			continue
		}

		for d = 0; d < len(digits)-1; d++ {
			if digits[d] == digits[d+1] {
				cnt++
				break
			}
		}
	}

	fmt.Println(cnt)
}
