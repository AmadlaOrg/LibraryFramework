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

/*import (
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLI", func() {
	Context("help command", func() {
		It("should print the help message", func() {
			// Get the current working directory
			cwd, err := os.Getwd()
			Expect(err).ToNot(HaveOccurred())
			// Construct the absolute path to cli.go
			cmdPath := filepath.Join(cwd, "./cli.go")

			// Run the command
			cmd := exec.Command("go", "run", cmdPath, "help")
			output, err := cmd.CombinedOutput()
			if err != nil {
				Fail(string(output)) // Provide more context on failure
			}
			Expect(err).ToNot(HaveOccurred(), string(output))

			expectedOutput := `HERY CLI application

Usage:
  hery [command]

Available Commands:
  client      HERY client
  collection  Collections
  completion  Generate the autocompletion script for the specified shell
  compose     Compose the specified entity
  entity      Entity commands
  help        Help about any command
  query       Query entities
  server      HERY Server
  settings    List the paths and other environment variables for HERY
  version     Print the version number of hery

Flags:
  -h, --help      help for hery
  -v, --version   version for hery

Use "hery [command] --help" for more information about a command.`

			Expect(string(output)).To(ContainSubstring(expectedOutput))
		})
	})
})

// TODO: Need to find a way to test
/*import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func captureOutput(f func()) string {
	// Save original stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute function
	f()

	// Restore stdout
	err := w.Close()
	if err != nil {
		panic(err)
		return ""
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestVersionCommand(t *testing.T) {
	output := captureOutput(func() {
		New("testcli", "Test CLI", "1.0.0", func(cmd *cobra.Command) {})
		// Simulate running `testcli version`
		//os.Args = []string{"testcli", "version"}
	})

	//expectedOutput := "testcli version 1.0.0\n"
	//assert.Contains(t, output, expectedOutput)

	println(output)
}*/
