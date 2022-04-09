// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"errors"
	"fmt"
	"os"

	arlo "github.com/jeffreydwalter/arlo-go"
	"github.com/phayes/permbits"
	"gopkg.in/AlecAivazis/survey.v1"
	yaml "gopkg.in/yaml.v2"
)

var conf = &Config{}

type Config struct {
	Username string `yaml:"username" survey:"username"`
	Password string `yaml:"password" survey:"password"`
}

func readConfig(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("file does not exist; have you run \"arlo-dl setup\"?")
		}

		return err
	}

	if perms := permbits.FileMode(fi.Mode()); perms != 0o600 {
		logger.Printf("warning: permissions of %q are insecure: %s, please use 0600", path, perms)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, conf); err != nil {
		return err
	}
	return nil
}

type CommandSetup struct{}

func (cmd *CommandSetup) Execute(args []string) error {
	logger.SetFlags(0)
	cli.Ensure()

login:
	questions := []*survey.Question{
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "What is your Arlo username/email?"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "What is your Arlo password?"},
			Validate: survey.Required,
		},
	}
	err := survey.Ask(questions, conf)
	if err != nil {
		return err
	}

	logger.Println("validating login...")
	if _, err = arlo.Login(conf.Username, conf.Password); err != nil {
		logger.Println(err)
		logger.Println("please enter new credentials")
		goto login
	}

	f, err := os.OpenFile(cli.ConfigFile, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		logger.Fatalf("error creating %q: %v", cli.ConfigFile, err)
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, "---\nusername: %s\npassword: %s\n", conf.Username, conf.Password); err != nil {
		logger.Fatalf("error writing to %q: %v", cli.ConfigFile, err)
	}
	logger.Printf("successfully wrote %q", cli.ConfigFile)

	return nil
}
