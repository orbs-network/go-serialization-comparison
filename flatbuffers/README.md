# Google Flatbuffers

https://github.com/google/flatbuffers

## Prerequisites

* Make sure [Flatbuffers](http://brewformulas.org/Flatbuffers) is installed (version 1.9 or later).

  > Install with `brew install flatbuffers`
  
  > Verify with `flatc --version`
  
* Make sure [Flatbuffers](https://github.com/google/flatbuffers) Go runtime library is installed.
 
  > Install with `go get github.com/google/flatbuffers/go`

  > Verify with ``ls `go env GOPATH`/src/github.com/google/flatbuffers``

## Build

* Compile all `.fbs` files to `.go` files from within this directory:
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-serialization-comparison/flatbuffers
flatc -g **/*.fbs
```

* You should push all generated `.go` files to git as well.

## Run

* To run (this will take you to project root):
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-serialization-comparison/protobuf
go test
```