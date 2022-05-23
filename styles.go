/*
This file is part of vdcsspp
Copyright 2022 epiccakeking
Licensed under 0BSD
*/

package main

import "strings"

func WriteSeveral(w *strings.Builder, s ...string) {
	for i := range s {
		w.WriteString(s[i])
	}
}

/*
All of the possible modules
*/
var modules = []module{
	{
		[]string{"#header"},
		func(w *strings.Builder, t *theme) { // Match content style
			w.WriteString("#header{margin: 20pt;")
			t.BgStyle(w, t)
			WriteSeveral(w, "}#header h1{color:", t.Fg.String(), "}")
		},
	}, {
		[]string{"#header"},
		func(w *strings.Builder, t *theme) { // Plain rectangle with background color
			WriteSeveral(w, "#header{background-color:", t.Bg.String(), ";border:none}")
		},
	}, {
		[]string{"#message"},
		func(w *strings.Builder, t *theme) { // Add backgrounds to the side elements
			w.WriteString("#message{")
			t.BgStyle(w, t)
			WriteSeveral(w, ";color:", t.Fg.String(), "}.blog-shout{")
			t.BgStyle(w, t)
			WriteSeveral(w, ";color:", t.Fg.String(), ";margin:10pt 0}")
		},
	}, {
		[]string{"#entry"},
		func(w *strings.Builder, t *theme) {
			w.WriteString("div .entry{border:none;margin: 30pt 10pt;")
			t.BgStyle(w, t)
			w.WriteString(";color:")
			w.WriteString(t.Fg.String())
			w.WriteString("}div .entrywrap{border:none;background:none}")
		},
	}, {
		[]string{"#entry"},
		func(w *strings.Builder, t *theme) {
			w.WriteString("div .entry{border:none;margin: 30pt 10pt}div .entrywrap{")
			t.BgStyle(w, t)
			w.WriteString(";color:")
			w.WriteString(t.Fg.String())
			w.WriteString("}")
		},
	}, {
		[]string{"#wrapper"},
		func(w *strings.Builder, t *theme) { // Use default layout
		},
	}, {
		[]string{"#wrapper"},
		func(w *strings.Builder, t *theme) { // Wide layout
			w.WriteString("#wrapper{width:100%;padding:0;border:none}#content{display:flex;flex-flow:nowrap row;}#main{flex-grow:1;padding:10pt}#side{flex-grow:0;float:none}")
		},
	},
}

/*
Backgrounds/border styles
*/
var backgroundStyles = []style{
	func(w *strings.Builder, t *theme) { // Rounded
		WriteSeveral(w, "border:none;border-radius:", t.CornerSize.String(), ";background:", t.Bg.String())
	},
	func(w *strings.Builder, t *theme) { // Sliced
		WriteSeveral(w,
			"border:none;background:linear-gradient(45deg,transparent ", t.CornerSize.String(), ",",
			t.Bg.String(), " ", t.CornerSize.String(), ",",
			t.Bg.String(), " calc(100% - ", t.CornerSize.String(), "),transparent calc(100% - ", t.CornerSize.String(), "))")
	},
	func(w *strings.Builder, t *theme) { // Rectangle
		WriteSeveral(w, "border:none;background:", t.Bg.String())
	},
}
