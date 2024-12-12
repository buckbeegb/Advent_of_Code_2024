package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day_eleven_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_eleven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eleven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	unique_rocks := make(map[int]int)
	for scanner.Scan() {
		split_input := strings.Split(scanner.Text(), " ")
		for _, item := range split_input {
			rock, _ := strconv.Atoi(item)
			unique_rocks[rock]++
		}
	}
	for i := 0; i < 25; i++ {
		new_unique_rocks := make(map[int]int)
		for rock := range unique_rocks {
			if rock == 0 {
				new_unique_rocks[1] += unique_rocks[0]
				continue
			}
			length := math.Log10(float64(rock))
			num_digits := int(length) + 1
			exp := int(math.Pow(10, float64(num_digits/2)))
			if num_digits%2 == 0 {
				new_rock := rock / exp
				new_unique_rocks[new_rock] += unique_rocks[rock]
				new_unique_rocks[rock-(new_rock*exp)] += unique_rocks[rock]
			} else {
				new_unique_rocks[rock*2024] += unique_rocks[rock]
			}
		}
		unique_rocks = new_unique_rocks
	}
	total_rocks := 0
	for key := range unique_rocks {
		total_rocks += unique_rocks[key]
	}
	fmt.Println(total_rocks)
}

func Day_eleven_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_eleven.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_eleven.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	unique_rocks := make(map[int]int)
	for scanner.Scan() {
		split_input := strings.Split(scanner.Text(), " ")
		for _, item := range split_input {
			rock, _ := strconv.Atoi(item)
			unique_rocks[rock] += 1
		}
	}
	for i := 0; i < 75; i++ {
		new_unique_rocks := make(map[int]int)
		for rock := range unique_rocks {
			if rock == 0 {
				new_unique_rocks[1] += unique_rocks[0]
				continue
			}
			length := math.Log10(float64(rock))
			num_digits := int(length) + 1
			exp := int(math.Pow(10, float64(num_digits/2)))
			if num_digits%2 == 0 {
				new_rock := rock / exp
				new_unique_rocks[new_rock] += unique_rocks[rock]
				new_unique_rocks[rock-(new_rock*exp)] += unique_rocks[rock]
			} else {
				new_unique_rocks[rock*2024] += unique_rocks[rock]
			}
		}
		unique_rocks = new_unique_rocks
	}
	total_rocks := 0
	for key := range unique_rocks {
		total_rocks += unique_rocks[key]
	}
	fmt.Println(total_rocks)
}
