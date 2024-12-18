package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_seventeen_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_seventeen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen_example_1.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen_example_2.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen_example_3.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen_example_4.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen_example_5.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen_example_6.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	registers := []int{}
	instruction_set := []int{}
	interpret_instructions := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			interpret_instructions = true
			continue
		}
		if !interpret_instructions {
			split_chars := strings.Split(scanner.Text(), ": ")
			reg_val, _ := strconv.Atoi(split_chars[1])
			registers = append(registers, reg_val)
		} else {
			split_chars := strings.Split(scanner.Text(), ": ")
			for _, char := range strings.Split(split_chars[1], ",") {
				instruction, _ := strconv.Atoi(char)
				instruction_set = append(instruction_set, instruction)
			}
		}
	}
	i := 0
	values := []int{}
	for i < len(instruction_set)-1 {
		switch instruction_set[i] {
		case 0:
			divisor := 1
			combo := find_combo(instruction_set[i+1], registers)
			for j := 0; j < combo; j++ {
				divisor = divisor * 2
			}
			registers[0] = registers[0] / divisor
		case 1:
			registers[1] = registers[1] ^ instruction_set[i+1]
		case 2:
			combo := find_combo(instruction_set[i+1], registers)
			registers[1] = combo & 7
		case 3:
			if registers[0] != 0 {
				i = instruction_set[i+1]
				continue
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			combo := find_combo(instruction_set[i+1], registers)
			values = append(values, combo&7)
		case 6:
			divisor := 1
			combo := find_combo(instruction_set[i+1], registers)
			for j := 0; j < combo; j++ {
				divisor = divisor * 2
			}
			registers[1] = registers[0] / divisor
		case 7:
			divisor := 1
			combo := find_combo(instruction_set[i+1], registers)
			for j := 0; j < combo; j++ {
				divisor = divisor * 2
			}
			registers[2] = registers[0] / divisor
		}
		i += 2
	}
	str_output := ""
	for i := 0; i < len(values); i++ {
		str_output += strconv.Itoa(values[i])
		if i != len(values)-1 {
			str_output += ","
		}
	}
	fmt.Println(str_output)
	// fmt.Println(values)
}

func Day_seventeen_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_seventeen.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_seventeen.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	registers := []int{}
	instruction_set := []int{}
	interpret_instructions := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			interpret_instructions = true
			continue
		}
		if !interpret_instructions {
			split_chars := strings.Split(scanner.Text(), ": ")
			reg_val, _ := strconv.Atoi(split_chars[1])
			registers = append(registers, reg_val)
		} else {
			split_chars := strings.Split(scanner.Text(), ": ")
			for _, char := range strings.Split(split_chars[1], ",") {
				instruction, _ := strconv.Atoi(char)
				instruction_set = append(instruction_set, instruction)
			}
		}
	}
	i := 0
	values := []int{}
	for i < len(instruction_set)-1 {
		switch instruction_set[i] {
		case 0:
			divisor := 1
			combo := find_combo(instruction_set[i+1], registers)
			for j := 0; j < combo; j++ {
				divisor = divisor * 2
			}
			registers[0] = registers[0] / divisor
		case 1:
			registers[1] = registers[1] ^ instruction_set[i+1]
		case 2:
			combo := find_combo(instruction_set[i+1], registers)
			registers[1] = combo & 7
		case 3:
			if registers[0] != 0 {
				i = instruction_set[i+1]
				continue
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			combo := find_combo(instruction_set[i+1], registers)
			values = append(values, combo&7)
		case 6:
			divisor := 1
			combo := find_combo(instruction_set[i+1], registers)
			for j := 0; j < combo; j++ {
				divisor = divisor * 2
			}
			registers[1] = registers[0] / divisor
		case 7:
			divisor := 1
			combo := find_combo(instruction_set[i+1], registers)
			for j := 0; j < combo; j++ {
				divisor = divisor * 2
			}
			registers[2] = registers[0] / divisor
		}
		i += 2
	}
	str_output := ""
	for i := 0; i < len(values); i++ {
		str_output += strconv.Itoa(values[i])
		if i != len(values)-1 {
			str_output += ","
		}
	}
	// fmt.Println(str_output)
	// fmt.Println(values)
}

func find_combo(combo int, registers []int) int {
	output_combo := 0
	switch combo {
	case 0:
		output_combo = 0
	case 1:
		output_combo = 1
	case 2:
		output_combo = 2
	case 3:
		output_combo = 3
	case 4:
		output_combo = registers[0]
	case 5:
		output_combo = registers[1]
	case 6:
		output_combo = registers[2]
	}
	return output_combo
}
