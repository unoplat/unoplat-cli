package command_executor

import (
	"sync"
	"testing"
)

func TestCmdExec(t *testing.T) {
	// Define commands to run in child processes
	commands := []string{"echo Hello, World!", "ls -e", "date"}
	errChan := make(chan error, len(commands))
	stdErrChan := make(chan string, len(commands))
	// Create a WaitGroup to wait for all child processes to finish
	var wg sync.WaitGroup

	// Iterate over the commands and run them in separate goroutines
	for _, cmdStr := range commands {
		wg.Add(1)
		go RunCommandWithColor(cmdStr, &wg, errChan, stdErrChan)
	}

	// Wait for all child processes to finish

	wg.Wait()

	close(errChan)
	close(stdErrChan)

	for err := range stdErrChan {
		// Handle the error from the child process
		if err != "" {
			PrintError("Error: \n" + err)
		}
	}
	for err := range errChan {
		// Handle the error from the child process
		PrintError("Error:\n" + err.Error())
	}
}
