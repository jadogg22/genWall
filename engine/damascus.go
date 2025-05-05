package engine

import (
	"genWall/cfg"
	"math"
	"math/rand"
	"github.com/fogleman/gg"

	"github.com/ojrac/opensimplex-go"

)

type Damascus struct {
	gen cfg.GeneralConfig
	opts   cfg.DamascusConfig
	noise opensimplex.Noise
}

func NewDamascus(gen cfg.GeneralConfig, opts cfg.DamascusConfig) *Damascus {
	return &Damascus{
		gen:   gen,
		opts:  opts,
		noise: opensimplex.NewNormalized(rand.Int63()),
	}
}

func (d Damascus) Name() string {
	return "Damascus"
}

func (d Damascus) Draw(c *Canvas) error {
	ctx := gg.NewContextForRGBA(c.img)
	// Set the background color
	backgound, _ := HexToRGBA(d.gen.BaseColor)
	ctx.SetColor(backgound)
	ctx.DrawRectangle(0, 0, float64(d.gen.Width), float64(d.gen.Height))
	ctx.Fill()

	for i := 0; i < d.opts.LineNum; i++ {

		cls := RandomColor(d.gen.Pallete)
		x := rand.Float64() * float64(d.gen.Width)
		y := rand.Float64() * float64(d.gen.Height)

		// set background color to have a gradient effect
		if i % 5 == 0 {
			ctx.SetRGBA255(int(backgound.R), int(backgound.G), int(backgound.B), 3)
			ctx.DrawRectangle(0, 0, float64(d.gen.Width), float64(d.gen.Height))
			ctx.Fill()
		}

		for j := 0; j < d.opts.LineLength; j++ {

			theta := d.noise.Eval2(x/d.opts.NoiseScale, y/d.opts.NoiseScale) * math.Pi * 2 	* d.opts.NoiseScale
			x += math.Cos(theta) * d.opts.DotStep
			y += math.Sin(theta) * d.opts.DotStep

			ctx.SetColor(cls)
			ctx.DrawEllipse(x, y, d.opts.DotRadius, d.opts.DotRadius)
			ctx.Fill()

			// if the point is out of bounds, reset it to a random position
			if x > float64(d.gen.Width) || x < 0 || y > float64(d.gen.Height) || y < 0 || rand.Float64() < 0.001 {
				x = rand.Float64() * float64(d.gen.Width)
				y = rand.Float64() * float64(d.gen.Height)
			}
		}
	}

	return nil
}




