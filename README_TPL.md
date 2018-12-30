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

Download [the zip here](https://liam.sh/ghr/arlo-dl_[[tag]]_windows_[[arch]].zip).

### Ubuntu/Debian

```bash
$ wget https://liam.sh/ghr/arlo-dl_[[tag]]_linux_[[arch]].deb
$ dpkg -i arlo-dl_[[tag]]_linux_[[arch]].deb
$ arlo-dl --help
```

### CentOS/Redhat

```bash
$ yum localinstall https://liam.sh/ghr/arlo-dl_[[tag]]_linux_[[arch]].rpm
$ arlo-dl --help
```

Some older CentOS versions may require (if you get `Cannot open: <url>. Skipping.`):

```console
$ wget https://liam.sh/ghr/arlo-dl_[[tag]]_linux_[[arch]].rpm
$ yum localinstall arlo-dl_[[tag]]_linux_[[arch]].rpm
```

### Manual Install

```bash
$ wget https://liam.sh/ghr/arlo-dl_[[tag]]_linux_[[arch]].tar.gz
$ tar -C /usr/bin/ -xzvf arlo-dl_[[tag]]_linux_[[arch]].tar.gz arlo-dl
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
$ arlo-dl
2018/12/30 02:41:34 reading config at "/home/myuser/.arlo-dl.yaml"
2018/12/30 02:41:34 logging into arlo with account: user@domain.com
2018/12/30 02:41:43 login successful
2018/12/30 02:41:43 looking for cameras on account
2018/12/30 02:41:43 found camera "Garage" (id: 5EM1847KA246A)
2018/12/30 02:41:43 found camera "Front door" (id: 5EM1847PA9866)
2018/12/30 02:41:43 renaming "Front door" to "Front_door"
2018/12/30 02:41:43 fetching library
2018/12/30 02:41:44 successfully fetched library; 7 items found
2018/12/30 02:41:44 skipping Front_door/1546120422020, already downloaded
2018/12/30 02:41:44 skipping Front_door/1546144499779, already downloaded
2018/12/30 02:41:44 skipping Front_door/1546119340985, already downloaded
2018/12/30 02:41:44 streaming recording Garage/1546112259159 to file: "arlo-recordings/Garage-2018.12.29-14.37.39-1546112259159.mp4"
2018/12/30 02:41:44 streaming recording Garage/1546112144203 to file: "arlo-recordings/Garage-2018.12.29-14.35.44-1546112144203.mp4"
2018/12/30 02:41:45 finished downloading "1546112144203"
2018/12/30 02:41:45 streaming recording Garage/1546107734692 to file: "arlo-recordings/Garage-2018.12.29-13.22.14-1546107734692.mp4"
2018/12/30 02:41:45 finished downloading "1546112259159"
2018/12/30 02:41:45 skipping Front_door/1546107714995, already downloaded
2018/12/30 02:41:46 finished downloading "1546107734692"
```

And, whala!

```console
$ $ tree arlo-recordings/
arlo-recordings/
├── Front_door-2018.12.29-13.21.54-1546107714995.mp4
├── Front_door-2018.12.29-16.35.40-1546119340985.mp4
├── Front_door-2018.12.29-16.53.42-1546120422020.mp4
├── Front_door-2018.12.29-23.34.59-1546144499779.mp4
├── Garage-2018.12.29-13.22.14-1546107734692.mp4
├── Garage-2018.12.29-14.35.44-1546112144203.mp4
└── Garage-2018.12.29-14.37.39-1546112259159.mp4

0 directories, 7 files
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
