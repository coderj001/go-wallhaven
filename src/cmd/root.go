package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-wallhaven",
	Version: "0.0.1",
	Short:   "go-wallhaven: Wallpaper download manager for wallhaven.com",
	Long:    `go-wallhaven: Wallpaper download manager for wallhaven.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the root command!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// GetRoot returns the root of all subcommands
func GetRoot() *cobra.Command {
	return rootCmd
}

func init() {}
