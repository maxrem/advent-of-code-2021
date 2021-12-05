package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diveCmd = &cobra.Command{
	Use:   "dive",
	Short: "Dive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dive!")
	},
}

func init() {
	RootCmd.AddCommand(diveCmd)
}