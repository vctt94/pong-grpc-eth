// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pong.proto

package pong

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

type GameStartedStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *GameStartedStreamRequest) Reset() {
	*x = GameStartedStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStartedStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStartedStreamRequest) ProtoMessage() {}

func (x *GameStartedStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStartedStreamRequest.ProtoReflect.Descriptor instead.
func (*GameStartedStreamRequest) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{0}
}

func (x *GameStartedStreamRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type GameStartedStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Started      bool   `protobuf:"varint,1,opt,name=started,proto3" json:"started,omitempty"`
	PlayerNumber int32  `protobuf:"varint,2,opt,name=player_number,json=playerNumber,proto3" json:"player_number,omitempty"` // Player 1 or Player 2
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`                                // Confirmation message or additional details
	ClientId     string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *GameStartedStreamResponse) Reset() {
	*x = GameStartedStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStartedStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStartedStreamResponse) ProtoMessage() {}

func (x *GameStartedStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStartedStreamResponse.ProtoReflect.Descriptor instead.
func (*GameStartedStreamResponse) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{1}
}

func (x *GameStartedStreamResponse) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *GameStartedStreamResponse) GetPlayerNumber() int32 {
	if x != nil {
		return x.PlayerNumber
	}
	return 0
}

func (x *GameStartedStreamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GameStartedStreamResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// SignalReadyRequest contains information about the client signaling readiness
type SignalReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *SignalReadyRequest) Reset() {
	*x = SignalReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalReadyRequest) ProtoMessage() {}

func (x *SignalReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalReadyRequest.ProtoReflect.Descriptor instead.
func (*SignalReadyRequest) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{2}
}

func (x *SignalReadyRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SignalReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignalReadyResponse) Reset() {
	*x = SignalReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalReadyResponse) ProtoMessage() {}

func (x *SignalReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalReadyResponse.ProtoReflect.Descriptor instead.
func (*SignalReadyResponse) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{3}
}

type GameUpdateBytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GameUpdateBytes) Reset() {
	*x = GameUpdateBytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameUpdateBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameUpdateBytes) ProtoMessage() {}

func (x *GameUpdateBytes) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameUpdateBytes.ProtoReflect.Descriptor instead.
func (*GameUpdateBytes) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{4}
}

func (x *GameUpdateBytes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PlayerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId     string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Input        string `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`                // e.g., "ArrowUp", "ArrowDown"
	PlayerNumber int32  `protobuf:"varint,3,opt,name=playerNumber,proto3" json:"playerNumber,omitempty"` // player 1 or player 2.
}

func (x *PlayerInput) Reset() {
	*x = PlayerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInput) ProtoMessage() {}

func (x *PlayerInput) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInput.ProtoReflect.Descriptor instead.
func (*PlayerInput) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerInput) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerInput) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *PlayerInput) GetPlayerNumber() int32 {
	if x != nil {
		return x.PlayerNumber
	}
	return 0
}

type GameUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameWidth     int32   `protobuf:"varint,13,opt,name=gameWidth,proto3" json:"gameWidth,omitempty"`
	GameHeight    int32   `protobuf:"varint,14,opt,name=gameHeight,proto3" json:"gameHeight,omitempty"`
	P1Width       int32   `protobuf:"varint,15,opt,name=p1Width,proto3" json:"p1Width,omitempty"`
	P1Height      int32   `protobuf:"varint,16,opt,name=p1Height,proto3" json:"p1Height,omitempty"`
	P2Width       int32   `protobuf:"varint,17,opt,name=p2Width,proto3" json:"p2Width,omitempty"`
	P2Height      int32   `protobuf:"varint,18,opt,name=p2Height,proto3" json:"p2Height,omitempty"`
	BallWidth     int32   `protobuf:"varint,19,opt,name=ballWidth,proto3" json:"ballWidth,omitempty"`
	BallHeight    int32   `protobuf:"varint,20,opt,name=ballHeight,proto3" json:"ballHeight,omitempty"`
	P1Score       int32   `protobuf:"varint,21,opt,name=p1Score,proto3" json:"p1Score,omitempty"`
	P2Score       int32   `protobuf:"varint,22,opt,name=p2Score,proto3" json:"p2Score,omitempty"`
	BallX         int32   `protobuf:"varint,1,opt,name=ballX,proto3" json:"ballX,omitempty"`
	BallY         int32   `protobuf:"varint,2,opt,name=ballY,proto3" json:"ballY,omitempty"`
	P1X           int32   `protobuf:"varint,3,opt,name=p1X,proto3" json:"p1X,omitempty"`
	P1Y           int32   `protobuf:"varint,4,opt,name=p1Y,proto3" json:"p1Y,omitempty"`
	P2X           int32   `protobuf:"varint,5,opt,name=p2X,proto3" json:"p2X,omitempty"`
	P2Y           int32   `protobuf:"varint,6,opt,name=p2Y,proto3" json:"p2Y,omitempty"`
	P1YVelocity   int32   `protobuf:"varint,7,opt,name=p1YVelocity,proto3" json:"p1YVelocity,omitempty"`
	P2YVelocity   int32   `protobuf:"varint,8,opt,name=p2YVelocity,proto3" json:"p2YVelocity,omitempty"`
	BallXVelocity int32   `protobuf:"varint,9,opt,name=ballXVelocity,proto3" json:"ballXVelocity,omitempty"`
	BallYVelocity int32   `protobuf:"varint,10,opt,name=ballYVelocity,proto3" json:"ballYVelocity,omitempty"`
	Fps           float32 `protobuf:"fixed32,11,opt,name=fps,proto3" json:"fps,omitempty"`
	Tps           float32 `protobuf:"fixed32,12,opt,name=tps,proto3" json:"tps,omitempty"`
	// Optional: if you want to send error messages or debug information
	Error string `protobuf:"bytes,23,opt,name=error,proto3" json:"error,omitempty"`
	Debug bool   `protobuf:"varint,24,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *GameUpdate) Reset() {
	*x = GameUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameUpdate) ProtoMessage() {}

