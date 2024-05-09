/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/LorenzHohermuth/origami/internal/registry"
	"github.com/LorenzHohermuth/origami/internal/splitter"
	"github.com/spf13/cobra"
)

// registryCmd represents the registry command
var registryCmd = &cobra.Command{
	Use:   "registry [file|direcotry]",
	Short: "Disributes tokens into HashMap",
	Long: `Distributes tokens from a file or a direcotry
	into a HashMap and prints the Map`,
	Run: func(cmd *cobra.Command, args []string) {
		tokens := splitter.SplitFile(args[0])
		tr := registry.TokenRegistry{ Map: make(map[string]int64)}
		tr.DistributeTokens(tokens)

		tw := new(tabwriter.Writer)
		tw.Init(os.Stdout, 0, 8, 0, '\t', 0)
		for v, k := range tr.Map {
			fmt.Fprintf(tw, "value: \"%s\" \t key: %b \n", v, k)
		}
		tw.Flush()
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
