package capnproto

import (
	"testing"
	"fmt"
	"zombiezen.com/go/capnproto2"
	"github.com/orbs-network/go-serialization-comparison/capnproto/direct"
	"github.com/orbs-network/go-serialization-comparison/util"
)

func TestDirect(t *testing.T) {

	fmt.Println("----------------------------------")
	fmt.Println(" TestDirect")
	fmt.Println("----------------------------------")

	// create the source transaction

	msg1, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		t.Fatal("new message error: ", err)
	}
	transactionData, err := direct.NewTransactionData(seg)
	if err != nil {
		t.Fatal("new message error: ", err)
	}
	transactionData.SetName("johnny")
	transaction, err := direct.NewRootTransaction(seg)
	transaction.SetData(transactionData)
	transaction.SetSignature([]byte{0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22})

	// prepare the wire

	encoded, err := msg1.Marshal()
	if err != nil {
		t.Fatal("marshal error: ", err)
	}

	fmt.Println("Encoded transaction:")
	util.PrettyPrintHex(encoded)

	// decode the wire

	msg2, err := capnp.Unmarshal(encoded)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	newDecoded, err := direct.ReadRootTransaction(msg2)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	signature, err := newDecoded.Signature()
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	fmt.Println("Decoded signature:")
	util.PrettyPrintHex(signature)

	newDecodedData, err := newDecoded.Data()
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	name, err := newDecodedData.Name()
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	fmt.Println("Decoded name:")
	fmt.Printf("\n\"%s\"\n\n", name)

	// Verify Transaction.signature over the raw serialized data

	fmt.Println("1. Verified signature over the transaction:")
	fmt.Println()
	fmt.Println("NOT SUPPORTED (cannot find field size without recursion)")
	fmt.Println()

	// Hash TransactionSender (eg. place it as a key in state)

	fmt.Println("2. Hash TransactionSender:")
	fmt.Println()
	fmt.Println("NOT SUPPORTED (cannot find field size without recursion)")
	fmt.Println()

	// Avoid decoding unneeded fields

	fmt.Println("3. Avoid decoding TransactionPayload:")
	fmt.Println()
	fmt.Println("CORE FEATURE (no decoding in general)")
	fmt.Println()
}