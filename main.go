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
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	samples := parseSamples(pipe, 44100*5) // 5 secs of audio

	if err := cmd.Process.Kill(); err != nil {
		panic(err)
	}

	log.Println("computing dft...")
	out := dft(samples)
	for _, v := range out {
		fmt.Printf("%.02f,%.02f\n", real(v)*100, imag(v)*100)
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
	for i := 0.0; i < N; i++ {
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
