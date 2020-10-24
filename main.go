package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"math/cmplx"
	"os"
	"os/exec"

	"gonum.org/v1/gonum/dsp/fourier"
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

	fft := fourier.NewFFT(44100)
	for {
		samples := parseSamples(pipe, 44100)

		out := fft.Coefficients(nil, samples)
		max := math.Inf(-1)
		freq := 0

		for i, v := range out {
			magnitude := cmplx.Abs(v)
			if magnitude >= max {
				max = magnitude
				freq = i
			}
		}

		// 30 is arbitrary, should be a value from a calibration of the microphone done beforehand
		if freq > 30 {
			fmt.Printf("Current note: %s\n", note(float64(freq)))
		}
	}
}

type Sample = float64

func parseSamples(r io.Reader, limit int) []Sample {
	v := []Sample{}

	b := make([]byte, 4) // 4 bytes because we're using float32 (32 bits = 4 bytes)
	for i := 0; i < limit; i++ {
		_, err := r.Read(b) // Read a single sample
		if err != nil {
			log.Println(err)

			break
		}
		v = append(v, float64(math.Float32frombits(binary.LittleEndian.Uint32(b))))
	}

	return v
}

// very slow (O(n^2))
// deprecated for gonum's fft
func dft(samples []Sample) []complex128 {
	output := []complex128{}
	input := []complex128{}

	for _, v := range samples {
		input = append(input, complex(v, 0))
	}

	N := float64(len(input))
	for k := 0.0; k < N/2; k++ {
		var c complex128
		for n, x := range input {
			c += x * complex(
				math.Cos(
					((2.0*math.Pi)/N)*k*float64(n),
				),
				-math.Sin(
					((2.0*math.Pi)/N)*k*float64(n),
				),
			)
		}
		output = append(output, c)
	}

	return output
}
