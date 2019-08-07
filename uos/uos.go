package uos

import (
	"bytes"
	"os"
	"os/exec"
)

func execShCmd(cmdStr string) error {
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func execShCmdWithOutput(cmdStr string) (string, error) {
	c := exec.Command("sh", "-c", cmdStr)
	var output bytes.Buffer
	c.Stdout = &output
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return "", err
	}
	return output.String(), nil
}
