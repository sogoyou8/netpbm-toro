package PGM

import (
	"fmt"
)

func (pgm *PGM) Rotate90CW() {
	// Create a new temporary matrix to store the rotated image data
	tempData := make([][]uint8, pgm.height)
	for y := 0; y < pgm.height; y++ {
		tempData[y] = make([]uint8, pgm.width)
	}

	for x := 0; x < pgm.width; x++ {
        for y := 0; y < pgm.height; y++ {
		
			tempData[x][pgm.width-1-y] = pgm.data[y][x]
			// print du caractere + espace
			print(tempData[x][pgm.width-1-y], " ")
		}
		// retour a la ligne
		println()
	}

	// Copy the rotated image data back to the original matrix
	//pgm.data = tempData
	fmt.Println()
}
