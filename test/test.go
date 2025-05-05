package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"genWall/cfg"
	"genWall/engine"
)

func TestCanvas(t *testing.T) {
	// Load the config file
	cfg := cfg.GrabConfig()
	// Create a new canvas
	canvas := engine.NewCanvas(cfg)
	assert.NotNil(t, canvas, "Canvas should not be nil")
	assert.NotNil(t, canvas.Img(), "Canvas name should not be nil")

	// Test setting a strategy
	strategy := engine.NewDamascus(cfg.General, cfg.Damascus)
	canvas.SetStrategy(strategy)
	assert.NotNil(t, canvas.Name, "Canvas strategy should not be nil")

	img := canvas.Img()
	assert.NotNil(t, img, "Canvas image should not be nil")
	assert.Equal(t, img.Bounds().Max.X, cfg.General.Width, "Canvas width should match config")
	assert.Equal(t, img.Bounds().Max.Y, cfg.General.Height, "Canvas height should match config")

	canvas.Draw()
	assert.NotNil(t, canvas.Img(), "Canvas image should not be nil after drawing")
	assert.Equal(t, canvas.Name, "Damascus", "Canvas name should match strategy name")
	assert.Equal(t, canvas.Img().Bounds().Max.X, cfg.General.Width, "Canvas width should match config after drawing")
	assert.Equal(t, canvas.Img().Bounds().Max.Y, cfg.General.Height, "Canvas height should match config after drawing")

	canvas.ToPNG("test.png")


}
