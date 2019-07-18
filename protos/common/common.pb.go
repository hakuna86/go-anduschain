// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAIL    Status = 1
)

var Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
}

var Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAIL":    1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type ProcessStatus int32

const (
	ProcessStatus_WAIT          ProcessStatus = 0
	ProcessStatus_MAKE_JOIN_TX  ProcessStatus = 1
	ProcessStatus_MAKE_BLOCK    ProcessStatus = 2
	ProcessStatus_VOTE_START    ProcessStatus = 3
	ProcessStatus_VOTE_COMPLETE ProcessStatus = 4
	ProcessStatus_REJECT        ProcessStatus = 5
)

var ProcessStatus_name = map[int32]string{
	0: "WAIT",
	1: "MAKE_JOIN_TX",
	2: "MAKE_BLOCK",
	3: "VOTE_START",
	4: "VOTE_COMPLETE",
	5: "REJECT",
}

var ProcessStatus_value = map[string]int32{
	"WAIT":          0,
	"MAKE_JOIN_TX":  1,
	"MAKE_BLOCK":    2,
	"VOTE_START":    3,
	"VOTE_COMPLETE": 4,
	"REJECT":        5,
}

func (x ProcessStatus) String() string {
	return proto.EnumName(ProcessStatus_name, int32(x))
}

