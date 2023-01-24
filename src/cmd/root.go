package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var page int
var categories string
var purity string

// var resolutions []string

var rootCmd = &cobra.Command{
	Use:     "go-wallhaven",
	Version: "0.0.1",
	Short:   "go-wallhaven - A command line tool to easily download wallpapers from wallhaven.cc",
	Long:    `go-wallhaven: A seamless solution for downloading and managing wallpapers from wallhaven.cc right from your command line`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Page: %d\n", page)
		fmt.Printf("Categories: %s\n", categories)
		fmt.Printf("Purity: %s\n", purity)
	},
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

func init() {
	GetRoot().Flags().IntVarP(&page, "page", "n", 1, "Total Numbers Of Pages")
	GetRoot().Flags().StringVarP(&categories, "categories", "c", "all", `
			Categories Filter:
			all     - Every wallpaper.
			general - For 'general' wallpapers only.
			anime   - For 'Anime' Wallpapers only.
			people  - For 'people' wallpapers only.
			ga      - For 'General' and 'Anime' wallpapers only.
			gp      - For 'General' and 'People' wallpapers only.
	`)
	GetRoot().Flags().StringVarP(&purity, "purity", "p", "all", `
			Purity Filter:
			sfw     - For 'Safe For Work'
			sketchy - For 'Sketchy'
			nsfw    - For 'Not Safe For Work'
			ws      - For 'SFW' and 'Sketchy'
			wn      - For 'SFW' and 'NSFW'
			sn      - For 'Sketchy' and 'NSFW'
			all     - For 'SFW', 'Sketchy' and 'NSFW'
	`)
}
