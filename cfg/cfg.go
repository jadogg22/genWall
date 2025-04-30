package cfg

import (
	"fmt"
	"os"
)

type Config struct {
	OutputPath string
}
	

func loadConfig() *Config {
	// Load the config file from ~/.config/genwall/config.toml
	// Parse the config file and return a Config struct
	// For now, we'll just return a default config
	return &Config{
		OutputPath: "~/.cache/genwall.png",
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
