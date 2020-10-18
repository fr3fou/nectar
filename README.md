# nectar - music note detector

Detecting musical notes from an audio input through the use of (Fast) Fourier Transforms in Go from scratch.

## TODO

- [x] Read mic output
- [x] Parse samples
- [x] Implement DFT
    - [x] Fix conjugate symmetry (ignore half the values)
        > Here we note that there is a symmetry to the graph. This is not a coincidence: if the input signal is real-valued, it will always be the case that the Fourier transform is symmetric about its center value. The reason for this goes back to our first primer on the Fourier series, in that the negative coefficients were complex conjugates of the positive ones. In any event, we only need concern ourselves with the first half of the values.
    - [ ] Fix values being twice as big
- [ ] Caclulate frequency
- [ ] Calculate the offset from A440

## References

- <https://www.youtube.com/watch?v=spUNpyF58BY>
- <https://www.wikiwand.com/en/Discrete_Fourier_transform>
- <https://jeremykun.com/2012/07/18/the-fast-fourier-transform/>
- <https://dsp.stackexchange.com/questions/4825/why-is-the-fft-mirrored>
