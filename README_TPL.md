<p align="center">arlo-dl -- cli tool for downloading arlo recordings and saving them to a file</p>
<p align="center">
  <a href="https://travis-ci.org/lrstanley/arlo-dl"><img src="https://travis-ci.org/lrstanley/arlo-dl.svg?branch=master" alt="Build Status"></a>
  <a href="https://byteirc.org/channel/%23%2Fdev%2Fnull"><img src="https://img.shields.io/badge/ByteIRC-%23%2Fdev%2Fnull-blue.svg" alt="IRC Chat"></a>
</p>

## Table of Contents
- [Installation](#installation)
  - [Ubuntu/Debian](#ubuntudebian)
  - [CentOS/Redhat](#centosredhat)
  - [Manual Install](#manual-install)
  - [Build from source](#build-from-source)
- [Usage](#usage)
- [Contributing](#contributing)
- [TODO](#todo)
- [License](#license)

## Installation

Check out the [releases](https://github.com/lrstanley/arlo-dl/releases)
page for prebuilt versions. arlo-dl should work on ubuntu/debian,
centos/redhat/fedora, etc. Below are example commands of how you would install
the utility.

### Ubuntu/Debian

```bash
$ wget https://liam.sh/ghr/arlo-dl_[[tag]]_[[os]]_[[arch]].deb
$ dpkg -i arlo-dl_[[tag]]_[[os]]_[[arch]].deb
$ arlo-dl --help
```

### CentOS/Redhat

```bash
$ yum localinstall https://liam.sh/ghr/arlo-dl_[[tag]]_[[os]]_[[arch]].rpm
$ arlo-dl --help
```

Some older CentOS versions may require (if you get `Cannot open: <url>. Skipping.`):

```console
$ wget https://liam.sh/ghr/arlo-dl_[[tag]]_[[os]]_[[arch]].rpm
$ yum localinstall arlo-dl_[[tag]]_[[os]]_[[arch]].rpm
```

### Manual Install

```bash
$ wget https://liam.sh/ghr/arlo-dl_[[tag]]_[[os]]_[[arch]].tar.gz
$ tar -C /usr/bin/ -xzvf arlo-dl_[[tag]]_[[os]]_[[arch]].tar.gz arlo-dl
$ chmod +x /usr/bin/arlo-dl
$ arlo-dl --help
```

### Source

Note that you must have [Go](https://golang.org/doc/install) installed (`v1.11.1` required).

    $ git clone https://github.com/lrstanley/arlo-dl.git && cd arlo-dl
    $ make
    $ ./arlo-dl --help

## Usage

The default configuration path is `/etc/arlo-dl.yaml` when using `deb`/`rpm`.
If you are not using these package formats, copy the example config file,
`example.arlo-dl.yaml`, to `arlo-dl.yaml`.

```
$ ./arlo-dl --help
Usage:
  arlo-dl [OPTIONS]

Application Options:
  -l, --log-path=PATH    Optional path to log output to
  -c, --config=PATH      Path to configuration file (default: ./arlo-dl.yaml)

Help Options:
  -h, --help             Show this help message
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
