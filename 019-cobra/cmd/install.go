package cmd

import (
	"fmt"
	"log"
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
		command := exec.Command("sh", "-c", `curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh | sh`)
		out, err := command.Output()
		if err != nil {
			log.Fatalf("command.Run() failed with %s\n", err)
		} else {
			fmt.Printf("Combined out:\n%s\n", string(out))
		}
	},
}
