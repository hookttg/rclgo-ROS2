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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/int16_multi_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Int16MultiArray", Int16MultiArrayTypeSupport)
	typemap.RegisterMessage("example_interfaces/msg/Int16MultiArray", Int16MultiArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewInt16MultiArray
// function instead.
type Int16MultiArray struct {
	Layout MultiArrayLayout `yaml:"layout"`// specification of data layout
	Data []int16 `yaml:"data"`// array of data
}

// NewInt16MultiArray creates a new Int16MultiArray with default values.
func NewInt16MultiArray() *Int16MultiArray {
	self := Int16MultiArray{}
	self.SetDefaults()
	return &self
}

func (t *Int16MultiArray) Clone() *Int16MultiArray {
	c := &Int16MultiArray{}
	c.Layout = *t.Layout.Clone()
	if t.Data != nil {
		c.Data = make([]int16, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Int16MultiArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Int16MultiArray) SetDefaults() {
	t.Layout.SetDefaults()
	t.Data = nil
}

// Int16MultiArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Int16MultiArrayPublisher struct {
	*rclgo.Publisher
}

// NewInt16MultiArrayPublisher creates and returns a new publisher for the
// Int16MultiArray
func NewInt16MultiArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Int16MultiArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Int16MultiArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Int16MultiArrayPublisher{pub}, nil
}

func (p *Int16MultiArrayPublisher) Publish(msg *Int16MultiArray) error {
	return p.Publisher.Publish(msg)
}

// Int16MultiArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Int16MultiArraySubscription struct {
	*rclgo.Subscription
}

// Int16MultiArraySubscriptionCallback type is used to provide a subscription
// handler function for a Int16MultiArraySubscription.
type Int16MultiArraySubscriptionCallback func(msg *Int16MultiArray, info *rclgo.RmwMessageInfo, err error)

// NewInt16MultiArraySubscription creates and returns a new subscription for the
// Int16MultiArray
func NewInt16MultiArraySubscription(node *rclgo.Node, topic_name string, subscriptionCallback Int16MultiArraySubscriptionCallback) (*Int16MultiArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Int16MultiArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Int16MultiArrayTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Int16MultiArraySubscription{sub}, nil
}

func (s *Int16MultiArraySubscription) TakeMessage(out *Int16MultiArray) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneInt16MultiArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInt16MultiArraySlice(dst, src []Int16MultiArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Int16MultiArrayTypeSupport types.MessageTypeSupport = _Int16MultiArrayTypeSupport{}

type _Int16MultiArrayTypeSupport struct{}

func (t _Int16MultiArrayTypeSupport) New() types.Message {
	return NewInt16MultiArray()
}

func (t _Int16MultiArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Int16MultiArray
	return (unsafe.Pointer)(C.example_interfaces__msg__Int16MultiArray__create())
}

func (t _Int16MultiArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Int16MultiArray__destroy((*C.example_interfaces__msg__Int16MultiArray)(pointer_to_free))
}

func (t _Int16MultiArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int16MultiArray)
	mem := (*C.example_interfaces__msg__Int16MultiArray)(dst)
	MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&mem.layout), &m.Layout)
	primitives.Int16__Sequence_to_C((*primitives.CInt16__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _Int16MultiArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int16MultiArray)
	mem := (*C.example_interfaces__msg__Int16MultiArray)(ros2_message_buffer)
	MultiArrayLayoutTypeSupport.AsGoStruct(&m.Layout, unsafe.Pointer(&mem.layout))
	primitives.Int16__Sequence_to_Go(&m.Data, *(*primitives.CInt16__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _Int16MultiArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Int16MultiArray())
}

type CInt16MultiArray = C.example_interfaces__msg__Int16MultiArray
type CInt16MultiArray__Sequence = C.example_interfaces__msg__Int16MultiArray__Sequence

func Int16MultiArray__Sequence_to_Go(goSlice *[]Int16MultiArray, cSlice CInt16MultiArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int16MultiArray, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Int16MultiArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Int16MultiArray__Sequence_to_C(cSlice *CInt16MultiArray__Sequence, goSlice []Int16MultiArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Int16MultiArray)(C.malloc(C.sizeof_struct_example_interfaces__msg__Int16MultiArray * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Int16MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Int16MultiArray__Array_to_Go(goSlice []Int16MultiArray, cSlice []CInt16MultiArray) {
	for i := 0; i < len(cSlice); i++ {
		Int16MultiArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int16MultiArray__Array_to_C(cSlice []CInt16MultiArray, goSlice []Int16MultiArray) {
	for i := 0; i < len(goSlice); i++ {
		Int16MultiArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
