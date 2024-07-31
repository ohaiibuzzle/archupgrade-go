package executors

import (
	"os"
	"os/exec"
)

func getUserShell() string {
	// TODO: Maybe we can use something else?
	return "/bin/sh"
}

func Shell(command string, input string) error {
	cmd := exec.Command(getUserShell(), "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func ShellWithOutput(command string) (string, error) {
	cmd := exec.Command(getUserShell(), "-c", command)
	out, err := cmd.Output()
	return string(out), err
}
