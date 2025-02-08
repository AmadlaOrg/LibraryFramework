package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Setup func(*cobra.Command)

// New attaches to the cli the basic details and command that all the cli application require to work
//
// Params:
// - 📇 name - Is the of the application (normally all lowercase)
// - 📜 title - Is the name but with uppercase letters where need be
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

	appName := fmt.Sprintf("auditor-%s", name)
	appTitleName := fmt.Sprintf("Auditor %s", title)

	var (
		rootCmd = &cobra.Command{
			Use:     appName,
			Short:   appTitleName + " CLI application",
			Version: version,
		}
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print the version number of " + appName,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println(appName + " version " + version)
			},
		}
	)

	rootCmd.AddCommand(versionCmd)

	// TODO: Test with and without `=`
	setup(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
