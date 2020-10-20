package main

import (
	"math"
	"strconv"
)

const twelfthrootof2 float64 = 1.059463094359

var (
	a4 = 440
	// https://github.com/fr3fou/gusic/blob/72a7e32d5644ed6d123e365d416fdca51a268161/gusic/step.go#L38
	// c0    = float64(a4) * math.Pow(twelfthrootof2, float64(-4*12-9))
	c0    = float64(a4) * math.Pow(2, -4.75)
	notes = []string{
		"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
	}
)

func note(freq float64) string {
	h := int(math.Round(12 * math.Log2(freq/c0)))
	octave := h / 12

	return notes[h%12] + strconv.Itoa(octave)
}
