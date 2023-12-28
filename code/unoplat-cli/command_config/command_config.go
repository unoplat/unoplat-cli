package command_config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Define the struct to represent the JSON data
type Config struct {
	CmdName        string          `json:"cmdName"`
	MappedCommands []MappedCommand `json:"mapped_commands"`
}

type MappedCommand struct {
	ID         string   `json:"id"`
	Command    string   `json:"command"`
	Dependency []string `json:"dependency"`
}

func ParseFromFile(jsonFilePath string) (*Config, error) {
	// Read the JSON file
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
		return nil, err
	}

	// Parse the JSON data into the Config struct
	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
		return nil, err
	}
	return &config, nil
}

func GetCommandConfig(cmdName string) *Config {
	folderPath := "command_declaration"

	pwd, _ := os.Getwd()

	fileList, err := os.ReadDir(path.Join(pwd, folderPath))
	if err != nil {
		log.Fatalf("Error listing files in folder: %v", err)
	}

	for _, file := range fileList {
		if filepath.Ext(file.Name()) == ".json" {
			config, err := ParseFromFile(path.Join(folderPath, file.Name()))
			if err != nil {
				log.Fatalf("Error Getting config: %v", err)
			}
			if config.CmdName == cmdName {
				return config
			}
		}
	}
	return nil
}
