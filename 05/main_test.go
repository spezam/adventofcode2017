package main

import (
	"fmt"
	"os"
	"testing"
)

var tt = []struct {
	name      string
	inputFile string
	steps     int
}{
	{"Maze exit", "input_test", 5},
	{"Maze exit two", "input_test", 10},
}

func TestMaze(t *testing.T) {
	f, err := os.Open(tt[0].inputFile)
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	e := ExitMaze(f)
	if e != tt[0].steps {
		t.Errorf("Steps should be %d, got: %d", tt[0].steps, e)
	}
}

func TestMazeTwo(t *testing.T) {
	f, err := os.Open(tt[1].inputFile)
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	e := ExitMazeTwo(f)
	if e != tt[1].steps {
		t.Errorf("Steps should be %d, got: %d", tt[1].steps, e)
	}
}
