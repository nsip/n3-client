package main

import (
	"time"

	g "github.com/nsip/n3-client/global"
	pub "github.com/nsip/n3-client/publish"
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

	// check context is a valid context in n3-transport
	if ctx == g.Cfg.RPC.CtxPrivID || ctx == g.Cfg.RPC.CtxPrivDef {
		fPln("error: the 2nd Param - [context] is invalid, Nothing Set")
		return
	}

	pub.Send(g.Cfg.RPC.CtxPrivID, objRoot, ctx, g.MARKDelID, 0)
	time.Sleep(200 * time.Millisecond)
}
