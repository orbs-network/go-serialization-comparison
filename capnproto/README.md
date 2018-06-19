# Cap'n Proto

https://github.com/capnproto/go-capnproto2

## Prerequisites

* Make sure [Capnproto](http://brewformulas.org/Capnp) is installed (version 0.6.1 or later).

  > Install with `brew install capnp`
  
  > Verify with `capnp --version`

* Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  
  > Install with ``export PATH=$PATH:`go env GOPATH`/bin``
  
  > Verify with `echo $PATH`

* Make sure [Capnproto](https://github.com/google/flatbuffers) Go runtime library is installed.
 
  > Install with `go get -u zombiezen.com/go/capnproto2/...`

  > Verify with ``ls `go env GOPATH`/src/zombiezen.com/go/capnproto2``

## Build

* Compile all `.capnp` files to `.go` files from within this directory:
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-serialization-comparison/capnproto
rm  `find . -name "*.capnp.go"`
capnp compile -I`go env GOPATH`/src/zombiezen.com/go/capnproto2/std -ogo ./**/*.capnp
```

* You should push all generated `.go` files to git as well.

## Run

* To run (this will take you to project root):
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-serialization-comparison/capnproto
go test
```