/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Opens Issue area in Support Center",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issue called")
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	issueCmd.Flags().BoolP("beta", "b", false, "Launches Beta site for this command")
}
