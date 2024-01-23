// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model/vendor_products.proto

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

type VendorProducts struct {
	Vid      int32      `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (m *VendorProducts) Reset()         { *m = VendorProducts{} }
func (m *VendorProducts) String() string { return proto.CompactTextString(m) }
func (*VendorProducts) ProtoMessage()    {}
func (*VendorProducts) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f57b2346708f393, []int{0}
}
func (m *VendorProducts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VendorProducts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VendorProducts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VendorProducts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VendorProducts.Merge(m, src)
}
func (m *VendorProducts) XXX_Size() int {
	return m.Size()
}
func (m *VendorProducts) XXX_DiscardUnknown() {
	xxx_messageInfo_VendorProducts.DiscardUnknown(m)
}

var xxx_messageInfo_VendorProducts proto.InternalMessageInfo

func (m *VendorProducts) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *VendorProducts) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*VendorProducts)(nil), "zigbeealliance.distributedcomplianceledger.model.VendorProducts")
}

func init() { proto.RegisterFile("model/vendor_products.proto", fileDescriptor_1f57b2346708f393) }

var fileDescriptor_1f57b2346708f393 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2f, 0x4b, 0xcd, 0x4b, 0xc9, 0x2f, 0x8a, 0x2f, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0xa8, 0xca, 0x4c, 0x4f, 0x4a, 0x4d, 0x4d,
	0xcc, 0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4b, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a,
	0x2d, 0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xe9,
	0x81, 0xcd, 0x91, 0x12, 0x86, 0x18, 0x07, 0x35, 0x07, 0x62, 0x8c, 0x52, 0x25, 0x17, 0x5f, 0x18,
	0xd8, 0xfc, 0x00, 0xa8, 0xf1, 0x42, 0x02, 0x5c, 0xcc, 0x65, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x20, 0xa6, 0x50, 0x28, 0x17, 0x07, 0xcc, 0x72, 0x09, 0x26, 0x05, 0x66, 0x0d,
	0x6e, 0x23, 0x4b, 0x3d, 0x52, 0x6d, 0xd7, 0x83, 0x9a, 0x1f, 0x04, 0x37, 0xca, 0x29, 0xe1, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xdc, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0x16, 0xe9, 0xc2, 0x6c, 0xd2, 0x47, 0xb2, 0x49, 0x17, 0x61,
	0x95, 0x2e, 0xc4, 0x2e, 0xfd, 0x0a, 0x7d, 0x88, 0x27, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0x7e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x20, 0xa1, 0xf6, 0xb1, 0x49, 0x01, 0x00,
	0x00,
}

func (m *VendorProducts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VendorProducts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VendorProducts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Products) > 0 {
		for iNdEx := len(m.Products) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Products[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVendorProducts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Vid != 0 {
		i = encodeVarintVendorProducts(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVendorProducts(dAtA []byte, offset int, v uint64) int {
	offset -= sovVendorProducts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VendorProducts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovVendorProducts(uint64(m.Vid))
	}
	if len(m.Products) > 0 {
		for _, e := range m.Products {
			l = e.Size()
			n += 1 + l + sovVendorProducts(uint64(l))
		}
	}
	return n
}

func sovVendorProducts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVendorProducts(x uint64) (n int) {
	return sovVendorProducts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VendorProducts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVendorProducts
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
			return fmt.Errorf("proto: VendorProducts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VendorProducts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorProducts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Products", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorProducts
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
				return ErrInvalidLengthVendorProducts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVendorProducts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Products = append(m.Products, &Product{})
			if err := m.Products[len(m.Products)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVendorProducts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVendorProducts
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
func skipVendorProducts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVendorProducts
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
					return 0, ErrIntOverflowVendorProducts
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
					return 0, ErrIntOverflowVendorProducts
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
				return 0, ErrInvalidLengthVendorProducts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVendorProducts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVendorProducts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVendorProducts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVendorProducts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVendorProducts = fmt.Errorf("proto: unexpected end of group")
)
