package uos

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecShCmd(cmdStr string) error {
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func ExecShCmdWithOutput(cmdStr string) (string, error) {
	c := exec.Command("sh", "-c", cmdStr)
	var output bytes.Buffer
	c.Stdout = &output
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return "", err
	}
	return output.String(), nil
}

func ExecShCmdWithStdoutAndStderr(cmdStr string) (stdout string, stderr string, retErr error) {
	c := exec.Command("sh", "-c", cmdStr)
	var stdoutBuf, stderrBuf bytes.Buffer
	c.Stdout = &stdoutBuf
	c.Stderr = &stderrBuf
	err := c.Run()
	return stdoutBuf.String(), stderrBuf.String(), err
}

func MustExecShCmd(cmdStr string) {
	if err := ExecShCmd(cmdStr); err != nil {
		panic(err)
	}
}

func MustExecShCmdQuietly(cmdStr string) {
	stdout, stderr, err := ExecShCmdWithStdoutAndStderr(cmdStr)
	if err != nil {
		panic(fmt.Sprintf(
			"Failed to execute %v: %v. Stdout: %v, Stderr: %v", cmdStr, err, stdout, stderr))
	}
}
