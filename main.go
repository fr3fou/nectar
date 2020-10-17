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

	samples := parseSamples(pipe, 44100*5) // 5 secs of audio

	if err := cmd.Process.Kill(); err != nil {
		panic(err)
	}

	file, err := os.Create("output.pcm")
	if err != nil {
		panic(err)
	}

	fmt.Println(samples)

	if err := binary.Write(file, binary.LittleEndian, samples); err != nil {
		panic(err)
	}
}

type Sample = float64

func parseSamples(r io.Reader, limit int) []Sample {
	v := []Sample{}

	for i := 0; i < limit; i += 44100 {
		b := make([]byte, 4) // 32 because we're using float32
		_, err := r.Read(b)
		if err != nil {
			log.Println(err)

			break
		}
		v = append(v, float64(math.Float32frombits(binary.LittleEndian.Uint32(b))))
	}

	return v
}
