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
	// create strategy based on the config
	strategy := engine.NewDamascus(cfg.General, cfg.Damascus)
	canvas.SetStrategy(strategy)


	fmt.Println("Generating wallpaper with " + canvas.Name())
	// Generate the wallpaper
	err := canvas.Draw()
	if err != nil {
		fmt.Println("Error generating wallpaper: ", err)
		return
	}
	canvas.ToPNG("output.png")
}

