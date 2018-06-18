package protobuf

import (
	"testing"
	"github.com/orbs-network/go-serialization-comparison/protobuf/nested"
	"time"
	"github.com/golang/protobuf/proto"
	"github.com/orbs-network/go-serialization-comparison/util"
	"fmt"
	"bytes"
	"crypto/md5"
)

func quickMarshal(errOut *error, pb proto.Message) []byte {
	if *errOut != nil {
		return nil
	}
	res, err := proto.Marshal(pb)
	if err != nil && *errOut == nil {
		*errOut = err
		return nil
	}
	return res
}

func hashBytes(buffer []byte) []byte {
	hash := md5.Sum(buffer)
	return hash[:]
}

func TestNested(t *testing.T) {

	fmt.Println("----------------------------------")
	fmt.Println(" TestNested")
	fmt.Println("----------------------------------")

	// create the source transaction

	var err error = nil;

	decoded := &nested.Transaction{
		DataBytes: quickMarshal(&err, &nested.TransactionData{
			Header: &nested.TransactionHeader{
				ProtocolVersion: 0x1,
				VirtualChain: 0x22,
				SenderBytes: quickMarshal(&err, &nested.TransactionSender{
					AddressScheme: nested.AddressScheme_SCHEME01,
					NetworkType: nested.NetworkType_MAIN_NET,
					PublicKey: []byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x18,0x21,0x22,0x23,0x24,0x25,0x26,0x27,0x28,0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38},
				}),
				TimeStamp: time.Now().Unix(),
			},
			PayloadBytes: quickMarshal(&err, &nested.TransactionPayload{
				ContractName: "ZincToken",
				MethodName: "Transfer",
				InputArgument: []*nested.MethodCallArgument{
					{Type: &nested.MethodCallArgument_Uint64{Uint64: 40}},
					{Type: &nested.MethodCallArgument_String_{String_: "arg"}},
				},
			}),
		}),
	}
	if err != nil {
		t.Fatal("marshal error: ", err)
	}

	// prepare the wire

	decoded.Signature = hashBytes(decoded.DataBytes)

	encoded, err := proto.Marshal(decoded)
	if err != nil {
		t.Fatal("marshal error: ", err)
	}

	fmt.Println("Encoded transaction:")
	util.PrettyPrintHex(encoded)

	fmt.Println("Orig transaction signature:")
	util.PrettyPrintHex(decoded.Signature)

	// decode the wire

	newDecoded := &nested.Transaction{}
	err = proto.Unmarshal(encoded, newDecoded)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	fmt.Println("Decoded transaction signature:")
	util.PrettyPrintHex(newDecoded.Signature)

	// Verify Transaction.signature over the raw serialized data

	fmt.Println("1. Verified signature over the transaction:")
	checkedSignature := hashBytes(newDecoded.DataBytes)
	util.PrettyPrintHex(checkedSignature)
	if bytes.Equal(checkedSignature, newDecoded.Signature) {
		fmt.Println("MATCHES!")
		fmt.Println()
	} else {
		t.Fatal("verified signature does not match")
	}

	// Hash TransactionSender (eg. place it as a key in state)

	newTransactionData := &nested.TransactionData{}
	err = proto.Unmarshal(newDecoded.DataBytes, newTransactionData)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}
	fmt.Println("2. Hash TransactionSender:")
	hashSender := hashBytes(newTransactionData.Header.SenderBytes)
	util.PrettyPrintHex(hashSender)

	// Avoid decoding unneeded fields

	newTransactionPayload := &nested.TransactionPayload{}
	err = proto.Unmarshal(newTransactionData.PayloadBytes, newTransactionPayload)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}
	fmt.Println("3. Avoid decoding TransactionPayload:")
	fmt.Printf("\n\"%s.%s\"\n\n", newTransactionPayload.ContractName, newTransactionPayload.MethodName)
}