package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Setup func(*cobra.Command)

// New
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
