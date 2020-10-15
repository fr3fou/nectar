package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"parec",
		"--format=float32le",
		"--rate=48000",
		"--channels=1",
	)

	cmd.Stderr = os.Stderr

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	file, err := os.Create("test.pcm")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, pipe)
	if err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
