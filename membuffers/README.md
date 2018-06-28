# MemBuffers

https://github.com/orbs-network/membuffers

## Prerequisites

* Make sure MemBuffers [compiler](https://github.com/orbs-network/membuffers/go/membufc) is installed (version 0.0.1 or later).

  > Install with `brew install orbs-network/membuffers/membufc`
  
  > Verify with `membufc --version`
  
* Make sure [MemBuffers](https://github.com/orbs-network/membuffers) Go runtime library is installed.
 
  > Install with `go get github.com/orbs-network/membuffers/go`

  > Verify with ``ls `go env GOPATH`/src/github.com/orbs-network/membuffers``

## Build

* Compile all `.proto` files to `.go` files from within this directory:
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-serialization-comparison/membuffers
rm  `find . -name "*.mb.go"`
membufc --go ./**/*.proto
```

* You should push all generated `.go` files to git as well.

## Run

* To run (this will take you to project root):
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-serialization-comparison/membuffers
go test
```