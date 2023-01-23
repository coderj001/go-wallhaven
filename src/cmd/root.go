package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-wallhaven",
	Version: "0.0.1",
	Short:   "go-wallhaven - A command line tool to easily download wallpapers from wallhaven.cc",
	Long:    `go-wallhaven: A seamless solution for downloading and managing wallpapers from wallhaven.cc right from your command line`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

// GetRoot returns the root of all subcommands
func GetRoot() *cobra.Command {
	return rootCmd
}

var pageCmd = &cobra.Command{
	Use:     "page",
	Aliases: []string{"p"},
	Short:   "Number of pages",
	Run: func(cmd *cobra.Command, args []string) {
		page, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: argument should be an integer")
			os.Exit(1)
		}
		fmt.Println(page)

	},
}

func init() {
	rootCmd.AddCommand(pageCmd)
}
