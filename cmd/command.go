package cmd

import (
	"fmt"
	"os/exec"

	"github.com/actions-go/toolkit/core"
)

func RunCommand(command string) ([]byte, error) {
	core.Info(fmt.Sprintf("running command %s", command))
	return exec.Command("git", "diff", "development", "feature/LSP-0000/checking-if-files-changed", "--", "./scripts", "./scripts").Output()
}
