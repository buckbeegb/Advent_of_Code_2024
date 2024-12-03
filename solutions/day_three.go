package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day_three_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_three.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_three.txt")
	if err != nil {
		panic(err)
	}
	regex := "mul\\([0-9]+,[0-9]+\\)"
	re, _ := regexp.Compile(regex)
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	total_product := 0
	for scanner.Scan() {
		text := scanner.Text()
		matching := re.FindAllStringSubmatch(text, -1)
		for _, match := range matching {
			mul_to_decode := match[0]
			reduced_mul := strings.ReplaceAll(strings.ReplaceAll(mul_to_decode, "mul(", ""), ")", "")
			numbers := strings.Split(reduced_mul, ",")
			left, _ := strconv.Atoi(numbers[0])
			right, _ := strconv.Atoi(numbers[1])
			product := left * right
			total_product += product
		}
	}
	fmt.Println(total_product)
}

func Day_three_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_three.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_three_part_2.txt")
	if err != nil {
		panic(err)
	}
	regex := "mul\\([0-9]+,[0-9]+\\)"
	re, _ := regexp.Compile(regex)
	regex2 := "don't\\(\\)"
	redont, _ := regexp.Compile(regex2)
	regex3 := "do\\(\\)"
	redo, _ := regexp.Compile(regex3)
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	total_product := 0
	end_line_state := true
	for scanner.Scan() {
		text := scanner.Text()
		matching := re.FindAllStringSubmatch(text, -1)
		matching_range := re.FindAllStringSubmatchIndex(text, -1)
		dos := redo.FindAllStringSubmatchIndex(text, -1)
		donts := redont.FindAllStringSubmatchIndex(text, -1)
		// fmt.Println(dos)
		// fmt.Println(donts)
		// fmt.Println(matching_range)
		for i := 0; i < len(matching); i++ {
			position_compare_execution := matching_range[i][0]
			var closest_do int
			var closest_dont int
			if end_line_state {
				closest_do = position_compare_execution
				closest_dont = 1000000
			} else {
				closest_dont = position_compare_execution
				closest_do = 1000000
			}
			for j := 0; j < len(dos); j++ {
				if dos[j][0] < position_compare_execution && position_compare_execution-dos[j][0] < closest_do {
					closest_do = position_compare_execution - dos[j][0]
				}
			}
			for j := 0; j < len(donts); j++ {
				if donts[j][0] < position_compare_execution && position_compare_execution-donts[j][0] < closest_dont {
					closest_dont = position_compare_execution - donts[j][0]
				}
			}
			if closest_do < closest_dont {
				mul_to_decode := matching[i][0]
				// fmt.Println(mul_to_decode)
				reduced_mul := strings.ReplaceAll(strings.ReplaceAll(mul_to_decode, "mul(", ""), ")", "")
				numbers := strings.Split(reduced_mul, ",")
				left, _ := strconv.Atoi(numbers[0])
				right, _ := strconv.Atoi(numbers[1])
				product := left * right
				total_product += product
			}
		}
		last_dont := donts[len(donts)-1][0]
		last_do := dos[len(dos)-1][0]
		if last_do > last_dont {
			end_line_state = true
		} else {
			end_line_state = false
		}
	}
	fmt.Println(total_product)
}
