/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var isSalesOrderBeta bool
var salesOrderEditPath string = "/QBIntegration/SalesOrderEdit.aspx"
var salesOrderQueryString string = ""

// salesOrderCmd represents the salesOrder command
var salesOrderCmd = &cobra.Command{
	Use:   "salesOrder",
	Short: "Opens sales order edit for provided ID or Number (guid or int)",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("No sales order id/number provided")
			return
		}

		supportCenterRoot := GetSupportCenterRootPath(isIssueBeta)
		salesOrderValue := "0"
		if _, err := strconv.Atoi(args[0]); err == nil {
			salesOrderValue = args[0]
			issueQueryString = "?SalesOrderID="
		}
		if IsValidUUID(args[0]){
			salesOrderValue = args[0]
			issueQueryString = "?QBNumber="
		}
		if salesOrderValue == "0"{
			fmt.Println("No valid sales order information provided")
			return
		}

		issueUrl := supportCenterRoot + salesOrderEditPath + salesOrderQueryString + salesOrderValue

		Openbrowser(issueUrl)
	},
}

func init() {
	rootCmd.AddCommand(salesOrderCmd)

	salesOrderCmd.Flags().BoolVarP(&isSalesOrderBeta,"beta", "b", false, "Launches Beta site for this command")
}
