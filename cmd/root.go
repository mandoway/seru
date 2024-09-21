package cmd

import (
	"github.com/mandoway/seru/reduction"
	"github.com/mandoway/seru/version"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type Flags struct {
	InputFile, TestScript, GivenLanguage string
	UseStrategyIsolation                 bool
}

var flags Flags

var rootCmd = &cobra.Command{
	Use:   "seru",
	Short: "A tool to reduce a program while maintaining a property",
	// TODO
	Long: `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := reduction.StartReductionProcess(flags.InputFile, flags.TestScript, flags.GivenLanguage, flags.UseStrategyIsolation)
		if err != nil {
			log.Fatal(err)
		}
	},
	Version: version.Version,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.InputFile, "input", "i", "", "-i <path to file>")
	rootCmd.PersistentFlags().StringVarP(&flags.TestScript, "test", "t", "", "-i <path to testscript>")
	rootCmd.PersistentFlags().StringVarP(&flags.GivenLanguage, "lang", "l", "", "-l <language of file>")
	rootCmd.PersistentFlags().BoolVarP(&flags.UseStrategyIsolation, "strategy-isolation", "s", false, "")

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
