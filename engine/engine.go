package engine

import (
	"fmt"
	"genWall/cfg"
	"image"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

type Canvas struct {
	w, h int
	img  *image.RGBA
	conf *cfg.Config
}

func CreateEngine(cfg *cfg.Config) {
	fmt.Println("Creating Image")
	genImage(cfg)
}

func genImage(cfg *cfg.Config) {
	c := generativeart.NewCanva(cfg.Width, cfg.Height)
	c.SetBackground(cfg.BaseColor)
	c.FillBackground()
	c.SetColorSchema(cfg.Pallete)
	c.Draw(arts.NewContourLine(500))
	c.ToPNG("contourline.png")
}
