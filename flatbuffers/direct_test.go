package flatbuffers

import (
	"testing"
	"fmt"
	"github.com/google/flatbuffers/go"
	"github.com/orbs-network/go-serialization-comparison/flatbuffers/direct"
	"github.com/orbs-network/go-serialization-comparison/util"
)

func TestDirect(t *testing.T) {

	fmt.Println("----------------------------------")
	fmt.Println(" TestDirect")
	fmt.Println("----------------------------------")

	// create the source transaction

	builder := flatbuffers.NewBuilder(0)
	signatureOffset := builder.CreateByteVector([]byte{0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22})
	nameOffset := builder.CreateString("johnny")
	direct.TransactionDataStart(builder)
	direct.TransactionDataAddName(builder, nameOffset)
	dataOffset := direct.TransactionDataEnd(builder)
	direct.TransactionStart(builder)
	direct.TransactionAddData(builder, dataOffset)
	direct.TransactionAddSignature(builder, signatureOffset)
	transactionOffset := direct.TransactionEnd(builder)
	builder.Finish(transactionOffset)

	// prepare the wire

	encoded := builder.FinishedBytes()

	fmt.Println("Encoded transaction:")
	util.PrettyPrintHex(encoded)

	// decode the wire

	newDecoded := direct.GetRootAsTransaction(encoded, 0)

	signature := newDecoded.SignatureBytes()

	fmt.Println("Decoded signature:")
	util.PrettyPrintHex(signature)

	fmt.Println("Decoded name:")
	fmt.Printf("\n\"%s\"\n\n", string(newDecoded.Data(nil).Name()))

	// Verify Transaction.signature over the raw serialized data

	fmt.Println("1. Verified signature over the transaction:")
	fmt.Println()
	fmt.Println("NOT SUPPORTED (transaction data is not sequential)")
	fmt.Println()

	// Hash TransactionSender (eg. place it as a key in state)

	fmt.Println("2. Hash TransactionSender:")
	fmt.Println()
	fmt.Println("NOT SUPPORTED (fields are not sequential)")
	fmt.Println()

	// Avoid decoding unneeded fields

	fmt.Println("3. Avoid decoding TransactionPayload:")
	fmt.Println()
	fmt.Println("CORE FEATURE (no decoding in general)")
	fmt.Println()
}