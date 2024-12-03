package main

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func day_three_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_three.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_three.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func day_three_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_three.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_three.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
