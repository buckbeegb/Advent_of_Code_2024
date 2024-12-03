package solutions

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_eight_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_eight.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eight.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func Day_eight_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_eight.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eight.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
