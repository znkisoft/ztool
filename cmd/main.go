package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	Execute()
}

var RootCmd = &cobra.Command{
	Use:   "ztool",
	Short: "util package for znki side-projects",
	Long:  `https://github.com/znkisoft/ztool`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
