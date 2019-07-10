package main

import (
	"time"

	pub "../publish"
	"github.com/spf13/cobra"
)

// activateCmd represents the activate command
var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate a objectRoot in a certain Context to a Privacy Policy File",
	Long: `
	e.g. [usage]: privacy activate 'ObjectRoot' 'Context' '4947ED1F-1E94-4850-8B8F-35C653F51E9F'
	Use this command line to link an objectRoot in a certain context to a privacy control file ID,
	objectRoot can be linked to different Privacy Files, but the latest linkage is working.
	`,
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		addCtxPrivid(args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(activateCmd)
}

func addCtxPrivid(params ...string) {
	objRoot, ctx, priid := params[0], params[1], params[2]
	if !S(priid).IsUUID() {
		fPln("error: the 3rd Param - [privacy id] is invalid UUID, Nothing Set")
		return
	}
	pub.Send(ctxid, objRoot, ctx, priid)
	time.Sleep(200 * time.Millisecond)
}
