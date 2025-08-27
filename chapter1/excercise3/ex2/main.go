package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF},
	color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xBB},
	color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xAA},
	color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xCC},
}

const (
	blackIndex  = 0 // first color in palette
	GreenIndex  = 1 // next color in palette
	BlueIndex   = 2 // next color in palette
	RedIndex    = 3 // next color in palette
	YellowIndex = 4 // next color in palette
)

func main() {
	f, err := os.Create("out.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create file: %v\n", err)
		return
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 7     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		ran := rand.Intn(4) + 1 // random color index between 1 and 4
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if ran == 1 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					RedIndex)
			} else if ran == 2 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					BlueIndex)
			} else if ran == 3 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					YellowIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					GreenIndex)
			}
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
