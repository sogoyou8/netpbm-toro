package PGM

import "fmt"

func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
	fmt.Printf("%s", pgm.magicNumber)
	return
}