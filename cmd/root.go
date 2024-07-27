package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"

	"github.com/junjie1999/exchange-cli/calc"
)

var Logo = `
    _______         __
   / _____/ _______/ /_  ____ _____  ____ ____
  / __/ | |/_/ ___/ __ \/ __ ` + "`" + `/ __ \/ __ ` + "`" + `/ _ \
 / /____>  </ /__/ / / / /_/ / / / / /_/ /  __/
/_____/_/|_|\___/_/ /_/\__,_/_/ /_/\__, /\___/
                                  /____/     
`

var rootCmd = &cobra.Command{
	Use:   "exchange-cli",
	Short: "exchange-cli is a CLI tool to convert currencies",
	Long: `A CLI tool to convert currencies.
    It scrapes the latest currency exchange rates from x-rates.com.`,
	Example: "  exchange-cli <amount> <base> <target>\n" +
		"  exchange-cli 100 USD EUR\n" +
		"  exchange-cli update",

	PersistentPreRun: logo(),

	Args: argsValidation(),

	RunE: rootRun(),

	Version: "v1.0.0",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.DisableDescriptions = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func logo() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(Logo)
	}
}

func isDigit(arg string) bool {
	for _, r := range arg {
		if !unicode.IsDigit(r) && r != '.' {
			return false
		}
	}
	return true
}

func argsValidation() cobra.PositionalArgs {

	return cobra.MatchAll(cobra.ExactArgs(3), func(cmd *cobra.Command, args []string) error {
		if !isDigit(args[0]) {
			return errors.New("invalid argument '" + args[0] + "', must be a digit.")
		}

		args[1] = strings.ToUpper(args[1])
		args[2] = strings.ToUpper(args[2])

		return nil
	})
}

func rootRun() func(cmd *cobra.Command, args []string) error {

	return func(cmd *cobra.Command, args []string) error {

		if len(args) != 3 {
			cmd.Help()
			return nil
		}

		amountStr := args[0]
		base := strings.ToUpper(args[1])
		target := strings.ToUpper(args[2])

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			return err
		}

		Ex, err := calc.Convert(amount, base, target)
		if err != nil {
			return err
		}

		output(Ex)
		return nil
	}
}

func output(Ex *calc.Exchange) {
	fmt.Printf("Timestamp: %s\n\n%12.2f %s\n%12.2f %s\n\n", calc.GetTimeStamp(),
		Ex.Amount, Ex.Base, Ex.Result, Ex.Target)
}
