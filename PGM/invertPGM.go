package PGM

import "fmt"

func (pgm *PGM) Invert() {
	for x := 0; x < pgm.width; x++ {
		for y := 0; y < pgm.height; y++ {
			if pgm.magicNumber == "P2" {
				fmt.Print(uint8(pgm.max)-pgm.data[x][y], " ")
				pgm.data[x][y] = uint8(pgm.max) - pgm.data[x][y]

			}
		}
		fmt.Println()
	}
}
