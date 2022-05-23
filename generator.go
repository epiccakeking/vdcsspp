/*
This file is part of vdcsspp
Copyright 2022 epiccakeking
Licensed under 0BSD
*/

package main

import (
	"math/big"
	"strings"

	"crypto/rand"
)

func choice[T any](l []T) T {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(l))))
	if err != nil {
		panic(err)
	}
	return l[n.Int64()]
}

type module struct {
	Affects []string                           // Elements/tags that this style affects, used to detect clashing styles
	Write   func(w *strings.Builder, t *theme) // Write the style to the string builder
}

type style func(w *strings.Builder, t *theme) // Writes a style to the string builder

type CssColor interface {
	String() string
}

type NamedColor string

func (n NamedColor) String() string {
	return string(n)
}

type theme struct {
	Fg, Bg, BodyBg CssColor
	CornerSize     Unit
	BgStyle        style
}

func Generate() string {
	builder := new(strings.Builder)
	t := new(theme)
	t.Fg, t.Bg = getColorPair()
	t.BodyBg = NamedColor("black")
	t.BgStyle = choice(backgroundStyles)
	t.CornerSize = Unit{10, "pt"}
	// Add random styles until we can't find any new ones
	affected := make(map[string]bool)
trials:
	for i := 0; i < 10; i++ {
		m := choice(modules)
		for _, a := range m.Affects {
			if affected[a] {
				continue trials
			}
		}
		m.Write(builder, t)
		for _, a := range m.Affects {
			affected[a] = true
		}
		i = 0 // Reset the remaining trial count
	}
	return builder.String()
}
