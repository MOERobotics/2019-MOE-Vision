# 2019-MOE-Vision

## Tech Stack

- GoCV Go wrapper for OpenCV
- GoNum for numerical computations

## Prerequisites

### Go

#### Windows

Download Go from [the Go downloads page](https://golang.org/dl/) for Windows. Run the installer.

#### Mac

Install [Homebrew](https://brew.sh/) and run `brew install go`

### GoCV

We're using [GoCV](https://github.com/hybridgroup/gocv) as the Go wrapper for OpenCV. Follow its instructions on how to install GoCV and OpenCV for your system. Ignore anything under "Cache Builds" or "Custom Environment"

- [Ubuntu / Linux](https://github.com/hybridgroup/gocv#ubuntulinux)
- [Raspbian](https://github.com/hybridgroup/gocv#raspbian)
- [Mac](https://github.com/hybridgroup/gocv#macos)
- [Windows](https://github.com/hybridgroup/gocv#windows)

### Visual Studio Code

You can use any editor to write Go code but [VS Code](https://code.visualstudio.com/) has a [great extension for Go](https://code.visualstudio.com/docs/languages/go). Be sure to click "Install All" when prompted by Visual Studio Code to install Go tools and utilities.

## Building

To build `2019-MOE-Vision.go` run `go build`. This generates a `2019-MOE-Vision` or `2019-MOE-Vision.exe` executable. You can run this in Terminal with `./2019-MOE-Vision` or in PowerShell `./2019-MOE-Vision.exe`. When you run `go build` using recent versions of Go, it should pull in packages automatically using `go mod`.

## Installing and Deployment

_TBD_

## References

- [GoCV docs](https://godoc.org/gocv.io/x/gocv)
- [GoNum docs](https://godoc.org/gonum.org/v1/gonum)
