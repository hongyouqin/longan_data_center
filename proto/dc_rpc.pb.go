// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dc_rpc.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 拉取人
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc_rpc_fca8d9355fc952ab, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type UpdateNotity struct {
	EmployeeTableUpdate  bool     `protobuf:"varint,1,opt,name=employee_table_update,json=employeeTableUpdate,proto3" json:"employee_table_update,omitempty"`
	StorageTableUpdate   bool     `protobuf:"varint,2,opt,name=storage_table_update,json=storageTableUpdate,proto3" json:"storage_table_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNotity) Reset()         { *m = UpdateNotity{} }
func (m *UpdateNotity) String() string { return proto.CompactTextString(m) }
func (*UpdateNotity) ProtoMessage()    {}
func (*UpdateNotity) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc_rpc_fca8d9355fc952ab, []int{1}
}
func (m *UpdateNotity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNotity.Unmarshal(m, b)
}
func (m *UpdateNotity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNotity.Marshal(b, m, deterministic)
}
func (dst *UpdateNotity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNotity.Merge(dst, src)
}
func (m *UpdateNotity) XXX_Size() int {
	return xxx_messageInfo_UpdateNotity.Size(m)
}
func (m *UpdateNotity) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNotity.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNotity proto.InternalMessageInfo

func (m *UpdateNotity) GetEmployeeTableUpdate() bool {
	if m != nil {
		return m.EmployeeTableUpdate
	}
	return false
}

func (m *UpdateNotity) GetStorageTableUpdate() bool {
	if m != nil {
		return m.StorageTableUpdate
	}
	return false
}

// 人脸特征
type Feature struct {
	Face                 *UserParam `protobuf:"bytes,1,opt,name=face,proto3" json:"face,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc_rpc_fca8d9355fc952ab, []int{2}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetFace() *UserParam {
	if m != nil {
		return m.Face
	}
	return nil
}

// 用户参数
type UserParam struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId               uint64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FacePhoto            string   `protobuf:"bytes,3,opt,name=face_photo,json=facePhoto,proto3" json:"face_photo,omitempty"`
	FaceFeature          []byte   `protobuf:"bytes,4,opt,name=face_feature,json=faceFeature,proto3" json:"face_feature,omitempty"`
	FaceSize             uint64   `protobuf:"varint,5,opt,name=face_size,json=faceSize,proto3" json:"face_size,omitempty"`
	AcquisitionTime      uint64   `protobuf:"varint,6,opt,name=acquisition_time,json=acquisitionTime,proto3" json:"acquisition_time,omitempty"`
	FaceScore            uint32   `protobuf:"varint,7,opt,name=face_score,json=faceScore,proto3" json:"face_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserParam) Reset()         { *m = UserParam{} }
func (m *UserParam) String() string { return proto.CompactTextString(m) }
func (*UserParam) ProtoMessage()    {}
func (*UserParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc_rpc_fca8d9355fc952ab, []int{3}
}
func (m *UserParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserParam.Unmarshal(m, b)
}
func (m *UserParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserParam.Marshal(b, m, deterministic)
}
func (dst *UserParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserParam.Merge(dst, src)
}
func (m *UserParam) XXX_Size() int {
	return xxx_messageInfo_UserParam.Size(m)
}
func (m *UserParam) XXX_DiscardUnknown() {
	xxx_messageInfo_UserParam.DiscardUnknown(m)
}

var xxx_messageInfo_UserParam proto.InternalMessageInfo

func (m *UserParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserParam) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserParam) GetFacePhoto() string {
	if m != nil {
		return m.FacePhoto
	}
	return ""
}

func (m *UserParam) GetFaceFeature() []byte {
	if m != nil {
		return m.FaceFeature
	}
	return nil
}

func (m *UserParam) GetFaceSize() uint64 {
	if m != nil {
		return m.FaceSize
	}
	return 0
}

func (m *UserParam) GetAcquisitionTime() uint64 {
	if m != nil {
		return m.AcquisitionTime
	}
	return 0
}

func (m *UserParam) GetFaceScore() uint32 {
	if m != nil {
		return m.FaceScore
	}
	return 0
}

// 储存结果
type StorageReply struct {
	IsSuccess            bool     `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageReply) Reset()         { *m = StorageReply{} }
