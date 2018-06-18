package protobuf

import (
	"log"
	"testing"
	"github.com/orbs-network/go-serialization-comparison/protobuf/direct"
	"time"
	"github.com/golang/protobuf/proto"
	"github.com/orbs-network/go-serialization-comparison/util"
	"crypto/md5"
	"fmt"
	"bytes"
)

func hashField(pb proto.Message) []byte {
	encoded, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("marshal error: ", err)
	}
	hash := md5.Sum(encoded)
	return hash[:]
}

func TestDirect(t *testing.T) {

	fmt.Println("----------------------------------")
	fmt.Println(" TestDirect")
	fmt.Println("----------------------------------")

	// source transaction

	decoded := &direct.Transaction{
		Data: &direct.TransactionData{
			Header: &direct.TransactionHeader{
				ProtocolVersion: 0x1,
				VirtualChain: 0x22,
				Sender: &direct.TransactionSender{
					AddressScheme: direct.AddressScheme_SCHEME01,
					NetworkType: direct.NetworkType_MAIN_NET,
					PublicKey: []byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x18,0x21,0x22,0x23,0x24,0x25,0x26,0x27,0x28,0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38},
				},
				TimeStamp: time.Now().Unix(),
			},
			Payload: &direct.TransactionPayload{
				ContractName: "ZincToken",
				MethodName: "Transfer",
				InputArgument: []*direct.MethodCallArgument{
					{Type: &direct.MethodCallArgument_Uint64{Uint64: 40}},
					{Type: &direct.MethodCallArgument_String_{String_: "arg"}},
				},
			},
		},
	}

	// prepare the wire

	decoded.Signature = hashField(decoded.Data)

	encoded, err := proto.Marshal(decoded)
	if err != nil {
		t.Fatal("marshal error: ", err)
	}

	fmt.Println("Encoded transaction:")
	util.PrettyPrintHex(encoded)

	fmt.Println("Orig transaction signature:")
	util.PrettyPrintHex(decoded.Signature)

	// decode the wire

	newDecoded := &direct.Transaction{}
	err = proto.Unmarshal(encoded, newDecoded)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	fmt.Println("Decoded transaction method:")
	fmt.Printf("\n\"%s.%s\"\n\n", newDecoded.Data.Payload.ContractName, newDecoded.Data.Payload.MethodName)

	fmt.Println("Decoded transaction signature:")
	util.PrettyPrintHex(newDecoded.Signature)

	// Verify Transaction.signature over the raw serialized data

	fmt.Println("1. Verified signature over the transaction:")
	checkedSignature := hashField(newDecoded.Data)
	util.PrettyPrintHex(checkedSignature)
	if bytes.Equal(checkedSignature, newDecoded.Signature) {
		fmt.Println("MATCHES!")
		fmt.Println()
	} else {
		t.Fatal("verified signature does not match")
	}

	// Hash TransactionSender (eg. place it as a key in state)

	fmt.Println("2. Hash TransactionSender:")
	hashSender := hashField(newDecoded.Data.Header.Sender)
	util.PrettyPrintHex(hashSender)

	// Avoid decoding unneeded fields

	fmt.Println("3. Avoid decoding TransactionPayload:")
	fmt.Println()
	fmt.Println("NOT SUPPORTED")
	fmt.Println()
}