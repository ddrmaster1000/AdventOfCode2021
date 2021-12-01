package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(file_path string) []int {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	my_list := []int{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		tmp := scanner.Text()
		intVal, err := strconv.Atoi(tmp)
		if err != nil {
			log.Fatal(err)
		}
		my_list = append(my_list, intVal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return my_list
}

func main() {
	// my_slice := []int{199,
	// 	200,
	// 	208,
	// 	210,
	// 	200,
	// 	207,
	// 	240,
	// 	269,
	// 	260,
	// 	263}
	// fmt.Println(my_slice)

	my_slice := readFile("inputfile.txt")

	// Iterate over slice
	prev := 0
	new_val := 0
	first_slice := true
	islarger_count := 0

	for i := 0; i < len(my_slice); i++ {
		// Skip the first iteration
		if first_slice {
			prev = new_val
			first_slice = false
			continue
		}
		// Check if it is larger or not
		new_val = my_slice[i]
		// fmt.Println(new_val)
		if new_val > prev {
			islarger_count++
		}
		prev = new_val
	}
	fmt.Println(islarger_count)
}
