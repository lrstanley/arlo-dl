<!-- template:begin:header -->
<!-- do not edit anything in this "template" block, its auto-generated -->
<p align="center">arlo-dl -- :movie_camera: :camera: cli tool for downloading arlo recordings and saving them to a file (add to a cron!)</p>
<p align="center">
  <a href="https://github.com/lrstanley/arlo-dl/releases">
    <img alt="Release Downloads" src="https://img.shields.io/github/downloads/lrstanley/arlo-dl/total?style=flat-square">
  </a>


  <a href="https://github.com/lrstanley/arlo-dl/actions?query=workflow%3Arelease+event%3Apush">
    <img alt="GitHub Workflow Status (release @ master)" src="https://img.shields.io/github/workflow/status/lrstanley/arlo-dl/release/master?label=release&style=flat-square&event=push">
  </a>


  <a href="https://github.com/lrstanley/arlo-dl/actions?query=workflow%3Atest+event%3Apush">
    <img alt="GitHub Workflow Status (test @ master)" src="https://img.shields.io/github/workflow/status/lrstanley/arlo-dl/test/master?label=test&style=flat-square&event=push">
  </a>

  <img alt="Code Coverage" src="https://img.shields.io/codecov/c/github/lrstanley/arlo-dl/master?style=flat-square">

  <a href="https://pkg.go.dev/github.com/lrstanley/arlo-dl">
    <img alt="Go Documentation" src="https://pkg.go.dev/badge/github.com/lrstanley/arlo-dl?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/lrstanley/arlo-dl">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/lrstanley/arlo-dl?style=flat-square">
  </a>
  <img alt="Bug reports" src="https://img.shields.io/github/issues/lrstanley/arlo-dl/bug?label=issues&style=flat-square">
  <img alt="Feature requests" src="https://img.shields.io/github/issues/lrstanley/arlo-dl/enhancement?label=feature%20requests&style=flat-square">
  <a href="https://github.com/lrstanley/arlo-dl/pulls">
    <img alt="Open Pull Requests" src="https://img.shields.io/github/issues-pr/lrstanley/arlo-dl?style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/arlo-dl/releases">
    <img alt="Latest Semver Release" src="https://img.shields.io/github/v/release/lrstanley/arlo-dl?style=flat-square">
    <img alt="Latest Release Date" src="https://img.shields.io/github/release-date/lrstanley/arlo-dl?style=flat-square">
  </a>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/lrstanley/arlo-dl?style=flat-square">
  <a href="https://github.com/lrstanley/arlo-dl/discussions/new?category=q-a">
    <img alt="Ask a Question" src="https://img.shields.io/badge/discussions-ask_a_question!-green?style=flat-square">
  </a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/discord-bytecord-blue.svg?style=flat-square" alt="Discord Chat"></a>
</p>
<!-- template:end:header -->

