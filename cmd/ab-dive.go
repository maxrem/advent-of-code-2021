package cmd

import (
	"fmt"
	"github.com/maxrem/advent-of-code-2021/lib/io"

	"github.com/spf13/cobra"
)

var diveCmd = &cobra.Command{
	Use:   "dive",
	Short: "Dive",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan string)
		fileReader := io.NewFileReader("ab-dive", true)

		go fileReader.Read(ch)

		for line := range ch {
			fmt.Println(line)
		}
	},
}

func init() {
	RootCmd.AddCommand(diveCmd)
}

type DiveReader struct{
	ch chan string 
}

func (r *DiveReader) Init(ch chan string) {
	
}
