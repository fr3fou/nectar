package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"time"

	"gonum.org/v1/gonum/dsp/fourier"
)

func _main() {
	// c := gusic.Chord{
	// 	gusic.A(4, time.Second, 1),
	// 	gusic.B(4, time.Second, 1),
	// }
}

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

	log.Println("recording...in 3 sec")
	time.Sleep(1 * time.Second)
	log.Println("recording...in 2 sec")
	time.Sleep(1 * time.Second)
	log.Println("recording...in 1 sec")
	time.Sleep(1 * time.Second)
	log.Println("recording...")
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	samples := parseSamples(pipe, 44100*2) // 2 secs of audio

	if err := cmd.Process.Kill(); err != nil {
		panic(err)
	}
	log.Println("finished recording...")

	f, err := os.Create("naive.csv")
	if err != nil {
		panic(err)
	}

	log.Println("computing naive dft...")
	out := dft(samples)
	for _, v := range out {
		fmt.Fprintf(f, "%.02f,%.02f\n", real(v)*100, imag(v)*100)
	}

	g, err := os.Create("gonum.csv")
	if err != nil {
		panic(err)
	}

	log.Println("computing gonum fft...")
	out = fourier.NewFFT(len(samples)).Coefficients(nil, samples)
	for _, v := range out {
		fmt.Fprintf(g, "%.02f,%.02f\n", real(v)*100, imag(v)*100)
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

func dft(samples []Sample) []complex128 {
	output := []complex128{}
	input := []complex128{}

	for _, v := range samples {
		input = append(input, complex(v, 0))
	}

	N := float64(len(input))
	for i := 0.0; i < N/2; i++ {
		var c complex128
		for n, x := range input {
			c += x * complex(
				math.Cos(((2.0*math.Pi)/N*i*float64(n))),
				-math.Sin(((2.0*math.Pi)/N*i*float64(n))),
			)
		}
		output = append(output, c)
	}

	return output
}
