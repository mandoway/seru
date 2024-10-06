package cmd

import (
	"github.com/mandoway/seru/reduction"
	"github.com/mandoway/seru/version"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var flags Flags

var rootCmd = &cobra.Command{
	Use:   "seru",
	Short: "A tool to reduce a program while maintaining a property",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := reduction.StartReductionProcess(flags.InputFile, flags.TestScript, flags.GivenLanguage, flags.UseStrategyIsolation, flags.EnableMetrics, flags.GetReducer())
		if err != nil {
			log.Fatal(err)
		}
	},
	Version: version.Version,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.InputFile, "input", "i", "", "-i <path to file>")
	_ = rootCmd.MarkPersistentFlagRequired("input")
	_ = rootCmd.MarkPersistentFlagFilename("input")
	rootCmd.PersistentFlags().StringVarP(&flags.TestScript, "test", "t", "", "-i <path to testscript>")
	_ = rootCmd.MarkPersistentFlagRequired("test")
	_ = rootCmd.MarkPersistentFlagFilename("test")

	rootCmd.PersistentFlags().StringVarP(&flags.GivenLanguage, "lang", "l", "", "-l <language of file>")
	rootCmd.PersistentFlags().StringVarP(&flags.SyntacticReducer, "reducer", "r", PersesReducer, "-r (perses|vulcan)")

	rootCmd.PersistentFlags().BoolVarP(&flags.UseStrategyIsolation, "strategy-isolation", "s", false, "only one strategy will be applied before next iteration, if enabled (otherwise combines all)")
	rootCmd.PersistentFlags().BoolVarP(&flags.EnableMetrics, "enable-metrics", "m", false, "stores metrics as a json file")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
