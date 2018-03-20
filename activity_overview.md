# Activity Overview: Go and SDL

# Presented By:

## Colin J. Mills
## Kyle A. Kreutzer

# Introduction:

## Assumptions:

This tutorial was built to run on Ubuntu 16.04. 


## Install Needed Packages

NOTE: All packages can be installed by using the `install.sh` script int he repo.

First need to add a ppa to be able to have the go runtime available to install.

`sudo add-apt-repository ppa:longsleep/golang-backports`

`sudo apt-get update`

`sudo apt-get install golang-go`

Now with go install install libsdl to be able to use SDL with Go.

`sudo apt-get install libsdl2-dev`

Now we need to be able to instal the bindings for Go to SDL. This has been provided by the lovely github package.

`go get github.com/veandco/go-sdl2/sdl`


# The Meat and Potatoes


## Step 1: Setting up a blank Go program

The first step is to build a blank Go program.

```

package main

import "github.com/veandco/go-sdl2/sdl"

func main() {

}

```

## Step 1: Initing SDL

The first step to making an application with SDL is to init SDL with the subset of functionality that one wants. In our case we want everything. Therefore we say we want everything. Plus we want to defer the quit to make sure we do it.

```

    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }

    defer sdl.Quit()
```


## Step 2: Opening A Window 

Next step is to open a window using SDL2. This involves creating a window and defining a couple variables that we will need to use afterwards as well. These include WINDOW_WIDTH and WINDOW_HEIGHT:

```

const WIDNOW_WIDTH int = 800
const WINDOW_HEIGHT int = 800

```


We then have to create the actual window, which follows the same convention as we saw before hand:

```
    window, err := sdl.CreateWindow("MAKING GAMES WITH GO MAN!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
    int32(WINDOW_WIDTH), int32(WINDOW_HEIGHT), sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()
```

Again we defer the destroy, this would be similar to using a free() call at the end of a C program

## Step 3: Creating a Renderer For the Window

We now have to create a renderer to be able to put pixels on the screen. A renderer is SDL\'s mechanism for being able to display textures.

```
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if (err != nil) {
		panic(err)
	}
	defer renderer.Destroy()
```

Again we see that the convention of Go returning an error code, is very common.


## Step 4: Creating a Texture to set pixels:

A texture is a way for SDL to be able to accurately display pixels. We specify our texture\'s layout, which in this case is a 4 byte layout with enough space for Alpha Red Green and Blue bytes. This makes it simple to be able to specify colors as bytes rather then any other conventional mechanism.

```
    texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(WINDOW_WIDTH), int32(WINDOW_HEIGHT))
    if err != nil {
        panic(err)
    }
    defer texture.Destroy()
```


## Step 5: Creating Our Pixel Buffer

Now we have to make a pixel buffer to be able to BLIT or DRAW to the screen. This will imposed over our texture which will render it in the same format as we specified (ARGB). This is why we allocate * 4 on the end of our pixel array size, to account for the actual byte size. If this was C it would be similar to something like:

``` WINDOW_WIDTH * WINDOW_HIEGHT * sizeof(uint32) ```

This creates a slice which you can think of as a dynamic pool of memory. An analogous to C would be a malloc(size)

```
    pixelArraySize := WINDOW_WIDTH * WINDOW_HEIGHT * 4
    pixels := make([]byte, pixelArraySize)
```


## Step 6: Setting Colors In Our Pixel Buffer

So the next step is to actually set some colors in our pixel buffer. To be able to do this we need to create a function which takes the x and y position of where we want to set, a new type color and our pixels.

The color type is defined as follows:

```
    type color struct {
    	r, g, b byte
    }
```



To actually get the postiion we can use some single array accessing math. We want whatever y postiion we are mutliplied by the width of that row. We then add the X postion and account for our byte size which again is 4.

We then do some simple checking to see if we are in bounds and we set the pixels accordingly, 3 at a time with the rgb values. Note we are ignoring alpha for now.

```
func SetPixel(x,y int,  currentColor color,  pixels []byte) {

	currentPixelIndex := (y * WINDOW_WIDTH + x) * 4

	if currentPixelIndex < len(pixels) - 4 && currentPixelIndex >= 0 {
		pixels[currentPixelIndex] = currentColor.r
		pixels[currentPixelIndex + 1] = currentColor.g
		pixels[currentPixelIndex + 2] = currentColor.b
	}
	
}
```


## Step 6: Setting up our basic game loop:

Our basic game loop  is pretty simple. We first set all of our pixels a desired color, we do this by using a loop to loop through all the pixels.   We Update our texture with the however oiur pixels have been process that time around, and then we copy the texture on to the renderer and then finally the renderer gets updated. 

```
c := color{255, 255, 255}

for {
 
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
```













