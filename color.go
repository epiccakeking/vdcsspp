/*
This file is part of vdcsspp
Copyright 2022 epiccakeking
Licensed under 0BSD
*/

package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

type rgb [3]byte

func (r rgb) String() string {
	return fmt.Sprintf("#%x", [3]byte(r))
}

func checkContrast(a rgb, b rgb, minContrast float64) (ok bool) {
	var (
		l1 = a.Luminance()
		l2 = b.Luminance()
	)
	if l1 < l2 {
		return (l2+.05)/(l1+.05) > minContrast
	}
	return (l1+.05)/(l2+.05) > minContrast
}

func colorRGBToStr(c rgb) string {
	return fmt.Sprintf("#%x", c)
}

func getColorPair() (rgb, rgb) {
	var t [2]rgb
	for {
		binary.Read(rand.Reader, binary.BigEndian, &t)
		if checkContrast(t[0], t[1], 7) {
			return t[0], t[1]
		}
	}
}

//getColorPair with another third which is compliant with the first
func getColorTriple() (rgb, rgb, rgb) {
	a, b := getColorPair()
	var c rgb
	for {
		binary.Read(rand.Reader, binary.BigEndian, &c)
		if checkContrast(a, c, 7) {
			return a, b, c
		}
	}
}
