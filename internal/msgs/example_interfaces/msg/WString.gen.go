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

#include <example_interfaces/msg/w_string.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/WString", WStringTypeSupport)
	typemap.RegisterMessage("example_interfaces/msg/WString", WStringTypeSupport)
}

// Do not create instances of this type directly. Always use NewWString
// function instead.
type WString struct {
	Data string `yaml:"data"`// This is an example message of using a primitive datatype, wstring.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewWString creates a new WString with default values.
func NewWString() *WString {
	self := WString{}
	self.SetDefaults()
	return &self
}

func (t *WString) Clone() *WString {
	c := &WString{}
	c.Data = t.Data
	return c
}

func (t *WString) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WString) SetDefaults() {
	t.Data = ""
}

func (t *WString) GetTypeSupport() types.MessageTypeSupport {
	return WStringTypeSupport
}

// WStringPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WStringPublisher struct {
	*rclgo.Publisher
}

// NewWStringPublisher creates and returns a new publisher for the
// WString
func NewWStringPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WStringPublisher, error) {
	pub, err := node.NewPublisher(topic_name, WStringTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WStringPublisher{pub}, nil
}

func (p *WStringPublisher) Publish(msg *WString) error {
	return p.Publisher.Publish(msg)
}

// WStringSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WStringSubscription struct {
	*rclgo.Subscription
}

// WStringSubscriptionCallback type is used to provide a subscription
// handler function for a WStringSubscription.
type WStringSubscriptionCallback func(msg *WString, info *rclgo.RmwMessageInfo, err error)

// NewWStringSubscription creates and returns a new subscription for the
// WString
func NewWStringSubscription(node *rclgo.Node, topic_name string, subscriptionCallback WStringSubscriptionCallback) (*WStringSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WString
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WStringTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &WStringSubscription{sub}, nil
}

func (s *WStringSubscription) TakeMessage(out *WString) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWStringSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWStringSlice(dst, src []WString) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WStringTypeSupport types.MessageTypeSupport = _WStringTypeSupport{}

type _WStringTypeSupport struct{}

func (t _WStringTypeSupport) New() types.Message {
	return NewWString()
}

func (t _WStringTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__WString
	return (unsafe.Pointer)(C.example_interfaces__msg__WString__create())
}

func (t _WStringTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__WString__destroy((*C.example_interfaces__msg__WString)(pointer_to_free))
}

func (t _WStringTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WString)
	mem := (*C.example_interfaces__msg__WString)(dst)
	primitives.U16StringAsCStruct(unsafe.Pointer(&mem.data), m.Data)
}

func (t _WStringTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WString)
	mem := (*C.example_interfaces__msg__WString)(ros2_message_buffer)
	primitives.U16StringAsGoStruct(&m.Data, unsafe.Pointer(&mem.data))
}

func (t _WStringTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__WString())
}

type CWString = C.example_interfaces__msg__WString
type CWString__Sequence = C.example_interfaces__msg__WString__Sequence

func WString__Sequence_to_Go(goSlice *[]WString, cSlice CWString__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WString, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WStringTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WString__Sequence_to_C(cSlice *CWString__Sequence, goSlice []WString) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__WString)(C.malloc(C.sizeof_struct_example_interfaces__msg__WString * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WStringTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WString__Array_to_Go(goSlice []WString, cSlice []CWString) {
	for i := 0; i < len(cSlice); i++ {
		WStringTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WString__Array_to_C(cSlice []CWString, goSlice []WString) {
	for i := 0; i < len(goSlice); i++ {
		WStringTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
