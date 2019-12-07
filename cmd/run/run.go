package run

import (
	"github.com/pchan37/grader/internal/grader"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run <program to test>",
	Short: "Run a single program and compare output",
	Long: `Compare the output after running the specified program against the
    expected output and display the results.`,
	Run:  test,
	Args: cobra.ExactArgs(1),
}

func test(cmd *cobra.Command, args []string) {
	program := args[0]
	grader.Grade(program)
}
