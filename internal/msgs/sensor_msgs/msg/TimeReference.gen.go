/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo/internal/msgs/builtin_interfaces/msg"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/time_reference.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/TimeReference", TimeReferenceTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/TimeReference", TimeReferenceTypeSupport)
}

// Do not create instances of this type directly. Always use NewTimeReference
// function instead.
type TimeReference struct {
	Header std_msgs_msg.Header `yaml:"header"`// stamp is system time for which measurement was valid
	TimeRef builtin_interfaces_msg.Time `yaml:"time_ref"`// corresponding time from this external source
	Source string `yaml:"source"`// (optional) name of time source
}

// NewTimeReference creates a new TimeReference with default values.
func NewTimeReference() *TimeReference {
	self := TimeReference{}
	self.SetDefaults()
	return &self
}

func (t *TimeReference) Clone() *TimeReference {
	c := &TimeReference{}
	c.Header = *t.Header.Clone()
	c.TimeRef = *t.TimeRef.Clone()
	c.Source = t.Source
	return c
}

func (t *TimeReference) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TimeReference) SetDefaults() {
	t.Header.SetDefaults()
	t.TimeRef.SetDefaults()
	t.Source = ""
}

func (t *TimeReference) GetTypeSupport() types.MessageTypeSupport {
	return TimeReferenceTypeSupport
}

// TimeReferencePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TimeReferencePublisher struct {
	*rclgo.Publisher
}

// NewTimeReferencePublisher creates and returns a new publisher for the
// TimeReference
func NewTimeReferencePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TimeReferencePublisher, error) {
	pub, err := node.NewPublisher(topic_name, TimeReferenceTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TimeReferencePublisher{pub}, nil
}

func (p *TimeReferencePublisher) Publish(msg *TimeReference) error {
	return p.Publisher.Publish(msg)
}

// TimeReferenceSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TimeReferenceSubscription struct {
	*rclgo.Subscription
}

// TimeReferenceSubscriptionCallback type is used to provide a subscription
// handler function for a TimeReferenceSubscription.
type TimeReferenceSubscriptionCallback func(msg *TimeReference, info *rclgo.RmwMessageInfo, err error)

// NewTimeReferenceSubscription creates and returns a new subscription for the
// TimeReference
func NewTimeReferenceSubscription(node *rclgo.Node, topic_name string, subscriptionCallback TimeReferenceSubscriptionCallback) (*TimeReferenceSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TimeReference
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TimeReferenceTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &TimeReferenceSubscription{sub}, nil
}

func (s *TimeReferenceSubscription) TakeMessage(out *TimeReference) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTimeReferenceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTimeReferenceSlice(dst, src []TimeReference) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TimeReferenceTypeSupport types.MessageTypeSupport = _TimeReferenceTypeSupport{}

type _TimeReferenceTypeSupport struct{}

func (t _TimeReferenceTypeSupport) New() types.Message {
	return NewTimeReference()
}

func (t _TimeReferenceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__TimeReference
	return (unsafe.Pointer)(C.sensor_msgs__msg__TimeReference__create())
}

func (t _TimeReferenceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__TimeReference__destroy((*C.sensor_msgs__msg__TimeReference)(pointer_to_free))
}

func (t _TimeReferenceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TimeReference)
	mem := (*C.sensor_msgs__msg__TimeReference)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.time_ref), &m.TimeRef)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.source), m.Source)
}

func (t _TimeReferenceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TimeReference)
	mem := (*C.sensor_msgs__msg__TimeReference)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.TimeRef, unsafe.Pointer(&mem.time_ref))
	primitives.StringAsGoStruct(&m.Source, unsafe.Pointer(&mem.source))
}

func (t _TimeReferenceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__TimeReference())
}

type CTimeReference = C.sensor_msgs__msg__TimeReference
type CTimeReference__Sequence = C.sensor_msgs__msg__TimeReference__Sequence

func TimeReference__Sequence_to_Go(goSlice *[]TimeReference, cSlice CTimeReference__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TimeReference, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TimeReferenceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func TimeReference__Sequence_to_C(cSlice *CTimeReference__Sequence, goSlice []TimeReference) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__TimeReference)(C.malloc(C.sizeof_struct_sensor_msgs__msg__TimeReference * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TimeReferenceTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func TimeReference__Array_to_Go(goSlice []TimeReference, cSlice []CTimeReference) {
	for i := 0; i < len(cSlice); i++ {
		TimeReferenceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TimeReference__Array_to_C(cSlice []CTimeReference, goSlice []TimeReference) {
	for i := 0; i < len(goSlice); i++ {
		TimeReferenceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
