package main

import (
	"fmt"
	"os"

	g "github.com/nsip/n3-client/global"
	"github.com/spf13/cobra"
)

func main() {
	g.Init()
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "privacy",
	Short: "N3 Client Commnad Line Interface Tool For Privacy Control",
	Long:  "Allows user interaction with n3 Privacy Control, such as activate/deactivate a context's privacy policy",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
