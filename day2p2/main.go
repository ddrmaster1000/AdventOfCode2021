package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(file_path string) []string {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	my_list := []string{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		tmp := scanner.Text()
		my_list = append(my_list, tmp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return my_list
}

func main() {

	// Read File
	my_slice := readFile("inputfile.txt")

	horizontal := 0
	depth := 0
	aim := 0

	// Have a SUMS slice of all the scans with moving frame
	for i := 0; i < len(my_slice); i++ {
		words := strings.Fields(my_slice[i])
		direction := words[0]
		distance, err := strconv.Atoi(words[1])

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontal += distance
			depth += aim * distance
		case "up":
			// depth -= distance
			aim -= distance
		case "down":
			// depth += distance
			aim += distance
		}

	}

	final_num := horizontal * depth
	fmt.Println(final_num)

}
