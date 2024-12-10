package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_nine_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_nine.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_nine.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	raw_storage := []int{}
	unique_file_count := -1
	for scanner.Scan() {
		for i, char := range scanner.Text() {
			var appended_item int
			if i%2 == 0 {
				unique_file_count++
				appended_item = unique_file_count
			} else {
				appended_item = -1
			}
			length, _ := strconv.Atoi(string(char))
			for j := 0; j < length; j++ {
				raw_storage = append(raw_storage, appended_item)
			}
		}
	}
	empty_counter := 0
	full_counter := len(raw_storage) - 1
	for empty_counter < full_counter {
		if raw_storage[full_counter] == -1 {
			full_counter--
		}
		if raw_storage[empty_counter] != -1 {
			empty_counter++
		}
		if full_counter >= 0 && raw_storage[full_counter] != -1 && empty_counter < len(raw_storage) && raw_storage[empty_counter] == -1 {
			raw_storage[empty_counter] = raw_storage[full_counter]
			raw_storage[full_counter] = -1
		}
	}
	pos_sum := 0
	for i := 0; i < len(raw_storage); i++ {
		if raw_storage[i] == -1 {
			break
		}
		pos_sum += raw_storage[i] * i
	}
	fmt.Println(pos_sum)
}

type aocMemory struct {
	id     int
	start  int
	length int
}

func Day_nine_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_nine.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_nine.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	raw_storage := []int{}
	mem_storage := []aocMemory{}
	unique_file_count := -1
	cur_pos := 0
	for scanner.Scan() {
		for i, char := range scanner.Text() {
			var appended_item int
			length, _ := strconv.Atoi(string(char))
			if i%2 == 0 {
				unique_file_count++
				mem_storage = append(mem_storage, aocMemory{id: unique_file_count, start: cur_pos, length: length})
				cur_pos += length
			} else {
				mem_storage = append(mem_storage, aocMemory{id: -1, start: cur_pos, length: length})
				cur_pos += length
			}
			for j := 0; j < length; j++ {
				raw_storage = append(raw_storage, appended_item)
			}
		}
	}
	for i := len(mem_storage) - 1; i >= 0; i-- {
		if mem_storage[i].id == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if mem_storage[j].id == -1 && mem_storage[j].length >= mem_storage[i].length {
				mem_storage[j].length -= mem_storage[i].length
				mem_storage[i].start = mem_storage[j].start
				mem_storage[j].start = mem_storage[j].start + mem_storage[i].length
				break
			}
		}
	}
	checksum := 0
	for i := 0; i < len(mem_storage); i++ {
		if mem_storage[i].id == -1 {
			continue
		}
		for j := 0; j < mem_storage[i].length; j++ {
			checksum += mem_storage[i].id * (mem_storage[i].start + j)
		}
	}
	fmt.Println(checksum)
}
