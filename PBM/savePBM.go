package PBM

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (pbm *PBM) Save(filename string) error {
	//Genere un nom a chaque asuvegarde "Monday, August 23, 2023 at 12:34""Jeudi, 4 Aout, 2001 à 12:34"
	fileDate := time.Now().Format("2006-01-02_15-04-05")
	// supprimer l'extension ".pbm" du nom de fichier fourni.
	newFilename := fmt.Sprintf("%s_%s", strings.TrimSuffix(filename, ".pbm"), fileDate)

	//Sauvegarde
	file, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	// Gestion des erreurs
	defer file.Close()
	// Création d'un objet de `bufio.Writer`
	writer := bufio.NewWriter(file)
	// Écriture des données de tête du fichier PBM
	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)
	// Écriture des données du fichier PBM 
	if pbm.magicNumber == "P1" {
		// Écrit les données du fichier PBM en format texte (un caractère par pixel)
		for x := 0; x < pbm.width; x++ {
			for y := 0; y < pbm.height; y++ {
				if pbm.data[x][y] {
					fmt.Fprint(file, "0 ")
				} else {
					fmt.Fprint(file, "1 ")
				}
			}
			// ligne par ligne
			fmt.Fprintln(file)
		}
	} else if pbm.magicNumber == "P4" {
		// Écrit les données du fichier PBM en format binaire (un octet par pixel)
		for x := 0; x < pbm.width; x++ {
			for y := 0; y < pbm.height; y++ {
				if pbm.data[x][y] {
    				 // Écrit un octet contenant la valeur "0"
					if _, err := file.Write([]byte{0}); err != nil {
						return err
					}
				} else {
    				// Écrit un octet contenant la valeur "1" 
					if _, err := file.Write([]byte{1}); err != nil {
						return err
					}
				}
			}
		}
	}
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
