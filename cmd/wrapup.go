package cmd

import (
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/phlow"
)

var verboseWrapup bool

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "Add work to index and commit with close# message",
	Long: `
Add all your work from working area to index, and makes an automatic close# commit, which closes
an issue on github
`,

	Run: func(cmd *cobra.Command, args []string) {
		phlow.WrapUp(verboseWrapup)
	},
}

func init() {
	RootCmd.AddCommand(wrapupCmd)

	//Enable verbose printers
	wrapupCmd.Flags().BoolVarP(&verboseWrapup, "verbose", "v", false, "print verbose output")

}
