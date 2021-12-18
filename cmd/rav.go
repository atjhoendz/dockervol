/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/atjhoendz/dockervol/dockervol"
	"github.com/spf13/cobra"
)

// ravCmd represents the rav command
var ravCmd = &cobra.Command{
	Use:   "rav",
	Short: "Remove anonymous volume",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add confirmation before removing
		err := dockervol.RemoveAnonymousVolume()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	// TODO: add max deleted options
	rootCmd.AddCommand(ravCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ravCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ravCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
