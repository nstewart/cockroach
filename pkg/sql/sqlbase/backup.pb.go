// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/sql/sqlbase/backup.proto
// DO NOT EDIT!

/*
	Package sqlbase is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/sql/sqlbase/backup.proto
		cockroach/pkg/sql/sqlbase/encoded_datum.proto
		cockroach/pkg/sql/sqlbase/privilege.proto
		cockroach/pkg/sql/sqlbase/structured.proto

	It has these top-level messages:
		BackupRangeDescriptor
		BackupDescriptor
		UserPrivileges
		PrivilegeDescriptor
		ColumnType
		ForeignKeyReference
		ColumnDescriptor
		ColumnFamilyDescriptor
		InterleaveDescriptor
		IndexDescriptor
		DescriptorMutation
		TableDescriptor
		DatabaseDescriptor
		Descriptor
*/
package sqlbase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BackupRangeDescriptor represents a file that contains the diff for a key
// range between two timestamps.
type BackupRangeDescriptor struct {
	// An empty path means the range is empty.
	Path      string                                           `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	StartKey  github_com_cockroachdb_cockroach_pkg_roachpb.Key `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.Key" json:"start_key,omitempty"`
	EndKey    github_com_cockroachdb_cockroach_pkg_roachpb.Key `protobuf:"bytes,3,opt,name=end_key,json=endKey,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.Key" json:"end_key,omitempty"`
	StartTime cockroach_util_hlc.Timestamp                     `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time"`
	CRC       uint32                                           `protobuf:"varint,5,opt,name=crc,proto3" json:"crc,omitempty"`
}

func (m *BackupRangeDescriptor) Reset()                    { *m = BackupRangeDescriptor{} }
func (m *BackupRangeDescriptor) String() string            { return proto.CompactTextString(m) }
func (*BackupRangeDescriptor) ProtoMessage()               {}
func (*BackupRangeDescriptor) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{0} }

