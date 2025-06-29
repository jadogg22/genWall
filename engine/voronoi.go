package engine

import (
	"genWall/cfg"
	"github.com/fogleman/gg"
)

type Voronoi struct {
	gen  cfg.GeneralConfig
	opts cfg.VoronoiConfig
}

func NewVoronoi(gen cfg.GeneralConfig, opts cfg.VoronoiConfig) *Voronoi {
	return &Voronoi{
		gen:  gen,
		opts: opts,
	}
}

func (v *Voronoi) Name() string {
	return "Voronoi"
}

func (v *Voronoi) Draw(c *Canvas) error {
	ctx := gg.NewContextForRGBA(c.img)
	// Set the background color
	backgound, _ := HexToRGBA(v.gen.BaseColor)
	ctx.SetColor(backgound)
	ctx.DrawRectangle(0, 0, float64(v.gen.Width), float64(v.gen.Height))
	ctx.Fill()

	return nil
}