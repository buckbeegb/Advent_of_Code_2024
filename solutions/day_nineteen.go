package solutions

import (
	"bufio"
	// "fmt"
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

func Day_nineteen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_nineteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_nineteen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	valid_towels := make(map[string]bool)
	towel_sets := []string{}
	set_valid := true
	max_len := 0
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
	// towel_subsets := make(map[string]towel_subset)
	// total_valid := 0
	// for i := 0; i < len(towel_sets); i++ {
	// 	_, found := towel_subsets[towel_sets[i]]
	// 	if !found {
	// 		towel_subsets[towel_sets[i]] = towel_subset{towel: towel_sets[i], valid_methods: 0}
	// 	}
	// 	search_strings := []string{towel_sets[i]}
	// 	complete_match := false
	// 	for len(search_strings) > 0 {
	// 		addl_strings := []string{}
	// 		for k := max_len; k > 0; k-- {
	// 			if k > len(search_strings[0]) {
	// 				continue
	// 			}
	// 			substring := search_strings[0][:k]
	// 			subset, found := towel_subsets[substring]
	// 			if found && subset.valid_methods > 0 {
	// 				total_valid++
	// 			}
	// 			if valid_towels[search_strings[0][:k]] {
	// 				if k == len(search_strings[0]) {
	// 					towel_subsets[search_strings[0]] = towel_subset{towel: search_strings[0], valid_methods: 1}
	// 				}
	// 				addl_strings = append(addl_strings, search_strings[0][k:])
	// 			}
	// 			if complete_match {
	// 				break
	// 			}
	// 		}
	// 		if complete_match {
	// 			break
	// 		}
	// 		search_strings = append(addl_strings, search_strings[1:]...)
	// 	}
	// 	if complete_match {
	// 		fmt.Printf("Valid: %s\n", towel_sets[i])
	// 		total_valid++
	// 	}
	// }
	// fmt.Println(total_valid)
}
func Day_nineteen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_nineteen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_nineteen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}
