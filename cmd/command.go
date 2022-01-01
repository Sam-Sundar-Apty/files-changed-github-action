package cmd

import (
	"fmt"
	"os/exec"

	"github.com/actions-go/toolkit/core"
)

func RunCommand(command string) ([]byte, error) {
	core.Info(fmt.Sprintf("running command %s", command))
	return exec.Command(command).Output()
}
