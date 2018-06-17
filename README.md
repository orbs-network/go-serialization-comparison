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

* Verify `Transaction.signature` over the raw serialized data
* Hash `TransactionSender` (eg. place it as a key in state)
* Avoid decoding `TransactionPayload` since most services don't need its internals and it can be expensive

## Encodings

#### Google Protobuf

> https://github.com/golang/protobuf

#### Amino (Protobuf for blockchain)

> https://github.com/tendermint/go-amino

#### Google Flatbuffers

> https://github.com/google/flatbuffers

#### Cap'n Proto

> https://github.com/capnproto/go-capnproto2
