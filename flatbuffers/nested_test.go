package flatbuffers

import (
	"testing"
	"fmt"
	"github.com/google/flatbuffers/go"
	"github.com/orbs-network/go-serialization-comparison/util"
	"github.com/orbs-network/go-serialization-comparison/flatbuffers/nested"
	"crypto/md5"
	"bytes"
)

func hashBytes(buffer []byte) []byte {
	hash := md5.Sum(buffer)
	return hash[:]
}

func TestNested(t *testing.T) {

	fmt.Println("----------------------------------")
	fmt.Println(" TestNested")
	fmt.Println("----------------------------------")

	// create the source transaction

	builder1 := flatbuffers.NewBuilder(0)
	nameOffset := builder1.CreateString("johnny")
	nested.TransactionDataStart(builder1)
	nested.TransactionDataAddName(builder1, nameOffset)
	dataOffset := nested.TransactionDataEnd(builder1)
	builder1.Finish(dataOffset)

	builder2 := flatbuffers.NewBuilder(0)
	signatureOffset := builder2.CreateByteVector(hashBytes(builder1.FinishedBytes()))
	transactionDataOffset := builder2.CreateByteVector(builder1.FinishedBytes())
	nested.TransactionStart(builder2)
	nested.TransactionAddData(builder2, transactionDataOffset)
	nested.TransactionAddSignature(builder2, signatureOffset)
	transactionOffset := nested.TransactionEnd(builder2)
	builder2.Finish(transactionOffset)

	// prepare the wire

	encoded := builder2.FinishedBytes()

	fmt.Println("Encoded transaction:")
	util.PrettyPrintHex(encoded)

	// decode the wire

	newDecoded := nested.GetRootAsTransaction(encoded, 0)

	signature := newDecoded.SignatureBytes()

	fmt.Println("Decoded signature:")
	util.PrettyPrintHex(signature)

	// Verify Transaction.signature over the raw serialized data

	fmt.Println("1. Verified signature over the transaction:")
	checkedSignature := hashBytes(newDecoded.DataBytes())
	util.PrettyPrintHex(checkedSignature)
	if bytes.Equal(checkedSignature, newDecoded.SignatureBytes()) {
		fmt.Println("MATCHES!")
		fmt.Println()
	} else {
		t.Fatal("verified signature does not match")
	}

	// Hash TransactionSender (eg. place it as a key in state)

	fmt.Println("2. Hash TransactionSender:")
	fmt.Println()
	fmt.Println("HARD (every field needs to be its own flatbuffer)")
	fmt.Println()

	// Avoid decoding unneeded fields

	fmt.Println("3. Avoid decoding TransactionPayload:")
	fmt.Printf("\n\"%s\"\n\n", string(nested.GetRootAsTransactionData(newDecoded.DataBytes(), 0).Name()))
}