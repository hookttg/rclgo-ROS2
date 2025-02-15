/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/srv/add_two_ints.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/AddTwoInts_Request", AddTwoInts_RequestTypeSupport)
	typemap.RegisterMessage("example_interfaces/srv/AddTwoInts_Request", AddTwoInts_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewAddTwoInts_Request
// function instead.
type AddTwoInts_Request struct {
	A int64 `yaml:"a"`
	B int64 `yaml:"b"`
}

// NewAddTwoInts_Request creates a new AddTwoInts_Request with default values.
func NewAddTwoInts_Request() *AddTwoInts_Request {
	self := AddTwoInts_Request{}
	self.SetDefaults()
	return &self
}

func (t *AddTwoInts_Request) Clone() *AddTwoInts_Request {
	c := &AddTwoInts_Request{}
	c.A = t.A
	c.B = t.B
	return c
}

func (t *AddTwoInts_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AddTwoInts_Request) SetDefaults() {
	t.A = 0
	t.B = 0
}

func (t *AddTwoInts_Request) GetTypeSupport() types.MessageTypeSupport {
	return AddTwoInts_RequestTypeSupport
}

// AddTwoInts_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type AddTwoInts_RequestPublisher struct {
	*rclgo.Publisher
}

// NewAddTwoInts_RequestPublisher creates and returns a new publisher for the
// AddTwoInts_Request
func NewAddTwoInts_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*AddTwoInts_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, AddTwoInts_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AddTwoInts_RequestPublisher{pub}, nil
}

func (p *AddTwoInts_RequestPublisher) Publish(msg *AddTwoInts_Request) error {
	return p.Publisher.Publish(msg)
}

// AddTwoInts_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type AddTwoInts_RequestSubscription struct {
	*rclgo.Subscription
}

// AddTwoInts_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a AddTwoInts_RequestSubscription.
type AddTwoInts_RequestSubscriptionCallback func(msg *AddTwoInts_Request, info *rclgo.RmwMessageInfo, err error)

// NewAddTwoInts_RequestSubscription creates and returns a new subscription for the
// AddTwoInts_Request
func NewAddTwoInts_RequestSubscription(node *rclgo.Node, topic_name string, subscriptionCallback AddTwoInts_RequestSubscriptionCallback) (*AddTwoInts_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg AddTwoInts_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, AddTwoInts_RequestTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &AddTwoInts_RequestSubscription{sub}, nil
}

func (s *AddTwoInts_RequestSubscription) TakeMessage(out *AddTwoInts_Request) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneAddTwoInts_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAddTwoInts_RequestSlice(dst, src []AddTwoInts_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AddTwoInts_RequestTypeSupport types.MessageTypeSupport = _AddTwoInts_RequestTypeSupport{}

type _AddTwoInts_RequestTypeSupport struct{}

func (t _AddTwoInts_RequestTypeSupport) New() types.Message {
	return NewAddTwoInts_Request()
}

func (t _AddTwoInts_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__srv__AddTwoInts_Request
	return (unsafe.Pointer)(C.example_interfaces__srv__AddTwoInts_Request__create())
}

func (t _AddTwoInts_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__srv__AddTwoInts_Request__destroy((*C.example_interfaces__srv__AddTwoInts_Request)(pointer_to_free))
}

func (t _AddTwoInts_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AddTwoInts_Request)
	mem := (*C.example_interfaces__srv__AddTwoInts_Request)(dst)
	mem.a = C.int64_t(m.A)
	mem.b = C.int64_t(m.B)
}

func (t _AddTwoInts_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AddTwoInts_Request)
	mem := (*C.example_interfaces__srv__AddTwoInts_Request)(ros2_message_buffer)
	m.A = int64(mem.a)
	m.B = int64(mem.b)
}

func (t _AddTwoInts_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__srv__AddTwoInts_Request())
}

type CAddTwoInts_Request = C.example_interfaces__srv__AddTwoInts_Request
type CAddTwoInts_Request__Sequence = C.example_interfaces__srv__AddTwoInts_Request__Sequence

func AddTwoInts_Request__Sequence_to_Go(goSlice *[]AddTwoInts_Request, cSlice CAddTwoInts_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AddTwoInts_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		AddTwoInts_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func AddTwoInts_Request__Sequence_to_C(cSlice *CAddTwoInts_Request__Sequence, goSlice []AddTwoInts_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__srv__AddTwoInts_Request)(C.malloc(C.sizeof_struct_example_interfaces__srv__AddTwoInts_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		AddTwoInts_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func AddTwoInts_Request__Array_to_Go(goSlice []AddTwoInts_Request, cSlice []CAddTwoInts_Request) {
	for i := 0; i < len(cSlice); i++ {
		AddTwoInts_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AddTwoInts_Request__Array_to_C(cSlice []CAddTwoInts_Request, goSlice []AddTwoInts_Request) {
	for i := 0; i < len(goSlice); i++ {
		AddTwoInts_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
