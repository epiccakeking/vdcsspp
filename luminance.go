// Code derived from https://www.w3.org/TR/WCAG21/#dfn-relative-luminance
// Copyright © 2017-2018 W3C® (MIT, ERCIM, Keio, Beihang). W3C liability, trademark and document use rules apply.
package main

import "math"

func (c rgb) Luminance() float64 {
	var (
		R8bit, G8bit, B8bit = float64(c[0]), float64(c[1]), float64(c[2])
		RsRGB               = R8bit / 255
		GsRGB               = G8bit / 255
		BsRGB               = B8bit / 255
		R, G, B             float64
	)
	if RsRGB <= 0.03928 {
		R = RsRGB / 12.92
	} else {
		R = math.Pow((RsRGB+0.055)/1.055, 2.4)
	}
	if GsRGB <= 0.03928 {
		G = GsRGB / 12.92
	} else {
		G = math.Pow((GsRGB+0.055)/1.055, 2.4)
	}
	if BsRGB <= 0.03928 {
		B = BsRGB / 12.92
	} else {
		B = math.Pow((BsRGB+0.055)/1.055, 2.4)
	}
	return 0.2126*R + 0.7152*G + 0.0722*B
}
