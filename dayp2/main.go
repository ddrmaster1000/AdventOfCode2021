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

func getThrees(my_slice []int, starting int) []int {
	return_list := []int{}
	// my_count := 0
	for i := starting; i < len(my_slice); i++ {
		if i >= starting+3 {
			return return_list
		} else {
			return_list = append(return_list, my_slice[i])
		}
	}
	return return_list
}

func addSlice(my_slice []int) int {
	total := 0
	for i := 0; i < len(my_slice); i++ {
		total += my_slice[i]
	}
	return total
}

func main() {

	// Read File
	my_slice := readFile("inputfile.txt")

	// Get 3s as a moving window into a slice
	overall_threes := [][]int{}
	for i := 0; i < len(my_slice)-2; i++ {
		overall_threes = append(overall_threes, getThrees(my_slice, i))
	}
	// fmt.Println(overall_threes)

	// Have a SUMS slice of all the scans with moving frame
	sum_of_threes_slice := []int{}
	for i := 0; i < len(overall_threes); i++ {
		sum_of_threes_slice = append(sum_of_threes_slice, addSlice(overall_threes[i]))
	}
	// fmt.Println(sum_of_threes_slice)

	// Get if larger
	prev := 0
	new_val := 0
	first_slice := true
	islarger_count := 0
	for i := 0; i < len(sum_of_threes_slice); i++ {
		// Skip the first iteration
		if first_slice {
			prev = new_val
			first_slice = false
			continue
		}
		// Check if it is larger or not
		new_val = sum_of_threes_slice[i]
		if new_val > prev {
			islarger_count++
		}
		prev = new_val
	}
	fmt.Println(islarger_count)

}
