package dag

import (
	"errors"
	"fmt"
	"log"

	commandconfig "github.com/unoplat/unoplat-cli/code/unoplat-cli/command_config"
)

type CommandStatus string

var CMD_PENDING CommandStatus = "Pending"
var CMD_COMPLETED CommandStatus = "Completed"
var CMD_ERRORED CommandStatus = "Errored"

type DAG struct {
	config        *commandconfig.Config
	IdToIndex     map[string]int
	IdToCmdStatus map[string]CommandStatus
}

func NewDag(config *commandconfig.Config) (*DAG, error) {
	IdToIndex := map[string]int{}
	IdToCmdStatus := map[string]CommandStatus{}
	for i, cmd := range config.MappedCommands {
		for _, d := range cmd.Dependency {
			_, ok := IdToIndex[d]
			if !ok {
				log.Fatal(fmt.Sprintf("Command with ID %s has invalid dependency %s", cmd.ID, d))
				return nil, errors.New("Ensure to have dependent command after the independent command in the commands array")
			}
		}
		IdToIndex[cmd.ID] = i
		IdToCmdStatus[cmd.ID] = CMD_PENDING
	}

	return &DAG{
		config:        config,
		IdToIndex:     IdToIndex,
		IdToCmdStatus: IdToCmdStatus,
	}, nil
}

func (d *DAG) GetCmdsToRun() []commandconfig.MappedCommand {
	cmdsToRun := []commandconfig.MappedCommand{}
	for _, cmd := range d.config.MappedCommands {
		isSchedulable := true
		if d.IdToCmdStatus[cmd.ID] != CMD_PENDING {
			continue
		}
		for _, dep := range cmd.Dependency {
			status := d.IdToCmdStatus[dep]
			if status != CMD_COMPLETED {
				isSchedulable = false
				break
			}
		}
		if isSchedulable {
			cmdsToRun = append(cmdsToRun, cmd)
		}
	}
	return cmdsToRun
}
