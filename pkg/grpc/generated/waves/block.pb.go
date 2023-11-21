// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: waves/block.proto

package waves

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *Block_Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Signature    []byte               `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Transactions []*SignedTransaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_waves_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_waves_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *Block_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Block) GetTransactions() []*SignedTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type MicroBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version               int32                `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Reference             []byte               `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	UpdatedBlockSignature []byte               `protobuf:"bytes,3,opt,name=updated_block_signature,json=updatedBlockSignature,proto3" json:"updated_block_signature,omitempty"`
	SenderPublicKey       []byte               `protobuf:"bytes,4,opt,name=sender_public_key,json=senderPublicKey,proto3" json:"sender_public_key,omitempty"`
	Transactions          []*SignedTransaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`
	StateHash             []byte               `protobuf:"bytes,6,opt,name=state_hash,json=stateHash,proto3" json:"state_hash,omitempty"`
}

func (x *MicroBlock) Reset() {
	*x = MicroBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroBlock) ProtoMessage() {}

func (x *MicroBlock) ProtoReflect() protoreflect.Message {
	mi := &file_waves_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroBlock.ProtoReflect.Descriptor instead.
func (*MicroBlock) Descriptor() ([]byte, []int) {
	return file_waves_block_proto_rawDescGZIP(), []int{1}
}

func (x *MicroBlock) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MicroBlock) GetReference() []byte {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *MicroBlock) GetUpdatedBlockSignature() []byte {
	if x != nil {
		return x.UpdatedBlockSignature
	}
	return nil
}

func (x *MicroBlock) GetSenderPublicKey() []byte {
	if x != nil {
		return x.SenderPublicKey
	}
	return nil
}

func (x *MicroBlock) GetTransactions() []*SignedTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *MicroBlock) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

type SignedMicroBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MicroBlock   *MicroBlock `protobuf:"bytes,1,opt,name=micro_block,json=microBlock,proto3" json:"micro_block,omitempty"`
	Signature    []byte      `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	TotalBlockId []byte      `protobuf:"bytes,3,opt,name=total_block_id,json=totalBlockId,proto3" json:"total_block_id,omitempty"`
}

func (x *SignedMicroBlock) Reset() {
	*x = SignedMicroBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMicroBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMicroBlock) ProtoMessage() {}

func (x *SignedMicroBlock) ProtoReflect() protoreflect.Message {
	mi := &file_waves_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMicroBlock.ProtoReflect.Descriptor instead.
func (*SignedMicroBlock) Descriptor() ([]byte, []int) {
	return file_waves_block_proto_rawDescGZIP(), []int{2}
}

func (x *SignedMicroBlock) GetMicroBlock() *MicroBlock {
	if x != nil {
		return x.MicroBlock
	}
	return nil
}

func (x *SignedMicroBlock) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedMicroBlock) GetTotalBlockId() []byte {
	if x != nil {
		return x.TotalBlockId
	}
	return nil
}

type Block_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId             int32                          `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Reference           []byte                         `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	BaseTarget          int64                          `protobuf:"varint,3,opt,name=base_target,json=baseTarget,proto3" json:"base_target,omitempty"`
	GenerationSignature []byte                         `protobuf:"bytes,4,opt,name=generation_signature,json=generationSignature,proto3" json:"generation_signature,omitempty"`
	FeatureVotes        []uint32                       `protobuf:"varint,5,rep,packed,name=feature_votes,json=featureVotes,proto3" json:"feature_votes,omitempty"`
	Timestamp           int64                          `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Version             int32                          `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	Generator           []byte                         `protobuf:"bytes,8,opt,name=generator,proto3" json:"generator,omitempty"`
	RewardVote          int64                          `protobuf:"varint,9,opt,name=reward_vote,json=rewardVote,proto3" json:"reward_vote,omitempty"`
	TransactionsRoot    []byte                         `protobuf:"bytes,10,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	StateHash           []byte                         `protobuf:"bytes,11,opt,name=state_hash,json=stateHash,proto3" json:"state_hash,omitempty"`
	ChallengedHeader    *Block_Header_ChallengedHeader `protobuf:"bytes,12,opt,name=challenged_header,json=challengedHeader,proto3" json:"challenged_header,omitempty"`
}

func (x *Block_Header) Reset() {
	*x = Block_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block_Header) ProtoMessage() {}

func (x *Block_Header) ProtoReflect() protoreflect.Message {
	mi := &file_waves_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block_Header.ProtoReflect.Descriptor instead.
func (*Block_Header) Descriptor() ([]byte, []int) {
	return file_waves_block_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Block_Header) GetChainId() int32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Block_Header) GetReference() []byte {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *Block_Header) GetBaseTarget() int64 {
	if x != nil {
		return x.BaseTarget
	}
	return 0
}

func (x *Block_Header) GetGenerationSignature() []byte {
	if x != nil {
		return x.GenerationSignature
	}
	return nil
}

func (x *Block_Header) GetFeatureVotes() []uint32 {
	if x != nil {
		return x.FeatureVotes
	}
	return nil
}

func (x *Block_Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block_Header) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Block_Header) GetGenerator() []byte {
	if x != nil {
		return x.Generator
	}
	return nil
}

func (x *Block_Header) GetRewardVote() int64 {
	if x != nil {
		return x.RewardVote
	}
	return 0
}

func (x *Block_Header) GetTransactionsRoot() []byte {
	if x != nil {
		return x.TransactionsRoot
	}
	return nil
}

