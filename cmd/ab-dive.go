package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/maxrem/advent-of-code-2021/lib/io"

	"github.com/spf13/cobra"
)

var diveCmd = &cobra.Command{
	Use:   "dive",
	Short: "Dive",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan string)
		fileReader := io.NewFileReader("ab-dive", false)
		diveReader := DiveReader{ch: ch}

		go fileReader.Read(ch)

		diveReader.HandleSecond()
		diveReader.Print()
	},
}

func init() {
	RootCmd.AddCommand(diveCmd)
}

type DiveReader struct{
	ch chan string
	distance int
	depth int
	aim int
}

func (r *DiveReader) HandleFirst() {
	for row := range r.ch {
		instruction := strings.Split(row, " ")
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}
		switch(instruction[0]) {
		case "forward":
			r.distance += value
		case "down":
			r.depth += value
		case "up":
			r.depth -= value
			if r.depth < 0 {
				r.depth = 0
			}
		}
	}
}

func (r *DiveReader) HandleSecond() {
	for row := range r.ch {
		instruction := strings.Split(row, " ")
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}
		switch(instruction[0]) {
		case "forward":
			r.distance += value
			addedDepth := r.aim * value
			r.depth += addedDepth
		case "down":
			r.aim += value
		case "up":
			r.aim -= value
		}
	}
}

func (r *DiveReader) Print() {
	fmt.Println(
		fmt.Sprintf(
			"Distance: %d, depth: %d, answer: %d",
			r.distance,
			r.depth,
			r.distance * r.depth,
			),
		)
}