func (ProcessStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type SealConfirmStatus int32

const (
	SealConfirmStatus_PENDING    SealConfirmStatus = 0
	SealConfirmStatus_SEND_BLOCK SealConfirmStatus = 1
	SealConfirmStatus_SEAL_FAIL  SealConfirmStatus = 2
	SealConfirmStatus_COMPLETE   SealConfirmStatus = 3
)

var SealConfirmStatus_name = map[int32]string{
	0: "PENDING",
	1: "SEND_BLOCK",
	2: "SEAL_FAIL",
	3: "COMPLETE",
}

var SealConfirmStatus_value = map[string]int32{
	"PENDING":    0,
	"SEND_BLOCK": 1,
	"SEAL_FAIL":  2,
	"COMPLETE":   3,
}

func (x SealConfirmStatus) String() string {
	return proto.EnumName(SealConfirmStatus_name, int32(x))
}

func (SealConfirmStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type HeartBeat struct {
	Enode                string   `protobuf:"bytes,1,opt,name=enode,proto3" json:"enode,omitempty"`
	MinerAddress         string   `protobuf:"bytes,2,opt,name=minerAddress,proto3" json:"minerAddress,omitempty"`
	ChainID              string   `protobuf:"bytes,3,opt,name=chainID,proto3" json:"chainID,omitempty"`
	NodeVersion          string   `protobuf:"bytes,4,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	Head                 string   `protobuf:"bytes,5,opt,name=head,proto3" json:"head,omitempty"`
	Sign                 []byte   `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeat.Unmarshal(m, b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return xxx_messageInfo_HeartBeat.Size(m)
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func (m *HeartBeat) GetEnode() string {
	if m != nil {
		return m.Enode
	}
	return ""
}

func (m *HeartBeat) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *HeartBeat) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *HeartBeat) GetNodeVersion() string {
	if m != nil {
		return m.NodeVersion
	}
	return ""
}

func (m *HeartBeat) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

func (m *HeartBeat) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type Participate struct {
	OtprnHash            []byte   `protobuf:"bytes,1,opt,name=otprnHash,proto3" json:"otprnHash,omitempty"`
	Enode                string   `protobuf:"bytes,2,opt,name=enode,proto3" json:"enode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Participate) Reset()         { *m = Participate{} }
func (m *Participate) String() string { return proto.CompactTextString(m) }
func (*Participate) ProtoMessage()    {}
func (*Participate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Participate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Participate.Unmarshal(m, b)
}
func (m *Participate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Participate.Marshal(b, m, deterministic)
}
func (m *Participate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Participate.Merge(m, src)
}
func (m *Participate) XXX_Size() int {
	return xxx_messageInfo_Participate.Size(m)
}
func (m *Participate) XXX_DiscardUnknown() {
	xxx_messageInfo_Participate.DiscardUnknown(m)
}

var xxx_messageInfo_Participate proto.InternalMessageInfo

func (m *Participate) GetOtprnHash() []byte {
	if m != nil {
		return m.OtprnHash
	}
	return nil
}

func (m *Participate) GetEnode() string {
	if m != nil {
		return m.Enode
	}
	return ""
}

type ProcessMessage struct {
	Code                 ProcessStatus `protobuf:"varint,1,opt,name=code,proto3,enum=common.ProcessStatus" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProcessMessage) Reset()         { *m = ProcessMessage{} }
func (m *ProcessMessage) String() string { return proto.CompactTextString(m) }
func (*ProcessMessage) ProtoMessage()    {}
func (*ProcessMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *ProcessMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessMessage.Unmarshal(m, b)
}
func (m *ProcessMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessMessage.Marshal(b, m, deterministic)
}
func (m *ProcessMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessMessage.Merge(m, src)
}
func (m *ProcessMessage) XXX_Size() int {
	return xxx_messageInfo_ProcessMessage.Size(m)
}
func (m *ProcessMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessMessage proto.InternalMessageInfo

func (m *ProcessMessage) GetCode() ProcessStatus {
	if m != nil {
		return m.Code
	}
	return ProcessStatus_WAIT
}

// OTRPN 요청
type ReqOtprn struct {
	Enode                string   `protobuf:"bytes,1,opt,name=enode,proto3" json:"enode,omitempty"`
	MinerAddress         string   `protobuf:"bytes,2,opt,name=minerAddress,proto3" json:"minerAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqOtprn) Reset()         { *m = ReqOtprn{} }
func (m *ReqOtprn) String() string { return proto.CompactTextString(m) }
func (*ReqOtprn) ProtoMessage()    {}
func (*ReqOtprn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *ReqOtprn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqOtprn.Unmarshal(m, b)
}
func (m *ReqOtprn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqOtprn.Marshal(b, m, deterministic)
}
func (m *ReqOtprn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqOtprn.Merge(m, src)
}
func (m *ReqOtprn) XXX_Size() int {
	return xxx_messageInfo_ReqOtprn.Size(m)
}
func (m *ReqOtprn) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqOtprn.DiscardUnknown(m)
}

var xxx_messageInfo_ReqOtprn proto.InternalMessageInfo

func (m *ReqOtprn) GetEnode() string {
	if m != nil {
		return m.Enode
	}
	return ""
}

func (m *ReqOtprn) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

type ResOtprn struct {
	Result               Status   `protobuf:"varint,1,opt,name=result,proto3,enum=common.Status" json:"result,omitempty"`
	Otprn                []byte   `protobuf:"bytes,2,opt,name=otprn,proto3" json:"otprn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResOtprn) Reset()         { *m = ResOtprn{} }
func (m *ResOtprn) String() string { return proto.CompactTextString(m) }
func (*ResOtprn) ProtoMessage()    {}
func (*ResOtprn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *ResOtprn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResOtprn.Unmarshal(m, b)
}
func (m *ResOtprn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResOtprn.Marshal(b, m, deterministic)
}
func (m *ResOtprn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResOtprn.Merge(m, src)
}
func (m *ResOtprn) XXX_Size() int {
	return xxx_messageInfo_ResOtprn.Size(m)
}
func (m *ResOtprn) XXX_DiscardUnknown() {
	xxx_messageInfo_ResOtprn.DiscardUnknown(m)
}

var xxx_messageInfo_ResOtprn proto.InternalMessageInfo

func (m *ResOtprn) GetResult() Status {
	if m != nil {
		return m.Result
	}
	return Status_SUCCESS
}

func (m *ResOtprn) GetOtprn() []byte {
	if m != nil {
		return m.Otprn
	}
	return nil
}

// 투표
type Vote struct {
	Header               []byte   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	VoterAddress         string   `protobuf:"bytes,2,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	VoterSign            []byte   `protobuf:"bytes,3,opt,name=voterSign,proto3" json:"voterSign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{5}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Vote) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *Vote) GetVoterSign() []byte {
	if m != nil {
		return m.VoterSign
	}
	return nil
}

// 투표 결과 요청
type ReqVoteResult struct {
	OtprnHash            []byte   `protobuf:"bytes,1,opt,name=otprnHash,proto3" json:"otprnHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVoteResult) Reset()         { *m = ReqVoteResult{} }
func (m *ReqVoteResult) String() string { return proto.CompactTextString(m) }
func (*ReqVoteResult) ProtoMessage()    {}
func (*ReqVoteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{6}
}

func (m *ReqVoteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqVoteResult.Unmarshal(m, b)
}
func (m *ReqVoteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqVoteResult.Marshal(b, m, deterministic)
}
func (m *ReqVoteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqVoteResult.Merge(m, src)
}
func (m *ReqVoteResult) XXX_Size() int {
	return xxx_messageInfo_ReqVoteResult.Size(m)
}
func (m *ReqVoteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqVoteResult.DiscardUnknown(m)
}

var xxx_messageInfo_ReqVoteResult proto.InternalMessageInfo

func (m *ReqVoteResult) GetOtprnHash() []byte {
	if m != nil {
		return m.OtprnHash
	}
	return nil
}

type ResVoteResult struct {
	Result               Status   `protobuf:"varint,1,opt,name=result,proto3,enum=common.Status" json:"result,omitempty"`
	BlockHash            string   `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Voters               []byte   `protobuf:"bytes,3,opt,name=voters,proto3" json:"voters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResVoteResult) Reset()         { *m = ResVoteResult{} }
func (m *ResVoteResult) String() string { return proto.CompactTextString(m) }
func (*ResVoteResult) ProtoMessage()    {}
func (*ResVoteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{7}
}

func (m *ResVoteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResVoteResult.Unmarshal(m, b)
}
func (m *ResVoteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResVoteResult.Marshal(b, m, deterministic)
}
func (m *ResVoteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResVoteResult.Merge(m, src)
}
func (m *ResVoteResult) XXX_Size() int {
	return xxx_messageInfo_ResVoteResult.Size(m)
}
func (m *ResVoteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ResVoteResult.DiscardUnknown(m)
}

var xxx_messageInfo_ResVoteResult proto.InternalMessageInfo

func (m *ResVoteResult) GetResult() Status {
	if m != nil {
		return m.Result
	}
	return Status_SUCCESS
}

func (m *ResVoteResult) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *ResVoteResult) GetVoters() []byte {
	if m != nil {
		return m.Voters
	}
	return nil
}

// 블록 실링 확인 요청
type ReqConfirmSeal struct {
	Enode                string   `protobuf:"bytes,1,opt,name=enode,proto3" json:"enode,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	VoteHash             []byte   `protobuf:"bytes,3,opt,name=voteHash,proto3" json:"voteHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqConfirmSeal) Reset()         { *m = ReqConfirmSeal{} }
func (m *ReqConfirmSeal) String() string { return proto.CompactTextString(m) }
func (*ReqConfirmSeal) ProtoMessage()    {}
func (*ReqConfirmSeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{8}
}

func (m *ReqConfirmSeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqConfirmSeal.Unmarshal(m, b)
}
func (m *ReqConfirmSeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqConfirmSeal.Marshal(b, m, deterministic)
}
func (m *ReqConfirmSeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqConfirmSeal.Merge(m, src)
}
func (m *ReqConfirmSeal) XXX_Size() int {
	return xxx_messageInfo_ReqConfirmSeal.Size(m)
}
func (m *ReqConfirmSeal) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqConfirmSeal.DiscardUnknown(m)
}

var xxx_messageInfo_ReqConfirmSeal proto.InternalMessageInfo

func (m *ReqConfirmSeal) GetEnode() string {
	if m != nil {
		return m.Enode
	}
	return ""
}

func (m *ReqConfirmSeal) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ReqConfirmSeal) GetVoteHash() []byte {
	if m != nil {
		return m.VoteHash
	}
	return nil
}

type ResConfirmSeal struct {
	Code                 SealConfirmStatus `protobuf:"varint,1,opt,name=code,proto3,enum=common.SealConfirmStatus" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResConfirmSeal) Reset()         { *m = ResConfirmSeal{} }
func (m *ResConfirmSeal) String() string { return proto.CompactTextString(m) }
func (*ResConfirmSeal) ProtoMessage()    {}
func (*ResConfirmSeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{9}
}

func (m *ResConfirmSeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResConfirmSeal.Unmarshal(m, b)
}
func (m *ResConfirmSeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResConfirmSeal.Marshal(b, m, deterministic)
}
func (m *ResConfirmSeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResConfirmSeal.Merge(m, src)
}
func (m *ResConfirmSeal) XXX_Size() int {
	return xxx_messageInfo_ResConfirmSeal.Size(m)
}
func (m *ResConfirmSeal) XXX_DiscardUnknown() {
	xxx_messageInfo_ResConfirmSeal.DiscardUnknown(m)
}

var xxx_messageInfo_ResConfirmSeal proto.InternalMessageInfo

func (m *ResConfirmSeal) GetCode() SealConfirmStatus {
	if m != nil {
		return m.Code
	}
	return SealConfirmStatus_PENDING
}

// 블록 요청
type ReqBlock struct {
	Enode                string   `protobuf:"bytes,1,opt,name=enode,proto3" json:"enode,omitempty"`
	Block                []byte   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBlock) Reset()         { *m = ReqBlock{} }
func (m *ReqBlock) String() string { return proto.CompactTextString(m) }
func (*ReqBlock) ProtoMessage()    {}
func (*ReqBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{10}
}

func (m *ReqBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBlock.Unmarshal(m, b)
}
func (m *ReqBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBlock.Marshal(b, m, deterministic)
}
func (m *ReqBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBlock.Merge(m, src)
}
func (m *ReqBlock) XXX_Size() int {
	return xxx_messageInfo_ReqBlock.Size(m)
}
func (m *ReqBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBlock proto.InternalMessageInfo

func (m *ReqBlock) GetEnode() string {
	if m != nil {
		return m.Enode
	}
	return ""
}

func (m *ReqBlock) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

// 블록 서명 요청
type ReqFairnodeSign struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	VoteHash             []byte   `protobuf:"bytes,2,opt,name=voteHash,proto3" json:"voteHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFairnodeSign) Reset()         { *m = ReqFairnodeSign{} }
func (m *ReqFairnodeSign) String() string { return proto.CompactTextString(m) }
func (*ReqFairnodeSign) ProtoMessage()    {}
func (*ReqFairnodeSign) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{11}
}

func (m *ReqFairnodeSign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFairnodeSign.Unmarshal(m, b)
}
func (m *ReqFairnodeSign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFairnodeSign.Marshal(b, m, deterministic)
}
func (m *ReqFairnodeSign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFairnodeSign.Merge(m, src)
}
func (m *ReqFairnodeSign) XXX_Size() int {
	return xxx_messageInfo_ReqFairnodeSign.Size(m)
}
func (m *ReqFairnodeSign) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFairnodeSign.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFairnodeSign proto.InternalMessageInfo

func (m *ReqFairnodeSign) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ReqFairnodeSign) GetVoteHash() []byte {
	if m != nil {
		return m.VoteHash
	}
	return nil
}

type ResFairnodeSign struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResFairnodeSign) Reset()         { *m = ResFairnodeSign{} }
func (m *ResFairnodeSign) String() string { return proto.CompactTextString(m) }
func (*ResFairnodeSign) ProtoMessage()    {}
func (*ResFairnodeSign) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{12}
}

