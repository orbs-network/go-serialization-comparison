package membuffers

import (
	"crypto/md5"
	"testing"
	"fmt"
	"time"
	"github.com/orbs-network/go-serialization-comparison/membuffers/direct"
	"github.com/orbs-network/go-serialization-comparison/util"
	"bytes"
)

func hashBytes(buffer []byte) []byte {
	hash := md5.Sum(buffer)
	return hash[:]
}

func TestDirect(t *testing.T) {

	fmt.Println("----------------------------------")
	fmt.Println(" TestDirect")
	fmt.Println("----------------------------------")

	// create the source transaction

	builder := &direct.TransactionBuilder{
		Data: &direct.TransactionDataBuilder{
			Header: &direct.TransactionHeaderBuilder{
				ProtocolVersion: 0x1,
				VirtualChain: 0x22,
				Sender: &direct.TransactionSenderBuilder{
					AddressScheme: direct.AddressScheme_SCHEME01,
					NetworkType: direct.NetworkType_MAIN_NET,
					PublicKey: []byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x18,0x21,0x22,0x23,0x24,0x25,0x26,0x27,0x28,0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38},
				},
				TimeStamp: uint64(time.Now().Unix()),
			},
			Payload: &direct.TransactionPayloadBuilder{
				ContractName: "ZincToken",
				MethodName: "Transfer",
				InputArgument: []*direct.MethodCallArgumentBuilder{
					{Type: direct.MethodCallArgumentTypeUint64, Uint64: 40},
					{Type: direct.MethodCallArgumentTypeString, String: "arg"},
				},
			},
		},
		Signature: []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,},
	}

	// prepare the wire

	encoded := make([]byte, builder.CalcRequiredSize())
	err := builder.Write(encoded)
	if err != nil {
		t.Fatal("marshal error: ", err)
	}

	fmt.Println("Encoded transaction:")
	util.PrettyPrintHex(encoded)

	// decode the wire

	newDecoded := direct.TransactionReader(encoded)
	newDecoded.MutateSignature(hashBytes(newDecoded.RawData()))

	fmt.Println("Encoded transaction with signature:")
	util.PrettyPrintHex(encoded)

	fmt.Println("Decoded transaction signature:")
	util.PrettyPrintHex(newDecoded.Signature())

	// Verify Transaction.signature over the raw serialized data

	fmt.Println("1. Verified signature over the transaction:")
	checkedSignature := hashBytes(newDecoded.RawData())
	util.PrettyPrintHex(checkedSignature)
	if bytes.Equal(checkedSignature, newDecoded.Signature()) {
		fmt.Println("MATCHES!")
		fmt.Println()
	} else {
		t.Fatal("verified signature does not match")
	}

	// Hash TransactionSender (eg. place it as a key in state)

	fmt.Println("2. Hash TransactionSender:")
	hashSender := hashBytes(newDecoded.Data().Header().RawSender())
	util.PrettyPrintHex(hashSender)

	// Avoid decoding unneeded fields

	fmt.Println("3. Avoid decoding TransactionPayload:")
	fmt.Printf("\n\"%s.%s\"\n\n", newDecoded.Data().Payload().ContractName(), newDecoded.Data().Payload().MethodName())
}