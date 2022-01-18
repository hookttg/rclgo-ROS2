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

#include <std_msgs/msg/float64.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Float64", Float64TypeSupport)
	typemap.RegisterMessage("std_msgs/msg/Float64", Float64TypeSupport)
}

// Do not create instances of this type directly. Always use NewFloat64
// function instead.
type Float64 struct {
	Data float64 `yaml:"data"`
}

// NewFloat64 creates a new Float64 with default values.
func NewFloat64() *Float64 {
	self := Float64{}
	self.SetDefaults()
	return &self
}

func (t *Float64) Clone() *Float64 {
	c := &Float64{}
	c.Data = t.Data
	return c
}

func (t *Float64) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Float64) SetDefaults() {
	t.Data = 0
}

// Float64Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type Float64Publisher struct {
	*rclgo.Publisher
}

// NewFloat64Publisher creates and returns a new publisher for the
// Float64
func NewFloat64Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Float64Publisher, error) {
	pub, err := node.NewPublisher(topic_name, Float64TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Float64Publisher{pub}, nil
}

func (p *Float64Publisher) Publish(msg *Float64) error {
	return p.Publisher.Publish(msg)
}

// Float64Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type Float64Subscription struct {
	*rclgo.Subscription
}

// Float64SubscriptionCallback type is used to provide a subscription
// handler function for a Float64Subscription.
type Float64SubscriptionCallback func(msg *Float64, info *rclgo.RmwMessageInfo, err error)

// NewFloat64Subscription creates and returns a new subscription for the
// Float64
func NewFloat64Subscription(node *rclgo.Node, topic_name string, subscriptionCallback Float64SubscriptionCallback) (*Float64Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Float64
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Float64TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Float64Subscription{sub}, nil
}

func (s *Float64Subscription) TakeMessage(out *Float64) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFloat64Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFloat64Slice(dst, src []Float64) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Float64TypeSupport types.MessageTypeSupport = _Float64TypeSupport{}

type _Float64TypeSupport struct{}

func (t _Float64TypeSupport) New() types.Message {
	return NewFloat64()
}

func (t _Float64TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Float64
	return (unsafe.Pointer)(C.std_msgs__msg__Float64__create())
}

func (t _Float64TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Float64__destroy((*C.std_msgs__msg__Float64)(pointer_to_free))
}

func (t _Float64TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Float64)
	mem := (*C.std_msgs__msg__Float64)(dst)
	mem.data = C.double(m.Data)
}

func (t _Float64TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Float64)
	mem := (*C.std_msgs__msg__Float64)(ros2_message_buffer)
	m.Data = float64(mem.data)
}

func (t _Float64TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Float64())
}

type CFloat64 = C.std_msgs__msg__Float64
type CFloat64__Sequence = C.std_msgs__msg__Float64__Sequence

func Float64__Sequence_to_Go(goSlice *[]Float64, cSlice CFloat64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Float64, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Float64TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Float64__Sequence_to_C(cSlice *CFloat64__Sequence, goSlice []Float64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Float64)(C.malloc(C.sizeof_struct_std_msgs__msg__Float64 * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Float64TypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Float64__Array_to_Go(goSlice []Float64, cSlice []CFloat64) {
	for i := 0; i < len(cSlice); i++ {
		Float64TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Float64__Array_to_C(cSlice []CFloat64, goSlice []Float64) {
	for i := 0; i < len(goSlice); i++ {
		Float64TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