func (x *GameUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameUpdate.ProtoReflect.Descriptor instead.
func (*GameUpdate) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{6}
}

func (x *GameUpdate) GetGameWidth() int32 {
	if x != nil {
		return x.GameWidth
	}
	return 0
}

func (x *GameUpdate) GetGameHeight() int32 {
	if x != nil {
		return x.GameHeight
	}
	return 0
}

func (x *GameUpdate) GetP1Width() int32 {
	if x != nil {
		return x.P1Width
	}
	return 0
}

func (x *GameUpdate) GetP1Height() int32 {
	if x != nil {
		return x.P1Height
	}
	return 0
}

func (x *GameUpdate) GetP2Width() int32 {
	if x != nil {
		return x.P2Width
	}
	return 0
}

func (x *GameUpdate) GetP2Height() int32 {
	if x != nil {
		return x.P2Height
	}
	return 0
}

func (x *GameUpdate) GetBallWidth() int32 {
	if x != nil {
		return x.BallWidth
	}
	return 0
}

func (x *GameUpdate) GetBallHeight() int32 {
	if x != nil {
		return x.BallHeight
	}
	return 0
}

func (x *GameUpdate) GetP1Score() int32 {
	if x != nil {
		return x.P1Score
	}
	return 0
}

func (x *GameUpdate) GetP2Score() int32 {
	if x != nil {
		return x.P2Score
	}
	return 0
}

func (x *GameUpdate) GetBallX() int32 {
	if x != nil {
		return x.BallX
	}
	return 0
}

func (x *GameUpdate) GetBallY() int32 {
	if x != nil {
		return x.BallY
	}
	return 0
}

func (x *GameUpdate) GetP1X() int32 {
	if x != nil {
		return x.P1X
	}
	return 0
}

func (x *GameUpdate) GetP1Y() int32 {
	if x != nil {
		return x.P1Y
	}
	return 0
}

func (x *GameUpdate) GetP2X() int32 {
	if x != nil {
		return x.P2X
	}
	return 0
}

func (x *GameUpdate) GetP2Y() int32 {
	if x != nil {
		return x.P2Y
	}
	return 0
}

func (x *GameUpdate) GetP1YVelocity() int32 {
	if x != nil {
		return x.P1YVelocity
	}
	return 0
}

func (x *GameUpdate) GetP2YVelocity() int32 {
	if x != nil {
		return x.P2YVelocity
	}
	return 0
}

func (x *GameUpdate) GetBallXVelocity() int32 {
	if x != nil {
		return x.BallXVelocity
	}
	return 0
}

func (x *GameUpdate) GetBallYVelocity() int32 {
	if x != nil {
		return x.BallYVelocity
	}
	return 0
}

