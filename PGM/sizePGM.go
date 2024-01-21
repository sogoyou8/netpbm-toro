package PGM

import "fmt"

func (pgm *PGM) Size() (int, int) {
	fmt.Println("Magic Number:", pgm.magicNumber)
	fmt.Println("Width:", pgm.width)
	fmt.Println("Heigt:", pgm.height)
	return pgm.height, pgm.width
}
