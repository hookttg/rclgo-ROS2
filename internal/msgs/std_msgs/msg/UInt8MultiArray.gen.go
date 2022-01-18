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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/u_int8_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/UInt8MultiArray", UInt8MultiArrayTypeSupport)
	typemap.RegisterMessage("std_msgs/msg/UInt8MultiArray", UInt8MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt8MultiArray
// function instead.
type UInt8MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []uint8 `yaml:"data"`// array of data
}

// NewUInt8MultiArray creates a new UInt8MultiArray with default values.
func NewUInt8MultiArray() *UInt8MultiArray {
	self := UInt8MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *UInt8MultiArray) Clone() *UInt8MultiArray {
	c := &UInt8MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]uint8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *UInt8MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt8MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// UInt8MultiArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type UInt8MultiArrayPublisher struct {
	*rclgo.Publisher
}

// NewUInt8MultiArrayPublisher creates and returns a new publisher for the
// UInt8MultiArray
func NewUInt8MultiArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*UInt8MultiArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, UInt8MultiArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &UInt8MultiArrayPublisher{pub}, nil
}

func (p *UInt8MultiArrayPublisher) Publish(msg *UInt8MultiArray) error {
	return p.Publisher.Publish(msg)
}

// UInt8MultiArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type UInt8MultiArraySubscription struct {
	*rclgo.Subscription
}

// UInt8MultiArraySubscriptionCallback type is used to provide a subscription
// handler function for a UInt8MultiArraySubscription.
type UInt8MultiArraySubscriptionCallback func(msg *UInt8MultiArray, info *rclgo.RmwMessageInfo, err error)

// NewUInt8MultiArraySubscription creates and returns a new subscription for the
// UInt8MultiArray
func NewUInt8MultiArraySubscription(node *rclgo.Node, topic_name string, subscriptionCallback UInt8MultiArraySubscriptionCallback) (*UInt8MultiArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg UInt8MultiArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, UInt8MultiArrayTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &UInt8MultiArraySubscription{sub}, nil
}

func (s *UInt8MultiArraySubscription) TakeMessage(out *UInt8MultiArray) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneUInt8MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt8MultiArraySlice(dst, src []UInt8MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt8MultiArrayTypeSupport types.MessageTypeSupport = _UInt8MultiArrayTypeSupport{}

type _UInt8MultiArrayTypeSupport struct{}

func (t _UInt8MultiArrayTypeSupport) New() types.Message {
	return NewUInt8MultiArray()
}

func (t _UInt8MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt8MultiArray
	return (unsafe.Pointer)(C.std_msgs__msg__UInt8MultiArray__create())
}

func (t _UInt8MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt8MultiArray__destroy((*C.std_msgs__msg__UInt8MultiArray)(pointer_to_free))
}

func (t _UInt8MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt8MultiArray)
	mem := (*C.std_msgs__msg__UInt8MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _UInt8MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt8MultiArray)
	mem := (*C.std_msgs__msg__UInt8MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Uint8__Sequence_to_Go(&m.Data, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _UInt8MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt8MultiArray())
}

type CUInt8MultiArray = C.std_msgs__msg__UInt8MultiArray
type CUInt8MultiArray__Sequence = C.std_msgs__msg__UInt8MultiArray__Sequence

func UInt8MultiArray__Sequence_to_Go(goSlice *[]UInt8MultiArray, cSlice CUInt8MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt8MultiArray, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		UInt8MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func UInt8MultiArray__Sequence_to_C(cSlice *CUInt8MultiArray__Sequence, goSlice []UInt8MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt8MultiArray)(C.malloc(C.sizeof_struct_std_msgs__msg__UInt8MultiArray * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		UInt8MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func UInt8MultiArray__Array_to_Go(goSlice []UInt8MultiArray, cSlice []CUInt8MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		UInt8MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt8MultiArray__Array_to_C(cSlice []CUInt8MultiArray, goSlice []UInt8MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		UInt8MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
