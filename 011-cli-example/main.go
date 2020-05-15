package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)
func main(){
	var echoTimes int

	var cmdPrint = &cobra.Command{
		Use: "Print [string to print]",
		Short: "Print anything to the screen",
		Args: cobra.MinimumNArgs(1),
		Run: func (cmd *cobra.Command, args []string)  {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use: "echo [string to echo]",
		Short: "Echo anything to the screen",
		Args: cobra.MaximumNArgs(1),
		Run: func (cmd *cobra.Command, args []string)  {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use: "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Args: cobra.MinimumNArgs(1),
		Run: func (cmd *cobra.Command, args []string)  {
			for i := 0; i < echoTimes; i++{
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}