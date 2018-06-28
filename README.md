# Go Serialization Comparison

Comparison of various serialization formats for Go

## Requirements

Assume this example Protobuf transaction format:

```proto
message Transaction {
  TransactionHeader header = 1; 
  TransactionPayload payload = 2;
  bytes signature = 3; // 64B, signature over the hash of the header and payload together
}
  
message TransactionHeader {
  uint32 protocol_version = 1;
  uint64 virtual_chain = 2;
  TransactionSender sender = 3;
  uint64 time_stamp = 4; 
}
  
message TransactionSender {
  AddressScheme address_scheme = 1;
  NetworkType network_type = 2;
  bytes public_key = 3; // 32B
}
  
enum AddressScheme {
  RESERVED_SCHEME = 0;
  SCHEME01 = 1;
  FUTURE_SCHEME = 255;
}
  
enum NetworkType {
  RESERVED_NETWORK = 0;
  MAIN_NET = 77;
  TEST_NET = 84;
  FUTURE_NETWORK = 255;
}
  
message TransactionPayload {
  string contract_name = 1;
  string method_name = 2;
  InputArguments input_arguments = 3;
}
  
message InputArguments {
  repeated MethodCallArgument argument = 1; // variable length
}
  
message MethodCallArgument {
  oneof type {
    uint32 uint32 = 1;
    uint64 uint64 = 2;
    string string = 3;
    bytes bytes = 4;
  }
}
```

#### Non-trivial tasks

1. Verify `Transaction.signature` over the raw serialized data
2. Hash `TransactionSender` (eg. place it as a key in state)
3. Avoid decoding `TransactionPayload` since most services don't need its internals and it can be expensive

## Compared Encodings

#### Google Protobuf

> https://github.com/golang/protobuf

Full details [here](protobuf/README.md).

#### Google FlatBuffers

> https://github.com/google/flatbuffers

Full details [here](flatbuffers/README.md).

#### Cap'n Proto

> https://github.com/capnproto/go-capnproto2

Full details [here](capnproto/README.md).

#### MemBuffers

> https://github.com/orbs-network/membuffers

Full details [here](membuffers/README.md).

## Prerequisites

* Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  
  > Verify with `go version`

* Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  
  > Install with ``export PATH=$PATH:`go env GOPATH`/bin``
  
  > Verify with `echo $PATH`

* Make sure Git is installed (version 2 or later).
  
  > Verify with `git --version`

## Build

* Clone the repo to your Go workspace:
```
cd `go env GOPATH`
go get github.com/orbs-network/go-serialization-comparison
cd src/github.com/orbs-network/go-serialization-comparison
git checkout master
```

* Follow the instructions inside each encoding folder.