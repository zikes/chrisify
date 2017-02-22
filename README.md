# chrisify

## Linux Install

1. Install the OpenCV Developer package. On Ubuntu systems that's `sudo apt install libopencv-dev`

2. `go get github.com/zikes/chrisify`

3. `go get github.com/lazywei/go-opencv`

4. `cd $GOPATH/src/github.com/zikes/chrisify && go build`

## Usage


Simplest: `chrisify path/to/image.jpg > output.jpg`

If executed from any location besides the repository, you must tell it where to find the
bundled Haar Cascade face recognition XML file. I tried to bundle it with the binary, but
it must be provided as a file to the OpenCV library, so a file path is necessary.

`chrisify --haar /path/to/haarcascade_frontalface_alt.xml /path/to/input.jpg > output.jpg`

If you'd like to use different faces, you can provide a directory of PNG files to be imported:

`chrisify --faces /path/to/faces /path/to/input.jpg > output.jpg`
