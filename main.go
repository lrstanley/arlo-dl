// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"
)

var (
	version = "master"
	commit  = "latest"
	date    = "-"

	cli    = &Flags{}
	logger = log.New(os.Stdout, "", log.LstdFlags)
)

type Flags struct {
	ConfigFile    string `short:"c" long:"config-file" description:"configuration file (see 'arlo-dl setup', default: $HOME/.arlo-dl.yaml)"`
	OutputDir     string `short:"o" long:"output-dir" description:"location to store recordings" default:"arlo-recordings"`
	History       int    `long:"history" description:"how many days back to download" default:"14"`
	Quiet         bool   `short:"q" long:"quiet" description:"don't log to stdout"`
	VersionFlag   bool   `short:"v" long:"version" description:"display the version of arlo-dl and exit"`
	MaxConcurrent int    `short:"C" long:"max-concurrent" description:"maximum amount of recordings to download concurrently" default:"2"`
	NameFormat    string `short:"f" long:"name-format" description:"go-template format for the file name" default:"{{.Camera.DeviceName}}-{{.Timestamp}}-{{.Recording.Name}}.mp4"`

	CommandSetup CommandSetup `command:"setup" description:"generate a config for use with arlo-dl"`
}

func (f *Flags) Ensure() {
	if cli.Quiet {
		logger.SetOutput(ioutil.Discard)
	}

	if cli.ConfigFile == "" {
		user, err := user.Current()
		if err != nil {
			logger.Fatal(err)
		}
		cli.ConfigFile = filepath.Join(user.HomeDir, ".arlo-dl.yaml")
	}
}

func main() {
	var err error
	parser := flags.NewParser(cli, flags.HelpFlag)
	parser.SubcommandsOptional = true
	if _, err = parser.Parse(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli.Ensure()

	if parser.Active != nil {
		os.Exit(0)
	}

	logger.Printf("reading config at %q", cli.ConfigFile)
	if err = readConfig(cli.ConfigFile); err != nil {
		logger.Fatalf("error reading config: %v", err)
	}

	fetch()
}
