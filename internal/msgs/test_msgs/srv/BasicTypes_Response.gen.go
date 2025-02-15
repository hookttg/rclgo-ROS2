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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/srv/basic_types.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/BasicTypes_Response", BasicTypes_ResponseTypeSupport)
	typemap.RegisterMessage("test_msgs/srv/BasicTypes_Response", BasicTypes_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewBasicTypes_Response
// function instead.
type BasicTypes_Response struct {
	BoolValue bool `yaml:"bool_value"`
	ByteValue byte `yaml:"byte_value"`
	CharValue byte `yaml:"char_value"`
	Float32Value float32 `yaml:"float32_value"`
	Float64Value float64 `yaml:"float64_value"`
	Int8Value int8 `yaml:"int8_value"`
	Uint8Value uint8 `yaml:"uint8_value"`
	Int16Value int16 `yaml:"int16_value"`
	Uint16Value uint16 `yaml:"uint16_value"`
	Int32Value int32 `yaml:"int32_value"`
	Uint32Value uint32 `yaml:"uint32_value"`
	Int64Value int64 `yaml:"int64_value"`
	Uint64Value uint64 `yaml:"uint64_value"`
	StringValue string `yaml:"string_value"`
}

// NewBasicTypes_Response creates a new BasicTypes_Response with default values.
func NewBasicTypes_Response() *BasicTypes_Response {
	self := BasicTypes_Response{}
	self.SetDefaults()
	return &self
}

func (t *BasicTypes_Response) Clone() *BasicTypes_Response {
	c := &BasicTypes_Response{}
	c.BoolValue = t.BoolValue
	c.ByteValue = t.ByteValue
	c.CharValue = t.CharValue
	c.Float32Value = t.Float32Value
	c.Float64Value = t.Float64Value
	c.Int8Value = t.Int8Value
	c.Uint8Value = t.Uint8Value
	c.Int16Value = t.Int16Value
	c.Uint16Value = t.Uint16Value
	c.Int32Value = t.Int32Value
	c.Uint32Value = t.Uint32Value
	c.Int64Value = t.Int64Value
	c.Uint64Value = t.Uint64Value
	c.StringValue = t.StringValue
	return c
}

func (t *BasicTypes_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BasicTypes_Response) SetDefaults() {
	t.BoolValue = false
	t.ByteValue = 0
	t.CharValue = 0
	t.Float32Value = 0
	t.Float64Value = 0
	t.Int8Value = 0
	t.Uint8Value = 0
	t.Int16Value = 0
	t.Uint16Value = 0
	t.Int32Value = 0
	t.Uint32Value = 0
	t.Int64Value = 0
	t.Uint64Value = 0
	t.StringValue = ""
}

func (t *BasicTypes_Response) GetTypeSupport() types.MessageTypeSupport {
	return BasicTypes_ResponseTypeSupport
}

// BasicTypes_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BasicTypes_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewBasicTypes_ResponsePublisher creates and returns a new publisher for the
// BasicTypes_Response
func NewBasicTypes_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BasicTypes_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, BasicTypes_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BasicTypes_ResponsePublisher{pub}, nil
}

func (p *BasicTypes_ResponsePublisher) Publish(msg *BasicTypes_Response) error {
	return p.Publisher.Publish(msg)
}

// BasicTypes_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type BasicTypes_ResponseSubscription struct {
	*rclgo.Subscription
}

// BasicTypes_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a BasicTypes_ResponseSubscription.
type BasicTypes_ResponseSubscriptionCallback func(msg *BasicTypes_Response, info *rclgo.RmwMessageInfo, err error)

// NewBasicTypes_ResponseSubscription creates and returns a new subscription for the
// BasicTypes_Response
func NewBasicTypes_ResponseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback BasicTypes_ResponseSubscriptionCallback) (*BasicTypes_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg BasicTypes_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, BasicTypes_ResponseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &BasicTypes_ResponseSubscription{sub}, nil
}

