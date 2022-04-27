package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
    "path/filepath"

	"go-rest-api/app"
)

// main is the entry point into our application. However, it provides poor
// usability since it does not allow us to return errors like most Go programs.
// Instead, we delegate most of our program to the Run() function.
func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()


	// If we have an application error (app.Error) then we can just display the
	// message. If we have any other error, print the raw error message.
	var e *app.Error
	if err := Run(ctx, os.Args[1:]); errors.As(err, &e) {
		fmt.Fprintln(os.Stderr, e.Message)
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Run executes the main program.
func Run(ctx context.Context, args []string) error {
	// Shift off subcommand from the argument list, if available.
	var cmd string
	if len(args) > 0 {
		cmd, args = args[0], args[1:]
	}

	fs := flag.NewFlagSet("", flag.ExitOnError)
	var configFileName string
	fs.StringVar(&configFileName, "config", "", "config file, required")

	// Lint disabled because this is set to "flag.ExitOnError".
	fs.Parse(args) //nolint:golangcilint,errcheck

	configFileName = strings.TrimSpace(configFileName)

	if configFileName == "" {
		fs.Usage()
		os.Exit(1)
	}

	configFileFullPath, err := filepath.Abs(configFileName)
	if err != nil {
		return err
	}

	cfg, err := app.NewConfig(configFileFullPath)
	if err != nil {
		return err
	}

	// Delegate subcommands to their own Run() methods.
	switch cmd {
	case "db-setup":
		switch cfg.Env {
		case app.ENV_DEVELOPMENT:
			if err := dbSQLiteSetup(cfg); err != nil {
				return err
			}
		default:
			panic("TODO")
		}

		return nil

	case "db-drop":
		switch cfg.Env {
		case app.ENV_DEVELOPMENT:
			if err := dbSQLiteDrop(cfg); err != nil {
				return err
			}
		default:
			panic("TODO")
		}

		return nil

	case "", "-h", "help":
		usage()
		return flag.ErrHelp
	default:
		return fmt.Errorf("app %s: unknown command", cmd)
	}
}

// usage prints the top-level CLI usage message.
func usage() {
	fmt.Println(`
Command line utility for interacting with the WTF Dial service.

Usage:

	appctl <command> --config config-file [args]

The commands are:

	db-setup	setup the database
	db-drop		drop the database
	db-reset	reset the database
`[1:])
}
