// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/syntax.proto

/*
Package syntax_proto is a generated protocol buffer package.

It is generated from these files:
	proto/v1/syntax.proto

It has these top-level messages:
	ParsedExpr
	Expr
	Constant
	SourceInfo
*/
package syntax_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ParsedExpr struct {
	Expr       *Expr       `protobuf:"bytes,2,opt,name=expr" json:"expr,omitempty"`
	SourceInfo *SourceInfo `protobuf:"bytes,3,opt,name=source_info,json=sourceInfo" json:"source_info,omitempty"`
}

func (m *ParsedExpr) Reset()                    { *m = ParsedExpr{} }
func (m *ParsedExpr) String() string            { return proto.CompactTextString(m) }
func (*ParsedExpr) ProtoMessage()               {}
func (*ParsedExpr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ParsedExpr) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *ParsedExpr) GetSourceInfo() *SourceInfo {
	if m != nil {
		return m.SourceInfo
	}
	return nil
}

type Expr struct {
	Id int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	// Types that are valid to be assigned to ExprKind:
	//	*Expr_ConstExpr
	//	*Expr_IdentExpr
	//	*Expr_SelectExpr
	//	*Expr_CallExpr
	//	*Expr_ListExpr
	//	*Expr_StructExpr
	//	*Expr_ComprehensionExpr
	ExprKind isExpr_ExprKind `protobuf_oneof:"expr_kind"`
}

func (m *Expr) Reset()                    { *m = Expr{} }
func (m *Expr) String() string            { return proto.CompactTextString(m) }
func (*Expr) ProtoMessage()               {}
func (*Expr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isExpr_ExprKind interface {
	isExpr_ExprKind()
}

type Expr_ConstExpr struct {
	ConstExpr *Constant `protobuf:"bytes,3,opt,name=const_expr,json=constExpr,oneof"`
}
type Expr_IdentExpr struct {
	IdentExpr *Expr_Ident `protobuf:"bytes,4,opt,name=ident_expr,json=identExpr,oneof"`
}
type Expr_SelectExpr struct {
	SelectExpr *Expr_Select `protobuf:"bytes,5,opt,name=select_expr,json=selectExpr,oneof"`
}
type Expr_CallExpr struct {
	CallExpr *Expr_Call `protobuf:"bytes,6,opt,name=call_expr,json=callExpr,oneof"`
}
type Expr_ListExpr struct {
	ListExpr *Expr_CreateList `protobuf:"bytes,7,opt,name=list_expr,json=listExpr,oneof"`
}
type Expr_StructExpr struct {
	StructExpr *Expr_CreateStruct `protobuf:"bytes,8,opt,name=struct_expr,json=structExpr,oneof"`
}
type Expr_ComprehensionExpr struct {
	ComprehensionExpr *Expr_Comprehension `protobuf:"bytes,9,opt,name=comprehension_expr,json=comprehensionExpr,oneof"`
}

func (*Expr_ConstExpr) isExpr_ExprKind()         {}
func (*Expr_IdentExpr) isExpr_ExprKind()         {}
func (*Expr_SelectExpr) isExpr_ExprKind()        {}
func (*Expr_CallExpr) isExpr_ExprKind()          {}
func (*Expr_ListExpr) isExpr_ExprKind()          {}
func (*Expr_StructExpr) isExpr_ExprKind()        {}
func (*Expr_ComprehensionExpr) isExpr_ExprKind() {}

func (m *Expr) GetExprKind() isExpr_ExprKind {
	if m != nil {
		return m.ExprKind
	}
	return nil
}

func (m *Expr) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Expr) GetConstExpr() *Constant {
	if x, ok := m.GetExprKind().(*Expr_ConstExpr); ok {
		return x.ConstExpr
	}
	return nil
}

func (m *Expr) GetIdentExpr() *Expr_Ident {
	if x, ok := m.GetExprKind().(*Expr_IdentExpr); ok {
		return x.IdentExpr
	}
	return nil
}

func (m *Expr) GetSelectExpr() *Expr_Select {
	if x, ok := m.GetExprKind().(*Expr_SelectExpr); ok {
		return x.SelectExpr
	}
	return nil
}

func (m *Expr) GetCallExpr() *Expr_Call {
	if x, ok := m.GetExprKind().(*Expr_CallExpr); ok {
		return x.CallExpr
	}
	return nil
}

func (m *Expr) GetListExpr() *Expr_CreateList {
	if x, ok := m.GetExprKind().(*Expr_ListExpr); ok {
		return x.ListExpr
	}
	return nil
}

func (m *Expr) GetStructExpr() *Expr_CreateStruct {
	if x, ok := m.GetExprKind().(*Expr_StructExpr); ok {
		return x.StructExpr
	}
	return nil
}

