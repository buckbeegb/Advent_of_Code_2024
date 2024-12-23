package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type towel_subset struct {
	towel         string
	valid_methods int
	parents       []towel_subset
	children      []towel_subset
}

var valid_towels map[string]bool
var cached_towels map[string]int

func Day_nineteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_nineteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_nineteen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	towel_sets := []string{}
	set_valid := true
	max_len := 0
	valid_towels = make(map[string]bool)
	cached_towels = make(map[string]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			set_valid = false
			continue
		}
		if set_valid {
			for _, towel := range strings.Split(scanner.Text(), ", ") {
				if len(towel) > max_len {
					max_len = len(towel)
				}
				valid_towels[towel] = true
			}
		} else {
			towel_sets = append(towel_sets, scanner.Text())
		}
	}
	possibility_sum := 0
	for _, ts := range towel_sets {
		if eval_towel(ts, max_len) > 0 {
			possibility_sum += 1
		}
	}
	fmt.Println(possibility_sum)
}

func Day_nineteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_nineteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_nineteen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	set_valid := true
	max_len := 0
	towel_sets := []string{}
	valid_towels = make(map[string]bool)
	cached_towels = make(map[string]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			set_valid = false
			continue
		}
		if set_valid {
			for _, towel := range strings.Split(scanner.Text(), ", ") {
				if len(towel) > max_len {
					max_len = len(towel)
				}
				valid_towels[towel] = true
			}
		} else {
			towel_sets = append(towel_sets, scanner.Text())
		}
	}
	possibility_sum := 0
	for _, ts := range towel_sets {
		possibility_sum += eval_towel(ts, max_len)
	}
	fmt.Println(possibility_sum)
}

func eval_towel(towel string, max_towel_size int) int {
	if len(towel) == 0 {
		return 1
	}
	if val, found := cached_towels[towel]; found {
		return val
	}
	num_valid := 0
	for i := max_towel_size; i > 0; i-- {
		if i > len(towel) {
			continue
		}
		cur_towel := towel[:i]
		if valid_towels[cur_towel] {
			num_valid += eval_towel(towel[i:], max_towel_size)
		}
	}
	cached_towels[towel] = num_valid
	return num_valid
}
