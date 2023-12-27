package command_config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestParser(t *testing.T) {
	folderPath := "../command_declaration"

	fileList, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Error listing files in folder: %v", err)
	}

	for _, file := range fileList {
		// Check if the file has a .json extension
		if filepath.Ext(file.Name()) == ".json" {
			jsonData, err := os.ReadFile(path.Join(folderPath, file.Name()))
			if err != nil {
				log.Fatalf("Error reading JSON file: %v", err)
			}

			var config Config
			if err := json.Unmarshal(jsonData, &config); err != nil {
				log.Fatalf("Error parsing JSON data: %v", err)
			}

			// Print the parts of the Config struct
			fmt.Printf("CmdName: %s\n", config.CmdName)

			fmt.Println("Mapped Commands:")
			for _, mappedCommand := range config.MappedCommands {
				fmt.Printf("  ID: %s\n", mappedCommand.ID)
				fmt.Printf("  Command: %s\n", mappedCommand.Command)
				fmt.Println("  Dependencies:")
				for _, dependency := range mappedCommand.Dependency {
					fmt.Printf("    %s\n", dependency)
				}
			}
		}
	}
}
