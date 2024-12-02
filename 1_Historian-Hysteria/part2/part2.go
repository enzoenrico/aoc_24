package part2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readFile() ([]int, []int) {
	file, _ := os.Open("./example2.txt")
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 2 {
			num1, err := strconv.Atoi(fields[0])
			if err == nil {
				col1 = append(col1, num1)
			}
			num2, err := strconv.Atoi(fields[1])
			if err == nil {
				col2 = append(col2, num2)
			}
		}
	}
	return col1, col2
}

// run with the 2 column
func classify(arr []int) map[int]int {
	frequency := make(map[int]int)
	for _, val := range arr {
		if _, ok := frequency[val]; ok {
			frequency[val] += 1
		} else {
			frequency[val] = 1
		}
	}
	return frequency
}

func main() {
	l_col, r_col := readFile()
	freq := classify(r_col)

	total_sum := 0

	for _, val := range l_col {
		multiplier := 0
	}
}