func (x *Block_Header) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *Block_Header) GetChallengedHeader() *Block_Header_ChallengedHeader {
	if x != nil {
		return x.ChallengedHeader
	}
	return nil
}

type Block_Header_ChallengedHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseTarget          int64    `protobuf:"varint,1,opt,name=base_target,json=baseTarget,proto3" json:"base_target,omitempty"`
	GenerationSignature []byte   `protobuf:"bytes,2,opt,name=generation_signature,json=generationSignature,proto3" json:"generation_signature,omitempty"`
	FeatureVotes        []uint32 `protobuf:"varint,3,rep,packed,name=feature_votes,json=featureVotes,proto3" json:"feature_votes,omitempty"`
	Timestamp           int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Generator           []byte   `protobuf:"bytes,5,opt,name=generator,proto3" json:"generator,omitempty"`
	RewardVote          int64    `protobuf:"varint,6,opt,name=reward_vote,json=rewardVote,proto3" json:"reward_vote,omitempty"`
	StateHash           []byte   `protobuf:"bytes,7,opt,name=state_hash,json=stateHash,proto3" json:"state_hash,omitempty"`
	HeaderSignature     []byte   `protobuf:"bytes,8,opt,name=header_signature,json=headerSignature,proto3" json:"header_signature,omitempty"`
}

func (x *Block_Header_ChallengedHeader) Reset() {
	*x = Block_Header_ChallengedHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block_Header_ChallengedHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block_Header_ChallengedHeader) ProtoMessage() {}

func (x *Block_Header_ChallengedHeader) ProtoReflect() protoreflect.Message {
	mi := &file_waves_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block_Header_ChallengedHeader.ProtoReflect.Descriptor instead.
func (*Block_Header_ChallengedHeader) Descriptor() ([]byte, []int) {
	return file_waves_block_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Block_Header_ChallengedHeader) GetBaseTarget() int64 {
	if x != nil {
		return x.BaseTarget
	}
	return 0
}

func (x *Block_Header_ChallengedHeader) GetGenerationSignature() []byte {
	if x != nil {
		return x.GenerationSignature
	}
	return nil
}

func (x *Block_Header_ChallengedHeader) GetFeatureVotes() []uint32 {
	if x != nil {
		return x.FeatureVotes
	}
	return nil
}

func (x *Block_Header_ChallengedHeader) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block_Header_ChallengedHeader) GetGenerator() []byte {
	if x != nil {
		return x.Generator
	}
	return nil
}

func (x *Block_Header_ChallengedHeader) GetRewardVote() int64 {
	if x != nil {
		return x.RewardVote
	}
	return 0
}

func (x *Block_Header_ChallengedHeader) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *Block_Header_ChallengedHeader) GetHeaderSignature() []byte {
	if x != nil {
		return x.HeaderSignature
	}
	return nil
}

var File_waves_block_proto protoreflect.FileDescriptor

var file_waves_block_proto_rawDesc = []byte{
	0x0a, 0x11, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x77, 0x61, 0x76, 0x65, 0x73, 0x1a, 0x17, 0x77, 0x61, 0x76, 0x65,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x98, 0x07, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2b, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x85, 0x06, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x51, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0xb2, 0x02, 0x0a, 0x10, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x31, 0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x76, 0x6f,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x85,
	0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x0b, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x0a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x42, 0x65, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x67, 0x6f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x77, 0x61, 0x76,
	0x65, 0x73, 0xaa, 0x02, 0x05, 0x57, 0x61, 0x76, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_waves_block_proto_rawDescOnce sync.Once
	file_waves_block_proto_rawDescData = file_waves_block_proto_rawDesc
)

func file_waves_block_proto_rawDescGZIP() []byte {
	file_waves_block_proto_rawDescOnce.Do(func() {
		file_waves_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_waves_block_proto_rawDescData)
	})
	return file_waves_block_proto_rawDescData
}

var file_waves_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_waves_block_proto_goTypes = []interface{}{
	(*Block)(nil),                         // 0: waves.Block
	(*MicroBlock)(nil),                    // 1: waves.MicroBlock
	(*SignedMicroBlock)(nil),              // 2: waves.SignedMicroBlock
	(*Block_Header)(nil),                  // 3: waves.Block.Header
	(*Block_Header_ChallengedHeader)(nil), // 4: waves.Block.Header.ChallengedHeader
	(*SignedTransaction)(nil),             // 5: waves.SignedTransaction
}
var file_waves_block_proto_depIdxs = []int32{
	3, // 0: waves.Block.header:type_name -> waves.Block.Header
	5, // 1: waves.Block.transactions:type_name -> waves.SignedTransaction
	5, // 2: waves.MicroBlock.transactions:type_name -> waves.SignedTransaction
	1, // 3: waves.SignedMicroBlock.micro_block:type_name -> waves.MicroBlock
	4, // 4: waves.Block.Header.challenged_header:type_name -> waves.Block.Header.ChallengedHeader
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_waves_block_proto_init() }
func file_waves_block_proto_init() {
	if File_waves_block_proto != nil {
		return
	}
	file_waves_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_waves_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waves_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waves_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMicroBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waves_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block_Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waves_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block_Header_ChallengedHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waves_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waves_block_proto_goTypes,
		DependencyIndexes: file_waves_block_proto_depIdxs,
		MessageInfos:      file_waves_block_proto_msgTypes,
	}.Build()
	File_waves_block_proto = out.File
	file_waves_block_proto_rawDesc = nil
	file_waves_block_proto_goTypes = nil
	file_waves_block_proto_depIdxs = nil
}