func (s *BasicTypes_ResponseSubscription) TakeMessage(out *BasicTypes_Response) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneBasicTypes_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBasicTypes_ResponseSlice(dst, src []BasicTypes_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BasicTypes_ResponseTypeSupport types.MessageTypeSupport = _BasicTypes_ResponseTypeSupport{}

type _BasicTypes_ResponseTypeSupport struct{}

func (t _BasicTypes_ResponseTypeSupport) New() types.Message {
	return NewBasicTypes_Response()
}

func (t _BasicTypes_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__srv__BasicTypes_Response
	return (unsafe.Pointer)(C.test_msgs__srv__BasicTypes_Response__create())
}

func (t _BasicTypes_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__srv__BasicTypes_Response__destroy((*C.test_msgs__srv__BasicTypes_Response)(pointer_to_free))
}

func (t _BasicTypes_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BasicTypes_Response)
	mem := (*C.test_msgs__srv__BasicTypes_Response)(dst)
	mem.bool_value = C.bool(m.BoolValue)
	mem.byte_value = C.uint8_t(m.ByteValue)
	mem.char_value = C.uchar(m.CharValue)
	mem.float32_value = C.float(m.Float32Value)
	mem.float64_value = C.double(m.Float64Value)
	mem.int8_value = C.int8_t(m.Int8Value)
	mem.uint8_value = C.uint8_t(m.Uint8Value)
	mem.int16_value = C.int16_t(m.Int16Value)
	mem.uint16_value = C.uint16_t(m.Uint16Value)
	mem.int32_value = C.int32_t(m.Int32Value)
	mem.uint32_value = C.uint32_t(m.Uint32Value)
	mem.int64_value = C.int64_t(m.Int64Value)
	mem.uint64_value = C.uint64_t(m.Uint64Value)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.string_value), m.StringValue)
}

func (t _BasicTypes_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BasicTypes_Response)
	mem := (*C.test_msgs__srv__BasicTypes_Response)(ros2_message_buffer)
	m.BoolValue = bool(mem.bool_value)
	m.ByteValue = byte(mem.byte_value)
	m.CharValue = byte(mem.char_value)
	m.Float32Value = float32(mem.float32_value)
	m.Float64Value = float64(mem.float64_value)
	m.Int8Value = int8(mem.int8_value)
	m.Uint8Value = uint8(mem.uint8_value)
	m.Int16Value = int16(mem.int16_value)
	m.Uint16Value = uint16(mem.uint16_value)
	m.Int32Value = int32(mem.int32_value)
	m.Uint32Value = uint32(mem.uint32_value)
	m.Int64Value = int64(mem.int64_value)
	m.Uint64Value = uint64(mem.uint64_value)
	primitives.StringAsGoStruct(&m.StringValue, unsafe.Pointer(&mem.string_value))
}

func (t _BasicTypes_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__srv__BasicTypes_Response())
}

type CBasicTypes_Response = C.test_msgs__srv__BasicTypes_Response
type CBasicTypes_Response__Sequence = C.test_msgs__srv__BasicTypes_Response__Sequence

func BasicTypes_Response__Sequence_to_Go(goSlice *[]BasicTypes_Response, cSlice CBasicTypes_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BasicTypes_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		BasicTypes_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func BasicTypes_Response__Sequence_to_C(cSlice *CBasicTypes_Response__Sequence, goSlice []BasicTypes_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__srv__BasicTypes_Response)(C.malloc(C.sizeof_struct_test_msgs__srv__BasicTypes_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		BasicTypes_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func BasicTypes_Response__Array_to_Go(goSlice []BasicTypes_Response, cSlice []CBasicTypes_Response) {
	for i := 0; i < len(cSlice); i++ {
		BasicTypes_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BasicTypes_Response__Array_to_C(cSlice []CBasicTypes_Response, goSlice []BasicTypes_Response) {
	for i := 0; i < len(goSlice); i++ {
		BasicTypes_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
