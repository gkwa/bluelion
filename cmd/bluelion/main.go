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

	flag.StringVar(&config.LogFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	// Assuming the order of arguments is: inputFilePath, outputFilePath
	config.InputFilePath = args[0]

	if len(args) > 1 {
		config.OutputFilePath = args[1]
	} else {
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

func printUsage() {
	println("Usage: program [-v|--verbose] [--log-format=text|json] inputFilePath [outputFilePath]")
}
