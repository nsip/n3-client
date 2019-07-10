package main

import (
	"time"

	pub "../publish"
	"github.com/spf13/cobra"
)

// deactivateCmd represents the deactivate command
var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate an objectRoot in a certain Context from a Privacy Policy File",
	Long: `
	e.g. [usage]: privacy deactivate 'ObjectRoot' 'Context'
	Use this command line to unlink an objectRoot in a context from its privacy control file ID,	
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		rmCtxPrivid(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)
}

func rmCtxPrivid(params ...string) {
	objRoot, ctx := params[0], params[1]
	pub.Send(ctxid, objRoot, ctx, "00000000-0000-0000-0000-000000000000")
	time.Sleep(200 * time.Millisecond)
}
