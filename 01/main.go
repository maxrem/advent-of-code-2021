package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	// testCount, err := readFileAndCountIncreases(path + "/01/test.txt")
	// if err != nil {
	// 	panic("test has an error" + err.Error())
	// }
	// fmt.Println(testCount)

	count, err := readFileAndCountIncreases(path + "/01/input.txt")
	if err != nil {
		panic("input has an error" + err.Error())
	}
	fmt.Println(count)
}

func readFileAndCountIncreases(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	last := -1
	counter := 0
	first := true
	for scanner.Scan() {
		if first {
			first = false
			continue
		}
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		if current > last {
			counter++
		}

		last = current
	}

	return counter, nil
}