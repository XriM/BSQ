//
// EPITECH PROJECT, 2021
// BSQ (Golang)
// File description: BSQ in Go
// maxime.sarres@epitech.eu
//

package main

import (
	"fmt"
)

func printHelp() {
	fmt.Println("Usage:       ./bsq \"myFile\"\n\nConstraints: File must exist, format with \".\" and \"o\" only")
}

func printMap(b [][]byte) {
	for i := 0; i < len(b); i++ {
		fmt.Println(string(b[i]))
	}
}