func (m *Expr) GetComprehensionExpr() *Expr_Comprehension {
	if x, ok := m.GetExprKind().(*Expr_ComprehensionExpr); ok {
		return x.ComprehensionExpr
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Expr) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Expr_OneofMarshaler, _Expr_OneofUnmarshaler, _Expr_OneofSizer, []interface{}{
		(*Expr_ConstExpr)(nil),
		(*Expr_IdentExpr)(nil),
		(*Expr_SelectExpr)(nil),
		(*Expr_CallExpr)(nil),
		(*Expr_ListExpr)(nil),
		(*Expr_StructExpr)(nil),
		(*Expr_ComprehensionExpr)(nil),
	}
}

func _Expr_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Expr)
	// expr_kind
	switch x := m.ExprKind.(type) {
	case *Expr_ConstExpr:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConstExpr); err != nil {
			return err
		}
	case *Expr_IdentExpr:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IdentExpr); err != nil {
			return err
		}
	case *Expr_SelectExpr:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SelectExpr); err != nil {
			return err
		}
	case *Expr_CallExpr:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CallExpr); err != nil {
			return err
		}
	case *Expr_ListExpr:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListExpr); err != nil {
			return err
		}
	case *Expr_StructExpr:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructExpr); err != nil {
			return err
		}
	case *Expr_ComprehensionExpr:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ComprehensionExpr); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Expr.ExprKind has unexpected type %T", x)
	}
	return nil
}

func _Expr_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Expr)
	switch tag {
	case 3: // expr_kind.const_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Constant)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_ConstExpr{msg}
		return true, err
	case 4: // expr_kind.ident_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Ident)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_IdentExpr{msg}
		return true, err
	case 5: // expr_kind.select_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Select)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_SelectExpr{msg}
		return true, err
	case 6: // expr_kind.call_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Call)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_CallExpr{msg}
		return true, err
	case 7: // expr_kind.list_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_CreateList)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_ListExpr{msg}
		return true, err
	case 8: // expr_kind.struct_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_CreateStruct)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_StructExpr{msg}
		return true, err
	case 9: // expr_kind.comprehension_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Comprehension)
		err := b.DecodeMessage(msg)
		m.ExprKind = &Expr_ComprehensionExpr{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Expr_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Expr)
	// expr_kind
	switch x := m.ExprKind.(type) {
	case *Expr_ConstExpr:
		s := proto.Size(x.ConstExpr)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_IdentExpr:
		s := proto.Size(x.IdentExpr)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_SelectExpr:
		s := proto.Size(x.SelectExpr)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_CallExpr:
		s := proto.Size(x.CallExpr)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_ListExpr:
		s := proto.Size(x.ListExpr)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_StructExpr:
		s := proto.Size(x.StructExpr)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_ComprehensionExpr:
		s := proto.Size(x.ComprehensionExpr)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Expr_Ident struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Expr_Ident) Reset()                    { *m = Expr_Ident{} }
func (m *Expr_Ident) String() string            { return proto.CompactTextString(m) }
func (*Expr_Ident) ProtoMessage()               {}
func (*Expr_Ident) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Expr_Ident) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Expr_Select struct {
	Operand  *Expr  `protobuf:"bytes,1,opt,name=operand" json:"operand,omitempty"`
	Field    string `protobuf:"bytes,2,opt,name=field" json:"field,omitempty"`
	TestOnly bool   `protobuf:"varint,3,opt,name=test_only,json=testOnly" json:"test_only,omitempty"`
}

func (m *Expr_Select) Reset()                    { *m = Expr_Select{} }
func (m *Expr_Select) String() string            { return proto.CompactTextString(m) }
func (*Expr_Select) ProtoMessage()               {}
func (*Expr_Select) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Expr_Select) GetOperand() *Expr {
	if m != nil {
		return m.Operand
	}
	return nil
}

func (m *Expr_Select) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Expr_Select) GetTestOnly() bool {
	if m != nil {
		return m.TestOnly
	}
	return false
}

type Expr_Call struct {
	Target   *Expr   `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Function string  `protobuf:"bytes,2,opt,name=function" json:"function,omitempty"`
	Args     []*Expr `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
}

func (m *Expr_Call) Reset()                    { *m = Expr_Call{} }
func (m *Expr_Call) String() string            { return proto.CompactTextString(m) }
func (*Expr_Call) ProtoMessage()               {}
func (*Expr_Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

func (m *Expr_Call) GetTarget() *Expr {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Expr_Call) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Expr_Call) GetArgs() []*Expr {
	if m != nil {
		return m.Args
	}
	return nil
}

type Expr_CreateList struct {
	Elements []*Expr `protobuf:"bytes,1,rep,name=elements" json:"elements,omitempty"`
}

func (m *Expr_CreateList) Reset()                    { *m = Expr_CreateList{} }
func (m *Expr_CreateList) String() string            { return proto.CompactTextString(m) }
func (*Expr_CreateList) ProtoMessage()               {}
func (*Expr_CreateList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 3} }

func (m *Expr_CreateList) GetElements() []*Expr {
	if m != nil {
		return m.Elements
	}
	return nil
}

