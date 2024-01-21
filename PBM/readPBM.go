package PBM

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func parsePixel(value string) (bool, error) {
	switch value {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("unexpected pixel value: %s", value)
	}
}

func ReadPBM(filename string) (*PBM, error) {
	//lire le fichier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//identifier le magic number
	var magicNumber string
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading magic number")
	}
	if _, err := fmt.Sscanf(scanner.Text(), "%s", &magicNumber); err != nil {
		return nil, err
	}

	//lire les taille width et height
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading width and height")
	}

	var w, h int
	if _, err := fmt.Sscanf(scanner.Text(), "%d %d", &w, &h); err != nil {
		return nil, err
	}
	// le nouveau pbm pour l'afficher
	pbm := &PBM{
		magicNumber: magicNumber,
		width:       w,
		height:      h,
	}

	//affiche l'image ( data h w )
	if pbm.magicNumber == "P1" {
		// Lecture du fichier ligne par ligne
		pbm.data = make([][]bool, pbm.height)
		for y := 0; y < pbm.height; y++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("error reading row %d: %v", y, scanner.Err())
			}
			row := make([]bool, pbm.width)
		// Extraction des valeurs des pixels
			values := strings.Fields(scanner.Text())
			if len(values) != pbm.width {
				return nil, fmt.Errorf("unexpected number of values in row %d, expected %d, got %d", y, pbm.width, len(values))
			}
			// Conversion des valeurs en booléens
			for x, val := range values {
				pixel, err := parsePixel(val)
				if err != nil {
					return nil, fmt.Errorf("error parsing pixel at (%d, %d): %v", x, y, err)
				}
				row[x] = pixel
			}
			pbm.data[y] = row
		}

	} else if pbm.magicNumber == "P4" {
		// Lecture des données binaires du fichier
		binaryData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("error reading binary data: %v", err)
		}

		// Calcul du nombre de bits attendus
		expectedBits := pbm.width * pbm.height
		// Vérification du nombre de bits réels
		actualBits := len(binaryData) * 8

		if actualBits != expectedBits {
			return nil, fmt.Errorf("error reading sizes, expected %d bits, got %d bits", expectedBits, actualBits)
		}

		    // Initialisation de la matrice de données
		pbm.data = make([][]bool, pbm.height)
		for y := 0; y < pbm.height; y++ {
			row := make([]bool, pbm.width)
		    // Extraction des bits correspondant aux pixels
			for x := 0; x < pbm.width; x++ {
				// Calcul de l'index du bit
				bitIndex := y*pbm.width + x
				// Calcul de l'index du byte et du bit dans le byte
				byteIndex := bitIndex / 8
				bitOffset := 7 - (bitIndex % 8)
				row[x] = (binaryData[byteIndex]>>bitOffset)&1 == 1
			}
			// Extraction du bit
			pbm.data[y] = row
		}
	} else {
		return nil, fmt.Errorf("invalid magic number : %s", pbm.magicNumber)
	}

	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()

	}
	return pbm, nil
}
