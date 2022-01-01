package main

import (
	"fmt"
	"os"

	"github.com/actions-go/go-action/cmd"
	"github.com/actions-go/toolkit/core"
)

func runMain() {

	targetBranch, _ := core.GetInput("TARGET_BRANCH")

	if targetBranch == "" {
		core.Error("TARGET_BRANCH is required")
		os.Exit(1)
	}

	currentBranch, _ := core.GetInput("CURRENT_BRANCH")
	if currentBranch == "" {
		core.Error("CURRENT_BRANCH is required")
		os.Exit(1)
	}

	comparePath, _ := core.GetInput("COMPARE_PATH")
	if comparePath == "" {
		core.Error("COMPARE_PATH is required")
		os.Exit(1)
	}

	command := fmt.Sprintf("git diff %s %s -- %s %s", targetBranch, currentBranch, comparePath, comparePath)

	cmd, err := cmd.RunCommand(command)
	if err != nil {
		core.Errorf("Error running command: %s", err.Error())
		return
	}

	outputLength := len(string(cmd))

	if outputLength == 0 {
		core.SetOutput("CHANGED", "false")
		return
	}
	core.SetOutput("CHANGED", "true")
}

func main() {
	runMain()
}
