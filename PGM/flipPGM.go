package PGM

import "fmt"

func (pgm *PGM) Flip() {
	for x := 0; x < pgm.width; x++ {
		for y := pgm.height - 1; y >= 0; y-- {
			if pgm.magicNumber == "P2" {
				fmt.Print(uint8(pgm.max)-pgm.data[x][y], " ")
				pgm.data[x][y] = uint8(pgm.max) - pgm.data[x][y]

			
			} else {
				if pgm.data[x][y] == 1 {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
		}
		fmt.Println()
	}
}
