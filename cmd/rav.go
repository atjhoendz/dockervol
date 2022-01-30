package cmd

import (
	"fmt"

	"github.com/atjhoendz/dockervol/dockervol"
	"github.com/spf13/cobra"
)

var count int

// ravCmd represents the rav command
var ravCmd = &cobra.Command{
	Use:   "rav",
	Short: "rav - Remove anonymous docker volume",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add confirmation before removing
		err := dockervol.RemoveAnonymousVolume(count)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ravCmd)
	ravCmd.Flags().IntVarP(&count, "count", "c", 5, "Max anonymous volume that want be deleted")
}
