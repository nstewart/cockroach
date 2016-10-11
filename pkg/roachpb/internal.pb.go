// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/roachpb/internal.proto
// DO NOT EDIT!

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// InternalTimeSeriesData is a collection of data samples for some
// measurable value, where each sample is taken over a uniform time
// interval.
//
// The collection itself contains a start timestamp (in seconds since the unix
// epoch) and a sample duration (in milliseconds). Each sample in the collection
// will contain a positive integer offset that indicates the length of time
// between the start_timestamp of the collection and the time when the sample
// began, expressed as an whole number of sample intervals. For example, if the
// sample duration is 60000 (indicating 1 minute), then a contained sample with
// an offset value of 5 begins (5*60000ms = 300000ms = 5 minutes) after the
// start timestamp of this data.
//
// This is meant to be an efficient internal representation of time series data,
// ensuring that very little redundant data is stored on disk. With this goal in
// mind, this message does not identify the variable which is actually being
// measured; that information is expected be encoded in the key where this
// message is stored.
type InternalTimeSeriesData struct {
	// Holds a wall time, expressed as a unix epoch time in nanoseconds. This
	// represents the earliest possible timestamp for a sample within the
	// collection.
	StartTimestampNanos int64 `protobuf:"varint,1,opt,name=start_timestamp_nanos,json=startTimestampNanos" json:"start_timestamp_nanos"`
	// The duration of each sample interval, expressed in nanoseconds.
	SampleDurationNanos int64 `protobuf:"varint,2,opt,name=sample_duration_nanos,json=sampleDurationNanos" json:"sample_duration_nanos"`
	// The actual data samples for this metric.
	Samples []InternalTimeSeriesSample `protobuf:"bytes,3,rep,name=samples" json:"samples"`
}

func (m *InternalTimeSeriesData) Reset()                    { *m = InternalTimeSeriesData{} }
func (m *InternalTimeSeriesData) String() string            { return proto.CompactTextString(m) }
func (*InternalTimeSeriesData) ProtoMessage()               {}
func (*InternalTimeSeriesData) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

// A InternalTimeSeriesSample represents data gathered from multiple
// measurements of a variable value over a given period of time. The
// length of that period of time is stored in an
// InternalTimeSeriesData message; a sample cannot be interpreted
// correctly without a start timestamp and sample duration.
//
// Each sample may contain data gathered from multiple measurements of the same
// variable, as long as all of those measurements occurred within the sample
// period. The sample stores several aggregated values from these measurements:
// - The sum of all measured values
// - A count of all measurements taken
// - The maximum individual measurement seen
// - The minimum individual measurement seen
//
// If zero measurements are present in a sample, then it should be omitted
// entirely from any collection it would be a part of.
//
// If the count of measurements is 1, then max and min fields may be omitted
// and assumed equal to the sum field.
type InternalTimeSeriesSample struct {
	// Temporal offset from the "start_timestamp" of the InternalTimeSeriesData
	// collection this data point is part in. The units of this value are
	// determined by the value of the "sample_duration_milliseconds" field of
	// the TimeSeriesData collection.
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// Sum of all measurements.
	Sum float64 `protobuf:"fixed64,7,opt,name=sum" json:"sum"`
	// Count of measurements taken within this sample.
	Count uint32 `protobuf:"varint,6,opt,name=count" json:"count"`
	// Maximum encountered measurement in this sample.
	Max *float64 `protobuf:"fixed64,8,opt,name=max" json:"max,omitempty"`
	// Minimum encountered measurement in this sample.
	Min *float64 `protobuf:"fixed64,9,opt,name=min" json:"min,omitempty"`
}

func (m *InternalTimeSeriesSample) Reset()                    { *m = InternalTimeSeriesSample{} }
func (m *InternalTimeSeriesSample) String() string            { return proto.CompactTextString(m) }
func (*InternalTimeSeriesSample) ProtoMessage()               {}
func (*InternalTimeSeriesSample) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func init() {
	proto.RegisterType((*InternalTimeSeriesData)(nil), "cockroach.roachpb.InternalTimeSeriesData")
	proto.RegisterType((*InternalTimeSeriesSample)(nil), "cockroach.roachpb.InternalTimeSeriesSample")
}
func (m *InternalTimeSeriesData) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InternalTimeSeriesData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternal(data, i, uint64(m.StartTimestampNanos))
	data[i] = 0x10
	i++
	i = encodeVarintInternal(data, i, uint64(m.SampleDurationNanos))
	if len(m.Samples) > 0 {
		for _, msg := range m.Samples {
			data[i] = 0x1a
			i++
			i = encodeVarintInternal(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *InternalTimeSeriesSample) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InternalTimeSeriesSample) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternal(data, i, uint64(m.Offset))
	data[i] = 0x30
	i++
	i = encodeVarintInternal(data, i, uint64(m.Count))
	data[i] = 0x39
	i++
	i = encodeFixed64Internal(data, i, uint64(math.Float64bits(float64(m.Sum))))
	if m.Max != nil {
		data[i] = 0x41
		i++
		i = encodeFixed64Internal(data, i, uint64(math.Float64bits(float64(*m.Max))))
	}
	if m.Min != nil {
		data[i] = 0x49
		i++
		i = encodeFixed64Internal(data, i, uint64(math.Float64bits(float64(*m.Min))))
	}
	return i, nil
}

