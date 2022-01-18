/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/int64.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Int64", Int64TypeSupport)
	typemap.RegisterMessage("example_interfaces/msg/Int64", Int64TypeSupport)
}

// Do not create instances of this type directly. Always use NewInt64
// function instead.
type Int64 struct {
	Data int64 `yaml:"data"`// This is an example message of using a primitive datatype, int64.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewInt64 creates a new Int64 with default values.
func NewInt64() *Int64 {
	self := Int64{}
	self.SetDefaults()
	return &self
}

func (t *Int64) Clone() *Int64 {
	c := &Int64{}
	c.Data = t.Data
	return c
}

func (t *Int64) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Int64) SetDefaults() {
	t.Data = 0
}

// Int64Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type Int64Publisher struct {
	*rclgo.Publisher
}

// NewInt64Publisher creates and returns a new publisher for the
// Int64
func NewInt64Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Int64Publisher, error) {
	pub, err := node.NewPublisher(topic_name, Int64TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Int64Publisher{pub}, nil
}

func (p *Int64Publisher) Publish(msg *Int64) error {
	return p.Publisher.Publish(msg)
}

// Int64Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type Int64Subscription struct {
	*rclgo.Subscription
}

// Int64SubscriptionCallback type is used to provide a subscription
// handler function for a Int64Subscription.
type Int64SubscriptionCallback func(msg *Int64, info *rclgo.RmwMessageInfo, err error)

// NewInt64Subscription creates and returns a new subscription for the
// Int64
func NewInt64Subscription(node *rclgo.Node, topic_name string, subscriptionCallback Int64SubscriptionCallback) (*Int64Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Int64
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Int64TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Int64Subscription{sub}, nil
}

func (s *Int64Subscription) TakeMessage(out *Int64) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneInt64Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInt64Slice(dst, src []Int64) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Int64TypeSupport types.MessageTypeSupport = _Int64TypeSupport{}

type _Int64TypeSupport struct{}

func (t _Int64TypeSupport) New() types.Message {
	return NewInt64()
}

func (t _Int64TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Int64
	return (unsafe.Pointer)(C.example_interfaces__msg__Int64__create())
}

func (t _Int64TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Int64__destroy((*C.example_interfaces__msg__Int64)(pointer_to_free))
}

func (t _Int64TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int64)
	mem := (*C.example_interfaces__msg__Int64)(dst)
	mem.data = C.int64_t(m.Data)
}

func (t _Int64TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int64)
	mem := (*C.example_interfaces__msg__Int64)(ros2_message_buffer)
	m.Data = int64(mem.data)
}

func (t _Int64TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Int64())
}

type CInt64 = C.example_interfaces__msg__Int64
type CInt64__Sequence = C.example_interfaces__msg__Int64__Sequence

func Int64__Sequence_to_Go(goSlice *[]Int64, cSlice CInt64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int64, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Int64TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Int64__Sequence_to_C(cSlice *CInt64__Sequence, goSlice []Int64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Int64)(C.malloc(C.sizeof_struct_example_interfaces__msg__Int64 * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Int64TypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Int64__Array_to_Go(goSlice []Int64, cSlice []CInt64) {
	for i := 0; i < len(cSlice); i++ {
		Int64TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int64__Array_to_C(cSlice []CInt64, goSlice []Int64) {
	for i := 0; i < len(goSlice); i++ {
		Int64TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
