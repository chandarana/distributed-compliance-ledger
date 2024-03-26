// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/rejected_certificate.proto

package types

import (
	fmt "fmt"
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

type RejectedCertificate struct {
	Subject       string         `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId  string         `protobuf:"bytes,2,opt,name=subjectKeyId,proto3" json:"subjectKeyId,omitempty"`
	Certs         []*Certificate `protobuf:"bytes,3,rep,name=certs,proto3" json:"certs,omitempty"`
	SchemaVersion uint32         `protobuf:"varint,4,opt,name=schemaVersion,proto3" json:"schemaVersion,omitempty"`
}

func (m *RejectedCertificate) Reset()         { *m = RejectedCertificate{} }
func (m *RejectedCertificate) String() string { return proto.CompactTextString(m) }
func (*RejectedCertificate) ProtoMessage()    {}
func (*RejectedCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_830b292fc8b89847, []int{0}
}
func (m *RejectedCertificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RejectedCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RejectedCertificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RejectedCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RejectedCertificate.Merge(m, src)
}
func (m *RejectedCertificate) XXX_Size() int {
	return m.Size()
}
func (m *RejectedCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_RejectedCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_RejectedCertificate proto.InternalMessageInfo

func (m *RejectedCertificate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *RejectedCertificate) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *RejectedCertificate) GetCerts() []*Certificate {
	if m != nil {
		return m.Certs
	}
	return nil
}

func (m *RejectedCertificate) GetSchemaVersion() uint32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*RejectedCertificate)(nil), "zigbeealliance.distributedcomplianceledger.pki.RejectedCertificate")
}

func init() { proto.RegisterFile("pki/rejected_certificate.proto", fileDescriptor_830b292fc8b89847) }

var fileDescriptor_830b292fc8b89847 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xc8, 0xce, 0xd4,
	0x2f, 0x4a, 0xcd, 0x4a, 0x4d, 0x2e, 0x49, 0x4d, 0x89, 0x4f, 0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb,
	0x4c, 0x4e, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0xab, 0xca, 0x4c, 0x4f,
	0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4b, 0xc9, 0x2c, 0x2e, 0x29,
	0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4,
	0xa7, 0x16, 0xe9, 0x15, 0x64, 0x67, 0x4a, 0x89, 0x82, 0xcc, 0xc3, 0x30, 0x46, 0xe9, 0x0c, 0x23,
	0x97, 0x70, 0x10, 0xd4, 0x16, 0x67, 0x84, 0xac, 0x90, 0x04, 0x17, 0x7b, 0x71, 0x69, 0x12, 0x48,
	0x5c, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x52, 0xe2, 0xe2, 0x81, 0x32, 0xbd,
	0x53, 0x2b, 0x3d, 0x53, 0x24, 0x98, 0xc0, 0xd2, 0x28, 0x62, 0x42, 0x81, 0x5c, 0xac, 0x20, 0xab,
	0x8a, 0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xac, 0x49, 0x74, 0xac, 0x1e, 0x92, 0x4b, 0x82,
	0x20, 0x26, 0x09, 0xa9, 0x70, 0xf1, 0x16, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x86, 0xa5, 0x16, 0x15,
	0x67, 0xe6, 0xe7, 0x49, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x06, 0xa1, 0x0a, 0x3a, 0xc5, 0x9d, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c,
	0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x4b, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x35, 0xba, 0x30, 0xe7, 0xe8, 0x23, 0x39, 0x47, 0x17, 0xe1,
	0x1e, 0x5d, 0x88, 0x83, 0xf4, 0x2b, 0xf4, 0x41, 0x41, 0x57, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0x0e, 0x35, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x95, 0x77, 0xb5, 0x9e, 0x01,
	0x00, 0x00,
}

func (m *RejectedCertificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RejectedCertificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RejectedCertificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SchemaVersion != 0 {
		i = encodeVarintRejectedCertificate(dAtA, i, uint64(m.SchemaVersion))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Certs) > 0 {
		for iNdEx := len(m.Certs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRejectedCertificate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintRejectedCertificate(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintRejectedCertificate(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRejectedCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovRejectedCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RejectedCertificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovRejectedCertificate(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovRejectedCertificate(uint64(l))
	}
	if len(m.Certs) > 0 {
		for _, e := range m.Certs {
			l = e.Size()
			n += 1 + l + sovRejectedCertificate(uint64(l))
		}
	}
	if m.SchemaVersion != 0 {
		n += 1 + sovRejectedCertificate(uint64(m.SchemaVersion))
	}
	return n
}

func sovRejectedCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRejectedCertificate(x uint64) (n int) {
	return sovRejectedCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RejectedCertificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRejectedCertificate
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
			return fmt.Errorf("proto: RejectedCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RejectedCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedCertificate
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
				return ErrInvalidLengthRejectedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedCertificate
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
				return ErrInvalidLengthRejectedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedCertificate
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
				return ErrInvalidLengthRejectedCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certs = append(m.Certs, &Certificate{})
			if err := m.Certs[len(m.Certs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedCertificate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRejectedCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRejectedCertificate
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
func skipRejectedCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRejectedCertificate
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
					return 0, ErrIntOverflowRejectedCertificate
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
					return 0, ErrIntOverflowRejectedCertificate
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
				return 0, ErrInvalidLengthRejectedCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRejectedCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRejectedCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRejectedCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRejectedCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRejectedCertificate = fmt.Errorf("proto: unexpected end of group")
)
