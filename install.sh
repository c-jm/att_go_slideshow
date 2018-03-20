
#!/bin/sh

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go

sudo apt-get install libsdl2-dev

go get github.com/veandco/go-sdl2/sdl
