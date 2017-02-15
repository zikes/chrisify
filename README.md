# chrisify

## Linux Install

1. Install the OpenCV Developer package. On Ubuntu systems that's `sudo apt install libopencv-dev`

2. `git clone git@git.nwaonline.com:jhutchinson/chrisify.git $GOPATH/chrisify`

3. `go get github.com/lazywei/go-opencv`

4. `cd $GOPATH/chrisify && go build`

5. `./chrisify path/to/image.jpg > output.jpg`
