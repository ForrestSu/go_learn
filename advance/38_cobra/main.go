package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sink = &Sinker{}
var rootCmd = &cobra.Command{
	Use: "app",
}

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print Version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("version 1.0.0")
		},
	}
	versionCmd.SetOut(sink)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	// 自定义输入，默认是从标准输入读取
	rootCmd.SetArgs([]string{"version"})
	// 自定义输出
	rootCmd.SetOut(&Sinker{})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("data: %s\n", sink)
}
