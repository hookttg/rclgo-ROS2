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

#include <example_interfaces/msg/bool.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Bool", BoolTypeSupport)
	typemap.RegisterMessage("example_interfaces/msg/Bool", BoolTypeSupport)
}

// Do not create instances of this type directly. Always use NewBool
// function instead.
type Bool struct {
	Data bool `yaml:"data"`// This is an example message of using a primitive datatype, bool.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewBool creates a new Bool with default values.
func NewBool() *Bool {
	self := Bool{}
	self.SetDefaults()
	return &self
}

func (t *Bool) Clone() *Bool {
	c := &Bool{}
	c.Data = t.Data
	return c
}

func (t *Bool) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Bool) SetDefaults() {
	t.Data = false
}

func (t *Bool) GetTypeSupport() types.MessageTypeSupport {
	return BoolTypeSupport
}

// BoolPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type BoolPublisher struct {
	*rclgo.Publisher
}

// NewBoolPublisher creates and returns a new publisher for the
// Bool
func NewBoolPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*BoolPublisher, error) {
	pub, err := node.NewPublisher(topic_name, BoolTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BoolPublisher{pub}, nil
}

func (p *BoolPublisher) Publish(msg *Bool) error {
	return p.Publisher.Publish(msg)
}

// BoolSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type BoolSubscription struct {
	*rclgo.Subscription
}

// BoolSubscriptionCallback type is used to provide a subscription
// handler function for a BoolSubscription.
type BoolSubscriptionCallback func(msg *Bool, info *rclgo.RmwMessageInfo, err error)

// NewBoolSubscription creates and returns a new subscription for the
// Bool
func NewBoolSubscription(node *rclgo.Node, topic_name string, subscriptionCallback BoolSubscriptionCallback) (*BoolSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Bool
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, BoolTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &BoolSubscription{sub}, nil
}

func (s *BoolSubscription) TakeMessage(out *Bool) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneBoolSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBoolSlice(dst, src []Bool) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BoolTypeSupport types.MessageTypeSupport = _BoolTypeSupport{}

type _BoolTypeSupport struct{}

func (t _BoolTypeSupport) New() types.Message {
	return NewBool()
}

func (t _BoolTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Bool
	return (unsafe.Pointer)(C.example_interfaces__msg__Bool__create())
}

func (t _BoolTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Bool__destroy((*C.example_interfaces__msg__Bool)(pointer_to_free))
}

func (t _BoolTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Bool)
	mem := (*C.example_interfaces__msg__Bool)(dst)
	mem.data = C.bool(m.Data)
}

func (t _BoolTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Bool)
	mem := (*C.example_interfaces__msg__Bool)(ros2_message_buffer)
	m.Data = bool(mem.data)
}

func (t _BoolTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Bool())
}

type CBool = C.example_interfaces__msg__Bool
type CBool__Sequence = C.example_interfaces__msg__Bool__Sequence

func Bool__Sequence_to_Go(goSlice *[]Bool, cSlice CBool__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Bool, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		BoolTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Bool__Sequence_to_C(cSlice *CBool__Sequence, goSlice []Bool) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Bool)(C.malloc(C.sizeof_struct_example_interfaces__msg__Bool * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		BoolTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Bool__Array_to_Go(goSlice []Bool, cSlice []CBool) {
	for i := 0; i < len(cSlice); i++ {
		BoolTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Bool__Array_to_C(cSlice []CBool, goSlice []Bool) {
	for i := 0; i < len(goSlice); i++ {
		BoolTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
