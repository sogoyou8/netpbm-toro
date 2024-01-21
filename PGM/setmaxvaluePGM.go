package PGM

import "fmt"

// SetMaxValue sets the maximum value of the PGM image pixels.
func (pgm *PGM) SetMaxValue(maxValue uint8) {
	oldMax := pgm.max
	pgm.max = int(maxValue)

	for y := 0; y < pgm.height; y++ {
		for x := 0; x < pgm.width; x++ {
			pgm.data[y][x] = uint8(float64(pgm.data[y][x]) * float64(pgm.max) / float64(oldMax))
		}
		fmt.Println(pgm.max)
	}
}
