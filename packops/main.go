package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "packops",
		Short:        "DevOps Starter Pack",
		SilenceUsage: true,
	}

	cmd.AddCommand(create())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
