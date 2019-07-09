package main

import (
	"time"

	pub "../publish"
	"github.com/spf13/cobra"
)

// deactivateCmd represents the deactivate command
var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate a Context's Privacy Policy File",
	Long: `
	e.g. [usage]: privacy deactivate context-demo 'A example of how to deactivate a context privacy'
	Use this command line to unlink a context from its privacy control file ID,	
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			rmCtxPrivid(args[0], args[1])
		} else {
			rmCtxPrivid(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)
}

func rmCtxPrivid(params ...string) {
	objRoot, cmt := params[0], time.Now().Format("2006-1-2 15:4:5.000")
	if len(params) == 2 {
		cmt = params[1]
	}

	pub.Send("ctxid", objRoot, "00000000-0000-0000-0000-000000000000", cmt)
	time.Sleep(200 * time.Millisecond)
}
