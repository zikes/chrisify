FROM golang:latest

RUN go get -u -d gocv.io/x/gocv
RUN cd src/gocv.io/x/gocv
RUN make deps
RUN make download
RUN make build
RUN make cleanup
RUN source ./env.sh
RUN go run ./cmd/version/main.go




CMD ["chrisify"]
