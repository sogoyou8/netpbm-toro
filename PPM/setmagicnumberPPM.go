package PPM

import "fmt"

func (ppm *PPM) SetMagicNumber(magicNumber string) {
	ppm.magicNumber = magicNumber
	fmt.Printf("%s", ppm.magicNumber)
}