func (m *ResFairnodeSign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResFairnodeSign.Unmarshal(m, b)
}
func (m *ResFairnodeSign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResFairnodeSign.Marshal(b, m, deterministic)
}
func (m *ResFairnodeSign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResFairnodeSign.Merge(m, src)
}
func (m *ResFairnodeSign) XXX_Size() int {
	return xxx_messageInfo_ResFairnodeSign.Size(m)
}
func (m *ResFairnodeSign) XXX_DiscardUnknown() {
	xxx_messageInfo_ResFairnodeSign.DiscardUnknown(m)
}

var xxx_messageInfo_ResFairnodeSign proto.InternalMessageInfo

func (m *ResFairnodeSign) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.ProcessStatus", ProcessStatus_name, ProcessStatus_value)
	proto.RegisterEnum("common.SealConfirmStatus", SealConfirmStatus_name, SealConfirmStatus_value)
	proto.RegisterType((*HeartBeat)(nil), "common.HeartBeat")
	proto.RegisterType((*Participate)(nil), "common.Participate")
	proto.RegisterType((*ProcessMessage)(nil), "common.ProcessMessage")
	proto.RegisterType((*ReqOtprn)(nil), "common.ReqOtprn")
	proto.RegisterType((*ResOtprn)(nil), "common.ResOtprn")
	proto.RegisterType((*Vote)(nil), "common.Vote")
	proto.RegisterType((*ReqVoteResult)(nil), "common.ReqVoteResult")
	proto.RegisterType((*ResVoteResult)(nil), "common.ResVoteResult")
	proto.RegisterType((*ReqConfirmSeal)(nil), "common.ReqConfirmSeal")
	proto.RegisterType((*ResConfirmSeal)(nil), "common.ResConfirmSeal")
	proto.RegisterType((*ReqBlock)(nil), "common.ReqBlock")
	proto.RegisterType((*ReqFairnodeSign)(nil), "common.ReqFairnodeSign")
	proto.RegisterType((*ResFairnodeSign)(nil), "common.ResFairnodeSign")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0x5e, 0xfa, 0x91, 0xb5, 0x67, 0x69, 0xdf, 0xcc, 0xef, 0x40, 0x01, 0x4d, 0x62, 0xca, 0x05,
	0x1a, 0x93, 0xb6, 0x0a, 0x26, 0x71, 0xc3, 0x05, 0x4a, 0xdb, 0x8c, 0x75, 0xeb, 0x97, 0x9c, 0x30,
	0x10, 0x37, 0x9d, 0xd7, 0x9a, 0x36, 0xd0, 0xc6, 0xab, 0xed, 0xf2, 0x9f, 0xf8, 0x97, 0xc8, 0x4e,
	0xd2, 0xa5, 0x9b, 0x36, 0x21, 0xae, 0xda, 0xe7, 0x9c, 0xa3, 0xe7, 0xe3, 0xd8, 0x31, 0xfc, 0x3f,
	0x66, 0x8b, 0x05, 0x8b, 0x1b, 0xc9, 0xcf, 0xc9, 0x2d, 0x67, 0x92, 0x21, 0x33, 0x41, 0xee, 0x6f,
	0x03, 0xaa, 0xe7, 0x94, 0x70, 0xd9, 0xa4, 0x44, 0xa2, 0x3d, 0x28, 0xd3, 0x98, 0x4d, 0xa8, 0x63,
	0x1c, 0x18, 0x87, 0x55, 0x9c, 0x00, 0xe4, 0x82, 0xb5, 0x88, 0x62, 0xca, 0xbd, 0xc9, 0x84, 0x53,
	0x21, 0x9c, 0x82, 0x6e, 0x6e, 0xd4, 0x90, 0x03, 0xdb, 0xe3, 0x19, 0x89, 0xe2, 0x4e, 0xdb, 0x29,
	0xea, 0x76, 0x06, 0xd1, 0x01, 0xec, 0x28, 0x96, 0x2b, 0xca, 0x45, 0xc4, 0x62, 0xa7, 0xa4, 0xbb,
	0xf9, 0x12, 0x42, 0x50, 0x9a, 0x51, 0x32, 0x71, 0xca, 0xba, 0xa5, 0xff, 0xab, 0x9a, 0x88, 0xa6,
	0xb1, 0x63, 0x1e, 0x18, 0x87, 0x16, 0xd6, 0xff, 0x5d, 0x0f, 0x76, 0x86, 0x84, 0xcb, 0x68, 0x1c,
	0xdd, 0x12, 0x49, 0xd1, 0x3e, 0x54, 0x99, 0xbc, 0xe5, 0xf1, 0x39, 0x11, 0x33, 0x6d, 0xd8, 0xc2,
	0x77, 0x85, 0xbb, 0x28, 0x85, 0x5c, 0x14, 0xf7, 0x03, 0xd4, 0x87, 0x9c, 0x8d, 0xa9, 0x10, 0x3d,
	0x2a, 0x04, 0x99, 0x52, 0xf4, 0x06, 0x4a, 0xe3, 0x2c, 0x71, 0xfd, 0xdd, 0xb3, 0x93, 0x74, 0x4b,
	0xe9, 0x54, 0x20, 0x89, 0x5c, 0x09, 0xac, 0x47, 0xdc, 0x36, 0x54, 0x30, 0x5d, 0x0e, 0x94, 0xc4,
	0xbf, 0x6f, 0xca, 0x3d, 0x57, 0x2c, 0x22, 0x61, 0x79, 0x0d, 0x26, 0xa7, 0x62, 0x35, 0x97, 0xa9,
	0x7c, 0x3d, 0x93, 0x4f, 0x75, 0xd3, 0xae, 0x52, 0xd3, 0xc9, 0x34, 0xa1, 0x85, 0x13, 0xe0, 0x5e,
	0x43, 0xe9, 0x8a, 0x49, 0x8a, 0x9e, 0x83, 0xa9, 0x76, 0x46, 0x79, 0xba, 0x85, 0x14, 0x29, 0x37,
	0xbf, 0x98, 0x7c, 0xe0, 0x26, 0x5f, 0x53, 0x4b, 0xd4, 0x38, 0x50, 0xcb, 0x2e, 0x26, 0x4b, 0x5c,
	0x17, 0xdc, 0x63, 0xa8, 0x61, 0xba, 0x54, 0x22, 0x38, 0x31, 0xf2, 0xe4, 0xce, 0xdd, 0x85, 0x1a,
	0x17, 0xb9, 0xf1, 0xbf, 0xcd, 0xb7, 0x0f, 0xd5, 0x9b, 0x39, 0x1b, 0xff, 0xd4, 0xb4, 0x89, 0xcd,
	0xbb, 0x82, 0xca, 0xa7, 0x2d, 0x89, 0xd4, 0x60, 0x8a, 0xdc, 0x6b, 0xa8, 0x63, 0xba, 0x6c, 0xb1,
	0xf8, 0x7b, 0xc4, 0x17, 0x01, 0x25, 0xf3, 0x47, 0x4e, 0xe5, 0x01, 0xbb, 0x95, 0x67, 0x7f, 0x09,
	0x15, 0xc5, 0xa7, 0x9b, 0x09, 0xff, 0x1a, 0xbb, 0x1f, 0x95, 0x82, 0xc8, 0x2b, 0x1c, 0x6f, 0x5c,
	0x97, 0x17, 0xeb, 0x3c, 0x94, 0xcc, 0xb3, 0xb1, 0xfc, 0x95, 0x79, 0xaf, 0xaf, 0x4c, 0x53, 0x89,
	0x3d, 0x62, 0x6e, 0x0f, 0xca, 0xda, 0x4b, 0x76, 0xb4, 0x1a, 0xb8, 0x97, 0xf0, 0x1f, 0xa6, 0xcb,
	0x33, 0x12, 0x71, 0x35, 0xa4, 0xce, 0x62, 0x33, 0x85, 0xf1, 0x54, 0x8a, 0xc2, 0xbd, 0x14, 0x0d,
	0x45, 0x26, 0xee, 0x93, 0xa9, 0x4f, 0x8a, 0xc8, 0x15, 0xa7, 0x19, 0xd9, 0xba, 0x70, 0xf4, 0x0a,
	0xcc, 0x24, 0x05, 0xda, 0x81, 0xed, 0xe0, 0x73, 0xab, 0xe5, 0x07, 0x81, 0xbd, 0x85, 0x2a, 0x50,
	0x3a, 0xf3, 0x3a, 0x5d, 0xdb, 0x38, 0xfa, 0x01, 0xb5, 0x8d, 0x0f, 0x44, 0xb5, 0xbe, 0x78, 0x9d,
	0xd0, 0xde, 0x42, 0x36, 0x58, 0x3d, 0xef, 0xd2, 0x1f, 0x5d, 0x0c, 0x3a, 0xfd, 0x51, 0xf8, 0xd5,
	0x36, 0x50, 0x1d, 0x40, 0x57, 0x9a, 0xdd, 0x41, 0xeb, 0xd2, 0x2e, 0x28, 0x7c, 0x35, 0x08, 0xfd,
	0x51, 0x10, 0x7a, 0x38, 0xb4, 0x8b, 0x68, 0x17, 0x6a, 0x1a, 0xb7, 0x06, 0xbd, 0x61, 0xd7, 0x0f,
	0x7d, 0xbb, 0x84, 0x00, 0x4c, 0xec, 0x5f, 0xf8, 0xad, 0xd0, 0x2e, 0x1f, 0xf5, 0x60, 0xf7, 0xc1,
	0x76, 0x95, 0xaf, 0xa1, 0xdf, 0x6f, 0x77, 0xfa, 0x9f, 0xec, 0x2d, 0x45, 0x18, 0xf8, 0xfd, 0x76,
	0x2a, 0x60, 0xa0, 0x1a, 0x54, 0x03, 0xdf, 0xeb, 0x8e, 0xb4, 0xd9, 0x02, 0xb2, 0xa0, 0xb2, 0xa6,
	0x2e, 0x36, 0x4f, 0xbf, 0xbd, 0x9d, 0x46, 0x72, 0xb6, 0xba, 0x51, 0x47, 0xd7, 0x20, 0xf1, 0x64,
	0x25, 0xf4, 0x4b, 0xd5, 0x98, 0xb2, 0xe3, 0x1c, 0xd2, 0xaf, 0xa4, 0x48, 0xdf, 0xcc, 0x1b, 0x53,
	0xc3, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xf0, 0x52, 0xe2, 0x4b, 0x05, 0x00, 0x00,
}
