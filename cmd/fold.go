/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// foldCmd represents the fold command
var foldCmd = &cobra.Command{
	Use:   "fold [direcotry]",
	Short: "Use this command to compress a directory",
	Long: `Use this command to compress a directory into a .ori file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fold called")
	},
}

func init() {
	rootCmd.AddCommand(foldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// foldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// foldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
