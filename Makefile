default: build

build:
	go build .

prepare:
	brew install sdl2 sdl2_gfx sdl2_image sdl2_mixer sdl2_ttf
