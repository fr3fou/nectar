package main

import (
	"fmt"

	"github.com/lawl/pulseaudio"
)

func main() {
	c, err := pulseaudio.NewClient()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	s, err := c.Sources()
	if err != nil {
		panic(err)
	}

	for _, source := range s {
		fmt.Println(source)
	}
}
