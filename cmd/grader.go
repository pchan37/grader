package cmd

import (
	"os"

	"github.com/PGo-Projects/output"
	"github.com/pchan37/grader/cmd/run"
	"github.com/pchan37/grader/cmd/test"
	"github.com/spf13/cobra"
)

var graderCmd = &cobra.Command{
	Use:   "grader",
	Short: "A grader for CSE 307 sbml language",
	Long: `A utility that compares the expected output with the given output for
    the sbml language in CSE 307.`,
	Run: grader,
}

func init() {
	graderCmd.AddCommand(run.Cmd)
	graderCmd.AddCommand(test.Cmd)
}

func grader(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}
}

func Execute() {
	if err := graderCmd.Execute(); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
}
