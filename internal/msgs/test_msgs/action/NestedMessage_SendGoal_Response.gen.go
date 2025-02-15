/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo/internal/msgs/builtin_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/nested_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/NestedMessage_SendGoal_Response", NestedMessage_SendGoal_ResponseTypeSupport)
	typemap.RegisterMessage("test_msgs/action/NestedMessage_SendGoal_Response", NestedMessage_SendGoal_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewNestedMessage_SendGoal_Response
// function instead.
type NestedMessage_SendGoal_Response struct {
	Accepted bool `yaml:"accepted"`
	Stamp builtin_interfaces_msg.Time `yaml:"stamp"`
}

// NewNestedMessage_SendGoal_Response creates a new NestedMessage_SendGoal_Response with default values.
func NewNestedMessage_SendGoal_Response() *NestedMessage_SendGoal_Response {
	self := NestedMessage_SendGoal_Response{}
	self.SetDefaults()
	return &self
}

func (t *NestedMessage_SendGoal_Response) Clone() *NestedMessage_SendGoal_Response {
	c := &NestedMessage_SendGoal_Response{}
	c.Accepted = t.Accepted
	c.Stamp = *t.Stamp.Clone()
	return c
}

func (t *NestedMessage_SendGoal_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NestedMessage_SendGoal_Response) SetDefaults() {
	t.Accepted = false
	t.Stamp.SetDefaults()
}

func (t *NestedMessage_SendGoal_Response) GetTypeSupport() types.MessageTypeSupport {
	return NestedMessage_SendGoal_ResponseTypeSupport
}
func (t *NestedMessage_SendGoal_Response) GetGoalAccepted() bool {
	return t.Accepted
}

// NestedMessage_SendGoal_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type NestedMessage_SendGoal_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewNestedMessage_SendGoal_ResponsePublisher creates and returns a new publisher for the
// NestedMessage_SendGoal_Response
func NewNestedMessage_SendGoal_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*NestedMessage_SendGoal_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, NestedMessage_SendGoal_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_SendGoal_ResponsePublisher{pub}, nil
}

func (p *NestedMessage_SendGoal_ResponsePublisher) Publish(msg *NestedMessage_SendGoal_Response) error {
	return p.Publisher.Publish(msg)
}

// NestedMessage_SendGoal_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type NestedMessage_SendGoal_ResponseSubscription struct {
	*rclgo.Subscription
}

// NestedMessage_SendGoal_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a NestedMessage_SendGoal_ResponseSubscription.
type NestedMessage_SendGoal_ResponseSubscriptionCallback func(msg *NestedMessage_SendGoal_Response, info *rclgo.RmwMessageInfo, err error)

// NewNestedMessage_SendGoal_ResponseSubscription creates and returns a new subscription for the
// NestedMessage_SendGoal_Response
func NewNestedMessage_SendGoal_ResponseSubscription(node *rclgo.Node, topic_name string, subscriptionCallback NestedMessage_SendGoal_ResponseSubscriptionCallback) (*NestedMessage_SendGoal_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg NestedMessage_SendGoal_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, NestedMessage_SendGoal_ResponseTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_SendGoal_ResponseSubscription{sub}, nil
}

func (s *NestedMessage_SendGoal_ResponseSubscription) TakeMessage(out *NestedMessage_SendGoal_Response) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneNestedMessage_SendGoal_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNestedMessage_SendGoal_ResponseSlice(dst, src []NestedMessage_SendGoal_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NestedMessage_SendGoal_ResponseTypeSupport types.MessageTypeSupport = _NestedMessage_SendGoal_ResponseTypeSupport{}

type _NestedMessage_SendGoal_ResponseTypeSupport struct{}

func (t _NestedMessage_SendGoal_ResponseTypeSupport) New() types.Message {
	return NewNestedMessage_SendGoal_Response()
}

func (t _NestedMessage_SendGoal_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__NestedMessage_SendGoal_Response
	return (unsafe.Pointer)(C.test_msgs__action__NestedMessage_SendGoal_Response__create())
}

func (t _NestedMessage_SendGoal_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__NestedMessage_SendGoal_Response__destroy((*C.test_msgs__action__NestedMessage_SendGoal_Response)(pointer_to_free))
}

func (t _NestedMessage_SendGoal_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NestedMessage_SendGoal_Response)
	mem := (*C.test_msgs__action__NestedMessage_SendGoal_Response)(dst)
	mem.accepted = C.bool(m.Accepted)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.stamp), &m.Stamp)
}

func (t _NestedMessage_SendGoal_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NestedMessage_SendGoal_Response)
	mem := (*C.test_msgs__action__NestedMessage_SendGoal_Response)(ros2_message_buffer)
	m.Accepted = bool(mem.accepted)
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Stamp, unsafe.Pointer(&mem.stamp))
}

func (t _NestedMessage_SendGoal_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__NestedMessage_SendGoal_Response())
}

type CNestedMessage_SendGoal_Response = C.test_msgs__action__NestedMessage_SendGoal_Response
type CNestedMessage_SendGoal_Response__Sequence = C.test_msgs__action__NestedMessage_SendGoal_Response__Sequence

func NestedMessage_SendGoal_Response__Sequence_to_Go(goSlice *[]NestedMessage_SendGoal_Response, cSlice CNestedMessage_SendGoal_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NestedMessage_SendGoal_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		NestedMessage_SendGoal_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func NestedMessage_SendGoal_Response__Sequence_to_C(cSlice *CNestedMessage_SendGoal_Response__Sequence, goSlice []NestedMessage_SendGoal_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__action__NestedMessage_SendGoal_Response)(C.malloc(C.sizeof_struct_test_msgs__action__NestedMessage_SendGoal_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		NestedMessage_SendGoal_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func NestedMessage_SendGoal_Response__Array_to_Go(goSlice []NestedMessage_SendGoal_Response, cSlice []CNestedMessage_SendGoal_Response) {
	for i := 0; i < len(cSlice); i++ {
		NestedMessage_SendGoal_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NestedMessage_SendGoal_Response__Array_to_C(cSlice []CNestedMessage_SendGoal_Response, goSlice []NestedMessage_SendGoal_Response) {
	for i := 0; i < len(goSlice); i++ {
		NestedMessage_SendGoal_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
