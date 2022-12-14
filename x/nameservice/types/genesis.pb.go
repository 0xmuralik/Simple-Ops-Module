// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vulcanize/nameservice/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the nameservice module's genesis state.
type GenesisState struct {
	// params defines all the params of nameservice module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// records
	Records []Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records" json:"records" yaml:"records"`
	// authorities
	Authorities []AuthorityEntry `protobuf:"bytes,3,rep,name=authorities,proto3" json:"authorities" json:"authorities" yaml:"authorities"`
	// names
	Names []NameEntry `protobuf:"bytes,4,rep,name=names,proto3" json:"names" json:"names" yaml:"names"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe7037a2b22e67ef, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRecords() []Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *GenesisState) GetAuthorities() []AuthorityEntry {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func (m *GenesisState) GetNames() []NameEntry {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "vulcanize.nameservice.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("vulcanize/nameservice/v1beta1/genesis.proto", fileDescriptor_fe7037a2b22e67ef)
}

var fileDescriptor_fe7037a2b22e67ef = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0xb1, 0x4a, 0xfb, 0x40,
	0x1c, 0x07, 0xf0, 0xe4, 0xdf, 0xfe, 0x2b, 0xa4, 0x4e, 0xc1, 0x21, 0x16, 0x9a, 0xd6, 0x42, 0xa1,
	0x20, 0xcd, 0x59, 0xdd, 0xdc, 0x8c, 0x88, 0xe0, 0x20, 0x12, 0x37, 0xb7, 0x6b, 0xfc, 0x99, 0x9c,
	0x34, 0xb9, 0x72, 0xf7, 0x6b, 0x31, 0x0e, 0x3e, 0x83, 0x4f, 0xe0, 0xf3, 0x74, 0xec, 0xe8, 0x54,
	0xa4, 0x7d, 0x03, 0x9f, 0x40, 0x7a, 0x97, 0x68, 0x5c, 0x5a, 0xb7, 0x5c, 0xf8, 0x7e, 0xbf, 0x9f,
	0x40, 0xce, 0x3a, 0x9c, 0x4e, 0x46, 0x21, 0x4d, 0xd9, 0x33, 0x90, 0x94, 0x26, 0x20, 0x41, 0x4c,
	0x59, 0x08, 0x64, 0x3a, 0x18, 0x02, 0xd2, 0x01, 0x89, 0x20, 0x05, 0xc9, 0xa4, 0x37, 0x16, 0x1c,
	0xb9, 0xdd, 0xfc, 0x0e, 0x7b, 0xa5, 0xb0, 0x97, 0x87, 0x1b, 0x7b, 0x11, 0x8f, 0xb8, 0x4a, 0x92,
	0xf5, 0x93, 0x2e, 0x35, 0xc8, 0x66, 0xa1, 0x3c, 0xa4, 0x0a, 0x9d, 0xb7, 0x8a, 0xb5, 0x7b, 0xa9,
	0xdd, 0x5b, 0xa4, 0x08, 0xf6, 0xb9, 0x55, 0x1b, 0x53, 0x41, 0x13, 0xe9, 0x98, 0x6d, 0xb3, 0x57,
	0x3f, 0xee, 0x7a, 0x1b, 0xbf, 0xc3, 0xbb, 0x51, 0x61, 0xbf, 0x3a, 0x5b, 0xb4, 0x8c, 0x20, 0xaf,
	0xda, 0x0f, 0xd6, 0x8e, 0x80, 0x90, 0x8b, 0x7b, 0xe9, 0xfc, 0x6b, 0x57, 0xfe, 0xb0, 0x12, 0xa8,
	0xb4, 0xdf, 0x5d, 0xaf, 0x7c, 0x2e, 0x5a, 0xcd, 0x47, 0xc9, 0xd3, 0xd3, 0x4e, 0xbe, 0xd1, 0x69,
	0x67, 0x34, 0x19, 0xfd, 0x1c, 0x83, 0x62, 0xdc, 0x7e, 0xb1, 0xea, 0x74, 0x82, 0x31, 0x17, 0x0c,
	0x19, 0x48, 0xa7, 0xa2, 0xac, 0xfe, 0x16, 0xeb, 0x2c, 0x6f, 0x64, 0x17, 0x29, 0x8a, 0xcc, 0xef,
	0xe7, 0x66, 0x57, 0x9b, 0xa5, 0xbd, 0xc2, 0x2d, 0xbf, 0x0a, 0xca, 0xa0, 0x4d, 0xad, 0xff, 0x4a,
	0x70, 0xaa, 0x4a, 0xee, 0x6d, 0x91, 0xaf, 0x69, 0x02, 0x1a, 0x3d, 0xc8, 0xd1, 0x7d, 0x8d, 0xaa,
	0x70, 0xc1, 0xe9, 0x43, 0xa0, 0x97, 0xfd, 0xab, 0xd9, 0xd2, 0x35, 0xe7, 0x4b, 0xd7, 0xfc, 0x58,
	0xba, 0xe6, 0xeb, 0xca, 0x35, 0xe6, 0x2b, 0xd7, 0x78, 0x5f, 0xb9, 0xc6, 0xdd, 0x51, 0xc4, 0x30,
	0x9e, 0x0c, 0xbd, 0x90, 0x27, 0x04, 0x63, 0x2a, 0x24, 0x93, 0x04, 0x30, 0x06, 0x91, 0xb0, 0x14,
	0xc9, 0xd3, 0xaf, 0x0b, 0x80, 0xd9, 0x18, 0xe4, 0xb0, 0xa6, 0xfe, 0xf9, 0xc9, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbe, 0x6b, 0x6e, 0x5b, 0x88, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Names) > 0 {
		for iNdEx := len(m.Names) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Names[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Authorities) > 0 {
		for iNdEx := len(m.Authorities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authorities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Authorities) > 0 {
		for _, e := range m.Authorities {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Names) > 0 {
		for _, e := range m.Names {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authorities = append(m.Authorities, AuthorityEntry{})
			if err := m.Authorities[len(m.Authorities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Names", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Names = append(m.Names, NameEntry{})
			if err := m.Names[len(m.Names)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
