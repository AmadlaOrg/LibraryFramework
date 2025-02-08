package cli

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
