package grader

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/PGo-Projects/output"
)

func Grade(program string) {
	out, err := run(program)
	if err != nil {
		output.Errorln(err)
		os.Exit(1)
	}

	expectedOutput, err := getOutputFileContent(program)
	if err != nil {
		output.Errorln(err)
		os.Exit(1)
	}

	output.Println(fmt.Sprintf("Running %s:", program), output.BLUE)

	isDiff := false
	for index, line := range out {
		if index >= len(expectedOutput) {
			isDiff = true
			fmt.Println()
			output.ErrorStringln("The output was longer than the expected output...")
			for offset, line := range out[index:] {
				output.ErrorStringln(fmt.Sprintf("On line %d: got %s", index+offset+1, line))
			}
			break
		}

		expectedLine := expectedOutput[index]
		if line != expectedLine {
			isDiff = true
			output.ErrorStringln(
				fmt.Sprintf("On line %d: got %s, but expected %s", index+1, line, expectedLine),
			)
		}
	}

	if len(out) < len(expectedOutput) {
		isDiff = true
		output.ErrorStringln("The output was shorter than the expected output...")
		for _, line := range expectedOutput[len(out):] {
			output.ErrorStringln(line)
		}
	}

	if !isDiff {
		output.Successln(fmt.Sprintf("The program %s ran successfully!", program))
	} else {
		fmt.Println("")
	}
}

func run(program string) ([]string, error) {
	rawOutput, err := exec.Command("python", "sbml.py", program).Output()

	out := make([]string, 0)
	for _, line := range strings.Split(string(rawOutput), "\n") {
		out = append(out, line)
	}

	return out, err
}

func getOutputFileContent(program string) ([]string, error) {
	extension := filepath.Ext(program)
	outputFilename := strings.TrimSuffix(program, extension) + ".out"

	file, err := os.Open(outputFilename)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out, nil
}
