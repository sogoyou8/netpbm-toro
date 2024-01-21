package PPM

import "fmt"

func (ppm *PPM) Size() (int, int) {
	fmt.Println("Magic Number:", ppm.magicNumber)
	fmt.Println("Width:", ppm.width)
	fmt.Println("Heigt:", ppm.height)
	return ppm.height, ppm.width
}
