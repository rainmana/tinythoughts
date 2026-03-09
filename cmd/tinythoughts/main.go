package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waldzellai/tinythoughts/internal/config"
	"github.com/waldzellai/tinythoughts/internal/engine"
	"github.com/waldzellai/tinythoughts/internal/types"
)

var (
	cfg    *config.Config
	eng    *engine.Engine
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "tinythoughts",
	Short: "A CLI tool for systematic thinking and mental models",
	Long:  `tinythoughts provides command-line access to thinking frameworks including sequential thinking, mental models, and debugging approaches.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, err = config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		sessionID := fmt.Sprintf("cli-session-%d", time.Now().Unix())
		eng = engine.New(sessionID)

		if err := eng.LoadFrameworks(cfg.FrameworksDir); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not load frameworks: %v\n", err)
		}

		if err := eng.LoadSession(cfg.SessionFile); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not load session: %v\n", err)
		}
	},
}

var thinkCmd = &cobra.Command{
	Use:   "think [thought]",
	Short: "Sequential thinking",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		thoughtNum, _ := cmd.Flags().GetInt("number")
		total, _ := cmd.Flags().GetInt("total")
		nextNeeded, _ := cmd.Flags().GetBool("next")

		data := types.ThoughtData{
			Thought:           args[0],
			ThoughtNumber:     thoughtNum,
			TotalThoughts:     total,
			NextThoughtNeeded: nextNeeded,
		}

		result := eng.SequentialThinking(data)
		output(result)
		saveSession()
	},
}

var modelCmd = &cobra.Command{
	Use:   "model [model-name] [problem]",
	Short: "Apply a mental model",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		steps, _ := cmd.Flags().GetStringSlice("steps")
		reasoning, _ := cmd.Flags().GetString("reasoning")
		conclusion, _ := cmd.Flags().GetString("conclusion")

		data := types.MentalModelData{
			ModelName:  args[0],
			Problem:    args[1],
			Steps:      steps,
			Reasoning:  reasoning,
			Conclusion: conclusion,
		}

		result := eng.MentalModel(data)
		output(result)
		saveSession()
	},
}

var debugCmd = &cobra.Command{
	Use:   "debug [approach] [issue]",
	Short: "Apply a debugging approach",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		steps, _ := cmd.Flags().GetStringSlice("steps")
		findings, _ := cmd.Flags().GetString("findings")
		resolution, _ := cmd.Flags().GetString("resolution")

		data := types.DebuggingApproachData{
			ApproachName: args[0],
			Issue:        args[1],
			Steps:        steps,
			Findings:     findings,
			Resolution:   resolution,
		}

		result := eng.DebuggingApproach(data)
		output(result)
		saveSession()
	},
}

var frameworksCmd = &cobra.Command{
	Use:   "frameworks",
	Short: "List available frameworks",
	Run: func(cmd *cobra.Command, args []string) {
		frameworks := eng.ListFrameworks()
		for _, name := range frameworks {
			fmt.Println(name)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/tinythoughts/tinythoughts.yaml)")
	rootCmd.PersistentFlags().String("output", "json", "output format (json|text)")
	viper.BindPFlag("output_format", rootCmd.PersistentFlags().Lookup("output"))

	thinkCmd.Flags().IntP("number", "n", 1, "thought number")
	thinkCmd.Flags().IntP("total", "t", 1, "total thoughts")
	thinkCmd.Flags().BoolP("next", "x", false, "next thought needed")

	modelCmd.Flags().StringSliceP("steps", "s", []string{}, "steps")
	modelCmd.Flags().StringP("reasoning", "r", "", "reasoning")
	modelCmd.Flags().StringP("conclusion", "c", "", "conclusion")

	debugCmd.Flags().StringSliceP("steps", "s", []string{}, "steps")
	debugCmd.Flags().StringP("findings", "f", "", "findings")
	debugCmd.Flags().StringP("resolution", "r", "", "resolution")

	rootCmd.AddCommand(thinkCmd)
	rootCmd.AddCommand(modelCmd)
	rootCmd.AddCommand(debugCmd)
	rootCmd.AddCommand(frameworksCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
}

func output(result types.Result) {
	format := viper.GetString("output_format")
	if format == "json" {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Printf("Operation: %s\nStatus: %s\n", result.Operation, result.Status)
	}
}

func saveSession() {
	if err := eng.SaveSession(cfg.SessionFile); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not save session: %v\n", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
