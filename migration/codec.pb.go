// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: migration/codec.proto

package migration

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_iov_one_weave "github.com/iov-one/weave"
	weave "github.com/iov-one/weave"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Configuration struct {
	// Admin holds the address of the administrator allowed to upgrade schema
	// versions. If you wish to permit more than one entity to be an admin, use
	// multisig.
	Admin github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=admin,proto3,casttype=github.com/iov-one/weave.Address" json:"admin,omitempty"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf669b5eede564b, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return m.Size()
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetAdmin() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Admin
	}
	return nil
}

// Schema declares the maxiumum supported schema version for a package.
type Schema struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Pkg holds the name of the package that this migration is declared for.
	// For example, for extension `x/myext` package value is `myext`
	Pkg string `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	// Version holds the highest supported schema version.
	Version uint32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf669b5eede564b, []int{1}
}
func (m *Schema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return m.Size()
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Schema) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *Schema) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// UpgradeSchemaMsg is a request to upgrade schema version of a given package
// by one version.
type UpgradeSchemaMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Name of the package that schema version upgrade is made for.
	Pkg string `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	// To version defines to which version upgrade this schema. It always must be
	// a value one higher than the currently registered schema or 1 if this is
	// the initial upgrade.
	// This value must be provided to avoid consecutive upgrades submitted by a
	// mistake. This ensures at most one delivery as upgrades to an invalid
	// version are rejected.
	ToVersion uint32 `protobuf:"varint,3,opt,name=to_version,json=toVersion,proto3" json:"to_version,omitempty"`
}

func (m *UpgradeSchemaMsg) Reset()         { *m = UpgradeSchemaMsg{} }
func (m *UpgradeSchemaMsg) String() string { return proto.CompactTextString(m) }
func (*UpgradeSchemaMsg) ProtoMessage()    {}
func (*UpgradeSchemaMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf669b5eede564b, []int{2}
}
func (m *UpgradeSchemaMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpgradeSchemaMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpgradeSchemaMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpgradeSchemaMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeSchemaMsg.Merge(m, src)
}
func (m *UpgradeSchemaMsg) XXX_Size() int {
	return m.Size()
}
func (m *UpgradeSchemaMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeSchemaMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeSchemaMsg proto.InternalMessageInfo

func (m *UpgradeSchemaMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UpgradeSchemaMsg) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *UpgradeSchemaMsg) GetToVersion() uint32 {
	if m != nil {
		return m.ToVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*Configuration)(nil), "migration.Configuration")
	proto.RegisterType((*Schema)(nil), "migration.Schema")
	proto.RegisterType((*UpgradeSchemaMsg)(nil), "migration.UpgradeSchemaMsg")
}

func init() { proto.RegisterFile("migration/codec.proto", fileDescriptor_ecf669b5eede564b) }

var fileDescriptor_ecf669b5eede564b = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x17, 0x87, 0xd3, 0x7e, 0x73, 0x38, 0x8a, 0x42, 0x19, 0x18, 0x4b, 0xf1, 0x50, 0x10,
	0x5b, 0xd0, 0x9b, 0x37, 0xe7, 0x51, 0x76, 0xa9, 0xe8, 0x55, 0xb2, 0x26, 0x66, 0x41, 0x9a, 0xaf,
	0xa4, 0x59, 0xfd, 0x1b, 0xfe, 0x2c, 0x8f, 0x3b, 0x7a, 0x12, 0x69, 0xff, 0x85, 0x27, 0x59, 0xa3,
	0x43, 0xcf, 0xde, 0xde, 0x3c, 0xe4, 0xe1, 0x81, 0x0f, 0x0e, 0x0b, 0x25, 0x0d, 0xb3, 0x0a, 0x75,
	0x9a, 0x23, 0x17, 0x79, 0x52, 0x1a, 0xb4, 0xe8, 0x7b, 0x1b, 0x3c, 0x19, 0xfe, 0xe2, 0x93, 0x03,
	0x89, 0x12, 0xbb, 0x99, 0xae, 0x97, 0xa3, 0xd1, 0x0d, 0x8c, 0xae, 0x51, 0x3f, 0x2a, 0xb9, 0x74,
	0x8e, 0x7f, 0x09, 0xdb, 0x8c, 0x17, 0x4a, 0x07, 0x5b, 0x21, 0x89, 0xf7, 0xa6, 0x27, 0x9f, 0xef,
	0xc7, 0xa1, 0x54, 0x76, 0xb1, 0x9c, 0x27, 0x39, 0x16, 0xa9, 0xc2, 0xfa, 0x0c, 0xb5, 0x48, 0x9f,
	0x05, 0xab, 0x45, 0x72, 0xc5, 0xb9, 0x11, 0x55, 0x95, 0x39, 0x25, 0x62, 0x30, 0xb8, 0xcd, 0x17,
	0xa2, 0x60, 0xfe, 0x29, 0xec, 0x16, 0xc2, 0x32, 0xce, 0x2c, 0x0b, 0x48, 0x48, 0xe2, 0xe1, 0xf9,
	0x7e, 0xe2, 0x94, 0xd9, 0x37, 0xce, 0x36, 0x1f, 0xfc, 0x31, 0xf4, 0xcb, 0x27, 0xd9, 0x05, 0xbd,
	0x6c, 0x3d, 0xfd, 0x00, 0x76, 0x6a, 0x61, 0x2a, 0x85, 0x3a, 0xe8, 0x87, 0x24, 0x1e, 0x65, 0x3f,
	0xcf, 0xa8, 0x84, 0xf1, 0x5d, 0x29, 0x0d, 0xe3, 0xc2, 0x95, 0x66, 0x95, 0xfc, 0x6f, 0xec, 0x08,
	0xc0, 0xe2, 0xc3, 0xdf, 0x9e, 0x67, 0xf1, 0xde, 0x81, 0x69, 0xf0, 0xda, 0x50, 0xb2, 0x6a, 0x28,
	0xf9, 0x68, 0x28, 0x79, 0x69, 0x69, 0x6f, 0xd5, 0xd2, 0xde, 0x5b, 0x4b, 0x7b, 0xf3, 0x41, 0x77,
	0xc2, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x95, 0x80, 0x6f, 0x89, 0x01, 0x00, 0x00,
}

func (m *Configuration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Configuration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Admin) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Admin)))
		i += copy(dAtA[i:], m.Admin)
	}
	return i, nil
}

func (m *Schema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schema) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Pkg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Pkg)))
		i += copy(dAtA[i:], m.Pkg)
	}
	if m.Version != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func (m *UpgradeSchemaMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpgradeSchemaMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n2, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Pkg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Pkg)))
		i += copy(dAtA[i:], m.Pkg)
	}
	if m.ToVersion != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ToVersion))
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Configuration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *Schema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Pkg)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovCodec(uint64(m.Version))
	}
	return n
}

func (m *UpgradeSchemaMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Pkg)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.ToVersion != 0 {
		n += 1 + sovCodec(uint64(m.ToVersion))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Configuration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Configuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Configuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = append(m.Admin[:0], dAtA[iNdEx:postIndex]...)
			if m.Admin == nil {
				m.Admin = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Schema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Schema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pkg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pkg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpgradeSchemaMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpgradeSchemaMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpgradeSchemaMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pkg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pkg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToVersion", wireType)
			}
			m.ToVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCodec
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)
