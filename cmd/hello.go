package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "say hello",
	Long:  "a simple command just print hello",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Println("hello, ", name)
	},
}

var name string

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().String("name", "guest", "a name to say hello")
}
