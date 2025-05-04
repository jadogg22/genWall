package main

import (
	//"fmt"
	//"math/rand"
	//"time"

	"genWall/cfg"
	"genWall/engine"
	//"github.com/fogleman/gg"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func main() {
	cfg := cfg.LoadConfig()
	engine.CreateEngine(cfg)

	// This is what we're looking to do here.
	// cfg := cfg.LoadConfig()
	// gen := SelectGenerator(cfg)

	// fmt.Println("Generating wallpaper with %s..." + gen.Name())
	// img := gen.Generate(cfg)
	// img.Save(cfg.OutputPath)
	// fmt.Println("Wallpaper saved to %s" + cfg.OutputPath)

	c := generativeart.NewCanva(cfg.Width, cfg.Height)
	c.SetBackground(cfg.BaseColor)
	c.FillBackground()
	c.SetColorSchema(cfg.Pallete)
	c.Draw(arts.NewContourLine(500))
	c.ToPNG("contourline.png")
}
