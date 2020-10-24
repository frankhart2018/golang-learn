// Lissajous generates GIF animations of random Lissajous figures and shows them in web browser
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var pallete = []color.Color{color.White, color.Black} // This is a slice

// Like var const can also appear at package level, the value of constant can only be string, boolean or number
const (
	whiteIndex = 0 // first color in pallete
	blackIndex = 1 // next color in pallete
)

func main() {
	// This is called a function literal
	handler := func(w http.ResponseWriter, r *http.Request) {
		keys, _ := r.URL.Query()["cycles"]
		key := keys[0]
		param, _ := strconv.Atoi(key)
		lissajous(w, param)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size . . +size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // this is a struct with LoopCount value = nframes
	phase := 0.0                        // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
