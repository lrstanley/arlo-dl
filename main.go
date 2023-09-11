// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"os/user"
	"path/filepath"

	"github.com/apex/log"
	"github.com/lrstanley/clix"
)

const (
	version = "master"
	commit  = "latest"
	date    = "-"
)

var (
	logger log.Interface
	cli    = &clix.CLI[Flags]{
		Links: clix.GithubLinks("github.com/lrstanley/arlo-dl", "master", "https://liam.sh"),
		VersionInfo: &clix.VersionInfo[Flags]{
			Version: version,
			Commit:  commit,
			Date:    date,
		},
	}
)

type Flags struct {
	ConfigFile    string `short:"c" long:"config-file" description:"configuration file (see 'arlo-dl setup', default: $HOME/.arlo-dl.yaml)"`
	OutputDir     string `short:"o" long:"output-dir" description:"location to store recordings" default:"arlo-recordings"`
	History       int    `long:"history" description:"how many days back to download" default:"14"`
	MaxConcurrent int    `short:"C" long:"max-concurrent" description:"maximum amount of recordings to download concurrently" default:"2"`
	NameFormat    string `short:"f" long:"name-format" description:"go-template format for the file name" default:"{{.Camera.DeviceName}}/{{.Time.Year}}/{{.Time.Month}}/{{.Timestamp}}-{{.Recording.Name}}.mp4"`

	CommandSetup CommandSetup `command:"setup" description:"generate a config for use with arlo-dl"`
}

func (f *Flags) Ensure() {
	if cli.Flags.ConfigFile == "" {
		usr, err := user.Current()
		if err != nil {
			logger.WithError(err).Fatal("failed to get current user")
		}
		cli.Flags.ConfigFile = filepath.Join(usr.HomeDir, ".arlo-dl.yaml")
	}
}

func main() {
	cli.LoggerConfig.Pretty = true
	_ = cli.ParseWithInit(func() error {
		logger = cli.Logger
		cli.Flags.Ensure()

		return nil
	}, clix.OptSubcommandsOptional)

	logger = cli.Logger
	cli.Flags.Ensure()

	logger.WithField("config", cli.Flags.ConfigFile).Info("reading config")
	if err := readConfig(cli.Flags.ConfigFile); err != nil {
		logger.WithError(err).Fatal("failed to read config")
	}

	fetch()
}
