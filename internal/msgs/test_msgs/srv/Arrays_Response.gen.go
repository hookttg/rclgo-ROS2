/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	test_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/test_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/srv/arrays.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Arrays_Response", Arrays_ResponseTypeSupport)
	typemap.RegisterMessage("test_msgs/srv/Arrays_Response", Arrays_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewArrays_Response
// function instead.
type Arrays_Response struct {
	BoolValues [3]bool `yaml:"bool_values"`
	ByteValues [3]byte `yaml:"byte_values"`
	CharValues [3]byte `yaml:"char_values"`
	Float32Values [3]float32 `yaml:"float32_values"`
	Float64Values [3]float64 `yaml:"float64_values"`
	Int8Values [3]int8 `yaml:"int8_values"`
	Uint8Values [3]uint8 `yaml:"uint8_values"`
	Int16Values [3]int16 `yaml:"int16_values"`
	Uint16Values [3]uint16 `yaml:"uint16_values"`
	Int32Values [3]int32 `yaml:"int32_values"`
	Uint32Values [3]uint32 `yaml:"uint32_values"`
	Int64Values [3]int64 `yaml:"int64_values"`
	Uint64Values [3]uint64 `yaml:"uint64_values"`
	StringValues [3]string `yaml:"string_values"`
	BasicTypesValues [3]test_msgs_msg.BasicTypes `yaml:"basic_types_values"`
	ConstantsValues [3]test_msgs_msg.Constants `yaml:"constants_values"`
	DefaultsValues [3]test_msgs_msg.Defaults `yaml:"defaults_values"`
	BoolValuesDefault [3]bool `yaml:"bool_values_default"`
	ByteValuesDefault [3]byte `yaml:"byte_values_default"`
	CharValuesDefault [3]byte `yaml:"char_values_default"`
	Float32ValuesDefault [3]float32 `yaml:"float32_values_default"`
	Float64ValuesDefault [3]float64 `yaml:"float64_values_default"`
	Int8ValuesDefault [3]int8 `yaml:"int8_values_default"`
	Uint8ValuesDefault [3]uint8 `yaml:"uint8_values_default"`
	Int16ValuesDefault [3]int16 `yaml:"int16_values_default"`
	Uint16ValuesDefault [3]uint16 `yaml:"uint16_values_default"`
	Int32ValuesDefault [3]int32 `yaml:"int32_values_default"`
	Uint32ValuesDefault [3]uint32 `yaml:"uint32_values_default"`
	Int64ValuesDefault [3]int64 `yaml:"int64_values_default"`
	Uint64ValuesDefault [3]uint64 `yaml:"uint64_values_default"`
	StringValuesDefault [3]string `yaml:"string_values_default"`
}

// NewArrays_Response creates a new Arrays_Response with default values.
func NewArrays_Response() *Arrays_Response {
	self := Arrays_Response{}
	self.SetDefaults()
	return &self
}

func (t *Arrays_Response) Clone() *Arrays_Response {
	c := &Arrays_Response{}
	c.BoolValues = t.BoolValues
	c.ByteValues = t.ByteValues
	c.CharValues = t.CharValues
	c.Float32Values = t.Float32Values
	c.Float64Values = t.Float64Values
	c.Int8Values = t.Int8Values
	c.Uint8Values = t.Uint8Values
	c.Int16Values = t.Int16Values
	c.Uint16Values = t.Uint16Values
	c.Int32Values = t.Int32Values
	c.Uint32Values = t.Uint32Values
	c.Int64Values = t.Int64Values
	c.Uint64Values = t.Uint64Values
	c.StringValues = t.StringValues
	test_msgs_msg.CloneBasicTypesSlice(c.BasicTypesValues[:], t.BasicTypesValues[:])
	test_msgs_msg.CloneConstantsSlice(c.ConstantsValues[:], t.ConstantsValues[:])
	test_msgs_msg.CloneDefaultsSlice(c.DefaultsValues[:], t.DefaultsValues[:])
	c.BoolValuesDefault = t.BoolValuesDefault
	c.ByteValuesDefault = t.ByteValuesDefault
	c.CharValuesDefault = t.CharValuesDefault
	c.Float32ValuesDefault = t.Float32ValuesDefault
	c.Float64ValuesDefault = t.Float64ValuesDefault
	c.Int8ValuesDefault = t.Int8ValuesDefault
	c.Uint8ValuesDefault = t.Uint8ValuesDefault
	c.Int16ValuesDefault = t.Int16ValuesDefault
	c.Uint16ValuesDefault = t.Uint16ValuesDefault
	c.Int32ValuesDefault = t.Int32ValuesDefault
	c.Uint32ValuesDefault = t.Uint32ValuesDefault
	c.Int64ValuesDefault = t.Int64ValuesDefault
	c.Uint64ValuesDefault = t.Uint64ValuesDefault
	c.StringValuesDefault = t.StringValuesDefault
	return c
}

func (t *Arrays_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Arrays_Response) SetDefaults() {
	t.BoolValues = [3]bool{}
	t.ByteValues = [3]byte{}
	t.CharValues = [3]byte{}
	t.Float32Values = [3]float32{}
	t.Float64Values = [3]float64{}
	t.Int8Values = [3]int8{}
	t.Uint8Values = [3]uint8{}
	t.Int16Values = [3]int16{}
	t.Uint16Values = [3]uint16{}
	t.Int32Values = [3]int32{}
	t.Uint32Values = [3]uint32{}
	t.Int64Values = [3]int64{}
	t.Uint64Values = [3]uint64{}
	t.StringValues = [3]string{}
	for i := range t.BasicTypesValues {
		t.BasicTypesValues[i].SetDefaults()
	}
	for i := range t.ConstantsValues {
		t.ConstantsValues[i].SetDefaults()
	}
	for i := range t.DefaultsValues {
		t.DefaultsValues[i].SetDefaults()
	}
	t.BoolValuesDefault = [3]bool{false,true,false}
	t.ByteValuesDefault = [3]byte{0,1,255}
	t.CharValuesDefault = [3]byte{0,1,127}
	t.Float32ValuesDefault = [3]float32{1.125,0.0,-1.125}
	t.Float64ValuesDefault = [3]float64{3.1415,0.0,-3.1415}
	t.Int8ValuesDefault = [3]int8{0,127,-128}
	t.Uint8ValuesDefault = [3]uint8{0,1,255}
	t.Int16ValuesDefault = [3]int16{0,32767,-32768}
	t.Uint16ValuesDefault = [3]uint16{0,1,65535}
	t.Int32ValuesDefault = [3]int32{0,2147483647,-2147483648}
	t.Uint32ValuesDefault = [3]uint32{0,1,4294967295}
	t.Int64ValuesDefault = [3]int64{0,9223372036854775807,-9223372036854775808}
	t.Uint64ValuesDefault = [3]uint64{0,1,18446744073709551615}
	t.StringValuesDefault = [3]string{"","max value","min value"}
}

func (t *Arrays_Response) GetTypeSupport() types.MessageTypeSupport {
	return Arrays_ResponseTypeSupport
}

// Arrays_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Arrays_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewArrays_ResponsePublisher creates and returns a new publisher for the
// Arrays_Response
func NewArrays_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Arrays_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, Arrays_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Arrays_ResponsePublisher{pub}, nil
}

