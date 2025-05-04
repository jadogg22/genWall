package cfg

import (
	"fmt"
	"image/color"
	"os"
)

const WIDTH = 2560
const HIGHT = 1440

var ROSE_PINE_BASE = color.RGBA{25, 23, 36, 255}
var ROSE_PINE_COLORS = []color.RGBA{
	{0xEB, 0x6F, 0x92, 0xFF},
	{0xF6, 0xC1, 0x77, 0xFF},
	{0xEB, 0xBC, 0xBA, 0xFF},
	{0x31, 0x74, 0x8F, 0xFF},
	{0x9C, 0xCF, 0xD8, 0xFF},
	{0xC4, 0xA7, 0xE7, 0xFF},
}

type Config struct {
	OutputPath string
	Algorthmn  string
	Width      int
	Height     int
	BaseColor  color.RGBA
	Pallete    []color.RGBA
}

func (cfg *Config) setSize(width int, height int) {
	cfg.Width = width
	cfg.Height = height
}

func LoadConfig() *Config {
	// Load the config file from ~/.config/genwall/config.toml
	// Parse the config file and return a Config struct
	// For now, we'll just return a default config
	return baseConfig()
}

func baseConfig() *Config {
	return &Config{
		OutputPath: "~/.cache/genwall.png",
		Algorthmn:  "Damascus",
		Width:      WIDTH,
		Height:     HIGHT,
		BaseColor:  ROSE_PINE_BASE,
		Pallete:    ROSE_PINE_COLORS,
	}

}

// Grab the config file from ~/.config/genwall/config.toml
func grabConfig() {
	path := "~/.config/genwall/config.toml"
	// Check if the config file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Config file does not exist at:", path)
		return
	}
	// Read the config file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	// Print the config file
	var config string
	_, err = fmt.Fscan(file, &config)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	fmt.Println("Config file contents:")

}

func genConfig() {
	// ~/.config/genwall/config.toml
	config := `
[output]
path = "~/.cache/genwall.png"
width = 1920
height = 1080

[style]
type = "noise"            # Options: "noise", "voronoi", "circle-pack", "shader", "plasma"
seed = "time"             # Options: "time", "hostname", "random"
palette = "solarized"     # Options: "solarized", "dracula", "monochrome", etc.

[noise]
scale = 0.01
octaves = 4
contrast = 1.2

[shader]
file = "~/wallfx/shaders/waves.frag"`

	// Check if the config file exists
	path := "~/.config/genwall/config.toml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the config file
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		defer file.Close()

		// Write the config to the file
		_, err = file.WriteString(config)
		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}
		fmt.Println("Config file created at:", path)
	} else {
		fmt.Println("Config file already exists at:", path)
		fmt.Println("Please edit it to your liking.")
		fmt.Println("You can use the following command to open it:")
		fmt.Println("nano", path)
		fmt.Println("or")
		fmt.Println("vim", path)
	}
}
