/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	L "github.com/aderrick318/sc-cli/lib"
	"github.com/spf13/cobra"
)

var isPersonBeta bool
const personEditPath string = "/SupportCenter/PersonEdit.aspx"
var personQueryString string = ""

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Opens person edit for provided personID",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("No person id provided")
			return
		}

		supportCenterRoot := L.GetSupportCenterRootPath(isPersonBeta)
		personValue := "0"
		if L.IsValidUUID(args[0]){
			personValue = args[0]
			issueQueryString = "?PersonID="
		}
		if personValue == "0"{
			fmt.Println("No valid person information provided")
			return
		}

		personUrl := supportCenterRoot + personEditPath + personQueryString + personValue

		L.Openbrowser(personUrl)
	},
}

func init() {
	rootCmd.AddCommand(personCmd)

	personCmd.Flags().BoolVarP(&isPersonBeta,"beta", "b", false, "Launches Beta site for this command")
}
