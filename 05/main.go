package main

// Jumps are relative: -1 moves to the previous instruction, and 2 skips the next
// one. Start at the first instruction in the list. The goal is to follow the
// jumps until one leads outside the list.
//
// --- Part Two ---
// After each jump, if the offset was three or more, instead decrease it by 1.
// Otherwise, increase it by 1 as before.

// Answer part one: 381680
// Answer part two: 29717847

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	fmt.Println("Steps to reach the Maze exit:", exitMaze(f))

	_, err = f.Seek(0, 0)
	fmt.Println("Steps to reach the Maze exit (part two):", exitMazeTwo(f))
}

func exitMaze(f *os.File) int {
	scanner := bufio.NewScanner(f)
	jumps := []int{}
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())

		n, _ := strconv.Atoi(strings.Join(r, ""))
		jumps = append(jumps, n)
	}

	steps, i, pi := 0, 0, 0
	for {
		steps++
		i += jumps[i]

		// previous index
		jumps[pi]++
		pi = i

		if i < 0 || i > len(jumps)-1 {
			// exit reached
			break
		}

	}

	return steps
}

func exitMazeTwo(f *os.File) int {
	scanner := bufio.NewScanner(f)
	jumps := []int{}
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())

		n, _ := strconv.Atoi(strings.Join(r, ""))
		jumps = append(jumps, n)
	}

	steps, i, pi := 0, 0, 0
	for {
		steps++
		i += jumps[i]

		// previous index
		if math.Abs(float64(i-pi)) > 2 {
			if math.Signbit(float64(jumps[pi])) {
				jumps[pi]++
			} else {
				jumps[pi]--
			}
		} else {
			jumps[pi]++
		}
		pi = i

		if i < 0 || i > len(jumps)-1 {
			// exit reached
			break
		}

	}

	return steps
}

// eof
