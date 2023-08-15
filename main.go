package main

import "fmt"

const binRule = "01101110" // 110

func main() {
	state := make([]bool, 150)
	state[149] = true
	for i := 0; i < 150; i++ {
		printState(state)
		state = nextState(state)
	}
}

func nextState(state []bool) []bool {
	next := make([]bool, len(state))
	for i := 0; i < len(state); i++ {
		next[i] = nextStateFor(i, state)
	}
	return next
}

func nextStateFor(i int, state []bool) bool {
	left := state[mod(i-1, len(state))]
	mid := state[i]
	right := state[mod(i+1, len(state))]
	return ruleFor(left, mid, right)
}

func ruleFor(left, mid, right bool) bool {
	if left && mid && right {
		return binRule[0] == '1'
	}
	if left && mid && !right {
		return binRule[1] == '1'
	}
	if left && !mid && right {
		return binRule[2] == '1'
	}
	if left && !mid && !right {
		return binRule[3] == '1'
	}
	if !left && mid && right {
		return binRule[4] == '1'
	}
	if !left && mid && !right {
		return binRule[5] == '1'
	}
	if !left && !mid && right {
		return binRule[6] == '1'
	}
	if !left && !mid && !right {
		return binRule[7] == '1'
	}
	return false
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func printState(state []bool) {
	for i := 0; i < len(state); i++ {
		if state[i] {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