func (p *Arrays_ResponsePublisher) Publish(msg *Arrays_Response) error {
	return p.Publisher.Publish(msg)
}

// Arrays_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Arrays_ResponseSubscription struct {
	*rclgo.Subscription
}

// Arrays_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a Arrays_ResponseSubscription.
type Arrays_ResponseSubscriptionCallback func(msg *Arrays_Response, info *rclgo.RmwMessageInfo, err error)

// NewArrays_ResponseSubscription creates and returns a new subscription for the
// Arrays_Response
func NewArrays_ResponseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback Arrays_ResponseSubscriptionCallback) (*Arrays_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Arrays_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Arrays_ResponseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Arrays_ResponseSubscription{sub}, nil
}

func (s *Arrays_ResponseSubscription) TakeMessage(out *Arrays_Response) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneArrays_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneArrays_ResponseSlice(dst, src []Arrays_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Arrays_ResponseTypeSupport types.MessageTypeSupport = _Arrays_ResponseTypeSupport{}

type _Arrays_ResponseTypeSupport struct{}

func (t _Arrays_ResponseTypeSupport) New() types.Message {
	return NewArrays_Response()
}

func (t _Arrays_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__srv__Arrays_Response
	return (unsafe.Pointer)(C.test_msgs__srv__Arrays_Response__create())
}

func (t _Arrays_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__srv__Arrays_Response__destroy((*C.test_msgs__srv__Arrays_Response)(pointer_to_free))
}

func (t _Arrays_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Arrays_Response)
	mem := (*C.test_msgs__srv__Arrays_Response)(dst)
	cSlice_bool_values := mem.bool_values[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values)), m.BoolValues[:])
	cSlice_byte_values := mem.byte_values[:]
	primitives.Byte__Array_to_C(*(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values)), m.ByteValues[:])
	cSlice_char_values := mem.char_values[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values)), m.CharValues[:])
	cSlice_float32_values := mem.float32_values[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values)), m.Float32Values[:])
	cSlice_float64_values := mem.float64_values[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values)), m.Float64Values[:])
	cSlice_int8_values := mem.int8_values[:]
	primitives.Int8__Array_to_C(*(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values)), m.Int8Values[:])
	cSlice_uint8_values := mem.uint8_values[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values)), m.Uint8Values[:])
	cSlice_int16_values := mem.int16_values[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values)), m.Int16Values[:])
	cSlice_uint16_values := mem.uint16_values[:]
	primitives.Uint16__Array_to_C(*(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values)), m.Uint16Values[:])
	cSlice_int32_values := mem.int32_values[:]
	primitives.Int32__Array_to_C(*(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values)), m.Int32Values[:])
	cSlice_uint32_values := mem.uint32_values[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values)), m.Uint32Values[:])
	cSlice_int64_values := mem.int64_values[:]
	primitives.Int64__Array_to_C(*(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values)), m.Int64Values[:])
	cSlice_uint64_values := mem.uint64_values[:]
	primitives.Uint64__Array_to_C(*(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values)), m.Uint64Values[:])
	cSlice_string_values := mem.string_values[:]
	primitives.String__Array_to_C(*(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values)), m.StringValues[:])
	cSlice_basic_types_values := mem.basic_types_values[:]
	test_msgs_msg.BasicTypes__Array_to_C(*(*[]test_msgs_msg.CBasicTypes)(unsafe.Pointer(&cSlice_basic_types_values)), m.BasicTypesValues[:])
	cSlice_constants_values := mem.constants_values[:]
	test_msgs_msg.Constants__Array_to_C(*(*[]test_msgs_msg.CConstants)(unsafe.Pointer(&cSlice_constants_values)), m.ConstantsValues[:])
	cSlice_defaults_values := mem.defaults_values[:]
	test_msgs_msg.Defaults__Array_to_C(*(*[]test_msgs_msg.CDefaults)(unsafe.Pointer(&cSlice_defaults_values)), m.DefaultsValues[:])
	cSlice_bool_values_default := mem.bool_values_default[:]
	primitives.Bool__Array_to_C(*(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values_default)), m.BoolValuesDefault[:])
	cSlice_byte_values_default := mem.byte_values_default[:]
	primitives.Byte__Array_to_C(*(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values_default)), m.ByteValuesDefault[:])
	cSlice_char_values_default := mem.char_values_default[:]
	primitives.Char__Array_to_C(*(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values_default)), m.CharValuesDefault[:])
	cSlice_float32_values_default := mem.float32_values_default[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values_default)), m.Float32ValuesDefault[:])
	cSlice_float64_values_default := mem.float64_values_default[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values_default)), m.Float64ValuesDefault[:])
	cSlice_int8_values_default := mem.int8_values_default[:]
	primitives.Int8__Array_to_C(*(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values_default)), m.Int8ValuesDefault[:])
	cSlice_uint8_values_default := mem.uint8_values_default[:]
	primitives.Uint8__Array_to_C(*(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values_default)), m.Uint8ValuesDefault[:])
	cSlice_int16_values_default := mem.int16_values_default[:]
	primitives.Int16__Array_to_C(*(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values_default)), m.Int16ValuesDefault[:])
	cSlice_uint16_values_default := mem.uint16_values_default[:]
	primitives.Uint16__Array_to_C(*(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values_default)), m.Uint16ValuesDefault[:])
	cSlice_int32_values_default := mem.int32_values_default[:]
	primitives.Int32__Array_to_C(*(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values_default)), m.Int32ValuesDefault[:])
	cSlice_uint32_values_default := mem.uint32_values_default[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values_default)), m.Uint32ValuesDefault[:])
	cSlice_int64_values_default := mem.int64_values_default[:]
	primitives.Int64__Array_to_C(*(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values_default)), m.Int64ValuesDefault[:])
	cSlice_uint64_values_default := mem.uint64_values_default[:]
	primitives.Uint64__Array_to_C(*(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values_default)), m.Uint64ValuesDefault[:])
	cSlice_string_values_default := mem.string_values_default[:]
	primitives.String__Array_to_C(*(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values_default)), m.StringValuesDefault[:])
}

