// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	arlo "github.com/jeffreydwalter/arlo-go"
	sempool "github.com/lrstanley/go-sempool"
)

var (
	reStripName      = regexp.MustCompile(`[^a-zA-Z0-9_~.-]+`)
	reTrimUnderscore = regexp.MustCompile(`_+`)
)

type RecordingTemplate struct {
	Recording *arlo.Recording
	Camera    *arlo.Camera
	Timestamp string
}

func fetch() {
	logger.Printf("logging into arlo with account: %s", conf.Username)
	api, err := arlo.Login(conf.Username, conf.Password)
	if err != nil {
		logger.Fatalf("failed to login: %s\n", err)
	}
	logger.Println("login successful")

	cmap := make(map[string]*arlo.Camera)

	logger.Println("looking for cameras on account")
	for _, camera := range api.Cameras {
		logger.Printf("found camera %q (id: %s)", camera.DeviceName, camera.DeviceId)
		name := reStripName.ReplaceAllString(camera.DeviceName, "_")
		name = reTrimUnderscore.ReplaceAllString(name, "_")
		name = strings.Trim(name, "_")
		if name != camera.DeviceName {
			logger.Printf("renaming %q to %q", camera.DeviceName, name)
			camera.DeviceName = name
		}

		cmap[camera.DeviceId] = &camera
	}

	now := time.Now()
	start := now.Add(-time.Duration(cli.History) * 24 * time.Hour)

	logger.Println("fetching library")
	library, err := api.GetLibrary(start, now)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("successfully fetched library; %d items found", len(*library))

	if len(*library) > 0 {
		if err := os.MkdirAll(cli.OutputDir, 0755); err != nil {
			logger.Fatalf("error creating dir %q: %v", cli.OutputDir, err)
		}
	}

	pool := sempool.New(cli.MaxConcurrent)

	for _, recording := range *library {
		pool.Slot()

		rtmpl := &RecordingTemplate{
			Recording: &recording,
			Camera:    cmap[recording.DeviceId],
			Timestamp: time.Unix(0, recording.UtcCreatedDate*int64(time.Millisecond)).Format(("2006.01.02-15.04.05")),
		}

		filename := strings.Builder{}
		tpl := template.Must(template.New("filename").Parse(cli.NameFormat))
		if err := tpl.Execute(&filename, rtmpl); err != nil {
			logger.Fatalf("error executing filename template for recording %q: %v", recording.UniqueId, err)
		}

		go func(r arlo.Recording, fn string) {
			defer pool.Free()

			fullFn := filepath.Join(cli.OutputDir, fn)

			if _, err := os.Stat(fullFn); err == nil {
				logger.Printf("skipping %s/%s, already downloaded", rtmpl.Camera.DeviceName, r.Name)
				return
			}

			f, err := os.Create(fullFn)
			if err != nil {
				logger.Fatal(err)
			}
			defer f.Close()

			logger.Printf("streaming recording %s/%s to file: %q", rtmpl.Camera.DeviceName, r.Name, fullFn)
			if err := api.DownloadFile(r.PresignedContentUrl, f); err != nil {
				logger.Fatal(err)
			}
			logger.Printf("finished downloading %q", r.Name)
		}(recording, filename.String())
	}

	pool.Wait()
}
