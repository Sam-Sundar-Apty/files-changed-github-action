package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	errorOutput = `::debug::Waiting  milliseconds
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::error::strconv.Atoi: parsing "": invalid syntax
`
	successOutput = `::debug::Waiting 10 milliseconds
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::set-output name=time::2020-01-10 20:10:20.000000001 +0000 UTC
`
)

func TestRunMain(t *testing.T) {
	outout, err := exec.Command("git", "diff", "development", "feature/LSP-0000/checking-if-files-changed", "--", "./scripts", "./scripts").Output()
	if err != nil {
		t.Errorf("Error running command: %s", err.Error())
		return
	}
	assert.Equal(t, errorOutput, string(outout))
}
