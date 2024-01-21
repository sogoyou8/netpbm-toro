package PGM

import "fmt"

func (pgm *PGM) Flop() {
    for x := pgm.width - 1; x >= 0; x-- {
        for y := 0; y < pgm.height; y++ {
            if pgm.magicNumber == "P2" {
				fmt.Print(uint8(pgm.max)-pgm.data[x][y], " ")
				pgm.data[x][y] = uint8(pgm.max) - pgm.data[x][y]
			}
        }
        fmt.Println()
    }
}