type Expr_CreateStruct struct {
	MessageName string                     `protobuf:"bytes,1,opt,name=message_name,json=messageName" json:"message_name,omitempty"`
	Entries     []*Expr_CreateStruct_Entry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *Expr_CreateStruct) Reset()                    { *m = Expr_CreateStruct{} }
func (m *Expr_CreateStruct) String() string            { return proto.CompactTextString(m) }
func (*Expr_CreateStruct) ProtoMessage()               {}
func (*Expr_CreateStruct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 4} }

func (m *Expr_CreateStruct) GetMessageName() string {
	if m != nil {
		return m.MessageName
	}
	return ""
}

func (m *Expr_CreateStruct) GetEntries() []*Expr_CreateStruct_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Expr_CreateStruct_Entry struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Types that are valid to be assigned to KeyKind:
	//	*Expr_CreateStruct_Entry_FieldKey
	//	*Expr_CreateStruct_Entry_MapKey
	KeyKind isExpr_CreateStruct_Entry_KeyKind `protobuf_oneof:"key_kind"`
	Value   *Expr                             `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Expr_CreateStruct_Entry) Reset()                    { *m = Expr_CreateStruct_Entry{} }
func (m *Expr_CreateStruct_Entry) String() string            { return proto.CompactTextString(m) }
func (*Expr_CreateStruct_Entry) ProtoMessage()               {}
func (*Expr_CreateStruct_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 4, 0} }

type isExpr_CreateStruct_Entry_KeyKind interface {
	isExpr_CreateStruct_Entry_KeyKind()
}

type Expr_CreateStruct_Entry_FieldKey struct {
	FieldKey string `protobuf:"bytes,2,opt,name=field_key,json=fieldKey,oneof"`
}
type Expr_CreateStruct_Entry_MapKey struct {
	MapKey *Expr `protobuf:"bytes,3,opt,name=map_key,json=mapKey,oneof"`
}

func (*Expr_CreateStruct_Entry_FieldKey) isExpr_CreateStruct_Entry_KeyKind() {}
func (*Expr_CreateStruct_Entry_MapKey) isExpr_CreateStruct_Entry_KeyKind()   {}

func (m *Expr_CreateStruct_Entry) GetKeyKind() isExpr_CreateStruct_Entry_KeyKind {
	if m != nil {
		return m.KeyKind
	}
	return nil
}

func (m *Expr_CreateStruct_Entry) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Expr_CreateStruct_Entry) GetFieldKey() string {
	if x, ok := m.GetKeyKind().(*Expr_CreateStruct_Entry_FieldKey); ok {
		return x.FieldKey
	}
	return ""
}

func (m *Expr_CreateStruct_Entry) GetMapKey() *Expr {
	if x, ok := m.GetKeyKind().(*Expr_CreateStruct_Entry_MapKey); ok {
		return x.MapKey
	}
	return nil
}

func (m *Expr_CreateStruct_Entry) GetValue() *Expr {
	if m != nil {
		return m.Value
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Expr_CreateStruct_Entry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Expr_CreateStruct_Entry_OneofMarshaler, _Expr_CreateStruct_Entry_OneofUnmarshaler, _Expr_CreateStruct_Entry_OneofSizer, []interface{}{
		(*Expr_CreateStruct_Entry_FieldKey)(nil),
		(*Expr_CreateStruct_Entry_MapKey)(nil),
	}
}

func _Expr_CreateStruct_Entry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Expr_CreateStruct_Entry)
	// key_kind
	switch x := m.KeyKind.(type) {
	case *Expr_CreateStruct_Entry_FieldKey:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FieldKey)
	case *Expr_CreateStruct_Entry_MapKey:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MapKey); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Expr_CreateStruct_Entry.KeyKind has unexpected type %T", x)
	}
	return nil
}

func _Expr_CreateStruct_Entry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Expr_CreateStruct_Entry)
	switch tag {
	case 2: // key_kind.field_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.KeyKind = &Expr_CreateStruct_Entry_FieldKey{x}
		return true, err
	case 3: // key_kind.map_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr)
		err := b.DecodeMessage(msg)
		m.KeyKind = &Expr_CreateStruct_Entry_MapKey{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Expr_CreateStruct_Entry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Expr_CreateStruct_Entry)
	// key_kind
	switch x := m.KeyKind.(type) {
	case *Expr_CreateStruct_Entry_FieldKey:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FieldKey)))
		n += len(x.FieldKey)
	case *Expr_CreateStruct_Entry_MapKey:
		s := proto.Size(x.MapKey)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Expr_Comprehension struct {
	IterVar       string `protobuf:"bytes,1,opt,name=iter_var,json=iterVar" json:"iter_var,omitempty"`
	IterRange     *Expr  `protobuf:"bytes,2,opt,name=iter_range,json=iterRange" json:"iter_range,omitempty"`
	AccuVar       string `protobuf:"bytes,3,opt,name=accu_var,json=accuVar" json:"accu_var,omitempty"`
	AccuInit      *Expr  `protobuf:"bytes,4,opt,name=accu_init,json=accuInit" json:"accu_init,omitempty"`
	LoopCondition *Expr  `protobuf:"bytes,5,opt,name=loop_condition,json=loopCondition" json:"loop_condition,omitempty"`
	LoopStep      *Expr  `protobuf:"bytes,6,opt,name=loop_step,json=loopStep" json:"loop_step,omitempty"`
	Result        *Expr  `protobuf:"bytes,7,opt,name=result" json:"result,omitempty"`
}

func (m *Expr_Comprehension) Reset()                    { *m = Expr_Comprehension{} }
func (m *Expr_Comprehension) String() string            { return proto.CompactTextString(m) }
func (*Expr_Comprehension) ProtoMessage()               {}
func (*Expr_Comprehension) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 5} }

func (m *Expr_Comprehension) GetIterVar() string {
	if m != nil {
		return m.IterVar
	}
	return ""
}

func (m *Expr_Comprehension) GetIterRange() *Expr {
	if m != nil {
		return m.IterRange
	}
	return nil
}

func (m *Expr_Comprehension) GetAccuVar() string {
	if m != nil {
		return m.AccuVar
	}
	return ""
}

func (m *Expr_Comprehension) GetAccuInit() *Expr {
	if m != nil {
		return m.AccuInit
	}
	return nil
}

func (m *Expr_Comprehension) GetLoopCondition() *Expr {
	if m != nil {
		return m.LoopCondition
	}
	return nil
}

func (m *Expr_Comprehension) GetLoopStep() *Expr {
	if m != nil {
		return m.LoopStep
	}
	return nil
}

func (m *Expr_Comprehension) GetResult() *Expr {
	if m != nil {
		return m.Result
	}
	return nil
}

type Constant struct {
	// Types that are valid to be assigned to ConstantKind:
	//	*Constant_NullValue
	//	*Constant_BoolValue
	//	*Constant_Int64Value
	//	*Constant_Uint64Value
	//	*Constant_DoubleValue
	//	*Constant_StringValue
	//	*Constant_BytesValue
	//	*Constant_DurationValue
	//	*Constant_TimestampValue
	ConstantKind isConstant_ConstantKind `protobuf_oneof:"constant_kind"`
}

func (m *Constant) Reset()                    { *m = Constant{} }
func (m *Constant) String() string            { return proto.CompactTextString(m) }
func (*Constant) ProtoMessage()               {}
func (*Constant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isConstant_ConstantKind interface {
	isConstant_ConstantKind()
}

type Constant_NullValue struct {
	NullValue google_protobuf1.NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,enum=google.protobuf.NullValue,oneof"`
}
type Constant_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,oneof"`
}
type Constant_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,oneof"`
}
type Constant_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,4,opt,name=uint64_value,json=uint64Value,oneof"`
}
type Constant_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,oneof"`
}
type Constant_StringValue struct {
	StringValue string `protobuf:"bytes,6,opt,name=string_value,json=stringValue,oneof"`
}
type Constant_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}
type Constant_DurationValue struct {
	DurationValue *google_protobuf.Duration `protobuf:"bytes,8,opt,name=duration_value,json=durationValue,oneof"`
}
type Constant_TimestampValue struct {
	TimestampValue *google_protobuf2.Timestamp `protobuf:"bytes,9,opt,name=timestamp_value,json=timestampValue,oneof"`
}

