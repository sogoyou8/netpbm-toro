package PBM

import "fmt"

func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
	fmt.Printf("%s", pbm.magicNumber)
	return
}
