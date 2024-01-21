package PPM

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (ppm *PPM) Save(filename string) error {
	//Genere un nom a chaque asuvegarde "Monday, August 23, 2023 at 12:34""Jeudi, 4 Aout, 2001 à 12:34"
	fileDate := time.Now().Format("2006-01-02_15-04-05")
	// supprimer l'extension ".ppm" du nom de fichier fourni.
	newFilename := fmt.Sprintf("%s_%s", filename, fileDate+".ppm")

	//Sauvegarde
	file, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	// Gestion des erreurs
	defer file.Close()
	// Création d'un objet de `bufio.Writer`
	writer := bufio.NewWriter(file)
	// Écriture des données de tête du fichier ppm
	fmt.Fprintf(file, "%s\n%d %d\n", ppm.magicNumber, ppm.width, ppm.height)
	// Écriture des données du fichier ppm
	if ppm.magicNumber == "P3" {
		max := strconv.Itoa(int(ppm.max))
		fmt.Fprintln(file, max)
		// Écrit les données du fichier ppm en format texte (un caractère par pixel)
		for x := 0; x < ppm.width; x++ {
			for y := 0; y < ppm.height; y++ {
				fmt.Fprint(file, ppm.data[x][y].R, " ")
				fmt.Fprint(file, ppm.data[x][y].G, " ")
				fmt.Fprint(file, ppm.data[x][y].B, " ")

			} // ligne par ligne
			fmt.Fprintln(file)
		}
	}
	// else if ppm.magicNumber == "P6" {
	// 	// Écrit les données du fichier ppm en format binaire (un octet par pixel)
	// 	for x := 0; x < ppm.width; x++ {
	// 		for y := 0; y < ppm.height; y++ {
	// 			if ppm.data[x][y] == 1 {
	// 				// Écrit un octet contenant la valeur "0"
	// 				if _, err := file.Write([]byte{0}); err != nil {
	// 					return err
	// 				}
	// 			} else {
	// 				// Écrit un octet contenant la valeur "1"
	// 				if _, err := file.Write([]byte{1}); err != nil {
	// 					return err
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	if err != nil {
		return err
	}
	//assure que toutes les données sont écrites dans le fichier.
	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Printf("nom du fichier:%s\n", newFilename)
	return nil
}
