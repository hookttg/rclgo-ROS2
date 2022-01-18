/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int64.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/UInt64", UInt64TypeSupport)
	typemap.RegisterMessage("std_msgs/msg/UInt64", UInt64TypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt64
// function instead.
type UInt64 struct {
	Data uint64 `yaml:"data"`
}

// NewUInt64 creates a new UInt64 with default values.
func NewUInt64() *UInt64 {
	self := UInt64{}
	self.SetDefaults()
	return &self
}

func (t *UInt64) Clone() *UInt64 {
	c := &UInt64{}
	c.Data = t.Data
	return c
}

func (t *UInt64) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt64) SetDefaults() {
	t.Data = 0
}

// UInt64Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type UInt64Publisher struct {
	*rclgo.Publisher
}

// NewUInt64Publisher creates and returns a new publisher for the
// UInt64
func NewUInt64Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*UInt64Publisher, error) {
	pub, err := node.NewPublisher(topic_name, UInt64TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &UInt64Publisher{pub}, nil
}

func (p *UInt64Publisher) Publish(msg *UInt64) error {
	return p.Publisher.Publish(msg)
}

// UInt64Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type UInt64Subscription struct {
	*rclgo.Subscription
}

// UInt64SubscriptionCallback type is used to provide a subscription
// handler function for a UInt64Subscription.
type UInt64SubscriptionCallback func(msg *UInt64, info *rclgo.RmwMessageInfo, err error)

// NewUInt64Subscription creates and returns a new subscription for the
// UInt64
func NewUInt64Subscription(node *rclgo.Node, topic_name string, subscriptionCallback UInt64SubscriptionCallback) (*UInt64Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg UInt64
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, UInt64TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &UInt64Subscription{sub}, nil
}

func (s *UInt64Subscription) TakeMessage(out *UInt64) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneUInt64Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt64Slice(dst, src []UInt64) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt64TypeSupport types.MessageTypeSupport = _UInt64TypeSupport{}

type _UInt64TypeSupport struct{}

func (t _UInt64TypeSupport) New() types.Message {
	return NewUInt64()
}

func (t _UInt64TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt64
	return (unsafe.Pointer)(C.std_msgs__msg__UInt64__create())
}

func (t _UInt64TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt64__destroy((*C.std_msgs__msg__UInt64)(pointer_to_free))
}

func (t _UInt64TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt64)
	mem := (*C.std_msgs__msg__UInt64)(dst)
	mem.data = C.uint64_t(m.Data)
}

func (t _UInt64TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt64)
	mem := (*C.std_msgs__msg__UInt64)(ros2_message_buffer)
	m.Data = uint64(mem.data)
}

func (t _UInt64TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt64())
}

type CUInt64 = C.std_msgs__msg__UInt64
type CUInt64__Sequence = C.std_msgs__msg__UInt64__Sequence

func UInt64__Sequence_to_Go(goSlice *[]UInt64, cSlice CUInt64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt64, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		UInt64TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func UInt64__Sequence_to_C(cSlice *CUInt64__Sequence, goSlice []UInt64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt64)(C.malloc(C.sizeof_struct_std_msgs__msg__UInt64 * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		UInt64TypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func UInt64__Array_to_Go(goSlice []UInt64, cSlice []CUInt64) {
	for i := 0; i < len(cSlice); i++ {
		UInt64TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt64__Array_to_C(cSlice []CUInt64, goSlice []UInt64) {
	for i := 0; i < len(goSlice); i++ {
		UInt64TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
