package test

import (
	"fmt"
	"os"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/pflags"
	"github.com/pchan37/grader/internal/grader"
	"github.com/spf13/cobra"
)

var (
	testSuiteName = "default"
	programs      []string
)

var Cmd = &cobra.Command{
	Use:   "test [<name of test suite>]",
	Short: "Run a test suite",
	Long: `Run your program against a set of programs as configured by
    config.pflags.`,
	Run:  test,
	Args: cobra.RangeArgs(0, 1),
}

func test(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		testSuiteName = args[0]
	}
	parseConfig()

	for _, program := range programs {
		grader.Grade(program)
	}
}

func parseConfig() {
	config, err := pflags.Parse("config.pflags", "test")
	if err == nil {
		output.Successln("Using configuration specified in config.pflags!")
		fmt.Println("")

		if filenames, ok := config.Array.Get(testSuiteName); ok {
			programs = make([]string, 0)
			for _, file := range filenames {
				programs = append(programs, file.(string))
			}
		}
	} else {
		output.Errorln(err)
		os.Exit(1)
	}
}
