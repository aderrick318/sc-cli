package cmd

import (
	"fmt"
	"strconv"

	L "github.com/aderrick318/sc-cli/lib"
	"github.com/spf13/cobra"
)

var isIssueBeta bool
var issueEditPath string = "/SupportCenter/IssueEdit.aspx"
var issueQueryString string = ""

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Opens Issue area in Support Center",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("No issue id/number provided")
			return
		}

		supportCenterRoot := L.GetSupportCenterRootPath(isIssueBeta)
		issueValue := "0"
		if _, err := strconv.Atoi(args[0]); err == nil {
			issueValue = args[0]
			issueQueryString = "?IssueNumber="
		}
		if L.IsValidUUID(args[0]){
			issueValue = args[0]
			issueQueryString = "?IssueID="
		}
		if issueValue == "0"{
			fmt.Println("No valid issue information provided")
			return
		}

		issueUrl := supportCenterRoot + issueEditPath + issueQueryString + issueValue

		L.Openbrowser(issueUrl)
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)

	issueCmd.Flags().BoolVarP(&isIssueBeta,"beta", "b", false, "Launches Beta site for this command")
}
