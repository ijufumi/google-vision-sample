package db

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "migration",
	}

	rootCmd.AddCommand(newDownCommand())
	rootCmd.AddCommand(newDownCommand())
	return rootCmd
}

func newUpCommand() *cobra.Command {
	return &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Up....")
		},
	}
}

func newDownCommand() *cobra.Command {
	return &cobra.Command{
		Use: "down",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Down....")
		},
	}
}
