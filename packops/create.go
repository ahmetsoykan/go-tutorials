package main

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func create() *cobra.Command {
	return &cobra.Command{
		Use: "print",
		RunE: func(cmd *cobra.Command, args []string) error {
			job := exec.Command("touch", "test.txt")
			if err := job.Run(); err != nil {
				log.Fatal(err)
			} else {
				cmd.Println("--done--")
			}
			return nil
		},
	}
}
