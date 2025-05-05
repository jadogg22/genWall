package cfg

import (
	"fmt"
	"image/color"
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




func baseConfig() *FullConfig {
	// Set the default config
	GeneralConfig := GeneralConfig{
		OutputPath: "~/.cache/genwall.png",
		Width:      WIDTH,
		Height:     HIGHT,
		BaseColor:  RGBAtoHex(ROSE_PINE_BASE),
		Pallete:    RGBAStoHex(ROSE_PINE_COLORS),
	}

	damascusConfig := DamascusConfig{
		Enabled:    true,
		LineNum:    500,
		DotRadius:  2.0,
		LineLength: 1500,
		NoiseScale: 800,
		DotStep:    0.4,
	}

	return &FullConfig{
		General:   GeneralConfig,
		Damascus:  damascusConfig,
		}
}
	


func GrabConfig() *FullConfig {
	// Check if the config file exists in the current directory
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config file:", err)
		cfg = baseConfig()
	}
	return cfg	
}


func RGBAtoHex(c color.RGBA) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func RGBAStoHex(c []color.RGBA) []string {
	hex := make([]string, len(c))
	for i, color := range c {
		hex[i] = RGBAtoHex(color)
	}
	return hex
}
