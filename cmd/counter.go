/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	Services "github.com/jadhamwi21/dota2-counterpick-cli-tool/services"
	"github.com/spf13/cobra"
)

var counterCmd = &cobra.Command{
	Use: "counter",
	Args: func(cmd *cobra.Command, args []string) error {
		if NumberOfArguments := len(args); NumberOfArguments == 0 {
			return errors.New("you have to pass a hero as argument")
		} else {
			return nil
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		counterService := &Services.CounterService{CounterList: make(map[string][]string)}
		counterService.Initialize()
		multipleFlag, _ := cmd.Flags().GetBool("multiple")
		if multipleFlag {
			MultipleHeroes := args
			fmt.Println(counterService.CounterMultipleHeroes(MultipleHeroes))
		} else {
			SingleHeroName := strings.ToLower(args[0])
			fmt.Println(counterService.CounterSingleHero(SingleHeroName))
		}
	},
}

func init() {
	rootCmd.AddCommand(counterCmd)

	counterCmd.Flags().BoolP("multiple", "m", false, "use this flag to denote you're countering multiple heroes")
}
