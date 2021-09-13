//
// EPITECH PROJECT, 2021
// BSQ (Golang)
// File description: BSQ in Go
// maxime.sarres@epitech.eu
//

package main

import (
	"log"
	"os"
	"strings"
)

/////////////////////////////////////////-
///////
///////      main
///////
/////////////////////////////////////////-

func main() {

	if len(os.Args) != 2 {
		log.Fatal("wrong args nbr")
	} else if os.Args[1] == "-h" {
		printHelp()
		os.Exit(0)
	}

	myMap := getArg()
	myMap = reHeading(myMap)
	oriMap := strings.Split(myMap, "\n")
	myMap = parsing(myMap)
	myMapS := strings.Split(myMap, "\n")
	myMapS = monitoring(myMapS, oriMap)
}
