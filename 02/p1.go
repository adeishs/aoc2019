package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(is *[]int) {
	for i := 0; i < len(*is); i = i + 4 {
		if (*is)[i] == 99 {
			return
		}

		p := i + 1
		op1 := (*is)[(*is)[p]]
		p++
		op2 := (*is)[(*is)[p]]
		p++
		dest := (*is)[p]
		if (*is)[i] == 1 {
			(*is)[dest] = op1 + op2
		} else if (*is)[i] == 2 {
			(*is)[dest] = op1 * op2
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	s := ""
	for scanner.Scan() {
		s += scanner.Text()
	}
	int_strs := strings.Split(s, ",")

	var is []int
	for i := range int_strs {
		t, _ := strconv.Atoi(int_strs[i])
		is = append(is, t)
	}
	is[1] = 12
	is[2] = 2
	run(&is)
	fmt.Println(is[0])
}
