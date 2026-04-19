package cli

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("CLI", func() {
	var (
		stdout  *bytes.Buffer
		oldExit func(int)
	)

	BeforeEach(func() {
		// Capture stdout
		stdout = new(bytes.Buffer)
		oldExit = osExit // Save original osExit

		// Mock osExit to prevent termination
		osExit = func(code int) {
			GinkgoWriter.Println("Intercepted os.Exit with code:", code)
		}
	})

	AfterEach(func() {
		// Restore original osExit
		osExit = oldExit
	})

	Describe("Version command", func() {
		It("should print the correct version", func() {
			// Simulate CLI with arguments
			os.Args = []string{"testcli", "version"}

			New("testcli", "Test CLI", "1.0.0", func(cmd *cobra.Command) {
				cmd.SetOut(stdout) // Redirect output
			})

			Expect(stdout.String()).To(ContainSubstring("testcli version 1.0.0"))
		})
	})

	Describe("Unknown command", func() {
		It("should print an error message and not exit", func() {
			// Simulate CLI with invalid command
			os.Args = []string{"testcli", "unknown"}

			New("testcli", "Test CLI", "1.0.0", func(cmd *cobra.Command) {
				cmd.SetErr(stdout) // Capture error output
			})

			Expect(stdout.String()).To(ContainSubstring("unknown command"))
		})
	})

	Describe("Custom setup function", func() {
		It("should add a custom command", func() {
			// Simulate CLI with a custom command
			os.Args = []string{"testcli", "custom"}

			New("testcli", "Test CLI", "1.0.0", func(cmd *cobra.Command) {
				customCmd := &cobra.Command{
					Use:   "custom",
					Short: "A custom command",
					Run: func(cmd *cobra.Command, args []string) {
						cmd.Println("Custom command executed")
					},
				}
				cmd.AddCommand(customCmd)
				cmd.SetOut(stdout) // Capture output
			})

			Expect(stdout.String()).To(ContainSubstring("Custom command executed"))
		})
	})
})
