/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/constants.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Constants", ConstantsTypeSupport)
	typemap.RegisterMessage("test_msgs/msg/Constants", ConstantsTypeSupport)
}
const (
	Constants_BOOL_CONST bool = true
	Constants_BYTE_CONST byte = 50
	Constants_CHAR_CONST byte = 100
	Constants_FLOAT32_CONST float32 = 1.125
	Constants_FLOAT64_CONST float64 = 1.125
	Constants_INT8_CONST int8 = -50
	Constants_UINT8_CONST uint8 = 200
	Constants_INT16_CONST int16 = -1000
	Constants_UINT16_CONST uint16 = 2000
	Constants_INT32_CONST int32 = -30000
	Constants_UINT32_CONST uint32 = 60000
	Constants_INT64_CONST int64 = -40000000
	Constants_UINT64_CONST uint64 = 50000000
)

// Do not create instances of this type directly. Always use NewConstants
// function instead.
type Constants struct {
}

// NewConstants creates a new Constants with default values.
func NewConstants() *Constants {
	self := Constants{}
	self.SetDefaults()
	return &self
}

func (t *Constants) Clone() *Constants {
	c := &Constants{}
	return c
}

func (t *Constants) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Constants) SetDefaults() {
}

// ConstantsPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ConstantsPublisher struct {
	*rclgo.Publisher
}

// NewConstantsPublisher creates and returns a new publisher for the
// Constants
func NewConstantsPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ConstantsPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ConstantsTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ConstantsPublisher{pub}, nil
}

func (p *ConstantsPublisher) Publish(msg *Constants) error {
	return p.Publisher.Publish(msg)
}

// ConstantsSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ConstantsSubscription struct {
	*rclgo.Subscription
}

// ConstantsSubscriptionCallback type is used to provide a subscription
// handler function for a ConstantsSubscription.
type ConstantsSubscriptionCallback func(msg *Constants, info *rclgo.RmwMessageInfo, err error)

// NewConstantsSubscription creates and returns a new subscription for the
// Constants
func NewConstantsSubscription(node *rclgo.Node, topic_name string, subscriptionCallback ConstantsSubscriptionCallback) (*ConstantsSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Constants
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ConstantsTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &ConstantsSubscription{sub}, nil
}

func (s *ConstantsSubscription) TakeMessage(out *Constants) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneConstantsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneConstantsSlice(dst, src []Constants) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ConstantsTypeSupport types.MessageTypeSupport = _ConstantsTypeSupport{}

type _ConstantsTypeSupport struct{}

func (t _ConstantsTypeSupport) New() types.Message {
	return NewConstants()
}

func (t _ConstantsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Constants
	return (unsafe.Pointer)(C.test_msgs__msg__Constants__create())
}

func (t _ConstantsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Constants__destroy((*C.test_msgs__msg__Constants)(pointer_to_free))
}

func (t _ConstantsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _ConstantsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _ConstantsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Constants())
}

type CConstants = C.test_msgs__msg__Constants
type CConstants__Sequence = C.test_msgs__msg__Constants__Sequence

func Constants__Sequence_to_Go(goSlice *[]Constants, cSlice CConstants__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Constants, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ConstantsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Constants__Sequence_to_C(cSlice *CConstants__Sequence, goSlice []Constants) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Constants)(C.malloc(C.sizeof_struct_test_msgs__msg__Constants * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ConstantsTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Constants__Array_to_Go(goSlice []Constants, cSlice []CConstants) {
	for i := 0; i < len(cSlice); i++ {
		ConstantsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Constants__Array_to_C(cSlice []CConstants, goSlice []Constants) {
	for i := 0; i < len(goSlice); i++ {
		ConstantsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
