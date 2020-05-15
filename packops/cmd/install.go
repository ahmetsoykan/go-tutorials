package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs packages",
}

var installBrew = &cobra.Command{
	Use:   "brew",
	Short: "Installs homebrew",
	Run: func(cmd *cobra.Command, args []string) {
		catCmd := exec.Command("curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh")
		wcCmd := exec.Command("sh")

		//make a pipe
		reader, writer := io.Pipe()
		var buf bytes.Buffer

		//set the output of "cat" command to pipe writer
		catCmd.Stdout = writer
		//set the input of the "wc" command pipe reader

		wcCmd.Stdin = reader

		//cache the output of "wc" to memory
		wcCmd.Stdout = &buf

		//start to execute "cat" command
		catCmd.Start()

		//start to execute "wc" command
		wcCmd.Start()

		//waiting for "cat" command complete and close the writer
		catCmd.Wait()
		writer.Close()

		//waiting for the "wc" command complete and close the reader
		wcCmd.Wait()
		reader.Close()
		//copy the buf to the standard output
		io.Copy(os.Stdout, &buf)

		fmt.Println("done")
	},
}
