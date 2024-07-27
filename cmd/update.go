package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"

	"github.com/junjie1999/exchange-cli/gobutil"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the currency exchange rates.",
	Long:  "Update the currenct exchange rates from x-rates.com.",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching currency exchange rates ...\n")
		err, timestamp := gobutil.UpdateCurrency()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Currency exchange rates updated - " + FormatUnixTimestamp(*timestamp) + "\n")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func FormatUnixTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
