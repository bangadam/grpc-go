// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// model
type Calculator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *Calculator) Reset() {
	*x = Calculator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calculator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calculator) ProtoMessage() {}

func (x *Calculator) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calculator.ProtoReflect.Descriptor instead.
func (*Calculator) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Calculator) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *Calculator) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type PrimeNumberDecomposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeNumberDecomposition) Reset() {
	*x = PrimeNumberDecomposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecomposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecomposition) ProtoMessage() {}

func (x *PrimeNumberDecomposition) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecomposition.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecomposition) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeNumberDecomposition) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

// request
type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculator *Calculator `protobuf:"bytes,1,opt,name=calculator,proto3" json:"calculator,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *SumRequest) GetCalculator() *Calculator {
	if x != nil {
		return x.Calculator
	}
	return nil
}

type PrimeNumberDecompositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumberDecomposition *PrimeNumberDecomposition `protobuf:"bytes,1,opt,name=primeNumberDecomposition,proto3" json:"primeNumberDecomposition,omitempty"`
}

func (x *PrimeNumberDecompositionRequest) Reset() {
	*x = PrimeNumberDecompositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionRequest) ProtoMessage() {}

func (x *PrimeNumberDecompositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeNumberDecompositionRequest) GetPrimeNumberDecomposition() *PrimeNumberDecomposition {
	if x != nil {
		return x.PrimeNumberDecomposition
	}
	return nil
}

// response
type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *SumResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumberDecompositionResponse) Reset() {
	*x = PrimeNumberDecompositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionResponse) ProtoMessage() {}

func (x *PrimeNumberDecompositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *PrimeNumberDecompositionResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x2c, 0x0a, 0x18,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x44, 0x0a, 0x0a, 0x53, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x83, 0x01, 0x0a, 0x1f, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x18, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3a, 0x0a,
	0x20, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xc8, 0x01, 0x0a, 0x11, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x18, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_calculator_calculatorpb_calculator_proto_rawDescData = file_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*Calculator)(nil),                       // 0: calculator.Calculator
	(*PrimeNumberDecomposition)(nil),         // 1: calculator.PrimeNumberDecomposition
	(*SumRequest)(nil),                       // 2: calculator.SumRequest
	(*PrimeNumberDecompositionRequest)(nil),  // 3: calculator.PrimeNumberDecompositionRequest
	(*SumResponse)(nil),                      // 4: calculator.SumResponse
	(*PrimeNumberDecompositionResponse)(nil), // 5: calculator.PrimeNumberDecompositionResponse
}
var file_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.SumRequest.calculator:type_name -> calculator.Calculator
	1, // 1: calculator.PrimeNumberDecompositionRequest.primeNumberDecomposition:type_name -> calculator.PrimeNumberDecomposition
	2, // 2: calculator.CalculatorService.Sum:input_type -> calculator.SumRequest
	3, // 3: calculator.CalculatorService.PrimeNumberDecomposition:input_type -> calculator.PrimeNumberDecompositionRequest
	4, // 4: calculator.CalculatorService.Sum:output_type -> calculator.SumResponse
	5, // 5: calculator.CalculatorService.PrimeNumberDecomposition:output_type -> calculator.PrimeNumberDecompositionResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_calculator_calculatorpb_calculator_proto_init() }
func file_calculator_calculatorpb_calculator_proto_init() {
	if File_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calculator); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecomposition); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionRequest); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionResponse); i {
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
			RawDescriptor: file_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_calculator_calculatorpb_calculator_proto = out.File
	file_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_calculator_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// unary api
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// server streaming api
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// unary api
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// server streaming api
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
