package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/taylormonacelli/bluelion"
	"github.com/taylormonacelli/goldbug"
)

var config bluelion.Config

func main() {
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.Verbose, "v", false, "Enable verbose output (shorthand)")

	flag.StringVar(&config.InputFilePath, "input", "data.txt", "Input file path")
	flag.StringVar(&config.OutputFilePath, "output", "", "Output file path (default: input file path)")

	flag.StringVar(&config.LogFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	if config.OutputFilePath == "" {
		config.OutputFilePath = config.InputFilePath
	}

	if config.Verbose || config.LogFormat != "" {
		if config.LogFormat == "json" {
			goldbug.SetDefaultLoggerJson(slog.LevelDebug)
		} else {
			goldbug.SetDefaultLoggerText(slog.LevelDebug)
		}
	}

	code := bluelion.Main(config)
	os.Exit(code)
}
