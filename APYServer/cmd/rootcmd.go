// Package cmd contains the cobra command line setup
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "superform-apy",
	Short: "superform-apy",
	Long:  `github.com/subhasishgoswami/superform`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("superform-apy")
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
