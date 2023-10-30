package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_file(file_name string) ([]string, error) {
	// Open the file for reading
	file, err := os.Open(file_name) // Replace "yourfile.txt" with the path to your file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close() // Close the file when we're done

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	file_lines := make([]string, 0)
	// Read and convert each line to an integer
	for scanner.Scan() {
		line := scanner.Text()

		file_lines = append(file_lines, line)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading the file:", scanner.Err())
	}
	return file_lines, nil
}

func main() {
	// num, err := strconv.Atoi(line)
	file_lines, _ := read_file("day1.txt")
	sum := 0

	for _, line := range file_lines {
		num, _ := strconv.Atoi(line)

		sum += (num / 3) - 2
	}
	fmt.Println(sum)

}
