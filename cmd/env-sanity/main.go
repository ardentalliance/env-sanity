package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ardentalliance/env-sanity/internal/envfile"
	"github.com/ardentalliance/env-sanity/internal/validator"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "check":
		runCheck(os.Args[2:])
	case "help", "-h", "--help", "-help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

// run env file check
func runCheck(args []string) {
	checkCommand := flag.NewFlagSet("check", flag.ExitOnError)

	envPath := checkCommand.String("env", ".env", "Path to .env file")
	schemaPath := checkCommand.String("schema", "env.schema.json", "Path to schema file")
	jsonOutput := checkCommand.Bool("json", false, "Print validation results as JSON output")

	if err := checkCommand.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	envValues, err := envfile.Parse(*envPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL ", err)
		os.Exit(1)
	}

	schema, err := validator.LoadSchema(*schemaPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL ", err)
		os.Exit(1)
	}

	results := validator.Validate(envValues, schema)

	if *jsonOutput {
		output, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "FAIL could not create JSON output:", err)
			os.Exit(1)
		}

		fmt.Println(string(output))

		if validator.HasFailures(results) {
			os.Exit(1)
		}

		return
	}

	fmt.Println()
	fmt.Println("env-sanity")
	fmt.Println()

	for _, result := range results {
		fmt.Printf("%-5s %-24s %s\n", result.Level, result.Key, result.Message)
	}

	if validator.HasFailures(results) {
		os.Exit(1)
	}
}

// usage manual
func printUsage() {
	fmt.Println(`env-sanity

Validate .env files against a simple JSON schema file.

Usage:
    env-sanity check [options]
	env-sanity help

Options: 
    -env string
		Path to .env file to validate.
		Default: .env

    -schema string
		Path to the JSON schema file that defines required variables,
		allowed values, min lengths, and simple types.
		Default: env.schema.json

	-json
		Print validation results as JSON output.
		Useful for scripts, CI jobs, or automated checks.

Expected files:
	By default, env-sanity looks for these files in the current directory:
		.env
		env.schema.json

Examples:
	Check the default .env file against the default schema file:

    	env-sanity check

	Check a different environment file:

		env-sanity check -env .env.local

	Check a specific .env file against a specific schema file:

    	env-sanity check -env .env.production -schema config/env.schema.json

	Print machine-readable JSON output:
		
		env-sanity check -json

Schema example:
	{
		"DATABASE_URL": {
			"required": true
		},
		"SESSION_SECRET": {
			"required": true,
			"minLength": 32,
			"placeholderAllowed": true
		},
		"NODE_ENV": {
			"required": false,
			"allowedValues": ["development", "test", "production"]
		},
		"SMTP_PORT": {
     		"required": false,
			"type": number
		}
	}

Exit Codes:
	0	No validation failures were found.
	1 	At least one validation failure was found, or a file could not be read.`)
}
