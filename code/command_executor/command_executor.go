package command_executor

import (
	"bufio"
	"io"
	"os/exec"
	"sync"

	"github.com/fatih/color"
)

var mu sync.Mutex

func GetCommand(cmdStr string) *exec.Cmd {
	mu.Lock()
	defer mu.Unlock()
	return exec.Command("sh", "-c", cmdStr)
}

func RunCommandWithColor(cmdStr string, wg *sync.WaitGroup, errChan chan error, stdErrChan chan string) {

	// Decrement the WaitGroup counter when the function exits
	defer wg.Done()

	// Define an array of colors for legibility
	colors := []*color.Color{
		color.New(color.FgGreen),     // Green
		color.New(color.FgBlue),      // Blue
		color.New(color.FgMagenta),   // Magenta
		color.New(color.FgYellow),    // Yellow
		color.New(color.FgCyan),      // Cyan
		color.New(color.FgWhite),     // White
		color.New(color.FgHiBlack),   // Black
		color.New(color.FgHiGreen),   // Bright Green
		color.New(color.FgHiYellow),  // Bright Yellow
		color.New(color.FgHiBlue),    // Bright Blue
		color.New(color.FgHiMagenta), // Bright Magenta
		color.New(color.FgHiCyan),    // Bright Cyan
		color.New(color.FgHiRed),     // Bright Red
	}

	// Pick a color based on the command string
	colorIndex := hashString(cmdStr) % uint32(len(colors))
	cmdColor := colors[colorIndex]
	erroredColor := color.New(color.FgRed)
	// Print the command in the selected color
	cmdColor.Printf("Running command: %s\n", cmdStr)

	// Create a command object
	cmd := GetCommand(cmdStr)
	// Capture the output and error streams
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// Start the command
	if err := cmd.Start(); err != nil {
		erroredColor.Printf("Error starting command: %v\n", err)
		errChan <- err
		return
	}

	// Create a goroutine to print the output and error in the same color
	go PrintColoredCmdOutput(stdout, cmdColor)

	// Create a goroutine to print the error in the same color
	go GetStdErr(stderr, cmdColor, stdErrChan)
	// Wait for the command to finish

	if err := cmd.Wait(); err != nil {
		erroredColor.Printf("Error running command: %v\n", cmdStr)
		errChan <- err
		return
	}
}

// Simple hash function to generate an index from a string
func hashString(s string) uint32 {
	var hash uint32
	for _, c := range s {
		hash = uint32(c) + ((hash << 5) - hash)
	}
	return hash
}

func PrintColoredCmdOutput(p io.ReadCloser, c *color.Color) {
	reader := bufio.NewReader(p)
	line, err := reader.ReadString('\n')
	for err == nil {

		c.Println(line)
		line, err = reader.ReadString('\n')
	}
}

func GetStdErr(p io.ReadCloser, c *color.Color, stdErrChan chan string) {
	reader := bufio.NewReader(p)
	line, err := reader.ReadString('\n')
	stdErr := ""
	for err == nil {
		stdErr += (line + "\n")
		line, err = reader.ReadString('\n')
	}
	if stdErr != "" {
		stdErrChan <- stdErr
	}
}

func PrintError(err string) {
	erroredColor := color.New(color.FgRed)
	erroredColor.Printf(err)
}
