package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unoplat/unoplat-cli/code/unoplat-cli/command_manager"
)

func GetInstallCmd() *cobra.Command {
	var verbose bool
	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "installs unoplat components",

		Run: func(cmd *cobra.Command, args []string) {
			Install()
		},
	}
	installCmd.Flags().BoolVarP(&verbose, "verbose", "v", true, "Verbose output")
	return installCmd
}

func Install() {
	command_manager.RunCommand("install")
}
