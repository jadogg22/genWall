package main

import (
	"fmt"
	//"math/rand"
	//"time"


	"genWall/cfg"
	"genWall/engine"
	//"github.com/fogleman/gg"

)

func main() {
	// Load the config file
	cfg := cfg.GrabConfig()
	// select the generator based on the config
	canvas := engine.NewCanvas(cfg)
	strategy := engine.NewDamascus(cfg.General, cfg.Damascus)
	canvas.SetStrategy(strategy)

	fmt.Println("Generating wallpaper with " + canvas.Name())
	err := canvas.Draw()
	if err != nil {
		fmt.Println("Error generating wallpaper: ", err)
		return
	}
	canvas.ToPNG("output.png")


	//canvas.ToPNG(cfg.General.OutputPath)
	// Generate the wallpaper using strategy
	

	// This is what we're looking to do here.
	// cfg := cfg.LoadConfig()
	// gen := SelectGenerator(cfg)

	// fmt.Println("Generating wallpaper with %s..." + gen.Name())
	// img := gen.Generate(cfg)
	// img.Save(cfg.OutputPath)
	// fmt.Println("Wallpaper saved to %s" + cfg.OutputPath)
}

