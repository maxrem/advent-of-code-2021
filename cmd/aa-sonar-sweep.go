package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var sonarSweepCmd = &cobra.Command{
	Use:   "sonar-sweep",
	Short: "Sonar sweep",
	Run: func(cmd *cobra.Command, args []string) {
		
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// testCount, err := readFileAndCountIncreases(path + "/01/test.txt")
	// if err != nil {
	// 	panic("test has an error " + err.Error())
	// }
	// fmt.Println(testCount)

	//count, err := readFileAndCountIncreases(path + "/01/input.txt")
	//if err != nil {
	//	panic("input has an error" + err.Error())
	//}
	//fmt.Println(count)

	//count, err := readBlocks(path + "/01/test.txt", 3)
	//if err != nil {
	//	panic("test has an error" + err.Error())
	//}
	//fmt.Println(count)

	count, err := readBlocks(path + "/input-files/aa-sonar-sweep.txt", 3)
	if err != nil {
		panic("input has an error" + err.Error())
	}
	
	fmt.Println(count)
	},
}

func init() {
	RootCmd.AddCommand(sonarSweepCmd)
}

func readBlocks(path string, blockSize int) (int, error) {
	var block []int

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	lastSum := -1
	count := -1
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		block = append(block, current)

		if len(block) >= blockSize {
			currentBlock := block[index:blockSize+index]
			currentSum := 0
			for _, measurement := range currentBlock {
				currentSum += measurement
			}
			//fmt.Println(currentBlock, currentSum)

			if lastSum < currentSum {
				count++
			}
			lastSum = currentSum

			index++
		}
	}

	//fmt.Println(block)

	return count, nil
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