func (*Constant_NullValue) isConstant_ConstantKind()      {}
func (*Constant_BoolValue) isConstant_ConstantKind()      {}
func (*Constant_Int64Value) isConstant_ConstantKind()     {}
func (*Constant_Uint64Value) isConstant_ConstantKind()    {}
func (*Constant_DoubleValue) isConstant_ConstantKind()    {}
func (*Constant_StringValue) isConstant_ConstantKind()    {}
func (*Constant_BytesValue) isConstant_ConstantKind()     {}
func (*Constant_DurationValue) isConstant_ConstantKind()  {}
func (*Constant_TimestampValue) isConstant_ConstantKind() {}

func (m *Constant) GetConstantKind() isConstant_ConstantKind {
	if m != nil {
		return m.ConstantKind
	}
	return nil
}

func (m *Constant) GetNullValue() google_protobuf1.NullValue {
	if x, ok := m.GetConstantKind().(*Constant_NullValue); ok {
		return x.NullValue
	}
	return google_protobuf1.NullValue_NULL_VALUE
}

func (m *Constant) GetBoolValue() bool {
	if x, ok := m.GetConstantKind().(*Constant_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Constant) GetInt64Value() int64 {
	if x, ok := m.GetConstantKind().(*Constant_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Constant) GetUint64Value() uint64 {
	if x, ok := m.GetConstantKind().(*Constant_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *Constant) GetDoubleValue() float64 {
	if x, ok := m.GetConstantKind().(*Constant_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Constant) GetStringValue() string {
	if x, ok := m.GetConstantKind().(*Constant_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Constant) GetBytesValue() []byte {
	if x, ok := m.GetConstantKind().(*Constant_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Constant) GetDurationValue() *google_protobuf.Duration {
	if x, ok := m.GetConstantKind().(*Constant_DurationValue); ok {
		return x.DurationValue
	}
	return nil
}

func (m *Constant) GetTimestampValue() *google_protobuf2.Timestamp {
	if x, ok := m.GetConstantKind().(*Constant_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Constant) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Constant_OneofMarshaler, _Constant_OneofUnmarshaler, _Constant_OneofSizer, []interface{}{
		(*Constant_NullValue)(nil),
		(*Constant_BoolValue)(nil),
		(*Constant_Int64Value)(nil),
		(*Constant_Uint64Value)(nil),
		(*Constant_DoubleValue)(nil),
		(*Constant_StringValue)(nil),
		(*Constant_BytesValue)(nil),
		(*Constant_DurationValue)(nil),
		(*Constant_TimestampValue)(nil),
	}
}

func _Constant_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Constant)
	// constant_kind
	switch x := m.ConstantKind.(type) {
	case *Constant_NullValue:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Constant_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Constant_Int64Value:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *Constant_Uint64Value:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *Constant_DoubleValue:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Constant_StringValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Constant_BytesValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesValue)
	case *Constant_DurationValue:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DurationValue); err != nil {
			return err
		}
	case *Constant_TimestampValue:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Constant.ConstantKind has unexpected type %T", x)
	}
	return nil
}

func _Constant_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Constant)
	switch tag {
	case 1: // constant_kind.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConstantKind = &Constant_NullValue{google_protobuf1.NullValue(x)}
		return true, err
	case 2: // constant_kind.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConstantKind = &Constant_BoolValue{x != 0}
		return true, err
	case 3: // constant_kind.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConstantKind = &Constant_Int64Value{int64(x)}
		return true, err
	case 4: // constant_kind.uint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConstantKind = &Constant_Uint64Value{x}
		return true, err
	case 5: // constant_kind.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ConstantKind = &Constant_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 6: // constant_kind.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConstantKind = &Constant_StringValue{x}
		return true, err
	case 7: // constant_kind.bytes_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ConstantKind = &Constant_BytesValue{x}
		return true, err
	case 8: // constant_kind.duration_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Duration)
		err := b.DecodeMessage(msg)
		m.ConstantKind = &Constant_DurationValue{msg}
		return true, err
	case 9: // constant_kind.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Timestamp)
		err := b.DecodeMessage(msg)
		m.ConstantKind = &Constant_TimestampValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Constant_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Constant)
	// constant_kind
	switch x := m.ConstantKind.(type) {
	case *Constant_NullValue:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Constant_BoolValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *Constant_Int64Value:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *Constant_Uint64Value:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *Constant_DoubleValue:
		n += proto.SizeVarint(5<<3 | proto.WireFixed64)
		n += 8
	case *Constant_StringValue:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Constant_BytesValue:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BytesValue)))
		n += len(x.BytesValue)
	case *Constant_DurationValue:
		s := proto.Size(x.DurationValue)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Constant_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SourceInfo struct {
	SyntaxVersion string          `protobuf:"bytes,1,opt,name=syntax_version,json=syntaxVersion" json:"syntax_version,omitempty"`
	Location      string          `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	LineOffsets   []int32         `protobuf:"varint,3,rep,packed,name=line_offsets,json=lineOffsets" json:"line_offsets,omitempty"`
	Positions     map[int64]int32 `protobuf:"bytes,4,rep,name=positions" json:"positions,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *SourceInfo) Reset()                    { *m = SourceInfo{} }
func (m *SourceInfo) String() string            { return proto.CompactTextString(m) }
func (*SourceInfo) ProtoMessage()               {}
func (*SourceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SourceInfo) GetSyntaxVersion() string {
	if m != nil {
		return m.SyntaxVersion
	}
	return ""
}

func (m *SourceInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SourceInfo) GetLineOffsets() []int32 {
	if m != nil {
		return m.LineOffsets
	}
	return nil
}

func (m *SourceInfo) GetPositions() map[int64]int32 {
	if m != nil {
		return m.Positions
	}
	return nil
}

func init() {
	proto.RegisterType((*ParsedExpr)(nil), "google.api.expr.v1.ParsedExpr")
	proto.RegisterType((*Expr)(nil), "google.api.expr.v1.Expr")
	proto.RegisterType((*Expr_Ident)(nil), "google.api.expr.v1.Expr.Ident")
	proto.RegisterType((*Expr_Select)(nil), "google.api.expr.v1.Expr.Select")
	proto.RegisterType((*Expr_Call)(nil), "google.api.expr.v1.Expr.Call")
	proto.RegisterType((*Expr_CreateList)(nil), "google.api.expr.v1.Expr.CreateList")
	proto.RegisterType((*Expr_CreateStruct)(nil), "google.api.expr.v1.Expr.CreateStruct")
	proto.RegisterType((*Expr_CreateStruct_Entry)(nil), "google.api.expr.v1.Expr.CreateStruct.Entry")
	proto.RegisterType((*Expr_Comprehension)(nil), "google.api.expr.v1.Expr.Comprehension")
	proto.RegisterType((*Constant)(nil), "google.api.expr.v1.Constant")
	proto.RegisterType((*SourceInfo)(nil), "google.api.expr.v1.SourceInfo")
}

func init() { proto.RegisterFile("proto/v1/syntax.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x2d, 0x7f, 0x45, 0x3a, 0x4a, 0xdc, 0x8e, 0xd8, 0x06, 0x57, 0xfd, 0x6e, 0xd1, 0xa1,
	0xd8, 0x87, 0xb3, 0xa6, 0xdd, 0x07, 0xb6, 0x0e, 0x05, 0x9c, 0x05, 0x70, 0xd6, 0xa1, 0x0d, 0x94,
	0x21, 0xbb, 0x14, 0x18, 0x99, 0xf6, 0x84, 0xc8, 0xa4, 0x20, 0x52, 0x41, 0x7c, 0xb9, 0x3d, 0xca,
	0x6e, 0xf6, 0x0c, 0x7b, 0x9a, 0xbd, 0xc2, 0x2e, 0x77, 0x33, 0x60, 0x38, 0x87, 0x94, 0x93, 0x34,
	0xb1, 0xbd, 0x3b, 0xf1, 0xf0, 0xf7, 0xff, 0x93, 0x3e, 0x3c, 0x3c, 0x34, 0x7c, 0x50, 0x94, 0xca,
	0xa8, 0xed, 0xd3, 0x67, 0xdb, 0x7a, 0x2e, 0x0d, 0x3f, 0x1b, 0xd0, 0x98, 0xb1, 0xa9, 0x52, 0xd3,
	0x5c, 0x0c, 0x78, 0x91, 0x0d, 0xc4, 0x59, 0x51, 0x0e, 0x4e, 0x9f, 0x45, 0xf7, 0x6c, 0x6c, 0x9b,
	0x88, 0xe3, 0x6a, 0xb2, 0x3d, 0xae, 0x4a, 0x6e, 0x32, 0x25, 0xad, 0x26, 0xba, 0xf3, 0xee, 0xbc,
	0x36, 0x65, 0x95, 0x1a, 0x37, 0x7b, 0xff, 0xdd, 0x59, 0x93, 0xcd, 0x84, 0x36, 0x7c, 0x56, 0x58,
	0xe0, 0xd1, 0xaf, 0x1e, 0xc0, 0x01, 0x2f, 0xb5, 0x18, 0xef, 0x9d, 0x15, 0x25, 0xfb, 0x14, 0xda,
	0xb8, 0x70, 0xbf, 0xf9, 0xc0, 0x7b, 0x1a, 0xee, 0xf4, 0x07, 0x57, 0x37, 0x34, 0x40, 0x2e, 0x26,
	0x8a, 0xbd, 0x82, 0x50, 0xab, 0xaa, 0x4c, 0x45, 0x92, 0xc9, 0x89, 0xea, 0xb7, 0x48, 0x74, 0xef,
	0x3a, 0xd1, 0x21, 0x61, 0xfb, 0x72, 0xa2, 0x62, 0xd0, 0x8b, 0xef, 0x1f, 0xda, 0xbe, 0x77, 0xb3,
	0xf9, 0xe8, 0xef, 0x10, 0xda, 0xb4, 0x7a, 0x0f, 0x9a, 0xd9, 0x98, 0xd6, 0x6e, 0xc5, 0xcd, 0x6c,
	0xcc, 0xbe, 0x03, 0x48, 0x95, 0xd4, 0x26, 0xa1, 0x3d, 0x59, 0xfb, 0x3b, 0xd7, 0xd9, 0xef, 0x22,
	0xc5, 0xa5, 0x19, 0x35, 0xe2, 0x80, 0x14, 0x7b, 0x76, 0x7b, 0x90, 0x8d, 0x85, 0x74, 0xf2, 0xf6,
	0xf2, 0xdd, 0x21, 0x3d, 0xd8, 0x47, 0x14, 0x0d, 0x48, 0x43, 0x06, 0x43, 0x08, 0xb5, 0xc8, 0x45,
	0xea, 0x1c, 0x3a, 0xe4, 0x70, 0x7f, 0xa9, 0xc3, 0x21, 0xb1, 0xa3, 0x46, 0x0c, 0x56, 0x45, 0x1e,
	0x2f, 0x21, 0x48, 0x79, 0x9e, 0x5b, 0x87, 0x2e, 0x39, 0xdc, 0x5d, 0xea, 0xb0, 0xcb, 0xf3, 0x7c,
	0xd4, 0x88, 0x7d, 0x54, 0xb8, 0x1d, 0x04, 0x79, 0x56, 0x27, 0x60, 0x83, 0xd4, 0x8f, 0x97, 0xab,
	0x4b, 0xc1, 0x8d, 0xf8, 0x31, 0xd3, 0xb8, 0x07, 0x1f, 0x75, 0xe4, 0x31, 0x82, 0xd0, 0xd6, 0x84,
	0x75, 0xf1, 0xc9, 0xe5, 0xc9, 0x1a, 0x97, 0x43, 0x52, 0xd0, 0x6f, 0xa1, 0x2f, 0x72, 0xfa, 0x19,
	0x58, 0xaa, 0x66, 0x45, 0x29, 0x7e, 0x11, 0x52, 0x67, 0x4a, 0x5a, 0xc3, 0x80, 0x0c, 0x3f, 0x5a,
	0x6e, 0x78, 0x51, 0x32, 0x6a, 0xc4, 0xef, 0x5d, 0xf2, 0x40, 0x24, 0xba, 0x0d, 0x1d, 0x4a, 0x3f,
	0x63, 0xd0, 0x96, 0x7c, 0x26, 0xfa, 0xde, 0x03, 0xef, 0x69, 0x10, 0xd3, 0x77, 0xa4, 0xa0, 0x6b,
	0x33, 0xcb, 0x76, 0x60, 0x43, 0x15, 0xa2, 0xe4, 0x72, 0x4c, 0xc0, 0xaa, 0x02, 0xad, 0x41, 0xf6,
	0x3e, 0x74, 0x26, 0x99, 0xc8, 0x6d, 0x59, 0x05, 0xb1, 0x1d, 0xb0, 0xdb, 0x10, 0x18, 0xa1, 0x4d,
	0xa2, 0x64, 0x3e, 0xa7, 0xc2, 0xf2, 0x63, 0x1f, 0x03, 0x6f, 0x65, 0x3e, 0x8f, 0x7e, 0xf3, 0xa0,
	0x8d, 0x27, 0xc1, 0x3e, 0x87, 0xae, 0xe1, 0xe5, 0x54, 0x98, 0xb5, 0xcb, 0x39, 0x8e, 0x45, 0xe0,
	0x4f, 0x2a, 0x99, 0xe2, 0xfd, 0x74, 0x0b, 0x2e, 0xc6, 0x78, 0xb7, 0x78, 0x39, 0xd5, 0xfd, 0xd6,
	0x83, 0xd6, 0xea, 0xbb, 0x85, 0x54, 0x34, 0x04, 0x38, 0x3f, 0x4f, 0xf6, 0x02, 0x7c, 0x91, 0x8b,
	0x99, 0x90, 0x46, 0xf7, 0xbd, 0x35, 0xfa, 0x05, 0x19, 0xfd, 0xde, 0x84, 0xcd, 0x8b, 0xc7, 0xc9,
	0x1e, 0xc2, 0xe6, 0x4c, 0x68, 0xcd, 0xa7, 0x22, 0xb9, 0x90, 0xe6, 0xd0, 0xc5, 0xde, 0xf0, 0x99,
	0x60, 0x7b, 0xb0, 0x21, 0xa4, 0x29, 0x33, 0xa1, 0xfb, 0x4d, 0x5a, 0xe8, 0x93, 0xff, 0x55, 0x29,
	0x83, 0x3d, 0x69, 0xca, 0x79, 0x5c, 0x6b, 0xa3, 0x3f, 0x3c, 0xe8, 0x50, 0xc8, 0x5d, 0x6a, 0x6f,
	0x71, 0xa9, 0xef, 0x42, 0x40, 0x67, 0x90, 0x9c, 0x88, 0xb9, 0xcd, 0x11, 0x56, 0x2b, 0x85, 0x5e,
	0x8b, 0x39, 0x7b, 0x0e, 0x1b, 0x33, 0x5e, 0xd0, 0x64, 0x6b, 0x75, 0xd2, 0x47, 0x8d, 0xb8, 0x3b,
	0xe3, 0x05, 0x8a, 0x06, 0xd0, 0x39, 0xe5, 0x79, 0x25, 0xdc, 0x25, 0x5f, 0x9e, 0x1b, 0x8b, 0x0d,
	0x01, 0xfc, 0x13, 0x31, 0x4f, 0x4e, 0x32, 0x39, 0x8e, 0xfe, 0x6a, 0xc2, 0xd6, 0xa5, 0x12, 0x65,
	0xb7, 0xc0, 0xcf, 0x8c, 0x28, 0x93, 0x53, 0x5e, 0xba, 0x0c, 0x6d, 0xe0, 0xf8, 0x88, 0x97, 0xec,
	0x2b, 0x00, 0x9a, 0x2a, 0xb9, 0x9c, 0x8a, 0xb5, 0x5d, 0x32, 0x40, 0x36, 0x46, 0x14, 0x3d, 0x79,
	0x9a, 0x56, 0xe4, 0xd9, 0xb2, 0x9e, 0x38, 0x46, 0xcf, 0x2f, 0x20, 0xa0, 0xa9, 0x4c, 0x66, 0x66,
	0xed, 0x0f, 0x20, 0x97, 0x7d, 0x99, 0x19, 0xf6, 0x0a, 0x7a, 0xb9, 0x52, 0x45, 0x92, 0x2a, 0x39,
	0xce, 0xa8, 0xe0, 0x3a, 0x6b, 0xb4, 0x5b, 0xc8, 0xef, 0xd6, 0x38, 0xae, 0x4b, 0x06, 0xda, 0x88,
	0xc2, 0x75, 0xa6, 0x15, 0xeb, 0x22, 0x7a, 0x68, 0x44, 0x81, 0x97, 0xa2, 0x14, 0xba, 0xca, 0x8d,
	0xeb, 0x47, 0x2b, 0x2e, 0x85, 0xe5, 0x86, 0x21, 0x04, 0x18, 0xa7, 0x74, 0xbb, 0x96, 0xff, 0x67,
	0x0b, 0xfc, 0xba, 0x69, 0xb3, 0x6f, 0x01, 0x64, 0x95, 0xe7, 0x89, 0x3d, 0x42, 0xcc, 0x78, 0x6f,
	0x27, 0xaa, 0x5d, 0xeb, 0x97, 0x6b, 0xf0, 0xa6, 0xca, 0xf3, 0x23, 0x24, 0xb0, 0x47, 0xcb, 0x7a,
	0xc0, 0xee, 0x03, 0x1c, 0x2b, 0x55, 0x8b, 0xf1, 0x44, 0x7c, 0x04, 0x30, 0x66, 0x81, 0x87, 0x10,
	0x66, 0xd2, 0x7c, 0xf9, 0xc2, 0x11, 0x98, 0xfc, 0x16, 0xf6, 0x35, 0x0a, 0x5a, 0xe4, 0x31, 0x6c,
	0x56, 0x17, 0x19, 0x3c, 0x84, 0xf6, 0xa8, 0x11, 0x87, 0xd5, 0x65, 0x68, 0xac, 0xaa, 0xe3, 0x5c,
	0x38, 0x08, 0xb3, 0xed, 0x21, 0x64, 0xa3, 0x0b, 0x48, 0x9b, 0x32, 0x93, 0x53, 0x07, 0x75, 0x5d,
	0x7d, 0x87, 0x36, 0xba, 0xd8, 0xd1, 0xf1, 0xdc, 0x08, 0xed, 0x18, 0x4c, 0xe3, 0x26, 0xee, 0x88,
	0x82, 0x16, 0x19, 0x42, 0xaf, 0x7e, 0xe7, 0x1d, 0x65, 0xdb, 0xf6, 0xad, 0x2b, 0x69, 0xf9, 0xde,
	0x61, 0xa3, 0x46, 0xbc, 0x55, 0x4b, 0xac, 0xc7, 0x1e, 0xdc, 0x58, 0xbc, 0xf6, 0xce, 0xc4, 0xb6,
	0xea, 0xab, 0xb9, 0xfd, 0xa9, 0xe6, 0x46, 0x8d, 0xb8, 0xb7, 0x10, 0x91, 0xcd, 0xf0, 0x06, 0x6c,
	0xa5, 0xee, 0xa4, 0xe8, 0x04, 0x1f, 0xfd, 0xeb, 0x01, 0x9c, 0xbf, 0xe7, 0xec, 0x09, 0xf4, 0xec,
	0x9f, 0x98, 0xe4, 0x54, 0x94, 0x78, 0x7f, 0xdc, 0x9d, 0xd9, 0xb2, 0xd1, 0x23, 0x1b, 0xc4, 0xce,
	0x98, 0xab, 0x94, 0x5f, 0xec, 0x8c, 0xf5, 0x18, 0xdb, 0x52, 0x9e, 0x49, 0x91, 0xa8, 0xc9, 0x44,
	0x0b, 0x63, 0x3b, 0x64, 0x27, 0x0e, 0x31, 0xf6, 0xd6, 0x86, 0xd8, 0x6b, 0x08, 0x0a, 0xa5, 0xa9,
	0x70, 0x75, 0xbf, 0x4d, 0x8d, 0xe9, 0xb3, 0xd5, 0x7f, 0x34, 0x06, 0x07, 0x35, 0x6f, 0x5b, 0xd3,
	0xb9, 0x3e, 0x7a, 0x09, 0xbd, 0xcb, 0x93, 0xec, 0x26, 0xb4, 0xb0, 0xe3, 0xd8, 0x2e, 0x85, 0x9f,
	0xf8, 0x6e, 0x9c, 0x97, 0x54, 0xc7, 0x35, 0x8e, 0x6f, 0x9a, 0x5f, 0x7b, 0xc3, 0x8f, 0xe1, 0xc3,
	0x54, 0xcd, 0xae, 0x59, 0x7c, 0x18, 0x1e, 0xd2, 0x4f, 0x3e, 0xc0, 0xb4, 0x1e, 0x78, 0xff, 0x78,
	0xde, 0x71, 0x97, 0x52, 0xfc, 0xfc, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x49, 0x74, 0xc9,
	0xf1, 0x09, 0x00, 0x00,
}
