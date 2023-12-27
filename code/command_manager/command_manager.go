package command_manager

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/unoplat/unoplat-cli/code/command_config"
	"github.com/unoplat/unoplat-cli/code/command_executor"
	"github.com/unoplat/unoplat-cli/code/dag"
)

func RunCommand(cmdName string) {
	cmdConfig := command_config.GetCommandConfig(cmdName)
	cmdDag, err := dag.NewDag(cmdConfig)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while creating DAG %v", err))
	}
	continueExecution := true
	for continueExecution {
		cmds := cmdDag.GetCmdsToRun()
		if len(cmds) == 0 {
			continueExecution = false
			break
		}
		errChan := make(chan error, len(cmds))
		stdErrChan := make(chan string, len(cmds))
		var wg sync.WaitGroup

		for _, cmd := range cmds {
			wg.Add(1)
			go command_executor.RunCommandWithColor(cmd.Command, &wg, errChan, stdErrChan)
			time.Sleep(10 * time.Millisecond)
		}

		wg.Wait()

		close(errChan)
		close(stdErrChan)

		continueExecution = continueExecution && !hasErrors(errChan, stdErrChan)
		for _, cmd := range cmds {
			cmdDag.IdToCmdStatus[cmd.ID] = dag.CMD_COMPLETED
		}
	}

}

func hasErrors(errChan chan error, stdErrChan chan string) bool {
	hasErrors := false
	for err := range stdErrChan {
		// Handle the error from the child process
		if err != "" {
			hasErrors = true
			command_executor.PrintError("Error: \n" + err)
		}
	}
	for err := range errChan {
		// Handle the error from the child process
		hasErrors = true
		command_executor.PrintError("Error:\n" + err.Error())
	}
	return hasErrors
}
