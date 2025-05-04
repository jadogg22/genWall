package engine

import (
	"fmt"
	"genWall/cfg"
)

func CreateEngine(cfg *cfg.Config) {
	fmt.Println("Swag")
	fmt.Printf("The size is %d x %d \n", cfg.Width, cfg.Height)
}
