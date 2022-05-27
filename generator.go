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
	perms := make([]*module, len(modules))
	for i := range modules {
		perms[i] = &modules[i]
	}
	// Shuffle modules
	for i := int64(len(perms)) - 1; i > 0; i-- {
		s, err := rand.Int(rand.Reader, big.NewInt(i))
		if err != nil {
			panic(err)
		}
		y := s.Int64()
		perms[y], perms[i] = perms[i], perms[y]
	}
trials:
	for _, m := range perms {
		for _, a := range m.Affects {
			if affected[a] {
				continue trials
			}
		}
		m.Write(builder, t)
	}
	return builder.String()
}
