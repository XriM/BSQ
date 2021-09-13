//
// EPITECH PROJECT, 2021
// BSQ (Golang)
// File description: BSQ in Go
// maxime.sarres@epitech.eu
//

package main

import (
	"io/ioutil"
	"log"
	"os"
)

/////////////////////////////////////////-
///////
///////      format
///////
/////////////////////////////////////////-

func getArg() string {
	arg := os.Args
	exist := realPath(string(arg[1]))
	if exist == false {
		log.Fatal("file wrong")
	}
	file, err := ioutil.ReadFile(arg[1])
	check(err)
	return string(file)
}

func reHeading(myMap string) string {
	j := 1
	for i := 0; i != len(myMap); i++ {
		if myMap[i] != 'o' && myMap[i] != '.' && myMap[i] != '\n' {
			j++
		}
	}

	myMap = myMap[j:]
	return myMap
}
