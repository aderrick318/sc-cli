/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// locationCmd represents the location command
var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("location called")
	},
}

func init() {
	rootCmd.AddCommand(locationCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	locationCmd.Flags().BoolP("beta", "b", false, "Launches Beta site for this command")
}
