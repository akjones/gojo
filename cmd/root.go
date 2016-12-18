package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:  "gojo",
	Long: `Enlighten yourself.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stuff!")
	},
}

func init() {
	RootCmd.PersistentFlags().StringP("author", "a", "akjones", "Author name for copyright attribution")
	RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "andrew jones hello@andrewjones.id.au")
	viper.SetDefault("license", "GPLV3")
}
