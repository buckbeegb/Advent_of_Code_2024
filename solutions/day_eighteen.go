package solutions

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_eighteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_eighteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eighteen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func Day_eighteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_eighteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eighteen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
