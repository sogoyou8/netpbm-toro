package PGM

import "fmt"

func (pgm *PGM) Set(x, y int, value uint8) {
	if y >= 0 && y < pgm.height && x >= 0 && x < pgm.width {
		fmt.Println("Valeur definit du pixel: ", value)
		//definitla valeur du pixel specifié
		pgm.data[x][y] = value;
		return
	}
}