func (t _Arrays_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Arrays_Response)
	mem := (*C.test_msgs__srv__Arrays_Response)(ros2_message_buffer)
	cSlice_bool_values := mem.bool_values[:]
	primitives.Bool__Array_to_Go(m.BoolValues[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values)))
	cSlice_byte_values := mem.byte_values[:]
	primitives.Byte__Array_to_Go(m.ByteValues[:], *(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values)))
	cSlice_char_values := mem.char_values[:]
	primitives.Char__Array_to_Go(m.CharValues[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values)))
	cSlice_float32_values := mem.float32_values[:]
	primitives.Float32__Array_to_Go(m.Float32Values[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values)))
	cSlice_float64_values := mem.float64_values[:]
	primitives.Float64__Array_to_Go(m.Float64Values[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values)))
	cSlice_int8_values := mem.int8_values[:]
	primitives.Int8__Array_to_Go(m.Int8Values[:], *(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values)))
	cSlice_uint8_values := mem.uint8_values[:]
	primitives.Uint8__Array_to_Go(m.Uint8Values[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values)))
	cSlice_int16_values := mem.int16_values[:]
	primitives.Int16__Array_to_Go(m.Int16Values[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values)))
	cSlice_uint16_values := mem.uint16_values[:]
	primitives.Uint16__Array_to_Go(m.Uint16Values[:], *(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values)))
	cSlice_int32_values := mem.int32_values[:]
	primitives.Int32__Array_to_Go(m.Int32Values[:], *(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values)))
	cSlice_uint32_values := mem.uint32_values[:]
	primitives.Uint32__Array_to_Go(m.Uint32Values[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values)))
	cSlice_int64_values := mem.int64_values[:]
	primitives.Int64__Array_to_Go(m.Int64Values[:], *(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values)))
	cSlice_uint64_values := mem.uint64_values[:]
	primitives.Uint64__Array_to_Go(m.Uint64Values[:], *(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values)))
	cSlice_string_values := mem.string_values[:]
	primitives.String__Array_to_Go(m.StringValues[:], *(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values)))
	cSlice_basic_types_values := mem.basic_types_values[:]
	test_msgs_msg.BasicTypes__Array_to_Go(m.BasicTypesValues[:], *(*[]test_msgs_msg.CBasicTypes)(unsafe.Pointer(&cSlice_basic_types_values)))
	cSlice_constants_values := mem.constants_values[:]
	test_msgs_msg.Constants__Array_to_Go(m.ConstantsValues[:], *(*[]test_msgs_msg.CConstants)(unsafe.Pointer(&cSlice_constants_values)))
	cSlice_defaults_values := mem.defaults_values[:]
	test_msgs_msg.Defaults__Array_to_Go(m.DefaultsValues[:], *(*[]test_msgs_msg.CDefaults)(unsafe.Pointer(&cSlice_defaults_values)))
	cSlice_bool_values_default := mem.bool_values_default[:]
	primitives.Bool__Array_to_Go(m.BoolValuesDefault[:], *(*[]primitives.CBool)(unsafe.Pointer(&cSlice_bool_values_default)))
	cSlice_byte_values_default := mem.byte_values_default[:]
	primitives.Byte__Array_to_Go(m.ByteValuesDefault[:], *(*[]primitives.CByte)(unsafe.Pointer(&cSlice_byte_values_default)))
	cSlice_char_values_default := mem.char_values_default[:]
	primitives.Char__Array_to_Go(m.CharValuesDefault[:], *(*[]primitives.CChar)(unsafe.Pointer(&cSlice_char_values_default)))
	cSlice_float32_values_default := mem.float32_values_default[:]
	primitives.Float32__Array_to_Go(m.Float32ValuesDefault[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_float32_values_default)))
	cSlice_float64_values_default := mem.float64_values_default[:]
	primitives.Float64__Array_to_Go(m.Float64ValuesDefault[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_float64_values_default)))
	cSlice_int8_values_default := mem.int8_values_default[:]
	primitives.Int8__Array_to_Go(m.Int8ValuesDefault[:], *(*[]primitives.CInt8)(unsafe.Pointer(&cSlice_int8_values_default)))
	cSlice_uint8_values_default := mem.uint8_values_default[:]
	primitives.Uint8__Array_to_Go(m.Uint8ValuesDefault[:], *(*[]primitives.CUint8)(unsafe.Pointer(&cSlice_uint8_values_default)))
	cSlice_int16_values_default := mem.int16_values_default[:]
	primitives.Int16__Array_to_Go(m.Int16ValuesDefault[:], *(*[]primitives.CInt16)(unsafe.Pointer(&cSlice_int16_values_default)))
	cSlice_uint16_values_default := mem.uint16_values_default[:]
	primitives.Uint16__Array_to_Go(m.Uint16ValuesDefault[:], *(*[]primitives.CUint16)(unsafe.Pointer(&cSlice_uint16_values_default)))
	cSlice_int32_values_default := mem.int32_values_default[:]
	primitives.Int32__Array_to_Go(m.Int32ValuesDefault[:], *(*[]primitives.CInt32)(unsafe.Pointer(&cSlice_int32_values_default)))
	cSlice_uint32_values_default := mem.uint32_values_default[:]
	primitives.Uint32__Array_to_Go(m.Uint32ValuesDefault[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_uint32_values_default)))
	cSlice_int64_values_default := mem.int64_values_default[:]
	primitives.Int64__Array_to_Go(m.Int64ValuesDefault[:], *(*[]primitives.CInt64)(unsafe.Pointer(&cSlice_int64_values_default)))
	cSlice_uint64_values_default := mem.uint64_values_default[:]
	primitives.Uint64__Array_to_Go(m.Uint64ValuesDefault[:], *(*[]primitives.CUint64)(unsafe.Pointer(&cSlice_uint64_values_default)))
	cSlice_string_values_default := mem.string_values_default[:]
	primitives.String__Array_to_Go(m.StringValuesDefault[:], *(*[]primitives.CString)(unsafe.Pointer(&cSlice_string_values_default)))
}

func (t _Arrays_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__srv__Arrays_Response())
}

type CArrays_Response = C.test_msgs__srv__Arrays_Response
type CArrays_Response__Sequence = C.test_msgs__srv__Arrays_Response__Sequence

func Arrays_Response__Sequence_to_Go(goSlice *[]Arrays_Response, cSlice CArrays_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Arrays_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Arrays_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Arrays_Response__Sequence_to_C(cSlice *CArrays_Response__Sequence, goSlice []Arrays_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__srv__Arrays_Response)(C.malloc(C.sizeof_struct_test_msgs__srv__Arrays_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Arrays_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Arrays_Response__Array_to_Go(goSlice []Arrays_Response, cSlice []CArrays_Response) {
	for i := 0; i < len(cSlice); i++ {
		Arrays_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Arrays_Response__Array_to_C(cSlice []CArrays_Response, goSlice []Arrays_Response) {
	for i := 0; i < len(goSlice); i++ {
		Arrays_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
