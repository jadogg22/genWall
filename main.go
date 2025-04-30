package main

import (
	//"fmt"
	"image/color"
	//"math/rand"
	//"time"


	//"github.com/fogleman/gg"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"

	
)

func main() {
	// cfg := loadConfig()
	// gen := SelectGenerator(cfg)

	// fmt.Println("Generating wallpaper with %s..." + gen.Name())
	// img := gen.Generate(cfg)
	// img.Save(cfg.OutputPath)
	// fmt.Println("Wallpaper saved to %s" + cfg.OutputPath)

	 colors := []color.RGBA{
	  {0x58, 0x18, 0x45, 0xFF},
	  {0x90, 0x0C, 0x3F, 0xFF},
	  {0xC7, 0x00, 0x39, 0xFF},
	  {0xFF, 0x57, 0x33, 0xFF},
	  {0xFF, 0xC3, 0x0F, 0xFF},
	 }
	 c := generativeart.NewCanva(1600, 1600)
	 c.SetBackground(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	 c.FillBackground()
	 c.SetColorSchema(colors)
	 c.Draw(arts.NewContourLine(500))
	 c.ToPNG("contourline.png")
}

