package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatal(err)
	}

	defer ttf.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	err = DrawTitle("Hello Gophers", renderer)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 1)
	err = DrawBackground(renderer)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)
}

func DrawBackground(renderer *sdl.Renderer) error {
	renderer.Clear()

	t, err := img.LoadTexture(renderer, "./resources/images/background.png")
	if err != nil {
		return err
	}
	defer t.Destroy()

	if err := renderer.Copy(t, nil, nil); err != nil {
		return err
	}

	renderer.Present()

	return nil
}

func DrawTitle(title string, renderer *sdl.Renderer) error {
	renderer.Clear()

	f, err := ttf.OpenFont("resources/fonts/Go-Regular.ttf", 14)
	if err != nil {
		return err
	}
	defer f.Close()

	surface, err := f.RenderUTF8_Solid(title, sdl.Color{
		R: 255,
		G: 100,
		B: 0,
		A: 255,
	})

	defer surface.Free()

	if err != nil {
		return err
	}

	t, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}
	defer t.Destroy()

	renderer.Copy(t, nil, nil)

	renderer.Present()

	return nil
}
