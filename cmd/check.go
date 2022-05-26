/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	fmt.Println("configuration checked")
	rootCmd.AddCommand(checkCmd)
	viper.SetConfigFile("splat.yaml")
	viper.ReadInConfig()

	fmt.Println(viper.Get("THING"))
	// viper.AddConfigPath("./")
	// viper.SetConfigName("splat") // Register config file name (no extension)
	// viper.SetConfigType("yaml")  // Look for specific type
	// viper.ReadInConfig()
	// fmt.Println(viper.ReadInConfig())
}
