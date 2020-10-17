package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"parec",
		"--format=float32le",
		"--rate=44100",
		"--channels=1",
	)

	cmd.Stderr = os.Stderr

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer pipe.Close()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
