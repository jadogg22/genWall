package engine

import (
	"genWall/cfg"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"math/rand"
	"sync"

	"github.com/ojrac/opensimplex-go"
)

type Damascus struct {
	gen   cfg.GeneralConfig
	opts  cfg.DamascusConfig
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

type Point struct {
	X, Y float64
	Cls  color.Color
}

func (d Damascus) Draw(c *Canvas) error {
	ctx := gg.NewContextForRGBA(c.img)
	// Set the background color
	background, _ := HexToRGBA(d.gen.BaseColor)
	ctx.SetColor(background)
	ctx.DrawRectangle(0, 0, float64(d.gen.Width), float64(d.gen.Height))
	ctx.Fill()

	points := make(chan Point, d.opts.LineNum*d.opts.LineLength)
	var wg sync.WaitGroup

	for i := 0; i < d.opts.LineNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cls := RandomColor(d.gen.Pallete)
			x := rand.Float64() * float64(d.gen.Width)
			y := rand.Float64() * float64(d.gen.Height)

			for j := 0; j < d.opts.LineLength; j++ {
				theta := d.noise.Eval2(x/d.opts.NoiseScale, y/d.opts.NoiseScale) * math.Pi * 2 * d.opts.NoiseScale
				x += math.Cos(theta) * d.opts.DotStep
				y += math.Sin(theta) * d.opts.DotStep

				points <- Point{X: x, Y: y, Cls: cls}

				// if the point is out of bounds, reset it to a random position
				if x > float64(d.gen.Width) || x < 0 || y > float64(d.gen.Height) || y < 0 || rand.Float64() < 0.001 {
					x = rand.Float64() * float64(d.gen.Width)
					y = rand.Float64() * float64(d.gen.Height)
				}
			}
		}()
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(points)
	}()

	// Draw the points from the channel
	for p := range points {
		ctx.SetColor(p.Cls)
		ctx.DrawEllipse(p.X, p.Y, d.opts.DotRadius, d.opts.DotRadius)
		ctx.Fill()
	}

	return nil
}
