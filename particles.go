package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type spot struct {
	X     float64
	Y     float64
	Life  float64
	Color color.RGBA
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, 512, 512),
		Undecorated: true,
	})
	if err != nil {
		panic(err)
	}

	canvas := pixelgl.NewCanvas(win.Bounds())

	spots := []spot{}

	for !win.Closed() {
		if win.Pressed(pixelgl.MouseButtonLeft) {
			clickX, clickY := win.MousePosition().XY()
			spots = append(spots, spot{X: clickX, Y: clickY, Life: 400, Color: colornames.Pink})
		}

		// draw on screen
		atlas := text.NewAtlas(
			basicfont.Face7x13,
			text.ASCII,
		)

		win.Clear(colornames.Black)

		for i, s := range spots {
			if s.Life > 0 {
				txt := text.New(pixel.V(s.X, s.Y), atlas)
				txt.Color = s.Color
				fmt.Fprintln(txt, ".")
				txt.Draw(win, pixel.IM)

				s.Color.R -= 5
				s.Color.G -= 5
				s.Color.B -= 5
				s.Life--

				spots[i] = s
			}
		}
		canvas.Draw(win, pixel.IM)
		win.Update()
	}

}

func main() {
	pixelgl.Run(run)
}
