/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/accel.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Accel", AccelTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/Accel", AccelTypeSupport)
}

// Do not create instances of this type directly. Always use NewAccel
// function instead.
type Accel struct {
	Linear Vector3 `yaml:"linear"`// This expresses acceleration in free space broken into its linear and angular parts.
	Angular Vector3 `yaml:"angular"`// This expresses acceleration in free space broken into its linear and angular parts.
}

// NewAccel creates a new Accel with default values.
func NewAccel() *Accel {
	self := Accel{}
	self.SetDefaults()
	return &self
}

func (t *Accel) Clone() *Accel {
	c := &Accel{}
	c.Linear = *t.Linear.Clone()
	c.Angular = *t.Angular.Clone()
	return c
}

func (t *Accel) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Accel) SetDefaults() {
	t.Linear.SetDefaults()
	t.Angular.SetDefaults()
}

// AccelPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type AccelPublisher struct {
	*rclgo.Publisher
}

// NewAccelPublisher creates and returns a new publisher for the
// Accel
func NewAccelPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*AccelPublisher, error) {
	pub, err := node.NewPublisher(topic_name, AccelTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AccelPublisher{pub}, nil
}

func (p *AccelPublisher) Publish(msg *Accel) error {
	return p.Publisher.Publish(msg)
}

// AccelSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type AccelSubscription struct {
	*rclgo.Subscription
}

// AccelSubscriptionCallback type is used to provide a subscription
// handler function for a AccelSubscription.
type AccelSubscriptionCallback func(msg *Accel, info *rclgo.RmwMessageInfo, err error)

// NewAccelSubscription creates and returns a new subscription for the
// Accel
func NewAccelSubscription(node *rclgo.Node, topic_name string, subscriptionCallback AccelSubscriptionCallback) (*AccelSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Accel
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, AccelTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &AccelSubscription{sub}, nil
}

func (s *AccelSubscription) TakeMessage(out *Accel) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneAccelSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAccelSlice(dst, src []Accel) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AccelTypeSupport types.MessageTypeSupport = _AccelTypeSupport{}

type _AccelTypeSupport struct{}

func (t _AccelTypeSupport) New() types.Message {
	return NewAccel()
}

func (t _AccelTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Accel
	return (unsafe.Pointer)(C.geometry_msgs__msg__Accel__create())
}

func (t _AccelTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Accel__destroy((*C.geometry_msgs__msg__Accel)(pointer_to_free))
}

func (t _AccelTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Accel)
	mem := (*C.geometry_msgs__msg__Accel)(dst)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.linear), &m.Linear)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.angular), &m.Angular)
}

func (t _AccelTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Accel)
	mem := (*C.geometry_msgs__msg__Accel)(ros2_message_buffer)
	Vector3TypeSupport.AsGoStruct(&m.Linear, unsafe.Pointer(&mem.linear))
	Vector3TypeSupport.AsGoStruct(&m.Angular, unsafe.Pointer(&mem.angular))
}

func (t _AccelTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Accel())
}

type CAccel = C.geometry_msgs__msg__Accel
type CAccel__Sequence = C.geometry_msgs__msg__Accel__Sequence

func Accel__Sequence_to_Go(goSlice *[]Accel, cSlice CAccel__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Accel, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		AccelTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Accel__Sequence_to_C(cSlice *CAccel__Sequence, goSlice []Accel) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Accel)(C.malloc(C.sizeof_struct_geometry_msgs__msg__Accel * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		AccelTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Accel__Array_to_Go(goSlice []Accel, cSlice []CAccel) {
	for i := 0; i < len(cSlice); i++ {
		AccelTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Accel__Array_to_C(cSlice []CAccel, goSlice []Accel) {
	for i := 0; i < len(goSlice); i++ {
		AccelTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
