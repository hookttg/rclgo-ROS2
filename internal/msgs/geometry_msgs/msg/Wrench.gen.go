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

#include <geometry_msgs/msg/wrench.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Wrench", WrenchTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/Wrench", WrenchTypeSupport)
}

// Do not create instances of this type directly. Always use NewWrench
// function instead.
type Wrench struct {
	Force Vector3 `yaml:"force"`
	Torque Vector3 `yaml:"torque"`
}

// NewWrench creates a new Wrench with default values.
func NewWrench() *Wrench {
	self := Wrench{}
	self.SetDefaults()
	return &self
}

func (t *Wrench) Clone() *Wrench {
	c := &Wrench{}
	c.Force = *t.Force.Clone()
	c.Torque = *t.Torque.Clone()
	return c
}

func (t *Wrench) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Wrench) SetDefaults() {
	t.Force.SetDefaults()
	t.Torque.SetDefaults()
}

func (t *Wrench) GetTypeSupport() types.MessageTypeSupport {
	return WrenchTypeSupport
}

// WrenchPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WrenchPublisher struct {
	*rclgo.Publisher
}

// NewWrenchPublisher creates and returns a new publisher for the
// Wrench
func NewWrenchPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WrenchPublisher, error) {
	pub, err := node.NewPublisher(topic_name, WrenchTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WrenchPublisher{pub}, nil
}

func (p *WrenchPublisher) Publish(msg *Wrench) error {
	return p.Publisher.Publish(msg)
}

// WrenchSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WrenchSubscription struct {
	*rclgo.Subscription
}

// WrenchSubscriptionCallback type is used to provide a subscription
// handler function for a WrenchSubscription.
type WrenchSubscriptionCallback func(msg *Wrench, info *rclgo.RmwMessageInfo, err error)

// NewWrenchSubscription creates and returns a new subscription for the
// Wrench
func NewWrenchSubscription(node *rclgo.Node, topic_name string, subscriptionCallback WrenchSubscriptionCallback) (*WrenchSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Wrench
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WrenchTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &WrenchSubscription{sub}, nil
}

func (s *WrenchSubscription) TakeMessage(out *Wrench) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWrenchSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWrenchSlice(dst, src []Wrench) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WrenchTypeSupport types.MessageTypeSupport = _WrenchTypeSupport{}

type _WrenchTypeSupport struct{}

func (t _WrenchTypeSupport) New() types.Message {
	return NewWrench()
}

func (t _WrenchTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Wrench
	return (unsafe.Pointer)(C.geometry_msgs__msg__Wrench__create())
}

func (t _WrenchTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Wrench__destroy((*C.geometry_msgs__msg__Wrench)(pointer_to_free))
}

func (t _WrenchTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Wrench)
	mem := (*C.geometry_msgs__msg__Wrench)(dst)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.force), &m.Force)
	Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.torque), &m.Torque)
}

func (t _WrenchTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Wrench)
	mem := (*C.geometry_msgs__msg__Wrench)(ros2_message_buffer)
	Vector3TypeSupport.AsGoStruct(&m.Force, unsafe.Pointer(&mem.force))
	Vector3TypeSupport.AsGoStruct(&m.Torque, unsafe.Pointer(&mem.torque))
}

func (t _WrenchTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Wrench())
}

type CWrench = C.geometry_msgs__msg__Wrench
type CWrench__Sequence = C.geometry_msgs__msg__Wrench__Sequence

func Wrench__Sequence_to_Go(goSlice *[]Wrench, cSlice CWrench__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Wrench, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WrenchTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Wrench__Sequence_to_C(cSlice *CWrench__Sequence, goSlice []Wrench) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Wrench)(C.malloc(C.sizeof_struct_geometry_msgs__msg__Wrench * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WrenchTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Wrench__Array_to_Go(goSlice []Wrench, cSlice []CWrench) {
	for i := 0; i < len(cSlice); i++ {
		WrenchTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Wrench__Array_to_C(cSlice []CWrench, goSlice []Wrench) {
	for i := 0; i < len(goSlice); i++ {
		WrenchTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
