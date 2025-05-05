package engine

import (
	"os"
	"image/png"
	"fmt"
	"genWall/cfg"
	"image"
	"strings"
	"image/color"
	"math/rand"
)

type Canvas struct {
	img  *image.RGBA
	conf *cfg.FullConfig
	strat ArtStrategy
}

func (c *Canvas) Name() string {
	return c.strat.Name()
}

func (c *Canvas) Img() *image.RGBA {
	return c.img
}

type ArtStrategy interface {
	Draw(c *Canvas) error
	Name() string
}

// CreateCanvas initializes a new Canvas with the given configuration. And selects the strategy based on the config.
func CreateCanvas(cfg *cfg.FullConfig) *Canvas {
	return &Canvas{
		img:    image.NewRGBA(image.Rect(0, 0, cfg.General.Width, cfg.General.Height)),
		conf:   cfg,
		strat:  selectStrategy(cfg),
	}
}

// NewCanvas initializes a new Canvas with the given configuration. It does not select a strategy.
func NewCanvas(cfg *cfg.FullConfig) *Canvas {
	return &Canvas{
		img:    image.NewRGBA(image.Rect(0, 0, cfg.General.Width, cfg.General.Height)),
		conf:   cfg,
		strat: nil,
	}
}

func selectStrategy(cfg *cfg.FullConfig) ArtStrategy {
	// Select the strategy based on the config
	// For now, we'll just return a dummy generator
	if cfg.Damascus.Enabled {
		return NewDamascus(cfg.General, cfg.Damascus)
	}
	return nil
}

func (c *Canvas) SetStrategy(strat ArtStrategy) {
	c.strat = strat
}



func (c *Canvas) Draw() error{
	// Placeholder for drawing logic
	fmt.Println("Drawing on canvas with strategy:", c.strat.Name())
	if c.strat == nil {
		return fmt.Errorf("no drawing strategy set")
	}
	// Call the strategy's Draw method
	err := c.strat.Draw(c)
	if err != nil {
		return fmt.Errorf("error drawing on canvas: %v", err)
	}
	return nil
}

func (c *Canvas) ToPNG(filename string) {
	// Placeholder for saving logic
	fmt.Println("Saving image to", filename)
	file, err := os.Create(filename)
	err = png.Encode(file, c.Img())
	if err != nil {
	 	panic(err)
	}
}

func HexToRGBA(hex string) (color.RGBA, error) {
	// Remove the '#' character if present
	if strings.HasPrefix(hex, "#") {
		hex = hex[1:]
	}

	// Parse the hex string into RGBA values
	var r, g, b, a uint8
	switch len(hex) {
	case 6: // RGB
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
		a = 255 // Fully opaque
	case 8: // RGBA
		fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
	default:
		return color.RGBA{}, fmt.Errorf("invalid hex color format")
	}

	return color.RGBA{r, g, b, a}, nil
}


func RandomColor(palette []string) color.RGBA {
	hex := palette[rand.Intn(len(palette))]
	rgba, err := HexToRGBA(hex)
	if err != nil {
		fmt.Println("Error converting hex to RGBA:", err)
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // fallback to black
	}
	return rgba
}




