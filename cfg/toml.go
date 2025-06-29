package cfg

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"image/color"
	"os"
	"strconv"
)

type FullConfig struct {
	General  GeneralConfig
	Damascus DamascusConfig
	Voronoi  VoronoiConfig
	Spray    SprayConfig
}

type GeneralConfig struct {
	OutputPath string   `toml:"output_path"`
	Width      int      `toml:"width"`
	Height     int      `toml:"height"`
	BaseColor  string   `toml:"background"`
	Pallete    []string `toml:"pallete"`
}

type DamascusConfig struct {
	Enabled    bool
	LineNum    int     `toml:"line_num"`
	DotRadius  float64 `toml:"dot_radius"`
	LineLength int     `toml:"line_length"`
	NoiseScale float64 `toml:"noise_scale"`
	DotStep    float64 `toml:"dot_step"`
}

type VoronoiConfig struct {
	Enabled     bool
	NumPoints   int     `toml:"num_points"`
	StrokeWidth float64 `toml:"stroke_width"`
	StrokeColor string  `toml:"stroke_color"`
}

type SprayConfig struct {
	Enabled   bool
	NumPoints int `toml:"num_points"`
}

// find the config file in ~/.config/genwall/config.toml
// or in the current directory
func LoadConfig() (*FullConfig, error) {
	var cfg FullConfig
	// Check if the config file exists in the current directory
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		// Check if the config file exists in the home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		configPath := fmt.Sprintf("%s/.config/genwall/config.toml", homeDir)
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("config file not found in current directory or ~/.config/genwall")
		}
		// Use the config file in the home directory
		_, err = toml.DecodeFile(configPath, &cfg)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
	} else {
		// Use the config file in the current directory
		_, err = toml.DecodeFile("config.toml", &cfg)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
	}
}

func HexToRGBA(hex string) (color.RGBA, error) {
	if len(hex) == 7 && hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return color.RGBA{}, fmt.Errorf("invalid color format: %s", hex)
	}
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}, nil
}