func (m *StorageReply) String() string { return proto.CompactTextString(m) }
func (*StorageReply) ProtoMessage()    {}
func (*StorageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc_rpc_fca8d9355fc952ab, []int{4}
}
func (m *StorageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageReply.Unmarshal(m, b)
}
func (m *StorageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageReply.Marshal(b, m, deterministic)
}
func (dst *StorageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageReply.Merge(dst, src)
}
func (m *StorageReply) XXX_Size() int {
	return xxx_messageInfo_StorageReply.Size(m)
}
func (m *StorageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageReply.DiscardUnknown(m)
}

var xxx_messageInfo_StorageReply proto.InternalMessageInfo

func (m *StorageReply) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*UpdateNotity)(nil), "proto.UpdateNotity")
	proto.RegisterType((*Feature)(nil), "proto.Feature")
	proto.RegisterType((*UserParam)(nil), "proto.UserParam")
	proto.RegisterType((*StorageReply)(nil), "proto.StorageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LonganDataCenterClient is the client API for LonganDataCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LonganDataCenterClient interface {
	// 存储员工人脸信息
	StorageEmployeeFace(ctx context.Context, in *UserParam, opts ...grpc.CallOption) (*StorageReply, error)
	// 存储陌生人脸信息
	StorageStrangerFace(ctx context.Context, in *UserParam, opts ...grpc.CallOption) (*StorageReply, error)
	// 拉取人脸注册表的所有数据
	ExtractFaceRegTableDatas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LonganDataCenter_ExtractFaceRegTableDatasClient, error)
	// 注册客户端，用于之后的更新通知接收
	RegisterClient(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LonganDataCenter_RegisterClientClient, error)
}

type longanDataCenterClient struct {
	cc *grpc.ClientConn
}

func NewLonganDataCenterClient(cc *grpc.ClientConn) LonganDataCenterClient {
	return &longanDataCenterClient{cc}
}

