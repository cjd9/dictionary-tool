package cmd

import (
	"fmt"
	"os"

	"github.com/example/merriam-webster/source/client"
	"github.com/spf13/cobra"
)

var defineCmd = &cobra.Command{
	Use:   "define <word>",
	Short: "Get the definition of a word from Merriam-Webster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]
		def, err := client.Lookup(word)
		if err != nil {
			if suggErr, ok := err.(*client.SuggestionError); ok {
				fmt.Fprintln(os.Stderr, "Couldn't find the word. Did you mean any of these words?")
				for _, s := range suggErr.Suggestions {
					fmt.Fprintln(os.Stderr, "-", s)
				}
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		fmt.Printf("`%s` (%s): %s\n", def.Pronunciation, def.PartOfSpeech, def.Meaning)
	},
}

func init() {
	RootCmd.AddCommand(defineCmd)
}
