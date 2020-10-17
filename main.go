package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
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

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	parseSamples(pipe)

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

// type Sample = float64

func parseSamples(r io.Reader) {
	for {
		b := make([]byte, 32) // 32 because we're using float32
		_, err := r.Read(b)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(math.Float32frombits(binary.LittleEndian.Uint32(b)))
	}
}
