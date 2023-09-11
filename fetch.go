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

	"github.com/apex/log"
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
	Time      time.Time
	Timestamp string
}

func fetch() {
	logger.WithField("account", conf.Username).Info("logging into arlo")
	api, err := arlo.Login(conf.Username, conf.Password)
	if err != nil {
		logger.Fatalf("failed to login: %s\n", err)
	}
	logger.Info("successfully logged in")

	cmap := make(map[string]*arlo.Camera)

	logger.Info("looking for cameras on account")
	for i := 0; i < len(api.Cameras); i++ {
		logger.WithFields(log.Fields{
			"name": api.Cameras[i].DeviceName,
			"id":   api.Cameras[i].DeviceId,
		}).Info("found camera")

		name := reStripName.ReplaceAllString(api.Cameras[i].DeviceName, "_")
		name = reTrimUnderscore.ReplaceAllString(name, "_")
		name = strings.Trim(name, "_")
		if name != api.Cameras[i].DeviceName {
			logger.WithFields(log.Fields{
				"old": api.Cameras[i].DeviceName,
				"new": name,
			}).Info("renaming camera")

			api.Cameras[i].DeviceName = name
		}

		cmap[api.Cameras[i].DeviceId] = &api.Cameras[i]
	}

	now := time.Now()
	start := now.Add(-time.Duration(cli.Flags.History) * 24 * time.Hour)

	logger.Info("fetching library")
	library, err := api.GetLibrary(start, now)
	if err != nil {
		logger.WithError(err).Fatal("failed to fetch library")
	}
	logger.WithField("count", len(*library)).Info("successfully fetched library")

	pool := sempool.New(cli.Flags.MaxConcurrent)

	for _, recording := range *library { // nolint:gocritic
		pool.Slot()

		go func(r *arlo.Recording) {
			defer pool.Free()

			rtmpl := &RecordingTemplate{
				Recording: r,
				Camera:    cmap[r.DeviceId],
				Time:      time.Unix(0, r.UtcCreatedDate*int64(time.Millisecond)),
			}
			rtmpl.Timestamp = rtmpl.Time.Format("2006.01.02-15.04.05")

			filename := strings.Builder{}
			tpl := template.Must(template.New("filename").Parse(cli.Flags.NameFormat))

			err = tpl.Execute(&filename, rtmpl)
			if err != nil {
				logger.Fatalf("error executing filename template for recording %q: %v", r.UniqueId, err)
			}

			fullFn := filepath.Join(cli.Flags.OutputDir, filename.String())

			err = os.MkdirAll(filepath.Dir(fullFn), 0o755)
			if err != nil {
				logger.Fatalf("error creating dir %q: %v", filepath.Dir(fullFn), err)
			}

			_, err = os.Stat(fullFn)
			if err == nil {
				logger.WithFields(log.Fields{
					"camera": rtmpl.Camera.DeviceName,
					"name":   r.Name,
				}).Info("skipping, already downloaded")
				return
			}

			var f *os.File
			f, err = os.Create(fullFn)
			if err != nil {
				logger.WithError(err).Fatal("failed to create file")
			}
			defer f.Close()

			logger.WithFields(log.Fields{
				"camera": rtmpl.Camera.DeviceName,
				"name":   r.Name,
			}).Info("downloading")

			err = api.DownloadFile(r.PresignedContentUrl, f)
			if err != nil {
				logger.WithError(err).Fatal("failed to download file")
			}
			logger.WithField("name", r.Name).Info("finished downloading")
		}(&recording) // nolint:gosec
	}

	pool.Wait()
}
