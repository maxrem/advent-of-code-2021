package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/maxrem/advent-of-code-2021/lib/io"
	"github.com/spf13/cobra"
)

var binaryDiagnosticCmd = &cobra.Command{
	Use:   "binary-diagnostic",
	Short: "Binary diagnostic",
	Run: func(cmd *cobra.Command, args []string) {
		isTest := false
		if args[0] == "test" {
			isTest = true
		}
		inputSize, err := strconv.Atoi(args[1]); if err != nil {
			log.Fatal(err)
		}
		

		ch := make(chan string)
		fileReader := io.NewFileReader("ac-binary-diagnostic", isTest)

		diagnosticReader := DiagnosticReader{
			ch:       ch,
			counters: make([]Iets, inputSize, inputSize),
		}

		go fileReader.Read(ch)

		diagnosticReader.HandleFirst()
		diagnosticReader.Print()
		gamma := diagnosticReader.GetBinary(true)
		epsilon := diagnosticReader.GetBinary(false)
		answer := diagnosticReader.GetUint(gamma) * diagnosticReader.GetUint(epsilon)
		fmt.Printf("Answer is %d\n", answer)

		// https://go.dev/play/p/lTKJdLxCIUT
	},
}

func init() {
	RootCmd.AddCommand(binaryDiagnosticCmd)
}

type Iets struct {
	ZeroCount int
	OneCount int
}

type DiagnosticReader struct {
	ch chan string
	counters []Iets
}

func (r *DiagnosticReader) Print() {
	for index, counter := range r.counters {
		fmt.Println(fmt.Sprintf("Counter: %d, zero: %d, one: %d", index, counter.ZeroCount, counter.OneCount))
	}
}

func (r *DiagnosticReader) GetBinary(isGamma bool) string {
	zero := "0"
	one := "1"
	if !isGamma {
		zero = "1"
		one = "0"
	}

	var result string
	for _, counter := range r.counters {
		if counter.ZeroCount > counter.OneCount {
			result += zero
			continue
		}

		result += one
	}

	return result
}

func (r *DiagnosticReader) GetUint(binary string) uint64 {
	i, err := strconv.ParseUint(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func (r *DiagnosticReader) HandleFirst() {
	for row := range r.ch {
		for pos, char := range row {
			if char == '0' {
				r.counters[pos].ZeroCount++
				continue
			}
			r.counters[pos].OneCount++
		}
	}
}