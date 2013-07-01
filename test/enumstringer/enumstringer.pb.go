// Code generated by protoc-gen-gogo.
// source: enumstringer.proto
// DO NOT EDIT!

package enumstringer

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import fmt "fmt"
import bytes "bytes"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type TheTestEnum int32

const (
	TheTestEnum_A TheTestEnum = 0
	TheTestEnum_B TheTestEnum = 1
	TheTestEnum_C TheTestEnum = 2
)

var TheTestEnum_name = map[int32]string{
	0: "A",
	1: "B",
	2: "C",
}
var TheTestEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
	"C": 2,
}

func (x TheTestEnum) Enum() *TheTestEnum {
	p := new(TheTestEnum)
	*p = x
	return p
}
func (x TheTestEnum) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(TheTestEnum_name, int32(x))
}
func (x *TheTestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TheTestEnum_value, data, "TheTestEnum")
	if err != nil {
		return err
	}
	*x = TheTestEnum(value)
	return nil
}

type NidOptEnum struct {
	Field1           TheTestEnum `protobuf:"varint,1,opt,enum=enumstringer.TheTestEnum" json:"Field1"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *NidOptEnum) Reset()         { *m = NidOptEnum{} }
func (m *NidOptEnum) String() string { return proto.CompactTextString(m) }
func (*NidOptEnum) ProtoMessage()    {}

type NinOptEnum struct {
	Field1           *TheTestEnum `protobuf:"varint,1,opt,enum=enumstringer.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NinOptEnum) Reset()         { *m = NinOptEnum{} }
func (m *NinOptEnum) String() string { return proto.CompactTextString(m) }
func (*NinOptEnum) ProtoMessage()    {}

func (m *NinOptEnum) GetField1() TheTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

type NidRepEnum struct {
	Field1           []TheTestEnum `protobuf:"varint,1,rep,enum=enumstringer.TheTestEnum" json:"Field1"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NidRepEnum) Reset()         { *m = NidRepEnum{} }
func (m *NidRepEnum) String() string { return proto.CompactTextString(m) }
func (*NidRepEnum) ProtoMessage()    {}

type NinRepEnum struct {
	Field1           []TheTestEnum `protobuf:"varint,1,rep,enum=enumstringer.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinRepEnum) Reset()         { *m = NinRepEnum{} }
func (m *NinRepEnum) String() string { return proto.CompactTextString(m) }
func (*NinRepEnum) ProtoMessage()    {}

func (m *NinRepEnum) GetField1() []TheTestEnum {
	if m != nil {
		return m.Field1
	}
	return nil
}

func init() {
	proto.RegisterEnum("enumstringer.TheTestEnum", TheTestEnum_name, TheTestEnum_value)
}
func NewPopulatedNidOptEnum(r randyEnumstringer, easy bool) *NidOptEnum {
	this := &NidOptEnum{}
	this.Field1 = TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

func NewPopulatedNinOptEnum(r randyEnumstringer, easy bool) *NinOptEnum {
	this := &NinOptEnum{}
	if r.Intn(10) != 0 {
		v1 := TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
		this.Field1 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

func NewPopulatedNidRepEnum(r randyEnumstringer, easy bool) *NidRepEnum {
	this := &NidRepEnum{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Field1 = make([]TheTestEnum, v2)
		for i := 0; i < v2; i++ {
			this.Field1[i] = TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

func NewPopulatedNinRepEnum(r randyEnumstringer, easy bool) *NinRepEnum {
	this := &NinRepEnum{}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.Field1 = make([]TheTestEnum, v3)
		for i := 0; i < v3; i++ {
			this.Field1[i] = TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

type randyEnumstringer interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEnumstringer(r randyEnumstringer) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringEnumstringer(r randyEnumstringer) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneEnumstringer(r)
	}
	return string(tmps)
}
func randUnrecognizedEnumstringer(r randyEnumstringer, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldEnumstringer(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldEnumstringer(data []byte, r randyEnumstringer, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateEnumstringer(data, uint64(key))
		data = encodeVarintPopulateEnumstringer(data, uint64(r.Int63()))
	case 1:
		data = encodeVarintPopulateEnumstringer(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateEnumstringer(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateEnumstringer(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateEnumstringer(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateEnumstringer(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (this *NidOptEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NidOptEnum)
	if !ok {
		return fmt.Errorf("that is not of type *NidOptEnum")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NidOptEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NidOptEnumbut is not nil && this == nil")
	}
	if this.Field1 != that1.Field1 {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NidOptEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NidOptEnum)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != that1.Field1 {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NinOptEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NinOptEnum)
	if !ok {
		return fmt.Errorf("that is not of type *NinOptEnum")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NinOptEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NinOptEnumbut is not nil && this == nil")
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", *this.Field1, *that1.Field1)
		}
	} else if this.Field1 != nil {
		return fmt.Errorf("this.Field1 == nil && that.Field1 != nil")
	} else if that1.Field1 != nil {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NinOptEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NinOptEnum)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NidRepEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NidRepEnum)
	if !ok {
		return fmt.Errorf("that is not of type *NidRepEnum")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NidRepEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NidRepEnumbut is not nil && this == nil")
	}
	if len(this.Field1) != len(that1.Field1) {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", len(this.Field1), len(that1.Field1))
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return fmt.Errorf("Field1 this[%v](%v) Not Equal that[%v](%v)", i, this.Field1[i], i, that1.Field1[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NidRepEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NidRepEnum)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Field1) != len(that1.Field1) {
		return false
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NinRepEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NinRepEnum)
	if !ok {
		return fmt.Errorf("that is not of type *NinRepEnum")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NinRepEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NinRepEnumbut is not nil && this == nil")
	}
	if len(this.Field1) != len(that1.Field1) {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", len(this.Field1), len(that1.Field1))
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return fmt.Errorf("Field1 this[%v](%v) Not Equal that[%v](%v)", i, this.Field1[i], i, that1.Field1[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NinRepEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NinRepEnum)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Field1) != len(that1.Field1) {
		return false
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