func encodeFixed64Internal(data []byte, offset int, v uint64) int {
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
func encodeFixed32Internal(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInternal(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *InternalTimeSeriesData) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.StartTimestampNanos))
	n += 1 + sovInternal(uint64(m.SampleDurationNanos))
	if len(m.Samples) > 0 {
		for _, e := range m.Samples {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func (m *InternalTimeSeriesSample) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.Offset))
	n += 1 + sovInternal(uint64(m.Count))
	n += 9
	if m.Max != nil {
		n += 9
	}
	if m.Min != nil {
		n += 9
	}
	return n
}

func sovInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternal(x uint64) (n int) {
	return sovInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalTimeSeriesData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: InternalTimeSeriesData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalTimeSeriesData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimestampNanos", wireType)
			}
			m.StartTimestampNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StartTimestampNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SampleDurationNanos", wireType)
			}
			m.SampleDurationNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.SampleDurationNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Samples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Samples = append(m.Samples, InternalTimeSeriesSample{})
			if err := m.Samples[len(m.Samples)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *InternalTimeSeriesSample) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: InternalTimeSeriesSample: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalTimeSeriesSample: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Sum = float64(math.Float64frombits(v))
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Max = &v2
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Min = &v2
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func skipInternal(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
				return 0, ErrInvalidLengthInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternal
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
				next, err := skipInternal(data[start:])
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
	ErrInvalidLengthInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/roachpb/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0x80, 0x97, 0xaf, 0xdf, 0x36, 0x8d, 0x08, 0x1a, 0x75, 0x84, 0x21, 0xb1, 0x0e, 0x0f, 0x05,
	0xa1, 0x05, 0x4f, 0x9e, 0xc7, 0x2e, 0x22, 0x78, 0xd8, 0x76, 0xf2, 0x52, 0x62, 0xcd, 0x6a, 0xd8,
	0x92, 0x94, 0x24, 0x05, 0x7f, 0xc6, 0x7e, 0x56, 0x8f, 0x1e, 0xf5, 0x22, 0x5a, 0xff, 0x88, 0x34,
	0xcd, 0x04, 0x19, 0xde, 0x5e, 0x9e, 0xf7, 0x79, 0x42, 0x5e, 0x78, 0x91, 0xa9, 0x6c, 0xa9, 0x15,
	0xcd, 0x9e, 0x92, 0x62, 0x99, 0x27, 0x6e, 0x2a, 0x1e, 0x12, 0x2e, 0x2d, 0xd3, 0x92, 0xae, 0xe2,
	0x42, 0x2b, 0xab, 0xd0, 0xe1, 0x8f, 0x15, 0x7b, 0x63, 0x78, 0x9c, 0xab, 0x5c, 0xb9, 0x6d, 0xd2,
	0x4c, 0xad, 0x38, 0x7a, 0x03, 0x70, 0x70, 0xe3, 0xdb, 0x39, 0x17, 0x6c, 0xc6, 0x34, 0x67, 0x66,
	0x42, 0x2d, 0x45, 0xd7, 0xf0, 0xc4, 0x58, 0xaa, 0x6d, 0x6a, 0xb9, 0x60, 0xc6, 0x52, 0x51, 0xa4,
	0x92, 0x4a, 0x65, 0x30, 0x08, 0x41, 0x14, 0x8c, 0xff, 0x57, 0xef, 0x67, 0x9d, 0xe9, 0x91, 0x53,
	0xe6, 0x1b, 0xe3, 0xae, 0x11, 0x5c, 0x49, 0x45, 0xb1, 0x62, 0xe9, 0x63, 0xa9, 0xa9, 0xe5, 0x4a,
	0xfa, 0xf2, 0xdf, 0xaf, 0xd2, 0x29, 0x13, 0x6f, 0xb4, 0xe5, 0x2d, 0xec, 0xb7, 0xd8, 0xe0, 0x20,
	0x0c, 0xa2, 0xbd, 0xab, 0xcb, 0x78, 0xeb, 0x92, 0x78, 0xfb, 0xbf, 0x33, 0xd7, 0xf8, 0x87, 0x37,
	0x2f, 0x8c, 0xd6, 0x00, 0xe2, 0xbf, 0x5c, 0x74, 0x0a, 0x7b, 0x6a, 0xb1, 0x30, 0xcc, 0xba, 0x73,
	0xba, 0xbe, 0xf5, 0x0c, 0x0d, 0x61, 0x37, 0x53, 0xa5, 0xb4, 0xb8, 0x17, 0x82, 0x68, 0xdf, 0x2f,
	0x5b, 0x84, 0x06, 0x30, 0x30, 0xa5, 0xc0, 0xfd, 0x10, 0x44, 0xc0, 0x6f, 0x1a, 0x80, 0x0e, 0x60,
	0x20, 0xe8, 0x33, 0xde, 0x69, 0xf8, 0xb4, 0x19, 0x1d, 0xe1, 0x12, 0xef, 0x7a, 0xc2, 0xe5, 0xf8,
	0xbc, 0xfa, 0x24, 0x9d, 0xaa, 0x26, 0xe0, 0xa5, 0x26, 0xe0, 0xb5, 0x26, 0xe0, 0xa3, 0x26, 0x60,
	0xfd, 0x45, 0x3a, 0xf7, 0x7d, 0x7f, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x9b, 0xe9,
	0xa7, 0xe1, 0x01, 0x00, 0x00,
}
