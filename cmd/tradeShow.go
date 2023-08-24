/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	L "github.com/aderrick318/sc-cli/lib"
	"github.com/spf13/cobra"
)

var isTradeShowBeta bool
var tradeShowEditPath string = "/SupportCenter/TradeShowEdit.aspx"
var tradeShowQueryString string = ""

// tradeShowCmd represents the tradeShow command
var tradeShowCmd = &cobra.Command{
	Use:   "tradeShow",
	Short: "Opens Trade Show edit for provided Trade ShowID (int)",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("No trade show id provided")
			return
		}

		supportCenterRoot := L.GetSupportCenterRootPath(isIssueBeta)
		tradeShowValue := "0"
		if _, err := strconv.Atoi(args[0]); err == nil {
			tradeShowValue = args[0]
			issueQueryString = "?TradeShowID="
		}
		
		if tradeShowValue == "0"{
			fmt.Println("No valid Trade Show information provided")
			return
		}

		issueUrl := supportCenterRoot + tradeShowEditPath + tradeShowQueryString + tradeShowValue

		L.Openbrowser(issueUrl)
	},
}

func init() {
	rootCmd.AddCommand(tradeShowCmd)

	tradeShowCmd.Flags().BoolVarP(&isTradeShowBeta,"beta", "b", false, "Launches Beta site for this command")
}
