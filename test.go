package main

import (
	"fmt"
	"netpbm_test/PBM"
	"netpbm_test/PGM"
	"netpbm_test/PPM"
)

func main() {
	testPBM()
	testPGM()
	testPPM()
}
func testPBM() {

	image, err := PBM.ReadPBM("test.pbm")
	if err != nil {
		fmt.Println("Error reading PBM file:", err)
		return
	}

	fmt.Println("taille de l'image ")
	image.Size()
	fmt.Print("\n")

	fmt.Println("verification pixel true/false ")
	image.At(0, 11)
	fmt.Print("\n")

	fmt.Println("attribution d'une valeur au pixel")
	image.Set(0, 0, true)
	fmt.Print("\n")

	fmt.Print(" inversion couleurs \n")
	image.Invert()
	//image.InvertB()
	fmt.Print(" \n")

	fmt.Println(" inversion horizontale")
	image.Flip()
	fmt.Print("\n")

	fmt.Println(" inversion verticale")
	image.Flop()
	fmt.Print("\n")

	fmt.Println(" definition du nombre magic")
	image.SetMagicNumber("P1")
	fmt.Print("\n")

	fmt.Println("sauvegarde")
	image.Save("test.pbm")
	fmt.Print("\n")
	//fmt.Println(image) //print true ou false
}

func testPGM() {
	image, err := PGM.ReadPGM("test.pgm")
	if err != nil {
		fmt.Println("Error reading PGM file:", err)
		return
	}

	//fmt.Println(image) //print true ou false
	fmt.Println("taille de l'image ")
	image.Size()
	fmt.Print("\n")

	fmt.Println("verification pixel true/false ")
	image.At(0, 11)
	fmt.Print("\n")

	fmt.Println("attribution d'une valeur au pixel")
	image.Set(1, 0, 5)
	fmt.Print("\n")

	fmt.Println(" inversion couleurs ")
	image.Invert()
	fmt.Print(" \n")

	fmt.Println(" inversion horizontale")
	image.Flip()
	fmt.Print("\n")

	fmt.Println(" inversion verticale")
	image.Flop()
	fmt.Print("\n")

	fmt.Println(" definition du nombre magic")
	image.SetMagicNumber("P2")
	fmt.Print("\n")

	fmt.Println("rotation 90 cw")
	image.Rotate90CW()
	fmt.Print("\n")

	fmt.Println(" definition valeur max")
	image.SetMaxValue(255)
	//image.SetMaxValue(uint8(strconv.Atoi("255")))
	fmt.Print("\n")

	fmt.Println("sauvegarde")
	image.Save("test.pgm")
	fmt.Print("\n")
}

func testPPM() {
	image, err := PPM.ReadPPM("test.ppm")
	if err != nil {
		fmt.Println("Error reading PGM file:", err)
		return
	}

	fmt.Println("taille de l'image : ")
	image.Size()
	fmt.Print("\n")

	fmt.Println("verification pixel true/false ")
	image.At(0, 11)
	fmt.Print("\n")

	fmt.Println("attribution d'une valeur au pixel")
	image.Set(1, 0, PPM.Pixel{R: 0, G: 23, B: 240})
	fmt.Print("\n")

	fmt.Println(" inversion horizontale")
	image.Flip()
	fmt.Print("\n")

	fmt.Println(" inversion verticale")
	image.Flop()
	fmt.Print("\n")

	fmt.Println("print data : ")
	image.PrintData()
	fmt.Print("\n")

	fmt.Println(" definition du nombre magic")
	image.SetMagicNumber("P2")
	fmt.Print("\n")

	fmt.Println(" definition du nombre magic")
	image.SetMagicNumber("P2")
	fmt.Print("\n")

	fmt.Println("rotation 90 cw")
	image.Rotate90CW()
	fmt.Print("\n")

	fmt.Println("sauvegarde")
	image.Save("test.ppm")
	fmt.Print("\n")

}