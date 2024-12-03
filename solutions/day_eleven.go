package solutions

import (
	"bufio"
	// "fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_eleven_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_eleven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eleven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
func Day_eleven_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_eleven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eleven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
