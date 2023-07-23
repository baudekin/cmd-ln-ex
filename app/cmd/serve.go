/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	top "github.com/baudekin/cmd-ln-ex"
	suba "github.com/baudekin/cmd-ln-ex/subpacka"
	subb "github.com/baudekin/cmd-ln-ex/subpackb"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Service",
	Long: `"start" service aka run hello commands.
and usage of using your command. For example:

go run main.go service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		top.HelloTop()
		suba.Hello()
		subb.Hello()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serveCmd.PersistentFlags().String("port", "-p", "Service Port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
