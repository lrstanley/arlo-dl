<p align="center">arlo-dl -- cli tool for downloading arlo recordings and saving them to a file</p>
<p align="center">
  <a href="https://travis-ci.org/lrstanley/arlo-dl"><img src="https://travis-ci.org/lrstanley/arlo-dl.svg?branch=master" alt="Build Status"></a>
  <a href="https://byteirc.org/channel/%23%2Fdev%2Fnull"><img src="https://img.shields.io/badge/ByteIRC-%23%2Fdev%2Fnull-blue.svg" alt="IRC Chat"></a>
</p>

## Table of Contents
- [Features](#features)
- [Installation](#installation)
  - [Windows](#windows)
  - [Ubuntu/Debian](#ubuntudebian)
  - [CentOS/Redhat](#centosredhat)
  - [Manual Install](#manual-install)
  - [Build from source](#build-from-source)
- [Usage](#usage)
  - [Example](#example)
- [Contributing](#contributing)
- [TODO](#todo)
- [License](#license)

## Features

   * Allows downloading Arlo recordings, for all cameras.
   * Efficient: concurrency support, and will not re-download a recording that
   was already downloaded.
   * Configurable filenames, download location and timeframe.

## Installation

Check out the [releases](https://github.com/lrstanley/arlo-dl/releases)
page for prebuilt versions. Below are example commands of how you would install
the utility.

### Windows

Download [the zip here](https://liam.sh/ghr/arlo-dl_0.0.3_windows_amd64.zip).

### Ubuntu/Debian

```bash
$ wget https://liam.sh/ghr/arlo-dl_0.0.3_linux_amd64.deb
$ dpkg -i arlo-dl_0.0.3_linux_amd64.deb
$ arlo-dl --help
```

### CentOS/Redhat

```bash
$ yum localinstall https://liam.sh/ghr/arlo-dl_0.0.3_linux_amd64.rpm
$ arlo-dl --help
```

Some older CentOS versions may require (if you get `Cannot open: <url>. Skipping.`):

```console
$ wget https://liam.sh/ghr/arlo-dl_0.0.3_linux_amd64.rpm
$ yum localinstall arlo-dl_0.0.3_linux_amd64.rpm
```

### Manual Install

```bash
$ wget https://liam.sh/ghr/arlo-dl_0.0.3_linux_amd64.tar.gz
$ tar -C /usr/bin/ -xzvf arlo-dl_0.0.3_linux_amd64.tar.gz arlo-dl
$ chmod +x /usr/bin/arlo-dl
$ arlo-dl --help
```

### Source

Note that you must have [Go](https://golang.org/doc/install) installed (`v1.11.1` required).

    $ git clone https://github.com/lrstanley/arlo-dl.git && cd arlo-dl
    $ make
    $ ./arlo-dl --help

## Usage

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

## Contributing

Please review the [CONTRIBUTING](CONTRIBUTING.md) doc for submitting issues/a guide
on submitting pull requests and helping out.

## TODO

 - [ ] add documentation flags for filename templating.

## License

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
