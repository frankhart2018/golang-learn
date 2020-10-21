// Change the Lissajous program's color pallete to green on black
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}} // This is a slice

// Like var const can also appear at package level, the value of constant can only be string, boolean or number
const (
	whiteIndex = 0 // first color in pallete
	greenIndex = 1 // next color in pallete
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
