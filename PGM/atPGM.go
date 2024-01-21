package PGM

import "fmt"

func (pgm *PGM) At(x, y int ) uint8 {
	if y >= 0 && y < pgm.height && x >= 0 && x < pgm.width {
		fmt.Println("Choix du pixel: ", pgm.data[x][y])
		return pgm.data[x][y]
	}
	return 0
}