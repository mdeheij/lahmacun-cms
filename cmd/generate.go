package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate static HTML content",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("TODO: Generating static HTML content")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
