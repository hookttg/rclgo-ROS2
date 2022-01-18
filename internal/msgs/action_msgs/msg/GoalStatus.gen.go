/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package action_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <action_msgs/msg/goal_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("action_msgs/GoalStatus", GoalStatusTypeSupport)
	typemap.RegisterMessage("action_msgs/msg/GoalStatus", GoalStatusTypeSupport)
}
const (
	GoalStatus_STATUS_UNKNOWN int8 = 0// Indicates status has not been properly set.
	GoalStatus_STATUS_ACCEPTED int8 = 1// The goal has been accepted and is awaiting execution.
	GoalStatus_STATUS_EXECUTING int8 = 2// The goal is currently being executed by the action server.
	GoalStatus_STATUS_CANCELING int8 = 3// The client has requested that the goal be canceled and the action server hasaccepted the cancel request.
	GoalStatus_STATUS_SUCCEEDED int8 = 4// The goal was achieved successfully by the action server.
	GoalStatus_STATUS_CANCELED int8 = 5// The goal was canceled after an external request from an action client.
	GoalStatus_STATUS_ABORTED int8 = 6// The goal was terminated by the action server without an external request.
)

// Do not create instances of this type directly. Always use NewGoalStatus
// function instead.
type GoalStatus struct {
	GoalInfo GoalInfo `yaml:"goal_info"`// Goal info (contains ID and timestamp).
	Status int8 `yaml:"status"`// Action goal state-machine status.
}

// NewGoalStatus creates a new GoalStatus with default values.
func NewGoalStatus() *GoalStatus {
	self := GoalStatus{}
	self.SetDefaults()
	return &self
}

func (t *GoalStatus) Clone() *GoalStatus {
	c := &GoalStatus{}
	c.GoalInfo = *t.GoalInfo.Clone()
	c.Status = t.Status
	return c
}

func (t *GoalStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GoalStatus) SetDefaults() {
	t.GoalInfo.SetDefaults()
	t.Status = 0
}
func (t *GoalStatus) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalInfo.GoalId.Uuid)
}

func (t *GoalStatus) SetGoalID(id *types.GoalID) {
	t.GoalInfo.GoalId.Uuid = *id
}

// GoalStatusPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GoalStatusPublisher struct {
	*rclgo.Publisher
}

// NewGoalStatusPublisher creates and returns a new publisher for the
// GoalStatus
func NewGoalStatusPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GoalStatusPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GoalStatusTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GoalStatusPublisher{pub}, nil
}

func (p *GoalStatusPublisher) Publish(msg *GoalStatus) error {
	return p.Publisher.Publish(msg)
}

// GoalStatusSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GoalStatusSubscription struct {
	*rclgo.Subscription
}

// GoalStatusSubscriptionCallback type is used to provide a subscription
// handler function for a GoalStatusSubscription.
type GoalStatusSubscriptionCallback func(msg *GoalStatus, info *rclgo.RmwMessageInfo, err error)

// NewGoalStatusSubscription creates and returns a new subscription for the
// GoalStatus
func NewGoalStatusSubscription(node *rclgo.Node, topic_name string, subscriptionCallback GoalStatusSubscriptionCallback) (*GoalStatusSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GoalStatus
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GoalStatusTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &GoalStatusSubscription{sub}, nil
}

func (s *GoalStatusSubscription) TakeMessage(out *GoalStatus) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGoalStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGoalStatusSlice(dst, src []GoalStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GoalStatusTypeSupport types.MessageTypeSupport = _GoalStatusTypeSupport{}

type _GoalStatusTypeSupport struct{}

func (t _GoalStatusTypeSupport) New() types.Message {
	return NewGoalStatus()
}

func (t _GoalStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.action_msgs__msg__GoalStatus
	return (unsafe.Pointer)(C.action_msgs__msg__GoalStatus__create())
}

func (t _GoalStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.action_msgs__msg__GoalStatus__destroy((*C.action_msgs__msg__GoalStatus)(pointer_to_free))
}

func (t _GoalStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GoalStatus)
	mem := (*C.action_msgs__msg__GoalStatus)(dst)
	GoalInfoTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_info), &m.GoalInfo)
	mem.status = C.int8_t(m.Status)
}

func (t _GoalStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GoalStatus)
	mem := (*C.action_msgs__msg__GoalStatus)(ros2_message_buffer)
	GoalInfoTypeSupport.AsGoStruct(&m.GoalInfo, unsafe.Pointer(&mem.goal_info))
	m.Status = int8(mem.status)
}

func (t _GoalStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__action_msgs__msg__GoalStatus())
}

type CGoalStatus = C.action_msgs__msg__GoalStatus
type CGoalStatus__Sequence = C.action_msgs__msg__GoalStatus__Sequence

func GoalStatus__Sequence_to_Go(goSlice *[]GoalStatus, cSlice CGoalStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalStatus, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GoalStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GoalStatus__Sequence_to_C(cSlice *CGoalStatus__Sequence, goSlice []GoalStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.action_msgs__msg__GoalStatus)(C.malloc(C.sizeof_struct_action_msgs__msg__GoalStatus * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GoalStatusTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GoalStatus__Array_to_Go(goSlice []GoalStatus, cSlice []CGoalStatus) {
	for i := 0; i < len(cSlice); i++ {
		GoalStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GoalStatus__Array_to_C(cSlice []CGoalStatus, goSlice []GoalStatus) {
	for i := 0; i < len(goSlice); i++ {
		GoalStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
