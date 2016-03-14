// Code generated by protoc-gen-go.
// source: expression.proto
// DO NOT EDIT!

package tipb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExpressionType int32

const (
	// children count 0.
	// values are in text format.
	ExpressionType_Null      ExpressionType = 0
	ExpressionType_Int64     ExpressionType = 1
	ExpressionType_Float64   ExpressionType = 2
	ExpressionType_String    ExpressionType = 3
	ExpressionType_ColumnRef ExpressionType = 21
	// children count 2.
	ExpressionType_AndAnd ExpressionType = 101
	ExpressionType_OrOr   ExpressionType = 102
	ExpressionType_Plus   ExpressionType = 103
	ExpressionType_Minus  ExpressionType = 104
)

var ExpressionType_name = map[int32]string{
	0:   "Null",
	1:   "Int64",
	2:   "Float64",
	3:   "String",
	21:  "ColumnRef",
	101: "AndAnd",
	102: "OrOr",
	103: "Plus",
	104: "Minus",
}
var ExpressionType_value = map[string]int32{
	"Null":      0,
	"Int64":     1,
	"Float64":   2,
	"String":    3,
	"ColumnRef": 21,
	"AndAnd":    101,
	"OrOr":      102,
	"Plus":      103,
	"Minus":     104,
}

func (x ExpressionType) Enum() *ExpressionType {
	p := new(ExpressionType)
	*p = x
	return p
}
func (x ExpressionType) String() string {
	return proto.EnumName(ExpressionType_name, int32(x))
}
func (x *ExpressionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExpressionType_value, data, "ExpressionType")
	if err != nil {
		return err
	}
	*x = ExpressionType(value)
	return nil
}
func (ExpressionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Evaluators should implement evaluation functions for every expression type.
type Expression struct {
	Tp               *ExpressionType `protobuf:"varint,1,opt,name=tp,enum=tipb.ExpressionType" json:"tp,omitempty"`
	Val              []byte          `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	Children         []*Expression   `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Expression) Reset()                    { *m = Expression{} }
func (m *Expression) String() string            { return proto.CompactTextString(m) }
func (*Expression) ProtoMessage()               {}
func (*Expression) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Expression) GetTp() ExpressionType {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return ExpressionType_Null
}

func (m *Expression) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Expression) GetChildren() []*Expression {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*Expression)(nil), "tipb.Expression")
	proto.RegisterEnum("tipb.ExpressionType", ExpressionType_name, ExpressionType_value)
}

var fileDescriptor1 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xdd, 0x4d, 0xd5, 0x76, 0x56, 0xcb, 0x10, 0x2c, 0xec, 0x71, 0xe9, 0x41, 0x8a, 0x87,
	0x1c, 0x8a, 0x78, 0xb7, 0xa2, 0xe0, 0x41, 0xbb, 0x58, 0xef, 0xb2, 0xcd, 0xa6, 0x6d, 0x60, 0x9a,
	0x84, 0x6c, 0x56, 0xea, 0xbf, 0x37, 0xe9, 0x41, 0xa1, 0x30, 0x87, 0x37, 0xbc, 0x37, 0x1f, 0x8f,
	0x01, 0x54, 0x07, 0xe7, 0x55, 0xd7, 0x69, 0x6b, 0x84, 0xf3, 0x36, 0x58, 0x3e, 0x08, 0xda, 0xad,
	0xa7, 0x5f, 0x00, 0xcf, 0x7f, 0x0e, 0xaf, 0x20, 0x0f, 0xae, 0xcc, 0xaa, 0x6c, 0x36, 0x9e, 0xdf,
	0x88, 0x14, 0x10, 0xff, 0xee, 0xe7, 0x8f, 0x53, 0xbc, 0x00, 0xf6, 0xdd, 0x50, 0x99, 0xc7, 0xc8,
	0x15, 0x9f, 0xc2, 0x50, 0xee, 0x34, 0xb5, 0x5e, 0x99, 0x92, 0x55, 0x6c, 0x56, 0xcc, 0xf1, 0xf4,
	0xe8, 0xee, 0x00, 0xe3, 0x13, 0xc4, 0x10, 0x06, 0xef, 0x3d, 0x11, 0x9e, 0xf1, 0x11, 0x9c, 0xbf,
	0x9a, 0xf0, 0x70, 0x8f, 0x59, 0xe4, 0x5e, 0xbe, 0x90, 0x6d, 0xd2, 0x92, 0x73, 0x80, 0x8b, 0x55,
	0xf0, 0xda, 0x6c, 0x91, 0xf1, 0x6b, 0x18, 0x3d, 0x59, 0xea, 0xf7, 0xe6, 0x43, 0x6d, 0x70, 0x92,
	0xac, 0x47, 0xd3, 0xc6, 0xc1, 0x23, 0x68, 0xe9, 0x97, 0x1e, 0x37, 0x49, 0xd5, 0xd4, 0x77, 0xb8,
	0x4d, 0xc8, 0x37, 0x6d, 0xa2, 0xdc, 0x2d, 0x6e, 0x61, 0x22, 0xed, 0x5e, 0xb8, 0xc8, 0x91, 0x8d,
	0x8b, 0xc5, 0xda, 0xf5, 0xb1, 0xdd, 0xa2, 0x58, 0x29, 0x52, 0x32, 0xd4, 0xe9, 0x0d, 0x75, 0xf6,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xb4, 0xae, 0xb3, 0x1b, 0x01, 0x00, 0x00,
}