package cmd

import (
	"fmt"

	L "github.com/aderrick318/sc-cli/lib"
	"github.com/spf13/cobra"
)

var isLocationBeta bool
const locationEditPath string = "/LocationEdit.aspx"
var locationQueryString string = ""

// locationCmd represents the location command
var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Opens Location Edit page for provided locationID",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("No location id provided")
			return
		}

		supportCenterRoot := L.GetSupportCenterRootPath(isLocationBeta)
		locationValue := "0"
		if L.IsValidUUID(args[0]){
			locationValue = args[0]
			issueQueryString = "?LocationID="
		}
		if locationValue == "0"{
			fmt.Println("No valid location information provided")
			return
		}

		locationUrl := supportCenterRoot + locationEditPath + locationQueryString + locationValue

		L.Openbrowser(locationUrl)
	},
}

func init() {
	rootCmd.AddCommand(locationCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	locationCmd.Flags().BoolVarP(&isLocationBeta,"beta", "b", false, "Launches Beta site for this command")
}
