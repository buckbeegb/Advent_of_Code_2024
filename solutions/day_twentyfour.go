package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Proc_Action struct {
	operation string
	regA      string
	regB      string
	outA      string
}

func Day_twentyfour_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentyfour.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentyfour.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	registers := make(map[string]int)
	total_registers := make(map[string]bool)
	actions := []Proc_Action{}
	xyz_vals := [][]string{{}, {}, {}}
	record_actions := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			record_actions = true
			continue
		}
		if record_actions {
			str_rem_arrow := strings.Replace(scanner.Text(), " ->", "", -1)
			split_str := strings.Split(str_rem_arrow, " ")
			actions = append(actions, Proc_Action{operation: split_str[1], regA: split_str[0], regB: split_str[2], outA: split_str[3]})
			for i := range []int{0, 1, 2, 3} {
				if i == 1 {
					continue
				}
				total_registers[split_str[i]] = true
				if rune(split_str[i][0]) == 'z' {
					xyz_vals[2] = append(xyz_vals[2], split_str[i])
				} else if rune(split_str[i][0]) == 'y' {
					xyz_vals[1] = append(xyz_vals[1], split_str[i])
				} else if rune(split_str[i][0]) == 'x' {
					xyz_vals[0] = append(xyz_vals[0], split_str[i])
				}
			}
		} else {
			split_str := strings.Split(scanner.Text(), ": ")
			registers[split_str[0]], _ = strconv.Atoi(split_str[1])
		}
	}
	sort.Strings(xyz_vals[0])
	sort.Strings(xyz_vals[1])
	sort.Strings(xyz_vals[2])
	output := 0
	multiplier := 1
	c_reg := evaluate_wire_pairs(duplicate_registers(registers), actions, len(total_registers))
	for _, z_val := range xyz_vals[2] {
		output += (c_reg[z_val] * multiplier)
		multiplier *= 2
	}
	fmt.Println(output)
}

func Day_twentyfour_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentyfour.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentyfour.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	registers := make(map[string]int)
	total_registers := make(map[string]bool)
	actions := []Proc_Action{}
	xy_action_lookup := make(map[string]Proc_Action)
	xyz_vals := [][]string{{}, {}, {}}
	record_actions := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			record_actions = true
			continue
		}
		if record_actions {
			str_rem_arrow := strings.Replace(scanner.Text(), " ->", "", -1)
			split_str := strings.Split(str_rem_arrow, " ")
			actions = append(actions, Proc_Action{operation: split_str[1], regA: split_str[0], regB: split_str[2], outA: split_str[3]})
			xy_action_lookup[strings.Join(split_str[:3], " ")] = actions[len(actions)-1]
			for i := range []int{0, 1, 2, 3} {
				if i == 1 {
					continue
				}
				total_registers[split_str[i]] = true
				if rune(split_str[i][0]) == 'z' {
					xyz_vals[2] = append(xyz_vals[2], split_str[i])
				}
			}
		} else {
			split_str := strings.Split(scanner.Text(), ": ")
			registers[split_str[0]], _ = strconv.Atoi(split_str[1])
		}
	}
	sort.Strings(xyz_vals[2])
	for i := 0; i < len(xyz_vals[2]); i++ {
		xyz_vals[0] = append(xyz_vals[0], "x"+xyz_vals[2][i][1:])
		xyz_vals[1] = append(xyz_vals[1], "y"+xyz_vals[2][i][1:])
	}
	sort.Strings(xyz_vals[0])
	sort.Strings(xyz_vals[1])
	swap_list := []string{}
	for i := 1; i < len(xyz_vals[2]); i++ {
		xor, and, xi, ai := find_relevant_leading_actions(xyz_vals[0][i], actions)
		if xi == 0 {
			continue
		}
		c2_actions, c2i := find_relevant_actions(and.outA, actions)
		o_actions, oi := find_relevant_actions(xor.outA, actions)
		if len(c2_actions) > 1 {
			actions[xi].outA = and.outA
			actions[ai].outA = xor.outA
			swap_list = append(swap_list, and.outA, xor.outA)
			i--
			continue
		}
		var oxor, oand Proc_Action
		var oxori, oandi int
		if o_actions[0].operation == "XOR" {
			oxor, oand = o_actions[0], o_actions[1]
			oxori, oandi = oi[0], oi[1]
		} else {
			oxor, oand = o_actions[1], o_actions[0]
			oxori, oandi = oi[1], oi[0]
		}
		if len(c2_actions) == 0 {
			actions[ai].outA = oxor.outA
			actions[oxori].outA = and.outA
			swap_list = append(swap_list, oxor.outA, and.outA)
			i--
			continue
		}
		if c2_actions[0].outA == xyz_vals[2][i] {
			actions[c2i[0]].outA = oxor.outA
			actions[oxori].outA = c2_actions[0].outA
			swap_list = append(swap_list, oxor.outA, c2_actions[0].outA)
			i--
			continue
		}
		if oand.outA == xyz_vals[2][i] {
			actions[oxori].outA = oand.outA
			actions[oandi].outA = oxor.outA
			swap_list = append(swap_list, oand.outA, oxor.outA)
			i--
			continue
		}
	}
	sort.Strings(swap_list)
	fmt.Println(strings.Join(swap_list, ","))
}

func evaluate_wire_pairs(registers map[string]int, actions []Proc_Action, total_registers_len int) map[string]int {
	for len(registers) < total_registers_len {
		for _, action := range actions {
			a, found_a := registers[action.regA]
			b, found_b := registers[action.regB]
			if !found_a || !found_b {
				continue
			}
			if action.operation == "AND" {
				registers[action.outA] = a & b
			} else if action.operation == "OR" {
				registers[action.outA] = a | b
			} else if action.operation == "XOR" {
				registers[action.outA] = a ^ b
			}
		}
	}
	return registers
}

func find_relevant_leading_actions(reg string, actions []Proc_Action) (Proc_Action, Proc_Action, int, int) {
	var xor Proc_Action
	var and Proc_Action
	xi := 0
	ai := 0
	for i := 0; i < len(actions); i++ {
		if reg == actions[i].regA || reg == actions[i].regB {
			if actions[i].operation == "XOR" {
				xor = actions[i]
				xi = i
			} else {
				and = actions[i]
				ai = i
			}
		}
	}
	return xor, and, xi, ai
}

func find_relevant_actions(reg string, actions []Proc_Action) ([]Proc_Action, []int) {
	action_list := []Proc_Action{}
	action_count := []int{}
	for i := 0; i < len(actions); i++ {
		if reg == actions[i].regA || reg == actions[i].regB {
			action_list = append(action_list, actions[i])
			action_count = append(action_count, i)
		}
	}
	return action_list, action_count
}

func duplicate_registers(registers map[string]int) map[string]int {
	dup := make(map[string]int)
	for key, value := range registers {
		dup[key] = value
	}
	return dup
}

func determine_validity(xyz_vals [][]string, c_reg map[string]int) bool {
	output := []int{0, 0, 0}
	for i := 0; i < len(xyz_vals); i++ {
		multiplier := 1
		for _, z_val := range xyz_vals[i] {
			output[i] += (c_reg[z_val] * multiplier)
			multiplier *= 2
		}
	}
	return (output[0] + output[1]) == output[2]
}
