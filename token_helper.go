// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func mainToken(session string, args []string) int {
	if len(args) == 0 {
		return 1
	}

	path := fmt.Sprintf("tmp/sessions/%s", session)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return 2
	}

	log.Printf("[DEBUG] token: %v %s", args, path)
	helper := &tokenCommand{Path: path}
	return helper.Run(args)
}

type tokenCommand struct {
	Path string
}

func (c *tokenCommand) Run(args []string) int {
	path := c.Path
	if path == "" {
		panic("Path must be set")
	}

	f := flag.NewFlagSet("token-disk", flag.ContinueOnError)
	f.StringVar(&path, "path", c.Path, "")
	if err := f.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "\n%s\n", err)
		return 1
	}

	path, err := homedir.Expand(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error expanding directory: %s\n", err)
		return 1
	}

	args = f.Args()
	switch args[0] {
	case "get":
		f, err := os.Open(path)
		if os.IsNotExist(err) {
			return 0
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return 1
		}
		defer f.Close()

		if _, err := io.Copy(os.Stdout, f); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return 1
		}
	case "store":
		f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return 1
		}
		defer f.Close()

		if _, err := io.Copy(f, os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return 1
		}
	case "erase":
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return 1
		}
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown subcommand: %s\n", args[0])
		return 1
	}

	return 0
}
