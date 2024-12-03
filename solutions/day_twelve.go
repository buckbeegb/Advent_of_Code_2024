package solutions

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_twelve_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twelve.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twelve.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func Day_twelve_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twelve.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twelve.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
