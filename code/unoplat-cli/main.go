package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra" // You can use a library like Cobra for CLI handling.
	"github.com/unoplat/unoplat-cli/code/unoplat-cli/cmd"
)

func main() {
	// Create a root command
	rootCmd := &cobra.Command{
		Use:   "unoplat-cli",
		Short: "unoplat-cli is a (CLI) application that allows you to install unoplat's components, manage and uninstall them in a simple way.",
	}

	rootCmd.AddCommand(cmd.GetInstallCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
