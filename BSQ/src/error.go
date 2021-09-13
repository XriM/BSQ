//
// EPITECH PROJECT, 2021
// BSQ (Golang)
// File description: BSQ in Go
// maxime.sarres@epitech.eu
//

package main

import "os"

/////////////////////////////////////////-
///////
///////      error/path checking
///////
/////////////////////////////////////////-

func realPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
