// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft_transfer/v1/trace.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClassTrace contains the base classID for ICS721 non-fungible tokens and the
// source tracing information path.
type ClassTrace struct {
	// path defines the chain of port/channel identifiers used for tracing the
	// source of the non-fungible token.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// base classID of the relayed non-fungible token.
	BaseClassId string `protobuf:"bytes,2,opt,name=base_class_id,json=baseClassId,proto3" json:"base_class_id,omitempty"`
}

func (m *ClassTrace) Reset()         { *m = ClassTrace{} }
func (m *ClassTrace) String() string { return proto.CompactTextString(m) }
func (*ClassTrace) ProtoMessage()    {}
func (*ClassTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4e6ac472424735f, []int{0}
}
func (m *ClassTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassTrace.Merge(m, src)
}
func (m *ClassTrace) XXX_Size() int {
	return m.Size()
}
func (m *ClassTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ClassTrace proto.InternalMessageInfo

func (m *ClassTrace) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ClassTrace) GetBaseClassId() string {
	if m != nil {
		return m.BaseClassId
	}
	return ""
}

func init() {
	proto.RegisterType((*ClassTrace)(nil), "nft_transfer.v1.ClassTrace")
}

func init() { proto.RegisterFile("nft_transfer/v1/trace.proto", fileDescriptor_f4e6ac472424735f) }

var fileDescriptor_f4e6ac472424735f = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x4b, 0x2b, 0x89,
	0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x29, 0x4a, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x96, 0xd4, 0x2b, 0x33, 0x54, 0x72,
	0xe1, 0xe2, 0x72, 0xce, 0x49, 0x2c, 0x2e, 0x0e, 0x01, 0x29, 0x12, 0x12, 0xe2, 0x62, 0x29, 0x48,
	0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x94, 0xb8, 0x78, 0x93,
	0x12, 0x8b, 0x53, 0xe3, 0x93, 0x41, 0xca, 0xe2, 0x33, 0x53, 0x24, 0x98, 0xc0, 0x92, 0xdc, 0x20,
	0x41, 0xb0, 0x56, 0xcf, 0x14, 0xa7, 0xd0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2e, 0xaa,
	0x2c, 0x28, 0xc9, 0xd7, 0xcd, 0x2f, 0x4a, 0xd7, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0x93,
	0xba, 0xb9, 0x20, 0x66, 0x85, 0x7e, 0x5e, 0x5a, 0x89, 0x2e, 0xdc, 0xd9, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x47, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x4f, 0x43, 0x48,
	0xd3, 0x00, 0x00, 0x00,
}

func (m *ClassTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BaseClassId) > 0 {
		i -= len(m.BaseClassId)
		copy(dAtA[i:], m.BaseClassId)
		i = encodeVarintTrace(dAtA, i, uint64(len(m.BaseClassId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTrace(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrace(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	l = len(m.BaseClassId)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	return n
}

func sovTrace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrace(x uint64) (n int) {
	return sovTrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: ClassTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTrace
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
func skipTrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrace
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
					return 0, ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrace
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
				return 0, ErrInvalidLengthTrace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTrace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTrace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTrace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTrace = fmt.Errorf("proto: unexpected end of group")
)
