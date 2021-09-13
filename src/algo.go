//
// EPITECH PROJECT, 2021
// BSQ (Golang)
// File description: BSQ in Go
// maxime.sarres@epitech.eu
//

package main

import (
	"strconv"
	"strings"
)

/////////////////////////////////////////-
///////
///////      algorithm
///////
/////////////////////////////////////////-

func parsing(myMap string) string {
	myMap = strings.Replace(myMap, "o", "0", -1)
	myMap = strings.Replace(myMap, ".", "1", -1)
	return myMap
}

func getLowerRect(i int, j int, a [][]int) int {
	bg, hd, hg := 0, 0, 0

	bg = a[i-1][j]
	hd = a[i][j-1]
	hg = a[i-1][j-1]

	if bg <= hd && bg <= hg {
		return bg
	}
	if hd <= bg && hd <= hg {
		return hd
	}
	if hg <= bg && hg <= hd {
		return hg
	}
	return -1
}

func mapToInt(maps []string, b [][]byte, a [][]int) ([][]int, [][]byte) {

	for i := 0; i < len(maps); i++ {
		b[i] = []byte(maps[i])
		for j := 0; j < len(maps[i]); j++ {
			nb, err := strconv.Atoi(string(maps[i][j]))
			if err != nil {
				panic(err)
			}
			a[i] = append(a[i], nb)
		}
	}
	return a, b
}

func algo(a [][]int, b [][]byte, res int) ([][]int, [][]byte, int) {

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			if a[i][j] != 0 {
				res := getLowerRect(i, j, a)
				a[i][j] += res
				b[i][j] = byte(a[i][j]) + 48
			}
		}
	}
	return a, b, res
}

func getRectPos(a [][]int, res int, posx int, posy int) ([][]int, int, int, int) {

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			if a[i][j] > res {
				res = a[i][j]
				posx = i
				posy = j
			}
		}
	}
	return a, res, posx, posy
}

func monitoring(maps []string, oriMaps []string) []string {
	a := make([][]int, len(maps))
	b := make([][]byte, len(maps))
	posx, posy, res := 0, 0, 0

	a, b = mapToInt(maps, b, a)
	a, b, res = algo(a, b, res)
	a, res, posx, posy = getRectPos(a, res, posx, posy)

	for i := 0; i < len(maps); i++ {
		b[i] = []byte(oriMaps[i])
	}
	for resx, resy, pos := res, res, posy; resx != 0; resx-- {
		for ; resy != 0; resy-- {
			b[posx][posy] = 'x'
			posy -= 1
		}
		posx -= 1
		posy = pos
		resy = res
	}

	printMap(b)

	return maps
}
