package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	osExit = os.Exit
)

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
		rootCmd = &cobra.Command{
			Use:     name,
			Short:   fmt.Sprintf("%s CLI application", title),
			Version: version,
		}
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
