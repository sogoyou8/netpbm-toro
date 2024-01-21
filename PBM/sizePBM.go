package PBM

import "fmt"

func (pbm *PBM) Size() (int, int) {
	fmt.Println("Magic Number:", pbm.magicNumber)
	fmt.Println("Width:", pbm.width)
	fmt.Println("Heigt:", pbm.height)
	return pbm.height, pbm.width
}
