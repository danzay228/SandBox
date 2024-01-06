package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func example1() {
	args := os.Args[1:]
	args2 := strings.Join(args, " ")
	fmt.Printf("%#v\n", args2)
}

var palette = []color.Color{color.White, color.Black}

func init() {
	fmt.Println("module initialize!")
}

func lissajous(w io.Writer) {
	const (
		whiteIndex = 0 // Первый цвет палитры
		blackIndex = 1 // Следующий цвет палитры

		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			у := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(у*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}

func example2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	http.ListenAndServe(":8000", nil)
}

func main() {
	example1()
	example2()
}
