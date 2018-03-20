/*

 Filename: scanner.go
 Authors: Kyle A. Kreutzer, Colin J. Mills
 Date: March 20th 2018
 */

package main 

import "github.com/veandco/go-sdl2/sdl"

type color struct {
	r, g, b byte
}

const WINDOW_WIDTH int = 800
const WINDOW_HEIGHT int = 600

func SetPixel(x,y int,  currentColor color,  pixels []byte) {

	currentPixelIndex := (y * WINDOW_WIDTH + x) * 4

	if currentPixelIndex < len(pixels) - 4 && currentPixelIndex >= 0 {
		pixels[currentPixelIndex] = currentColor.r
		pixels[currentPixelIndex + 1] = currentColor.g
		pixels[currentPixelIndex + 2] = currentColor.b
	}
	
}

func main() {
	PITCH := 4 // We have four bytes per pixel because we are using int32 

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow("MAKING GAMES WITH GO MAN!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	int32(WINDOW_WIDTH), int32(WINDOW_HEIGHT), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if (err != nil) {
		panic(err)
	}
	defer renderer.Destroy()
	 

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(WINDOW_WIDTH), int32(WINDOW_HEIGHT))
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	pixelArraySize := WINDOW_WIDTH * WINDOW_HEIGHT * 4
	pixels := make([]byte, pixelArraySize)

	// Get Keyboard State
	keyboardState := sdl.GetKeyboardState()

	for {

	// Events 
	for  e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {

		case *sdl.QuitEvent:
			return
		}
	}

	bPressed := keyboardState[sdl.SCANCODE_B] != 0
	rPressed := keyboardState[sdl.SCANCODE_R] != 0
	gPressed := keyboardState[sdl.SCANCODE_G] != 0

	c := color{255, 255, 255}

	if bPressed {
		c = color{0, 0, 255}
	}
	
	if rPressed {
		c = color{255, 0, 0}
	}
	
	if gPressed {
		c = color{0, 255, 0 }
	}

	for y := 0; y < WINDOW_HEIGHT; y++ {
		for x := 0; x < WINDOW_WIDTH; x++ {
			SetPixel(x,  y, c, pixels) 
		}
	}

	texture.Update(nil, pixels, WINDOW_WIDTH * PITCH)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	sdl.Delay(16)
	}
	
}
