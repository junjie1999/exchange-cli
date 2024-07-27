package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/junjie1999/exchange-cli/unit"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all country units.",
	Long:  "List all country units that are available.",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ListUnits()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "List all currency rates.")
}

func ListUnits() {
	fmt.Println("Available currencies:\n")

	keys := make([]string, 0, len(unit.X_Rate_CountryUnits))
	for k := range unit.X_Rate_CountryUnits {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%7s: %s\n", k, unit.X_Rate_CountryUnits[k])
	}
	fmt.Println()
}
