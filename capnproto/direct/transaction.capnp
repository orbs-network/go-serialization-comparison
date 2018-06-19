using Go = import "/go.capnp";
@0xd8ebca068ad880fd;
$Go.package("direct");
$Go.import("capnproto/direct");

struct Transaction {
    data @0 :TransactionData;
    signature @1 :Data;
}

struct TransactionData {
    name @0 :Text;
}