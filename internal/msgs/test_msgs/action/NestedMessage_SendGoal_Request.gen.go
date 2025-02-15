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
	unique_identifier_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/unique_identifier_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/nested_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/NestedMessage_SendGoal_Request", NestedMessage_SendGoal_RequestTypeSupport)
	typemap.RegisterMessage("test_msgs/action/NestedMessage_SendGoal_Request", NestedMessage_SendGoal_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewNestedMessage_SendGoal_Request
// function instead.
type NestedMessage_SendGoal_Request struct {
	GoalID unique_identifier_msgs_msg.UUID `yaml:"goal_id"`
	Goal NestedMessage_Goal `yaml:"goal"`
}

// NewNestedMessage_SendGoal_Request creates a new NestedMessage_SendGoal_Request with default values.
func NewNestedMessage_SendGoal_Request() *NestedMessage_SendGoal_Request {
	self := NestedMessage_SendGoal_Request{}
	self.SetDefaults()
	return &self
}

func (t *NestedMessage_SendGoal_Request) Clone() *NestedMessage_SendGoal_Request {
	c := &NestedMessage_SendGoal_Request{}
	c.GoalID = *t.GoalID.Clone()
	c.Goal = *t.Goal.Clone()
	return c
}

func (t *NestedMessage_SendGoal_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NestedMessage_SendGoal_Request) SetDefaults() {
	t.GoalID.SetDefaults()
	t.Goal.SetDefaults()
}

func (t *NestedMessage_SendGoal_Request) GetTypeSupport() types.MessageTypeSupport {
	return NestedMessage_SendGoal_RequestTypeSupport
}
func (t *NestedMessage_SendGoal_Request) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalID.Uuid)
}

func (t *NestedMessage_SendGoal_Request) SetGoalID(id *types.GoalID) {
	t.GoalID.Uuid = *id
}
func (t *NestedMessage_SendGoal_Request) GetGoalDescription() types.Message {
	return &t.Goal
}

func (t *NestedMessage_SendGoal_Request) SetGoalDescription(desc types.Message) {
	t.Goal = *desc.(*NestedMessage_Goal)
}

// NestedMessage_SendGoal_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type NestedMessage_SendGoal_RequestPublisher struct {
	*rclgo.Publisher
}

// NewNestedMessage_SendGoal_RequestPublisher creates and returns a new publisher for the
// NestedMessage_SendGoal_Request
func NewNestedMessage_SendGoal_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*NestedMessage_SendGoal_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, NestedMessage_SendGoal_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_SendGoal_RequestPublisher{pub}, nil
}

func (p *NestedMessage_SendGoal_RequestPublisher) Publish(msg *NestedMessage_SendGoal_Request) error {
	return p.Publisher.Publish(msg)
}

// NestedMessage_SendGoal_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type NestedMessage_SendGoal_RequestSubscription struct {
	*rclgo.Subscription
}

// NestedMessage_SendGoal_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a NestedMessage_SendGoal_RequestSubscription.
type NestedMessage_SendGoal_RequestSubscriptionCallback func(msg *NestedMessage_SendGoal_Request, info *rclgo.RmwMessageInfo, err error)

// NewNestedMessage_SendGoal_RequestSubscription creates and returns a new subscription for the
// NestedMessage_SendGoal_Request
func NewNestedMessage_SendGoal_RequestSubscription(node *rclgo.Node, topic_name string, subscriptionCallback NestedMessage_SendGoal_RequestSubscriptionCallback) (*NestedMessage_SendGoal_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg NestedMessage_SendGoal_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, NestedMessage_SendGoal_RequestTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_SendGoal_RequestSubscription{sub}, nil
}

func (s *NestedMessage_SendGoal_RequestSubscription) TakeMessage(out *NestedMessage_SendGoal_Request) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneNestedMessage_SendGoal_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNestedMessage_SendGoal_RequestSlice(dst, src []NestedMessage_SendGoal_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NestedMessage_SendGoal_RequestTypeSupport types.MessageTypeSupport = _NestedMessage_SendGoal_RequestTypeSupport{}

type _NestedMessage_SendGoal_RequestTypeSupport struct{}

func (t _NestedMessage_SendGoal_RequestTypeSupport) New() types.Message {
	return NewNestedMessage_SendGoal_Request()
}

func (t _NestedMessage_SendGoal_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__NestedMessage_SendGoal_Request
	return (unsafe.Pointer)(C.test_msgs__action__NestedMessage_SendGoal_Request__create())
}

func (t _NestedMessage_SendGoal_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__NestedMessage_SendGoal_Request__destroy((*C.test_msgs__action__NestedMessage_SendGoal_Request)(pointer_to_free))
}

func (t _NestedMessage_SendGoal_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NestedMessage_SendGoal_Request)
	mem := (*C.test_msgs__action__NestedMessage_SendGoal_Request)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_id), &m.GoalID)
	NestedMessage_GoalTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal), &m.Goal)
}

func (t _NestedMessage_SendGoal_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NestedMessage_SendGoal_Request)
	mem := (*C.test_msgs__action__NestedMessage_SendGoal_Request)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalID, unsafe.Pointer(&mem.goal_id))
	NestedMessage_GoalTypeSupport.AsGoStruct(&m.Goal, unsafe.Pointer(&mem.goal))
}

func (t _NestedMessage_SendGoal_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__NestedMessage_SendGoal_Request())
}

type CNestedMessage_SendGoal_Request = C.test_msgs__action__NestedMessage_SendGoal_Request
type CNestedMessage_SendGoal_Request__Sequence = C.test_msgs__action__NestedMessage_SendGoal_Request__Sequence

func NestedMessage_SendGoal_Request__Sequence_to_Go(goSlice *[]NestedMessage_SendGoal_Request, cSlice CNestedMessage_SendGoal_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NestedMessage_SendGoal_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		NestedMessage_SendGoal_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func NestedMessage_SendGoal_Request__Sequence_to_C(cSlice *CNestedMessage_SendGoal_Request__Sequence, goSlice []NestedMessage_SendGoal_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__action__NestedMessage_SendGoal_Request)(C.malloc(C.sizeof_struct_test_msgs__action__NestedMessage_SendGoal_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		NestedMessage_SendGoal_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func NestedMessage_SendGoal_Request__Array_to_Go(goSlice []NestedMessage_SendGoal_Request, cSlice []CNestedMessage_SendGoal_Request) {
	for i := 0; i < len(cSlice); i++ {
		NestedMessage_SendGoal_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NestedMessage_SendGoal_Request__Array_to_C(cSlice []CNestedMessage_SendGoal_Request, goSlice []NestedMessage_SendGoal_Request) {
	for i := 0; i < len(goSlice); i++ {
		NestedMessage_SendGoal_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
