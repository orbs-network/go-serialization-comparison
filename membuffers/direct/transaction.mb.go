// AUTO GENERATED FILE (by membufc proto compiler)
package direct

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Transaction

// reader

type Transaction struct {
	message membuffers.Message
}

var m_Transaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_Transaction_Unions = [][]membuffers.FieldType{}

func TransactionReader(buf []byte) *Transaction {
	x := &Transaction{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Transaction_Scheme, m_Transaction_Unions)
	return x
}

func (x *Transaction) IsValid() bool {
	return x.message.IsValid()
}

func (x *Transaction) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Transaction) Data() *TransactionData {
	b, s := x.message.GetMessage(0)
	return TransactionDataReader(b[:s])
}

func (x *Transaction) RawData() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Transaction) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *Transaction) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *Transaction) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type TransactionBuilder struct {
	builder membuffers.Builder
	Data *TransactionDataBuilder
	Signature []byte
}

func (w *TransactionBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Data)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *TransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionData

// reader

type TransactionData struct {
	message membuffers.Message
}

var m_TransactionData_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,}
var m_TransactionData_Unions = [][]membuffers.FieldType{}

func TransactionDataReader(buf []byte) *TransactionData {
	x := &TransactionData{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionData_Scheme, m_TransactionData_Unions)
	return x
}

func (x *TransactionData) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionData) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionData) Header() *TransactionHeader {
	b, s := x.message.GetMessage(0)
	return TransactionHeaderReader(b[:s])
}

func (x *TransactionData) RawHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionData) Payload() *TransactionPayload {
	b, s := x.message.GetMessage(1)
	return TransactionPayloadReader(b[:s])
}

func (x *TransactionData) RawPayload() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionDataBuilder struct {
	builder membuffers.Builder
	Header *TransactionHeaderBuilder
	Payload *TransactionPayloadBuilder
}

func (w *TransactionDataBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Header)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.Payload)
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionHeader

// reader

type TransactionHeader struct {
	message membuffers.Message
}

var m_TransactionHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessage,membuffers.TypeUint64,}
var m_TransactionHeader_Unions = [][]membuffers.FieldType{}

func TransactionHeaderReader(buf []byte) *TransactionHeader {
	x := &TransactionHeader{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionHeader_Scheme, m_TransactionHeader_Unions)
	return x
}

func (x *TransactionHeader) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionHeader) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionHeader) ProtocolVersion() uint32 {
	return x.message.GetUint32(0)
}

func (x *TransactionHeader) RawProtocolVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionHeader) MutateProtocolVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *TransactionHeader) VirtualChain() uint64 {
	return x.message.GetUint64(1)
}

func (x *TransactionHeader) RawVirtualChain() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionHeader) MutateVirtualChain(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *TransactionHeader) Sender() *TransactionSender {
	b, s := x.message.GetMessage(2)
	return TransactionSenderReader(b[:s])
}

func (x *TransactionHeader) RawSender() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionHeader) TimeStamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *TransactionHeader) RawTimeStamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *TransactionHeader) MutateTimeStamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type TransactionHeaderBuilder struct {
	builder membuffers.Builder
	ProtocolVersion uint32
	VirtualChain uint64
	Sender *TransactionSenderBuilder
	TimeStamp uint64
}

func (w *TransactionHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ProtocolVersion)
	w.builder.WriteUint64(buf, w.VirtualChain)
	err = w.builder.WriteMessage(buf, w.Sender)
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.TimeStamp)
	return nil
}

func (w *TransactionHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSender

// reader

type TransactionSender struct {
	message membuffers.Message
}

var m_TransactionSender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeBytes,}
var m_TransactionSender_Unions = [][]membuffers.FieldType{}

func TransactionSenderReader(buf []byte) *TransactionSender {
	x := &TransactionSender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionSender_Scheme, m_TransactionSender_Unions)
	return x
}

func (x *TransactionSender) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionSender) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionSender) AddressScheme() AddressScheme {
	return AddressScheme(x.message.GetUint16(0))
}

func (x *TransactionSender) RawAddressScheme() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionSender) MutateAddressScheme(v AddressScheme) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *TransactionSender) NetworkType() NetworkType {
	return NetworkType(x.message.GetUint16(1))
}

func (x *TransactionSender) RawNetworkType() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionSender) MutateNetworkType(v NetworkType) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *TransactionSender) PublicKey() []byte {
	return x.message.GetBytes(2)
}

func (x *TransactionSender) RawPublicKey() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionSender) MutatePublicKey(v []byte) error {
	return x.message.SetBytes(2, v)
}

// builder

type TransactionSenderBuilder struct {
	builder membuffers.Builder
	AddressScheme AddressScheme
	NetworkType NetworkType
	PublicKey []byte
}

func (w *TransactionSenderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.AddressScheme))
	w.builder.WriteUint16(buf, uint16(w.NetworkType))
	w.builder.WriteBytes(buf, w.PublicKey)
	return nil
}

func (w *TransactionSenderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionSenderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionPayload

// reader

type TransactionPayload struct {
	message membuffers.Message
}

var m_TransactionPayload_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeString,membuffers.TypeMessageArray,}
var m_TransactionPayload_Unions = [][]membuffers.FieldType{}

func TransactionPayloadReader(buf []byte) *TransactionPayload {
	x := &TransactionPayload{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionPayload_Scheme, m_TransactionPayload_Unions)
	return x
}

func (x *TransactionPayload) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionPayload) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionPayload) ContractName() string {
	return x.message.GetString(0)
}

