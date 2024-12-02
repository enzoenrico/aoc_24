package part1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile() ([]int, []int) {
	file, _ := os.Open("./input.txt")
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
	// i need to pair numbers according to their position in a ascending order
	// i can write a helper to get a specific number like the 3rd smallest
	// or i can just sort them
	return col1, col2
}

func subRight(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getPoints(col1, col2 []int) int {
	var total int
	total = 0

	for i := 0; i < len(col1); i++ {
		// absolute subtraction (no negatives >:[ )
		diff := subRight(col1[i], col2[i])
		// fmt.Printf("idx %d - total %d\n", i, total)
		total += diff
	}
	return total
}

func main() {
	col1, col2 := readFile()

	sort.Slice(col1, func(i, j int) bool {
		return col1[i] < col1[j]
	})

	sort.Slice(col2, func(i, j int) bool {
		return col2[i] < col2[j]
	})
	fmt.Println(len(col1))
	fmt.Println(len(col2))

	// for i := range len(col1) {
	// fmt.Printf("Col1: %d \t Col2: %d\n", col1[i], col2[i]) // banging to fallen angels - bvb rn
	// WE SCREAM, WE SHOUT, WE ARE THE FALLEN ANGELSSSS
	// *SICK ASS GUITAR SOLO*
	// }

	fmt.Printf("> %d", getPoints(col1, col2))
}