func (c *longanDataCenterClient) StorageEmployeeFace(ctx context.Context, in *UserParam, opts ...grpc.CallOption) (*StorageReply, error) {
	out := new(StorageReply)
	err := c.cc.Invoke(ctx, "/proto.LonganDataCenter/StorageEmployeeFace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longanDataCenterClient) StorageStrangerFace(ctx context.Context, in *UserParam, opts ...grpc.CallOption) (*StorageReply, error) {
	out := new(StorageReply)
	err := c.cc.Invoke(ctx, "/proto.LonganDataCenter/StorageStrangerFace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longanDataCenterClient) ExtractFaceRegTableDatas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LonganDataCenter_ExtractFaceRegTableDatasClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LonganDataCenter_serviceDesc.Streams[0], "/proto.LonganDataCenter/ExtractFaceRegTableDatas", opts...)
	if err != nil {
		return nil, err
	}
	x := &longanDataCenterExtractFaceRegTableDatasClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LonganDataCenter_ExtractFaceRegTableDatasClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type longanDataCenterExtractFaceRegTableDatasClient struct {
	grpc.ClientStream
}

func (x *longanDataCenterExtractFaceRegTableDatasClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *longanDataCenterClient) RegisterClient(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LonganDataCenter_RegisterClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LonganDataCenter_serviceDesc.Streams[1], "/proto.LonganDataCenter/RegisterClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &longanDataCenterRegisterClientClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LonganDataCenter_RegisterClientClient interface {
	Recv() (*UpdateNotity, error)
	grpc.ClientStream
}

type longanDataCenterRegisterClientClient struct {
	grpc.ClientStream
}

func (x *longanDataCenterRegisterClientClient) Recv() (*UpdateNotity, error) {
	m := new(UpdateNotity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LonganDataCenterServer is the server API for LonganDataCenter service.
type LonganDataCenterServer interface {
	// 存储员工人脸信息
	StorageEmployeeFace(context.Context, *UserParam) (*StorageReply, error)
	// 存储陌生人脸信息
	StorageStrangerFace(context.Context, *UserParam) (*StorageReply, error)
	// 拉取人脸注册表的所有数据
	ExtractFaceRegTableDatas(*Empty, LonganDataCenter_ExtractFaceRegTableDatasServer) error
	// 注册客户端，用于之后的更新通知接收
	RegisterClient(*Empty, LonganDataCenter_RegisterClientServer) error
}

func RegisterLonganDataCenterServer(s *grpc.Server, srv LonganDataCenterServer) {
	s.RegisterService(&_LonganDataCenter_serviceDesc, srv)
}

func _LonganDataCenter_StorageEmployeeFace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonganDataCenterServer).StorageEmployeeFace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LonganDataCenter/StorageEmployeeFace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonganDataCenterServer).StorageEmployeeFace(ctx, req.(*UserParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonganDataCenter_StorageStrangerFace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonganDataCenterServer).StorageStrangerFace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LonganDataCenter/StorageStrangerFace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonganDataCenterServer).StorageStrangerFace(ctx, req.(*UserParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonganDataCenter_ExtractFaceRegTableDatas_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LonganDataCenterServer).ExtractFaceRegTableDatas(m, &longanDataCenterExtractFaceRegTableDatasServer{stream})
}

type LonganDataCenter_ExtractFaceRegTableDatasServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type longanDataCenterExtractFaceRegTableDatasServer struct {
	grpc.ServerStream
}

func (x *longanDataCenterExtractFaceRegTableDatasServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _LonganDataCenter_RegisterClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LonganDataCenterServer).RegisterClient(m, &longanDataCenterRegisterClientServer{stream})
}

type LonganDataCenter_RegisterClientServer interface {
	Send(*UpdateNotity) error
	grpc.ServerStream
}

type longanDataCenterRegisterClientServer struct {
	grpc.ServerStream
}

func (x *longanDataCenterRegisterClientServer) Send(m *UpdateNotity) error {
	return x.ServerStream.SendMsg(m)
}

var _LonganDataCenter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LonganDataCenter",
	HandlerType: (*LonganDataCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StorageEmployeeFace",
			Handler:    _LonganDataCenter_StorageEmployeeFace_Handler,
		},
		{
			MethodName: "StorageStrangerFace",
			Handler:    _LonganDataCenter_StorageStrangerFace_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExtractFaceRegTableDatas",
			Handler:       _LonganDataCenter_ExtractFaceRegTableDatas_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterClient",
			Handler:       _LonganDataCenter_RegisterClient_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dc_rpc.proto",
}

func init() { proto.RegisterFile("dc_rpc.proto", fileDescriptor_dc_rpc_fca8d9355fc952ab) }

var fileDescriptor_dc_rpc_fca8d9355fc952ab = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0x92, 0x3f, 0xcd, 0xd4, 0x84, 0x68, 0x02, 0xc2, 0x2a, 0x42, 0x0a, 0x16, 0x87,
	0x70, 0x20, 0x54, 0xe1, 0xc0, 0x05, 0x71, 0xa0, 0xa4, 0x12, 0x12, 0x42, 0x91, 0xdd, 0x9e, 0xad,
	0xed, 0x66, 0x6a, 0x56, 0xb2, 0xbd, 0x66, 0x77, 0x2c, 0x91, 0xbe, 0x01, 0x8f, 0xc9, 0x9b, 0xa0,
	0x5d, 0x3b, 0x95, 0x55, 0x4e, 0x9c, 0xec, 0xfd, 0xbe, 0xf9, 0xcd, 0x7c, 0x9a, 0x5d, 0x08, 0x77,
	0x32, 0x33, 0xb5, 0x5c, 0xd5, 0x46, 0xb3, 0xc6, 0xa1, 0xff, 0xc4, 0x63, 0x18, 0x6e, 0xca, 0x9a,
	0xf7, 0x31, 0x43, 0x78, 0x5d, 0xef, 0x04, 0xd3, 0x77, 0xcd, 0x8a, 0xf7, 0xb8, 0x86, 0x67, 0x54,
	0xd6, 0x85, 0xde, 0x13, 0x65, 0x2c, 0x6e, 0x0a, 0xca, 0x1a, 0x6f, 0x47, 0xc1, 0x22, 0x58, 0x9e,
	0x24, 0xf3, 0x83, 0x79, 0xe5, 0xbc, 0x96, 0xc4, 0x73, 0x78, 0x6a, 0x59, 0x1b, 0x91, 0x3f, 0x40,
	0x8e, 0x3d, 0x82, 0x9d, 0xd7, 0x23, 0xe2, 0x77, 0x30, 0xbe, 0x24, 0xc1, 0x8d, 0x21, 0x7c, 0x0d,
	0x83, 0x5b, 0x21, 0xdb, 0xfe, 0xa7, 0xeb, 0x59, 0x1b, 0x73, 0x75, 0x6d, 0xc9, 0x6c, 0x85, 0x11,
	0x65, 0xe2, 0xdd, 0xf8, 0x4f, 0x00, 0x93, 0x7b, 0x0d, 0x11, 0x06, 0x95, 0x28, 0x5b, 0x66, 0x92,
	0xf8, 0x7f, 0x7c, 0x0e, 0xe3, 0xc6, 0x92, 0xc9, 0xd4, 0xce, 0xcf, 0x1d, 0x24, 0x23, 0x77, 0xfc,
	0xba, 0xc3, 0x97, 0x00, 0xae, 0x45, 0x56, 0xff, 0xd0, 0xac, 0xa3, 0x47, 0x1e, 0x99, 0x38, 0x65,
	0xeb, 0x04, 0x7c, 0x05, 0xa1, 0xb7, 0x6f, 0xdb, 0x3c, 0xd1, 0x60, 0x11, 0x2c, 0xc3, 0xe4, 0xd4,
	0x69, 0x87, 0x88, 0x2f, 0xc0, 0xd7, 0x67, 0x56, 0xdd, 0x51, 0x34, 0xf4, 0xcd, 0x4f, 0x9c, 0x90,
	0xaa, 0x3b, 0xc2, 0x37, 0x30, 0x13, 0xf2, 0x67, 0xa3, 0xac, 0x62, 0xa5, 0xab, 0x8c, 0x55, 0x49,
	0xd1, 0xc8, 0xd7, 0x3c, 0xe9, 0xe9, 0x57, 0xaa, 0xa4, 0xfb, 0x24, 0x56, 0x6a, 0x43, 0xd1, 0x78,
	0x11, 0x2c, 0x1f, 0xb7, 0x49, 0x52, 0x27, 0xc4, 0x6f, 0x21, 0x4c, 0xdb, 0x55, 0x25, 0x54, 0x17,
	0x7b, 0x57, 0xae, 0x6c, 0x66, 0x1b, 0x29, 0xc9, 0xda, 0x6e, 0xff, 0x13, 0x65, 0xd3, 0x56, 0x58,
	0xff, 0x3e, 0x86, 0xd9, 0x37, 0x5d, 0xe5, 0xa2, 0xfa, 0x22, 0x58, 0x5c, 0x50, 0xc5, 0x64, 0xf0,
	0x13, 0xcc, 0xbb, 0x1e, 0x9b, 0xee, 0xa2, 0x2e, 0x85, 0x24, 0xfc, 0x67, 0xad, 0x67, 0xf3, 0x4e,
	0xe9, 0x4f, 0x8c, 0x8f, 0x7a, 0x7c, 0xca, 0x46, 0x54, 0x39, 0x99, 0xff, 0xe3, 0x3f, 0x42, 0xb4,
	0xf9, 0xc5, 0x46, 0x48, 0x76, 0x5c, 0x42, 0xb9, 0xbf, 0x75, 0x17, 0xd0, 0x62, 0xd8, 0x21, 0xfe,
	0xe1, 0x9d, 0x4d, 0xbb, 0x53, 0xb7, 0xe4, 0xf8, 0xe8, 0x3c, 0xc0, 0x0f, 0x30, 0x4d, 0x28, 0x57,
	0x96, 0xc9, 0x5c, 0x14, 0x8a, 0x2a, 0x7e, 0xc0, 0x1c, 0x86, 0xf6, 0x5f, 0xac, 0x03, 0x3f, 0x23,
	0x4c, 0xa5, 0x2e, 0x57, 0x85, 0x5f, 0xc7, 0xca, 0xd4, 0x72, 0x1b, 0xdc, 0x8c, 0x7c, 0xed, 0xfb,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x30, 0x47, 0x8c, 0x00, 0x03, 0x00, 0x00,
}
