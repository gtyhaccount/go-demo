// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payload.proto

/*
	Package message is a generated protocol buffer package.

	It is generated from these files:
		payload.proto

	It has these top-level messages:
		Payload
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Payload is the message payload saves to db/message table and transmits on wire
type Payload struct {
	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Offset    int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type      int32  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Subtype   int32  `protobuf:"varint,6,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Body      string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	Extra     string `protobuf:"bytes,8,opt,name=extra,proto3" json:"extra,omitempty"`
	SenderId  string `protobuf:"bytes,9,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ExpiresAt int64  `protobuf:"varint,10,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{0} }

func (m *Payload) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Payload) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Payload) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Payload) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Payload) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Payload) GetSubtype() int32 {
	if m != nil {
		return m.Subtype
	}
	return 0
}

func (m *Payload) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Payload) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *Payload) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

func (m *Payload) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Payload)(nil), "message.Payload")
}
func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Id))
	}
	if len(m.Topic) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Topic)))
		i += copy(dAtA[i:], m.Topic)
	}
	if m.Offset != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Offset))
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Timestamp))
	}
	if m.Type != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Type))
	}
	if m.Subtype != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Subtype))
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if len(m.Extra) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Extra)))
		i += copy(dAtA[i:], m.Extra)
	}
	if len(m.SenderId) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.SenderId)))
		i += copy(dAtA[i:], m.SenderId)
	}
	if m.ExpiresAt != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.ExpiresAt))
	}
	return i, nil
}

func encodeFixed64Payload(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Payload(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Payload) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPayload(uint64(m.Id))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovPayload(uint64(m.Offset))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPayload(uint64(m.Timestamp))
	}
	if m.Type != 0 {
		n += 1 + sovPayload(uint64(m.Type))
	}
	if m.Subtype != 0 {
		n += 1 + sovPayload(uint64(m.Subtype))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	l = len(m.SenderId)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovPayload(uint64(m.ExpiresAt))
	}
	return n
}

func sovPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
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
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subtype", wireType)
			}
			m.Subtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subtype |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
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
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
				return 0, ErrInvalidLengthPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayload
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
				next, err := skipPayload(dAtA[start:])
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
	ErrInvalidLengthPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("payload.proto", fileDescriptorPayload) }

var fileDescriptorPayload = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x4d, 0x67, 0xda, 0x4e, 0x1f, 0x28, 0xf2, 0x10, 0x79, 0xa0, 0x96, 0xe2, 0xaa, 0x2b,
	0x37, 0x9e, 0x40, 0x77, 0xee, 0xa4, 0x17, 0x18, 0x52, 0xf3, 0x46, 0x02, 0xd6, 0x84, 0x26, 0xc2,
	0xf4, 0x26, 0x1e, 0xc9, 0xa5, 0x47, 0x90, 0x7a, 0x0a, 0x77, 0x32, 0x2f, 0x15, 0x77, 0xef, 0xfb,
	0xfe, 0x1f, 0xf2, 0x13, 0x38, 0xf6, 0x7a, 0x7a, 0x71, 0xda, 0xdc, 0xf8, 0xd1, 0x45, 0x87, 0xe5,
	0xc0, 0x21, 0xe8, 0x67, 0xbe, 0xfe, 0x51, 0x50, 0x3e, 0xa6, 0x08, 0x4f, 0x20, 0xb3, 0x86, 0x54,
	0xa3, 0xda, 0x55, 0x97, 0x59, 0x83, 0x67, 0x90, 0x47, 0xe7, 0xed, 0x13, 0x65, 0x8d, 0x6a, 0xab,
	0x2e, 0x01, 0x9e, 0x43, 0xe1, 0x76, 0xbb, 0xc0, 0x91, 0x56, 0xd2, 0x5c, 0x08, 0x2f, 0xa1, 0x8a,
	0x76, 0xe0, 0x10, 0xf5, 0xe0, 0x69, 0x2d, 0xd1, 0xbf, 0x40, 0x84, 0x75, 0x9c, 0x3c, 0x53, 0xde,
	0xa8, 0x36, 0xef, 0xe4, 0x46, 0x82, 0x32, 0xbc, 0xf5, 0xa2, 0x0b, 0xd1, 0x7f, 0x78, 0x68, 0xf7,
	0xce, 0x4c, 0x54, 0xca, 0xc3, 0x72, 0x1f, 0xd6, 0xf0, 0x3e, 0x8e, 0x9a, 0x36, 0x69, 0x8d, 0x00,
	0x5e, 0x40, 0x15, 0xf8, 0xd5, 0xf0, 0xb8, 0xb5, 0x86, 0x2a, 0x49, 0x36, 0x49, 0x3c, 0x18, 0xbc,
	0x02, 0xe0, 0xbd, 0xb7, 0x23, 0x87, 0xad, 0x8e, 0x04, 0x69, 0xd3, 0x62, 0xee, 0xe2, 0xfd, 0xe9,
	0xc7, 0x5c, 0xab, 0xcf, 0xb9, 0x56, 0x5f, 0x73, 0xad, 0xde, 0xbf, 0xeb, 0xa3, 0xbe, 0x90, 0xdf,
	0xb9, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x15, 0x06, 0x32, 0xb6, 0x2e, 0x01, 0x00, 0x00,
}
