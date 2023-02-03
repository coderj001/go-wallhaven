package cmd

import (
	"fmt"
	"os"

	"github.com/coderj001/go-wallheven/src/app"

	"github.com/spf13/cobra"
)

var (
	page       int
	query      string
	categories string
	purity     string
	sorting    string
	colors     string
	dir        string
)

var rootCmd = &cobra.Command{
	Use:     "go-wallhaven",
	Version: "0.0.1",
	Short:   "go-wallhaven - A command line tool to easily download wallpapers from wallhaven.cc",
	Long:    `go-wallhaven: A seamless solution for downloading and managing wallpapers from wallhaven.cc right from your command line`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := app.GetFullURL(page, categories, purity, sorting, colors, query)
		fmt.Println(url)
		// fmt.Println(app.FetchAPI(url))
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
	GetRoot().Flags().StringVarP(&categories, "categories", "t", "all", `
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
	GetRoot().Flags().StringVarP(&sorting, "sorting", "s", "relevance", `
		Sorting Filter:
		relevance    - For sorting the results based on relevance to the search query
		random       - For displaying random wallpapers
		date_added   - For sorting the results based on date added
		favorites    - For sorting the results based on number of favorites
		views        - For sorting the results based on number of views
		toplist      - For sorting the results based on toplist
		hot          - For sorting the results based on hotness

	`)
	GetRoot().Flags().StringVarP(&colors, "colors", "c", "", `Select Color Hex Code eg. 663399`)
	GetRoot().Flags().StringVarP(&query, "query", "q", "", `Search Image For Query eg. Dark Knight`)
	GetRoot().Flags().StringVarP(&dir, "dir", "d", "", `Path to the directory`)
}
