package solutions

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func Day_twentyfive_part_one() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentyfive.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentyfive.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	long_scan := [][]string{}
	locks := [][]int{}
	keys := [][]int{}
	for scanner.Scan() {
		if len(long_scan) == 7 {
			long_scan = [][]string{}
			continue
		}
		long_scan = append(long_scan, []string{scanner.Text()})
		if len(long_scan) == 7 {
			builder := []int{-1, -1, -1, -1, -1}
			lock := true
			for i := 0; i < len(long_scan); i++ {
				for j, r := range long_scan[i][0] {
					if i == 0 && string(r) != "#" {
						lock = false
						break
					}
					if string(r) == "#" {
						builder[j]++
					}
				}
			}
			if lock {
				locks = append(locks, builder)
			} else {
				keys = append(keys, builder)
			}
		}
	}
	total_pairs := 0
	for i := 0; i < len(locks); i++ {
		for j := 0; j < len(keys); j++ {
			if eval_key_lock_pair(keys[j], locks[i]) {
				total_pairs++
			}
		}
	}
	fmt.Println(total_pairs)
}

func Day_twentyfive_part_two() {
	dat, err := os.ReadFile("./Full_Inputs/day_twentyfive.txt")
	// dat, err := os.ReadFile("./Test_Inputs/day_twentyfive.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
	}
}

func eval_key_lock_pair(key []int, lock []int) bool {
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}