<!-- template:begin:toc -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :link: Table of Contents

  - [Features](#sparkles-features)
  - [Installation](#computer-installation)
    - [Windows](#windows)
    - [Ubuntu/Debian](#ubuntudebian)
    - [CentOS/Redhat](#centosredhat)
    - [Manual Install](#manual-install)
    - [Source](#gear-source)
  - [Usage](#toolbox-usage)
    - [Example](#example)
  - [Support & Assistance](#raising_hand_man-support-assistance)
  - [Contributing](#handshake-contributing)
  - [License](#balance_scale-license)
<!-- template:end:toc -->

## :sparkles: Features

   * Allows downloading Arlo recordings, for all cameras.
   * Efficient: concurrency support, and will not re-download a recording that
   was already downloaded.
   * Configurable filenames, download location and timeframe.

## :computer: Installation

Check out the [releases](https://github.com/lrstanley/arlo-dl/releases)
page for prebuilt versions.

<!-- template:begin:ghcr -->
<!-- do not edit anything in this "template" block, its auto-generated -->

<!-- template:end:ghcr -->

## :toolbox: Usage

```
$ ./arlo-dl -h
Usage:
  arlo-dl [OPTIONS] [setup]

Application Options:
  -c, --config-file=    configuration file (see 'arlo-dl setup', default: $HOME/.arlo-dl.yaml)
  -o, --output-dir=     location to store recordings (default: arlo-recordings)
      --history=        how many days back to download (default: 14)
  -q, --quiet           don't log to stdout
  -v, --version         display the version of arlo-dl and exit
  -C, --max-concurrent= maximum amount of recordings to download concurrently (default: 2)
  -f, --name-format=    go-template format for the file name (default: {{.Camera.DeviceName}}-{{.Timestamp}}-{{.Recording.Name}}.mp4)

Help Options:
  -h, --help            Show this help message

Available commands:
  setup  generate a config for use with arlo-dl
```

### Example

```console
$ arlo-dl setup
? What is your Arlo username/email? user@domain.com
? What is your Arlo password? *******
validating login...
successfully wrote "/home/myuser/.arlo-dl.yaml"
```

```console
$ ./arlo-dl
2019/01/02 23:08:59 reading config at "/home/myuser/.arlo-dl.yaml"
2019/01/02 23:08:59 logging into arlo with account: user@domain.com
2019/01/02 23:09:08 login successful
2019/01/02 23:09:08 looking for cameras on account
2019/01/02 23:09:08 found camera "Back yard" (id: SOMEID1234)
2019/01/02 23:09:08 renaming "Back yard" to "Back_yard"
2019/01/02 23:09:08 found camera "Garage" (id: SOMEID1234)
2019/01/02 23:09:08 found camera "Front door" (id: SOMEID1234)
2019/01/02 23:09:08 renaming "Front door" to "Front_door"
2019/01/02 23:09:08 found camera "Living room" (id: SOMEID1234)
2019/01/02 23:09:08 renaming "Living room" to "Living_room"
2019/01/02 23:09:08 fetching library
2019/01/02 23:09:11 successfully fetched library; 84 items found
2019/01/02 23:09:11 skipping Back_yard/1546462326993, already downloaded
2019/01/02 23:09:13 streaming recording Living_room/1546460905942 to file: "arlo-recordings/Living_room/2019/January/2019.01.02-15.28.25-1546460905942.mp4"
2019/01/02 23:09:14 finished downloading "1546460905942"
2019/01/02 23:09:14 skipping Garage/1546261944033, already downloaded
2019/01/02 23:09:14 skipping Front_door/1546261735945, already downloaded
2019/01/02 23:09:14 skipping Front_door/1546254747308, already downloaded
2019/01/02 23:09:14 skipping Front_door/1546254712723, already downloaded
2019/01/02 23:09:14 skipping Front_door/1546254674837, already downloaded
2019/01/02 23:09:14 skipping Front_door/1546234068210, already downloaded
2019/01/02 23:09:14 streaming recording Front_door/1546217634687 to file: "arlo-recordings/Front_door/2018/December/2018.12.30-19.53.54-1546217634687.mp4"
2019/01/02 23:09:14 finished downloading "1546461044310"
2019/01/02 23:09:16 skipping Front_door/1546107714995, already downloaded
2019/01/02 23:09:16 finished downloading "1546186896127"
```

And, whala!

```console
$ tree arlo-recordings/
arlo-recordings/
├── Back_yard
│   └── 2019
│       └── January
│           ├── 2019.01.01-15.56.33-1546376193394.mp4
│           ├── 2019.01.01-15.58.32-1546376312277.mp4
│           ├── 2019.01.01-15.58.48-1546376328088.mp4
│           ├── 2019.01.01-16.04.10-1546376650187.mp4
            └── [...]
├── Front_door
│   ├── 2018
│   │   └── December
│   │       ├── 2018.12.29-13.21.54-1546107714995.mp4
│   │       ├── 2018.12.29-16.35.40-1546119340985.mp4
│   │       ├── 2018.12.29-16.53.42-1546120422020.mp4
            └── [...]
│   └── 2019
│       └── January
│           ├── 2019.01.01-06.04.24-1546340664937.mp4
            └── [...]
├── Garage
│   └── 2019
│       └── January
│           ├── 2019.01.01-14.04.05-1546369445712.mp4
│           ├── 2019.01.01-14.07.04-1546369624175.mp4
│           └── 2019.01.01-15.42.49-1546375369041.mp4
└── Living_room
    └── 2019
        └── January
            ├── 2019.01.01-17.31.56-1546381916614.mp4
            ├── 2019.01.01-17.32.42-1546381962426.mp4
            ├── 2019.01.01-17.47.18-1546382838942.mp4
            └── [...]

16 directories, 91 files
```

<!-- template:begin:support -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :raising_hand_man: Support & Assistance

   * :heart: Please review the [Code of Conduct](CODE_OF_CONDUCT.md) for
     guidelines on ensuring everyone has the best experience interacting with
     the community.
   * :raising_hand_man: Take a look at the [support](SUPPORT.md) document on
     guidelines for tips on how to ask the right questions.
   * :lady_beetle: For all features/bugs/issues/questions/etc, [head over here](https://github.com/lrstanley/arlo-dl/issues/new/choose).
<!-- template:end:support -->

<!-- template:begin:contributing -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :handshake: Contributing

   * :heart: Please review the [Code of Conduct](CODE_OF_CONDUCT.md) for guidelines
     on ensuring everyone has the best experience interacting with the
	   community.
   * :clipboard: Please review the [contributing](CONTRIBUTING.md) doc for submitting
     issues/a guide on submitting pull requests and helping out.
   * :old_key: For anything security related, please review this repositories [security policy](https://github.com/lrstanley/arlo-dl/security/policy).
<!-- template:end:contributing -->

<!-- template:begin:license -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :balance_scale: License

```
MIT License

Copyright (c) 2018 Liam Stanley <me@liamstanley.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

_Also located [here](LICENSE)_
<!-- template:end:license -->
