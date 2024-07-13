package cmd

import (
	"fmt"
	"strings"

	"github.com/Shhwip/bgate-scraper/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List all books of the Bible and how many chapters they have",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("translation", cmd.Flag("translation"))
		viper.BindPFlag("padding", cmd.Flag("padding"))
		viper.BindPFlag("abbreviations", cmd.Flag("abbreviations"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		filter, _ := cmd.Flags().GetString("filter")
		translation := viper.GetString("translation")
		padding := viper.GetInt("padding")
		abbreviations := viper.GetBool("abbreviations")
		local, err := search.TranslationHasLocal(translation)
		cobra.CheckErr(err)

		var searcher search.Searcher
		if local {
			searcher, err = search.NewLocal(translation)
			cobra.CheckErr(err)
		} else {
			searcher = search.NewRemote(translation)
		}

		books, err := searcher.Booklist()
		cobra.CheckErr(err)

		if abbreviations {
			for _, book := range books {
				fmt.Printf("%s%s\n", strings.Repeat(" ", padding), book.String())
				for abbr, name := range search.Abbreviations {
					if name == book.Name {
						fmt.Printf("%s%s,\n", strings.Repeat(" ", padding*2), abbr)
					}
				}
			}
		} else {
			for _, book := range books {
				if strings.Contains(strings.ToLower(book.Name), strings.ToLower(filter)) {
					fmt.Printf("%s%s\n", strings.Repeat(" ", padding), book.String())
				}
			}
		}
	},
}

func init() {
	list.Flags().StringP("filter", "f", "", "Filter the list of books by name. (Case insensitive)")
	list.Flags().StringP("translation", "t", "ESV", "The translation of the Bible to search for.")
	list.Flags().IntP("padding", "p", 0, "Horizontal padding in character count.")
	list.Flags().BoolP("abbreviations", "a", false, "Display the abbreviations of the books of the bible.")
	root.AddCommand(list)
}
