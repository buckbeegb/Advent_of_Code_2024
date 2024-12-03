package solutions

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_seven_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_seven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func Day_seven_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_seven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
