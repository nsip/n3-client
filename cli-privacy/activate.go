package main

import (
	"time"

	pub "../publish"
	"github.com/spf13/cobra"
)

// activateCmd represents the activate command
var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate a Context Privacy Policy File",
	Long: `
	e.g. [usage]: privacy activate context-demo 4947ED1F-1E94-4850-8B8F-35C653F51E9F 'A example of how to activate a context privacy'
	Use this command line to link a context and a privacy control file ID,
	Context can be linked to Privacy File multiple times, but the latest linkage is valid.
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			addCtxPrivid(args[0], args[1])
		} else {
			addCtxPrivid(args[0], args[1], args[2])
		}
	},
}

func init() {
	rootCmd.AddCommand(activateCmd)
}

func addCtxPrivid(params ...string) {
	objRoot, priid, cmt := params[0], params[1], time.Now().Format("2006-1-2 15:4:5.000")
	if len(params) == 3 {
		cmt = params[2]
	}

	if !S(priid).IsUUID() {
		fPln("error: the 2nd Param - [privacy id] is invalid UUID, Nothing Set")
		return
	}

	pub.Send("ctxid", objRoot, priid, cmt)
	time.Sleep(200 * time.Millisecond)
}
