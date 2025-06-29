package engine

import (
	"genWall/cfg"
	"github.com/fogleman/gg"
)

type Spray struct {
	gen  cfg.GeneralConfig
	opts cfg.SprayConfig
}

func NewSpray(gen cfg.GeneralConfig, opts cfg.SprayConfig) *Spray {
	return &Spray{
		gen:  gen,
		opts: opts,
	}
}

func (s *Spray) Name() string {
	return "Spray"
}

func (s *Spray) Draw(c *Canvas) error {
	ctx := gg.NewContextForRGBA(c.img)
	// Set the background color
	backgound, _ := HexToRGBA(s.gen.BaseColor)
	ctx.SetColor(backgound)
	ctx.DrawRectangle(0, 0, float64(s.gen.Width), float64(s.gen.Height))
	ctx.Fill()

	return nil
}
