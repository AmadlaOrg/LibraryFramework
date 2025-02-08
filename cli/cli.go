package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// For mocking
var (
	osExit = os.Exit
)

// Setup function type definition is for attaching commands to the root of the cli
// This is for the implementation of a decoration pattern
//
// Params:
// - 🐍 *cobra.Command - Is for passing the root Command so that other commands can be attached to the main one
type Setup func(*cobra.Command)

// New attaches to the cli the basic details and command that all the cli application require to work
//
// Params:
// - 📇 name - Is the of the application (normally all lowercase)
// - 📜 title - Is the name but with uppercase letters where need be
// - ♻️ version - The version of the application
// - 🚀 setup - Is a callback/decoration-pattern so that other commands can be attached
//
// Example:
//
//	func main() {
//		cli.New(
//			"hery",
//			"HERY",
//			"1.0.0",
//			func(rootCmd *cobra.Command) {
//				rootCmd.AddCommand(cmd.SettingsCmd)
//				rootCmd.AddCommand(cmd.CollectionCmd)
//			})
//	}
func New(name, title, version string, setup Setup) {
	var (
		// rootCmd the main command setup
		rootCmd = &cobra.Command{
			Use:     name,
			Short:   fmt.Sprintf("%s CLI application", title),
			Version: version,
		}

		// versionCmd is for the command setup for displaying the application version
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print the version number of " + name,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println(name + " version " + version)
			},
		}
	)

	rootCmd.AddCommand(versionCmd)

	setup(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		osExit(1)
	}
}
