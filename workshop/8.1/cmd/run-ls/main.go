package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l", "/tmp/kubebuilder-tools-1.21.4-darwin-amd64.tar.gz")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
