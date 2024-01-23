// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dclgenutil/genesis.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the dclgenutil module's genesis state.
type GenesisState struct {
	// gen_txs defines the genesis transactions.
	GenTxs []encoding_json.RawMessage `protobuf:"bytes,1,rep,name=gen_txs,json=genTxs,proto3,casttype=encoding/json.RawMessage" json:"gentxs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_72454a08e2a56bab, []int{0}
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

func (m *GenesisState) GetGenTxs() []encoding_json.RawMessage {
	if m != nil {
		return m.GenTxs
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zigbeealliance.distributedcomplianceledger.dclgenutil.GenesisState")
}

func init() { proto.RegisterFile("dclgenutil/genesis.proto", fileDescriptor_72454a08e2a56bab) }

var fileDescriptor_72454a08e2a56bab = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc0, 0xf1, 0x1e, 0x42, 0x85, 0xa3, 0x53, 0x71, 0x38, 0x1c, 0x52, 0x71, 0x10, 0x97, 0x4b,
	0x06, 0x71, 0x16, 0xba, 0x38, 0x88, 0x08, 0xd5, 0xc9, 0x45, 0x72, 0xc9, 0xe3, 0x19, 0x49, 0x93,
	0xe3, 0xde, 0x3b, 0x3c, 0xfd, 0x14, 0x7e, 0x2c, 0xc7, 0x8e, 0x4e, 0x45, 0xee, 0x36, 0x3f, 0x82,
	0x93, 0xb4, 0x41, 0xda, 0x2d, 0x84, 0xc7, 0x9f, 0xf7, 0x7b, 0x79, 0x61, 0x8d, 0x47, 0x08, 0x2d,
	0x3b, 0xaf, 0x10, 0x02, 0x90, 0x23, 0x59, 0x37, 0x91, 0xe3, 0xf4, 0xf2, 0xdd, 0x61, 0x05, 0xa0,
	0xbd, 0x77, 0x3a, 0x18, 0x90, 0xd6, 0x11, 0x37, 0xae, 0x6a, 0x19, 0xac, 0x89, 0xcb, 0x3a, 0xfd,
	0x7a, 0xb0, 0x08, 0x8d, 0xdc, 0x45, 0x8e, 0x8f, 0x30, 0x62, 0xdc, 0x16, 0xd4, 0xe6, 0x95, 0x62,
	0xa7, 0x77, 0xf9, 0xe4, 0x3a, 0xd5, 0xef, 0x59, 0x33, 0x4c, 0xaf, 0xf2, 0x43, 0x84, 0xf0, 0xc4,
	0x1d, 0x15, 0xd9, 0xc9, 0xc1, 0xf9, 0x64, 0x7e, 0xf6, 0xb3, 0x9e, 0x8d, 0x11, 0x02, 0x77, 0xf4,
	0xbb, 0x9e, 0x15, 0x10, 0x4c, 0xb4, 0x2e, 0xa0, 0x7a, 0xa1, 0x18, 0xe4, 0x42, 0xbf, 0xde, 0x02,
	0x91, 0x46, 0x58, 0x6c, 0x66, 0x1e, 0x3a, 0x9a, 0xc3, 0x67, 0x2f, 0xb2, 0x55, 0x2f, 0xb2, 0xef,
	0x5e, 0x64, 0x1f, 0x83, 0x18, 0xad, 0x06, 0x31, 0xfa, 0x1a, 0xc4, 0xe8, 0xf1, 0x06, 0x1d, 0x3f,
	0xb7, 0x95, 0x34, 0x71, 0xa9, 0x12, 0xa1, 0xfc, 0x37, 0xa8, 0x3d, 0x43, 0xb9, 0x43, 0x94, 0x49,
	0xa1, 0x3a, 0xb5, 0x77, 0x0c, 0x7e, 0xab, 0x81, 0xaa, 0xf1, 0x76, 0xfd, 0x8b, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x14, 0xce, 0x8b, 0x18, 0x27, 0x01, 0x00, 0x00,
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
	if len(m.GenTxs) > 0 {
		for iNdEx := len(m.GenTxs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GenTxs[iNdEx])
			copy(dAtA[i:], m.GenTxs[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.GenTxs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.GenTxs) > 0 {
		for _, b := range m.GenTxs {
			l = len(b)
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
				return fmt.Errorf("proto: wrong wireType = %d for field GenTxs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenTxs = append(m.GenTxs, make([]byte, postIndex-iNdEx))
			copy(m.GenTxs[len(m.GenTxs)-1], dAtA[iNdEx:postIndex])
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
