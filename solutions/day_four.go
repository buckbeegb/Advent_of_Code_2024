package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_four_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_four.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_four.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	var word_search [][]string
	line_number := 0
	for scanner.Scan() {
		word_search = append(word_search, []string{})
		for _, char := range scanner.Text() {
			word_search[line_number] = append(word_search[line_number], string(char))
		}
		line_number++
	}
	mas_letters := []string{"M", "A", "S"}
	matched_xmases := 0
	for row, line := range word_search {
		for col, char := range line {
			if char == "X" {
				valid_xmas := []int{1, 1, 1, 1, 1, 1, 1, 1}
				for i := 1; i < 4; i++ {
					mas_letter := mas_letters[i-1]
					if col-i < 0 {
						valid_xmas[0] = 0
						valid_xmas[1] = 0
						valid_xmas[2] = 0
					}
					if col+i >= len(word_search) {
						valid_xmas[5] = 0
						valid_xmas[6] = 0
						valid_xmas[7] = 0
					}
					if row-i < 0 {
						valid_xmas[1] = 0
						valid_xmas[3] = 0
						valid_xmas[6] = 0
					}
					if row+i >= len(line) {
						valid_xmas[2] = 0
						valid_xmas[4] = 0
						valid_xmas[7] = 0
					}
					if valid_xmas[0] == 1 && word_search[row][col-i] != mas_letter {
						valid_xmas[0] = 0
					}
					if valid_xmas[1] == 1 && word_search[row-i][col-i] != mas_letter {
						valid_xmas[1] = 0
					}
					if valid_xmas[2] == 1 && word_search[row+i][col-i] != mas_letter {
						valid_xmas[2] = 0
					}
					if valid_xmas[3] == 1 && word_search[row-i][col] != mas_letter {
						valid_xmas[3] = 0
					}
					if valid_xmas[4] == 1 && word_search[row+i][col] != mas_letter {
						valid_xmas[4] = 0
					}
					if valid_xmas[5] == 1 && word_search[row][col+i] != mas_letter {
						valid_xmas[5] = 0
					}
					if valid_xmas[6] == 1 && word_search[row-i][col+i] != mas_letter {
						valid_xmas[6] = 0
					}
					if valid_xmas[7] == 1 && word_search[row+i][col+i] != mas_letter {
						valid_xmas[7] = 0
					}
				}
				for _, val := range valid_xmas {
					matched_xmases += val
				}
			}
		}
	}
	fmt.Println(matched_xmases)
}

func Day_four_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_four.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_four.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	var word_search [][]string
	line_number := 0
	for scanner.Scan() {
		word_search = append(word_search, []string{})
		for _, char := range scanner.Text() {
			word_search[line_number] = append(word_search[line_number], string(char))
		}
		line_number++
	}
	matched_xmases := 0
	for row, line := range word_search {
		for col, char := range line {
			if col == 0 || col == len(line)-1 || row == 0 || row == len(word_search)-1 {
				continue
			}
			m_count := 0
			s_count := 0
			if char == "A" {
				if word_search[row-1][col-1] == "M" && word_search[row+1][col+1] != "M" {
					m_count++
				}
				if word_search[row-1][col+1] == "M" && word_search[row+1][col-1] != "M" {
					m_count++
				}
				if word_search[row+1][col-1] == "M" && word_search[row-1][col+1] != "M" {
					m_count++
				}
				if word_search[row+1][col+1] == "M" && word_search[row-1][col-1] != "M" {
					m_count++
				}
				if word_search[row-1][col-1] == "S" && word_search[row+1][col+1] != "S" {
					s_count++
				}
				if word_search[row-1][col+1] == "S" && word_search[row+1][col-1] != "S" {
					s_count++
				}
				if word_search[row+1][col-1] == "S" && word_search[row-1][col+1] != "S" {
					s_count++
				}
				if word_search[row+1][col+1] == "S" && word_search[row-1][col-1] != "S" {
					s_count++
				}
			}
			if m_count == 2 && s_count == 2 {
				matched_xmases++
			}
		}
	}
	fmt.Println(matched_xmases)
}
