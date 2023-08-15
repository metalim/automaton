package main

import "fmt"

const binRule = "01101110" // 110
const N = 300

func main() {
	state := make([]bool, N)
	state[N-1] = true
	for i := 0; i < N; i++ {
		printStateCompact(state)
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
	for _, cell := range state {
		if cell {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

var prevState []bool

func printStateCompact(state []bool) {
	if prevState == nil {
		prevState = state
		return
	}

	for i := 0; i < len(state); i += 2 {
		tl := prevState[i]
		tr := prevState[i+1]
		bl := state[i]
		br := state[i+1]
		switch {
		case tl && tr && bl && br:
			fmt.Print("█")
		case tl && tr && bl && !br:
			fmt.Print("▛")
		case tl && tr && !bl && br:
			fmt.Print("▜")
		case !tl && tr && bl && br:
			fmt.Print("▟")
		case tl && !tr && bl && br:
			fmt.Print("▙")
		case tl && tr && !bl && !br:
			fmt.Print("▀")
		case !tl && !tr && bl && br:
			fmt.Print("▄")
		case tl && !tr && bl && !br:
			fmt.Print("▌")
		case !tl && tr && !bl && br:
			fmt.Print("▐")
		case !tl && tr && bl && !br:
			fmt.Print("▞")
		case tl && !tr && !bl && br:
			fmt.Print("▚")
		case tl && !tr && !bl && !br:
			fmt.Print("▘")
		case !tl && tr && !bl && !br:
			fmt.Print("▝")
		case !tl && !tr && bl && !br:
			fmt.Print("▖")
		case !tl && !tr && !bl && br:
			fmt.Print("▗")
		case !bl && !br && !tl && !tr:
			fmt.Print(" ")
		}
	}
	fmt.Println()
	prevState = nil
}