func (x *GameUpdate) GetFps() float32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *GameUpdate) GetTps() float32 {
	if x != nil {
		return x.Tps
	}
	return 0
}

func (x *GameUpdate) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GameUpdate) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type GameStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *GameStreamRequest) Reset() {
	*x = GameStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStreamRequest) ProtoMessage() {}

func (x *GameStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStreamRequest.ProtoReflect.Descriptor instead.
func (*GameStreamRequest) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{7}
}

func (x *GameStreamRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

var File_pong_proto protoreflect.FileDescriptor

var file_pong_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f,
	0x6e, 0x67, 0x22, 0x37, 0x0a, 0x18, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x19,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x31, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x61, 0x6d,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x63, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xfc, 0x04, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x31, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x31, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x31, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x31, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x32, 0x57, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x32, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x32, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x32, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x61, 0x6c, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x62, 0x61, 0x6c, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70,
	0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x32, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x32, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x62, 0x61, 0x6c, 0x6c, 0x58, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x59, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x59, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x31, 0x58, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x31, 0x58, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x31, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x31, 0x59,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x32, 0x58, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x32, 0x58, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x32, 0x59, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x32, 0x59, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x31, 0x59, 0x56, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x31, 0x59, 0x56, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x32, 0x59, 0x56, 0x65, 0x6c,
	0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x32, 0x59,
	0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x6c,
	0x58, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x62, 0x61, 0x6c, 0x6c, 0x58, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x24,
	0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x6c, 0x59, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x61, 0x6c, 0x6c, 0x59, 0x56, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x70, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x32, 0x9b, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x6e, 0x67, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x11, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x6e,
	0x67, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pong_proto_rawDescOnce sync.Once
	file_pong_proto_rawDescData = file_pong_proto_rawDesc
)

func file_pong_proto_rawDescGZIP() []byte {
	file_pong_proto_rawDescOnce.Do(func() {
		file_pong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pong_proto_rawDescData)
	})
	return file_pong_proto_rawDescData
}

var file_pong_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pong_proto_goTypes = []interface{}{
	(*GameStartedStreamRequest)(nil),  // 0: pong.GameStartedStreamRequest
	(*GameStartedStreamResponse)(nil), // 1: pong.GameStartedStreamResponse
	(*SignalReadyRequest)(nil),        // 2: pong.SignalReadyRequest
	(*SignalReadyResponse)(nil),       // 3: pong.SignalReadyResponse
	(*GameUpdateBytes)(nil),           // 4: pong.GameUpdateBytes
	(*PlayerInput)(nil),               // 5: pong.PlayerInput
	(*GameUpdate)(nil),                // 6: pong.GameUpdate
	(*GameStreamRequest)(nil),         // 7: pong.GameStreamRequest
}
var file_pong_proto_depIdxs = []int32{
	5, // 0: pong.PongGame.SendInput:input_type -> pong.PlayerInput
	7, // 1: pong.PongGame.StreamUpdates:input_type -> pong.GameStreamRequest
	2, // 2: pong.PongGame.SignalReady:input_type -> pong.SignalReadyRequest
	0, // 3: pong.PongGame.StartNotifier:input_type -> pong.GameStartedStreamRequest
	6, // 4: pong.PongGame.SendInput:output_type -> pong.GameUpdate
	4, // 5: pong.PongGame.StreamUpdates:output_type -> pong.GameUpdateBytes
	3, // 6: pong.PongGame.SignalReady:output_type -> pong.SignalReadyResponse
	1, // 7: pong.PongGame.StartNotifier:output_type -> pong.GameStartedStreamResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pong_proto_init() }
func file_pong_proto_init() {
	if File_pong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStartedStreamRequest); i {
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
		file_pong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStartedStreamResponse); i {
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
		file_pong_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalReadyRequest); i {
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
		file_pong_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalReadyResponse); i {
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
		file_pong_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameUpdateBytes); i {
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
		file_pong_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInput); i {
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
		file_pong_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameUpdate); i {
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
		file_pong_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStreamRequest); i {
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
			RawDescriptor: file_pong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pong_proto_goTypes,
		DependencyIndexes: file_pong_proto_depIdxs,
		MessageInfos:      file_pong_proto_msgTypes,
	}.Build()
	File_pong_proto = out.File
	file_pong_proto_rawDesc = nil
	file_pong_proto_goTypes = nil
	file_pong_proto_depIdxs = nil
}
