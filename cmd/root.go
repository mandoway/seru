package cmd

import (
	"github.com/mandoway/seru/reduction"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	inputFile, testScript string
	rootCmd               = &cobra.Command{
		Use:   "seru",
		Short: "A tool to reduce a program while maintaining a property",
		// TODO
		Long: `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			err := reduction.StartReductionProcess(inputFile, testScript, "")
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "-i <path to file>")
	rootCmd.PersistentFlags().StringVarP(&testScript, "test", "t", "", "-i <path to testscript>")

	_ = rootCmd.MarkPersistentFlagRequired("input")
	_ = rootCmd.MarkPersistentFlagRequired("test")

	_ = rootCmd.MarkPersistentFlagFilename("input")
	_ = rootCmd.MarkPersistentFlagFilename("test")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
