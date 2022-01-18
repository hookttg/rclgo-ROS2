/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package builtin_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <builtin_interfaces/msg/duration.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("builtin_interfaces/Duration", DurationTypeSupport)
	typemap.RegisterMessage("builtin_interfaces/msg/Duration", DurationTypeSupport)
}

// Do not create instances of this type directly. Always use NewDuration
// function instead.
type Duration struct {
	Sec int32 `yaml:"sec"`// Seconds component, range is valid over any possible int32 value.
	Nanosec uint32 `yaml:"nanosec"`// Nanoseconds component in the range of [0, 10e9).
}

// NewDuration creates a new Duration with default values.
func NewDuration() *Duration {
	self := Duration{}
	self.SetDefaults()
	return &self
}

func (t *Duration) Clone() *Duration {
	c := &Duration{}
	c.Sec = t.Sec
	c.Nanosec = t.Nanosec
	return c
}

func (t *Duration) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Duration) SetDefaults() {
	t.Sec = 0
	t.Nanosec = 0
}

// DurationPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type DurationPublisher struct {
	*rclgo.Publisher
}

// NewDurationPublisher creates and returns a new publisher for the
// Duration
func NewDurationPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*DurationPublisher, error) {
	pub, err := node.NewPublisher(topic_name, DurationTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &DurationPublisher{pub}, nil
}

func (p *DurationPublisher) Publish(msg *Duration) error {
	return p.Publisher.Publish(msg)
}

// DurationSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type DurationSubscription struct {
	*rclgo.Subscription
}

// DurationSubscriptionCallback type is used to provide a subscription
// handler function for a DurationSubscription.
type DurationSubscriptionCallback func(msg *Duration, info *rclgo.RmwMessageInfo, err error)

// NewDurationSubscription creates and returns a new subscription for the
// Duration
func NewDurationSubscription(node *rclgo.Node, topic_name string, subscriptionCallback DurationSubscriptionCallback) (*DurationSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Duration
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, DurationTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &DurationSubscription{sub}, nil
}

func (s *DurationSubscription) TakeMessage(out *Duration) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneDurationSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDurationSlice(dst, src []Duration) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DurationTypeSupport types.MessageTypeSupport = _DurationTypeSupport{}

type _DurationTypeSupport struct{}

func (t _DurationTypeSupport) New() types.Message {
	return NewDuration()
}

func (t _DurationTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.builtin_interfaces__msg__Duration
	return (unsafe.Pointer)(C.builtin_interfaces__msg__Duration__create())
}

func (t _DurationTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.builtin_interfaces__msg__Duration__destroy((*C.builtin_interfaces__msg__Duration)(pointer_to_free))
}

func (t _DurationTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Duration)
	mem := (*C.builtin_interfaces__msg__Duration)(dst)
	mem.sec = C.int32_t(m.Sec)
	mem.nanosec = C.uint32_t(m.Nanosec)
}

func (t _DurationTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Duration)
	mem := (*C.builtin_interfaces__msg__Duration)(ros2_message_buffer)
	m.Sec = int32(mem.sec)
	m.Nanosec = uint32(mem.nanosec)
}

func (t _DurationTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__builtin_interfaces__msg__Duration())
}

type CDuration = C.builtin_interfaces__msg__Duration
type CDuration__Sequence = C.builtin_interfaces__msg__Duration__Sequence

func Duration__Sequence_to_Go(goSlice *[]Duration, cSlice CDuration__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Duration, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		DurationTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Duration__Sequence_to_C(cSlice *CDuration__Sequence, goSlice []Duration) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.builtin_interfaces__msg__Duration)(C.malloc(C.sizeof_struct_builtin_interfaces__msg__Duration * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		DurationTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Duration__Array_to_Go(goSlice []Duration, cSlice []CDuration) {
	for i := 0; i < len(cSlice); i++ {
		DurationTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Duration__Array_to_C(cSlice []CDuration, goSlice []Duration) {
	for i := 0; i < len(goSlice); i++ {
		DurationTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
