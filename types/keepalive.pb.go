// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: keepalive.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A KeepaliveRecord is a tuple of an entity name and the time at which the
// entity's keepalive will next expire.
type KeepaliveRecord struct {
	// Metadata contains the name (of the entity), and namespace, labels and annotations of the keepalive record
	ObjectMeta           `protobuf:"bytes,1,opt,name=metadata,embedded=metadata" json:"metadata"`
	Time                 int64    `protobuf:"varint,4,opt,name=time,proto3" json:"time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeepaliveRecord) Reset()         { *m = KeepaliveRecord{} }
func (m *KeepaliveRecord) String() string { return proto.CompactTextString(m) }
func (*KeepaliveRecord) ProtoMessage()    {}
func (*KeepaliveRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_keepalive_090392187a0c6395, []int{0}
}
func (m *KeepaliveRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeepaliveRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeepaliveRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *KeepaliveRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepaliveRecord.Merge(dst, src)
}
func (m *KeepaliveRecord) XXX_Size() int {
	return m.Size()
}
func (m *KeepaliveRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepaliveRecord.DiscardUnknown(m)
}

var xxx_messageInfo_KeepaliveRecord proto.InternalMessageInfo

func (m *KeepaliveRecord) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*KeepaliveRecord)(nil), "sensu.types.KeepaliveRecord")
}
func (this *KeepaliveRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KeepaliveRecord)
	if !ok {
		that2, ok := that.(KeepaliveRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if this.Time != that1.Time {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *KeepaliveRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeepaliveRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintKeepalive(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Time != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintKeepalive(dAtA, i, uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintKeepalive(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedKeepaliveRecord(r randyKeepalive, easy bool) *KeepaliveRecord {
	this := &KeepaliveRecord{}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	this.Time = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Time *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedKeepalive(r, 5)
	}
	return this
}

type randyKeepalive interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneKeepalive(r randyKeepalive) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringKeepalive(r randyKeepalive) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneKeepalive(r)
	}
	return string(tmps)
}
func randUnrecognizedKeepalive(r randyKeepalive, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldKeepalive(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldKeepalive(dAtA []byte, r randyKeepalive, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateKeepalive(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateKeepalive(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateKeepalive(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateKeepalive(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateKeepalive(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateKeepalive(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateKeepalive(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *KeepaliveRecord) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovKeepalive(uint64(l))
	if m.Time != 0 {
		n += 1 + sovKeepalive(uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKeepalive(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKeepalive(x uint64) (n int) {
	return sovKeepalive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeepaliveRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeepalive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeepaliveRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeepaliveRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeepalive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKeepalive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeepalive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeepalive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeepalive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKeepalive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeepalive
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
					return 0, ErrIntOverflowKeepalive
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
					return 0, ErrIntOverflowKeepalive
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthKeepalive
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKeepalive
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
				next, err := skipKeepalive(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthKeepalive = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeepalive   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("keepalive.proto", fileDescriptor_keepalive_090392187a0c6395) }

var fileDescriptor_keepalive_090392187a0c6395 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4e, 0x4d, 0x2d,
	0x48, 0xcc, 0xc9, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x4e, 0xcd,
	0x2b, 0x2e, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03,
	0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x57, 0x8a, 0x2b, 0x37, 0xb5, 0x24, 0x11, 0xc2, 0x56, 0xaa,
	0xe2, 0xe2, 0xf7, 0x86, 0x19, 0x1d, 0x94, 0x9a, 0x9c, 0x5f, 0x94, 0x22, 0xe4, 0xc9, 0xc5, 0x01,
	0x52, 0x90, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xae, 0x87, 0x64,
	0x9b, 0x9e, 0x7f, 0x52, 0x56, 0x6a, 0x72, 0x89, 0x6f, 0x6a, 0x49, 0xa2, 0x93, 0xc8, 0x89, 0x7b,
	0xf2, 0x0c, 0x17, 0xee, 0xc9, 0x33, 0xbe, 0xba, 0x27, 0x0f, 0xd7, 0x14, 0x04, 0x67, 0x09, 0xc9,
	0x70, 0xb1, 0x94, 0x64, 0xe6, 0xa6, 0x4a, 0xb0, 0x28, 0x30, 0x6a, 0x30, 0x3b, 0x71, 0xbc, 0xba,
	0x27, 0x0f, 0xe6, 0x07, 0x81, 0x49, 0x27, 0xe5, 0x1f, 0x0f, 0xe5, 0x18, 0x57, 0x3c, 0x92, 0x63,
	0xdc, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x15, 0x6c, 0x5b, 0x12, 0x1b, 0xd8, 0x9d, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x07, 0xb2, 0x7c, 0x02, 0x01, 0x00, 0x00,
}
