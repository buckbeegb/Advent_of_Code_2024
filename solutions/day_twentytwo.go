package solutions

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_twentytwo_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentytwo.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentytwo.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func Day_twentytwo_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentytwo.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentytwo.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