func (x *TransactionPayload) RawContractName() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionPayload) MutateContractName(v string) error {
	return x.message.SetString(0, v)
}

func (x *TransactionPayload) MethodName() string {
	return x.message.GetString(1)
}

func (x *TransactionPayload) RawMethodName() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionPayload) MutateMethodName(v string) error {
	return x.message.SetString(1, v)
}

func (x *TransactionPayload) InputArgumentIterator() *TransactionPayloadInputArgumentIterator {
	return &TransactionPayloadInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(2)}
}

type TransactionPayloadInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionPayloadInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionPayloadInputArgumentIterator) NextInputArgument() *MethodCallArgument {
	b, s := i.iterator.NextMessage()
	return MethodCallArgumentReader(b[:s])
}

func (x *TransactionPayload) RawInputArgumentArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type TransactionPayloadBuilder struct {
	builder membuffers.Builder
	ContractName string
	MethodName string
	InputArgument []*MethodCallArgumentBuilder
}

func (w *TransactionPayloadBuilder) arrayOfInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.InputArgument))
	for i, v := range w.InputArgument {
		res[i] = v
	}
	return res
}

func (w *TransactionPayloadBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.ContractName)
	w.builder.WriteString(buf, w.MethodName)
	err = w.builder.WriteMessageArray(buf, w.arrayOfInputArgument())
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionPayloadBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionPayloadBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message MethodCallArgument

// reader

type MethodCallArgument struct {
	message membuffers.Message
}

var m_MethodCallArgument_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var m_MethodCallArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeString,membuffers.TypeBytes,}}

func MethodCallArgumentReader(buf []byte) *MethodCallArgument {
	x := &MethodCallArgument{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_MethodCallArgument_Scheme, m_MethodCallArgument_Unions)
	return x
}

func (x *MethodCallArgument) IsValid() bool {
	return x.message.IsValid()
}

func (x *MethodCallArgument) Raw() []byte {
	return x.message.RawBuffer()
}

type MethodCallArgumentType uint16

const (
	MethodCallArgumentTypeUint32 MethodCallArgumentType = 0
	MethodCallArgumentTypeUint64 MethodCallArgumentType = 1
	MethodCallArgumentTypeString MethodCallArgumentType = 2
	MethodCallArgumentTypeBytes MethodCallArgumentType = 3
)

func (x *MethodCallArgument) Type() MethodCallArgumentType {
	return MethodCallArgumentType(x.message.GetUint16(0))
}

func (x *MethodCallArgument) IsTypeUint32() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *MethodCallArgument) TypeUint32() uint32 {
	_, off := x.message.IsUnionIndex(0, 0, 0)
	return x.message.GetUint32InOffset(off)
}

func (x *MethodCallArgument) MutateTypeUint32(v uint32) error {
	is, off := x.message.IsUnionIndex(0, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint32InOffset(off, v)
	return nil
}

func (x *MethodCallArgument) IsTypeUint64() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *MethodCallArgument) TypeUint64() uint64 {
	_, off := x.message.IsUnionIndex(0, 0, 1)
	return x.message.GetUint64InOffset(off)
}

func (x *MethodCallArgument) MutateTypeUint64(v uint64) error {
	is, off := x.message.IsUnionIndex(0, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint64InOffset(off, v)
	return nil
}

func (x *MethodCallArgument) IsTypeString() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *MethodCallArgument) TypeString() string {
	_, off := x.message.IsUnionIndex(0, 0, 2)
	return x.message.GetStringInOffset(off)
}

func (x *MethodCallArgument) MutateTypeString(v string) error {
	is, off := x.message.IsUnionIndex(0, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetStringInOffset(off, v)
	return nil
}

func (x *MethodCallArgument) IsTypeBytes() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 3)
	return is
}

func (x *MethodCallArgument) TypeBytes() []byte {
	_, off := x.message.IsUnionIndex(0, 0, 3)
	return x.message.GetBytesInOffset(off)
}

func (x *MethodCallArgument) MutateTypeBytes(v []byte) error {
	is, off := x.message.IsUnionIndex(0, 0, 3)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetBytesInOffset(off, v)
	return nil
}

func (x *MethodCallArgument) RawType() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type MethodCallArgumentBuilder struct {
	builder membuffers.Builder
	Type MethodCallArgumentType
	Uint32 uint32
	Uint64 uint64
	String string
	Bytes []byte
}

func (w *MethodCallArgumentBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case MethodCallArgumentTypeUint32:
		w.builder.WriteUint32(buf, w.Uint32)
	case MethodCallArgumentTypeUint64:
		w.builder.WriteUint64(buf, w.Uint64)
	case MethodCallArgumentTypeString:
		w.builder.WriteString(buf, w.String)
	case MethodCallArgumentTypeBytes:
		w.builder.WriteBytes(buf, w.Bytes)
	}
	return nil
}

func (w *MethodCallArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MethodCallArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// enums

type AddressScheme uint16

const (
	AddressScheme_RESERVED_SCHEME AddressScheme = 0
	AddressScheme_SCHEME01 AddressScheme = 1
	AddressScheme_FUTURE_SCHEME AddressScheme = 255
)

type NetworkType uint16

const (
	NetworkType_RESERVED_NETWORK NetworkType = 0
	NetworkType_MAIN_NET NetworkType = 77
	NetworkType_TEST_NET NetworkType = 84
	NetworkType_FUTURE_NETWORK NetworkType = 255
)