// BackupDescriptor represents a consistent snapshot of ranges.
//
// Each range snapshot includes a path to data that is a diff of the data in
// that key range between a start and end timestamp. The end timestamp of all
// ranges in a backup is the same, but the start may vary (to allow individual
// tables to be backed up on different schedules).
type BackupDescriptor struct {
	EndTime cockroach_util_hlc.Timestamp `protobuf:"bytes,1,opt,name=end_time,json=endTime" json:"end_time"`
	Ranges  []BackupRangeDescriptor      `protobuf:"bytes,2,rep,name=ranges" json:"ranges"`
	SQL     []Descriptor                 `protobuf:"bytes,3,rep,name=sql" json:"sql"`
	// TODO(dan): Consider also including total file size and per-range data and
	// file size.
	DataSize int64 `protobuf:"varint,4,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
}

func (m *BackupDescriptor) Reset()                    { *m = BackupDescriptor{} }
func (m *BackupDescriptor) String() string            { return proto.CompactTextString(m) }
func (*BackupDescriptor) ProtoMessage()               {}
func (*BackupDescriptor) Descriptor() ([]byte, []int) { return fileDescriptorBackup, []int{1} }

func init() {
	proto.RegisterType((*BackupRangeDescriptor)(nil), "cockroach.sql.sqlbase.BackupRangeDescriptor")
	proto.RegisterType((*BackupDescriptor)(nil), "cockroach.sql.sqlbase.BackupDescriptor")
}
func (m *BackupRangeDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BackupRangeDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintBackup(data, i, uint64(len(m.Path)))
		i += copy(data[i:], m.Path)
	}
	if len(m.StartKey) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintBackup(data, i, uint64(len(m.StartKey)))
		i += copy(data[i:], m.StartKey)
	}
	if len(m.EndKey) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintBackup(data, i, uint64(len(m.EndKey)))
		i += copy(data[i:], m.EndKey)
	}
	data[i] = 0x22
	i++
	i = encodeVarintBackup(data, i, uint64(m.StartTime.Size()))
	n1, err := m.StartTime.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.CRC != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintBackup(data, i, uint64(m.CRC))
	}
	return i, nil
}

func (m *BackupDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BackupDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintBackup(data, i, uint64(m.EndTime.Size()))
	n2, err := m.EndTime.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Ranges) > 0 {
		for _, msg := range m.Ranges {
			data[i] = 0x12
			i++
			i = encodeVarintBackup(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.SQL) > 0 {
		for _, msg := range m.SQL {
			data[i] = 0x1a
			i++
			i = encodeVarintBackup(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.DataSize != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintBackup(data, i, uint64(m.DataSize))
	}
	return i, nil
}

func encodeFixed64Backup(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Backup(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBackup(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *BackupRangeDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	l = len(m.StartKey)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	l = len(m.EndKey)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	l = m.StartTime.Size()
	n += 1 + l + sovBackup(uint64(l))
	if m.CRC != 0 {
		n += 1 + sovBackup(uint64(m.CRC))
	}
	return n
}

func (m *BackupDescriptor) Size() (n int) {
	var l int
	_ = l
	l = m.EndTime.Size()
	n += 1 + l + sovBackup(uint64(l))
	if len(m.Ranges) > 0 {
		for _, e := range m.Ranges {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if len(m.SQL) > 0 {
		for _, e := range m.SQL {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if m.DataSize != 0 {
		n += 1 + sovBackup(uint64(m.DataSize))
	}
	return n
}

func sovBackup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBackup(x uint64) (n int) {
	return sovBackup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupRangeDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupRangeDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupRangeDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = append(m.StartKey[:0], data[iNdEx:postIndex]...)
			if m.StartKey == nil {
				m.StartKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndKey = append(m.EndKey[:0], data[iNdEx:postIndex]...)
			if m.EndKey == nil {
				m.EndKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartTime.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CRC", wireType)
			}
			m.CRC = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CRC |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func (m *BackupDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndTime.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ranges = append(m.Ranges, BackupRangeDescriptor{})
			if err := m.Ranges[len(m.Ranges)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SQL", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SQL = append(m.SQL, Descriptor{})
			if err := m.SQL[len(m.SQL)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSize", wireType)
			}
			m.DataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DataSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
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
func skipBackup(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBackup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipBackup(data[start:])
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
	ErrInvalidLengthBackup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBackup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/sql/sqlbase/backup.proto", fileDescriptorBackup) }

var fileDescriptorBackup = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x27, 0xd3, 0x71, 0xfe, 0x64, 0x14, 0x24, 0xb8, 0x50, 0x47, 0xec, 0x74, 0xf7, 0x20,
	0x45, 0x24, 0x91, 0xd5, 0xb3, 0x48, 0xd7, 0x93, 0xab, 0x87, 0xcd, 0x7a, 0xf2, 0xb2, 0xa4, 0x69,
	0x68, 0xcb, 0x74, 0x26, 0x9d, 0x24, 0x3d, 0xcc, 0x7e, 0x0a, 0x3f, 0x88, 0x1f, 0x64, 0x8e, 0x1e,
	0x3d, 0x0d, 0x5a, 0xbf, 0x85, 0x20, 0x48, 0xd2, 0xba, 0xae, 0xb2, 0x03, 0x82, 0x87, 0xc2, 0xcb,
	0x4b, 0xde, 0xdf, 0xf3, 0x3e, 0x4f, 0x5f, 0xf8, 0x88, 0x4b, 0xbe, 0x50, 0x92, 0xf1, 0x9c, 0x54,
	0x8b, 0x8c, 0xe8, 0x75, 0x69, 0xbf, 0x84, 0x69, 0x41, 0x12, 0xc6, 0x17, 0x75, 0x85, 0x2b, 0x25,
	0x8d, 0x44, 0x07, 0x57, 0xef, 0xb0, 0x5e, 0x97, 0xb8, 0x7b, 0x33, 0x7b, 0xbc, 0x7f, 0x5c, 0x1b,
	0x55, 0x73, 0x53, 0x2b, 0x91, 0xb6, 0x88, 0xd9, 0x5f, 0x52, 0xb5, 0x29, 0x4a, 0x92, 0x97, 0x9c,
	0x98, 0x62, 0x29, 0xb4, 0x61, 0xcb, 0x4e, 0x6a, 0x76, 0x2f, 0x93, 0x99, 0x74, 0x25, 0xb1, 0x55,
	0xdb, 0x3d, 0xfa, 0xd8, 0x87, 0x07, 0xb1, 0xdb, 0x88, 0xb2, 0x55, 0x26, 0x5e, 0x09, 0xcd, 0x55,
	0x51, 0x19, 0xa9, 0x10, 0x82, 0x83, 0x8a, 0x99, 0xdc, 0x07, 0x21, 0x88, 0x26, 0xd4, 0xd5, 0xe8,
	0x0c, 0x4e, 0xb4, 0x61, 0xca, 0x5c, 0x2c, 0xc4, 0xc6, 0xef, 0x87, 0x20, 0xba, 0x1d, 0x3f, 0xff,
	0xbe, 0x9b, 0x3f, 0xcd, 0x0a, 0x93, 0xd7, 0x09, 0xe6, 0x72, 0x49, 0xae, 0xb6, 0x49, 0x13, 0xf2,
	0xe7, 0x66, 0xae, 0xaa, 0x12, 0x7c, 0x2a, 0x36, 0x74, 0xec, 0x30, 0xa7, 0x62, 0x83, 0xde, 0xc2,
	0x91, 0x58, 0xa5, 0x0e, 0xe8, 0xfd, 0x07, 0x70, 0x28, 0x56, 0xa9, 0xc5, 0xc5, 0x10, 0xb6, 0x1b,
	0x5a, 0xfb, 0xfe, 0x20, 0x04, 0xd1, 0xf4, 0xf8, 0x21, 0xfe, 0x9d, 0xb2, 0x8d, 0x07, 0xe7, 0x25,
	0xc7, 0xef, 0x7e, 0xc5, 0x13, 0x0f, 0xb6, 0xbb, 0x79, 0x8f, 0xb6, 0xc6, 0x6c, 0x17, 0xdd, 0x87,
	0x1e, 0x57, 0xdc, 0xbf, 0x15, 0x82, 0xe8, 0x4e, 0x3c, 0x6a, 0x76, 0x73, 0xef, 0x84, 0x9e, 0x50,
	0xdb, 0x3b, 0xfa, 0x01, 0xe0, 0xdd, 0x36, 0xae, 0x6b, 0x49, 0xbd, 0x80, 0x63, 0x6b, 0xc1, 0x29,
	0x82, 0x7f, 0x57, 0xb4, 0xbe, 0x9d, 0xde, 0x6b, 0x38, 0x54, 0x36, 0x7c, 0xed, 0xf7, 0x43, 0x2f,
	0x9a, 0x1e, 0x3f, 0xc1, 0x37, 0x5e, 0x05, 0xbe, 0xf1, 0x3f, 0x75, 0xb0, 0x8e, 0x80, 0x5e, 0x42,
	0x4f, 0xaf, 0x4b, 0xdf, 0x73, 0xa0, 0xc3, 0x3d, 0xa0, 0x6b, 0xd3, 0x53, 0x3b, 0x6d, 0x2d, 0x9e,
	0x9f, 0xbd, 0xa1, 0x76, 0x14, 0x3d, 0x80, 0x93, 0x94, 0x19, 0x76, 0xa1, 0x8b, 0xcb, 0x36, 0x40,
	0x8f, 0x8e, 0x6d, 0xe3, 0xbc, 0xb8, 0x14, 0xf1, 0xe1, 0xf6, 0x6b, 0xd0, 0xdb, 0x36, 0x01, 0xf8,
	0xd4, 0x04, 0xe0, 0x73, 0x13, 0x80, 0x2f, 0x4d, 0x00, 0x3e, 0x7c, 0x0b, 0x7a, 0xef, 0x47, 0x1d,
	0x3c, 0x19, 0xba, 0xc3, 0x7a, 0xf6, 0x33, 0x00, 0x00, 0xff, 0xff, 0xac, 0xd5, 0x2b, 0x3f, 0x03,
	0x03, 0x00, 0x00,
}
