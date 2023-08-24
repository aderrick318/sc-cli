/*
Copyright © 2023 NAME HERE aderrick318@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sc-cli",
	Short: "cli tool for opening various areas of Support Center directly",
	Long: `Opens the corresponding edit pages for Issues, Location, Person records, and Sales Orders.`,
	// Run: func(cmd *cobra.Command, args []string) { 

	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

}


var productionRootUrl string = "https://support.pioneerrx.com"
var betaRootUrl string = "https://beta.pioneerrx.com"

func GetSupportCenterRootPath(isBeta bool) string {
	rootUrl := ""
	if isBeta {
		rootUrl = betaRootUrl
	} else {
		rootUrl = productionRootUrl
	}

	return rootUrl
}

func IsValidUUID(u string) bool {
    _, err := uuid.Parse(u)
    return err == nil
 }

 func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}