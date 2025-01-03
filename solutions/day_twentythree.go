package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	// "strconv"
	"strings"
)

func Day_twentythree_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentythree.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentythree.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	computers := make(map[string]map[string]int)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "-")
		_, found := computers[split[0]]
		if !found {
			computers[split[0]] = make(map[string]int)
			computers[split[0]][split[1]] = 1
		} else {
			computers[split[0]][split[1]] = 1
		}
		_, found = computers[split[1]]
		if !found {
			computers[split[1]] = make(map[string]int)
			computers[split[1]][split[0]] = 1
		} else {
			computers[split[1]][split[0]] = 1
		}
	}
	unique_sets := make(map[string]bool)
	for key, value := range computers {
		if string(key[0]) == "t" {
			key_list := make([]string, 0, len(value))
			for key2 := range value {
				key_list = append(key_list, key2)
			}
			for i := 0; i < len(key_list)-1; i++ {
				for j := 0; j < len(key_list); j++ {
					if computers[key_list[i]][key_list[j]] == 1 {
						strs := []string{key, key_list[i], key_list[j]}
						sort.Strings(strs)
						strsj := strings.Join(strs, "")
						unique_sets[strsj] = true
					}
				}
			}
		}
	}
	fmt.Println(len(unique_sets))
}

func Day_twentythree_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentythree.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentythree.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	computers := [][]bool{}
	found_comps := []int{}
	conversion := make([]string, 676)
	for i := 0; i < 676; i++ {
		computers = append(computers, make([]bool, 676))
		found_comps = append(found_comps, i)
	}
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "-")
		first_lookup := int((26 * (split[0][0] - 'a')) + (split[0][1] - 'a'))
		second_lookup := int((26 * (split[1][0] - 'a')) + (split[1][1] - 'a'))
		computers[first_lookup][second_lookup] = true
		computers[second_lookup][first_lookup] = true
		conversion[first_lookup] = split[0]
		conversion[second_lookup] = split[1]
	}
	max_clique := make(map[int]bool)
	// cliques := []map[int]bool{}
	for _, k := range found_comps {
		clique := make(map[int]bool)
		seen_vertices := make(map[int]bool)
		current_vertex := k
		for len(seen_vertices) < len(found_comps) {
			current_contains := true
			for key := range clique {
				if !computers[current_vertex][key] {
					current_contains = false
				}
			}
			if current_contains {
				clique[current_vertex] = true
			}
			seen_vertices[current_vertex] = true
			found_new := false
			for key := range computers[current_vertex] {
				if !seen_vertices[key] {
					current_vertex = key
					found_new = true
				}
			}
			if !found_new {
				for key := range computers {
					if !seen_vertices[key] {
						current_vertex = key
						break
					}
				}
			}
		}
		if len(clique) > len(max_clique) {
			max_clique = clique
		}
	}
	clique_list := make([]string, 0, len(max_clique))
	for key := range max_clique {
		back_convert := conversion[key]
		clique_list = append(clique_list, back_convert)
	}
	sort.Strings(clique_list)
	fmt.Println(strings.Join(clique_list, ","))
}
