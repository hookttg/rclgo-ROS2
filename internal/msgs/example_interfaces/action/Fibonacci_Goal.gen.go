/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/action/fibonacci.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/Fibonacci_Goal", Fibonacci_GoalTypeSupport)
	typemap.RegisterMessage("example_interfaces/action/Fibonacci_Goal", Fibonacci_GoalTypeSupport)
}

// Do not create instances of this type directly. Always use NewFibonacci_Goal
// function instead.
type Fibonacci_Goal struct {
	Order int32 `yaml:"order"`// Goal
}

// NewFibonacci_Goal creates a new Fibonacci_Goal with default values.
func NewFibonacci_Goal() *Fibonacci_Goal {
	self := Fibonacci_Goal{}
	self.SetDefaults()
	return &self
}

func (t *Fibonacci_Goal) Clone() *Fibonacci_Goal {
	c := &Fibonacci_Goal{}
	c.Order = t.Order
	return c
}

func (t *Fibonacci_Goal) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Fibonacci_Goal) SetDefaults() {
	t.Order = 0
}

func (t *Fibonacci_Goal) GetTypeSupport() types.MessageTypeSupport {
	return Fibonacci_GoalTypeSupport
}

// Fibonacci_GoalPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Fibonacci_GoalPublisher struct {
	*rclgo.Publisher
}

// NewFibonacci_GoalPublisher creates and returns a new publisher for the
// Fibonacci_Goal
func NewFibonacci_GoalPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Fibonacci_GoalPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Fibonacci_GoalTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_GoalPublisher{pub}, nil
}

func (p *Fibonacci_GoalPublisher) Publish(msg *Fibonacci_Goal) error {
	return p.Publisher.Publish(msg)
}

// Fibonacci_GoalSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Fibonacci_GoalSubscription struct {
	*rclgo.Subscription
}

// Fibonacci_GoalSubscriptionCallback type is used to provide a subscription
// handler function for a Fibonacci_GoalSubscription.
type Fibonacci_GoalSubscriptionCallback func(msg *Fibonacci_Goal, info *rclgo.RmwMessageInfo, err error)

// NewFibonacci_GoalSubscription creates and returns a new subscription for the
// Fibonacci_Goal
func NewFibonacci_GoalSubscription(node *rclgo.Node, topic_name string, subscriptionCallback Fibonacci_GoalSubscriptionCallback) (*Fibonacci_GoalSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Fibonacci_Goal
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Fibonacci_GoalTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_GoalSubscription{sub}, nil
}

func (s *Fibonacci_GoalSubscription) TakeMessage(out *Fibonacci_Goal) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFibonacci_GoalSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFibonacci_GoalSlice(dst, src []Fibonacci_Goal) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Fibonacci_GoalTypeSupport types.MessageTypeSupport = _Fibonacci_GoalTypeSupport{}

type _Fibonacci_GoalTypeSupport struct{}

func (t _Fibonacci_GoalTypeSupport) New() types.Message {
	return NewFibonacci_Goal()
}

func (t _Fibonacci_GoalTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__action__Fibonacci_Goal
	return (unsafe.Pointer)(C.example_interfaces__action__Fibonacci_Goal__create())
}

func (t _Fibonacci_GoalTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__action__Fibonacci_Goal__destroy((*C.example_interfaces__action__Fibonacci_Goal)(pointer_to_free))
}

func (t _Fibonacci_GoalTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Fibonacci_Goal)
	mem := (*C.example_interfaces__action__Fibonacci_Goal)(dst)
	mem.order = C.int32_t(m.Order)
}

func (t _Fibonacci_GoalTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Fibonacci_Goal)
	mem := (*C.example_interfaces__action__Fibonacci_Goal)(ros2_message_buffer)
	m.Order = int32(mem.order)
}

func (t _Fibonacci_GoalTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__action__Fibonacci_Goal())
}

type CFibonacci_Goal = C.example_interfaces__action__Fibonacci_Goal
type CFibonacci_Goal__Sequence = C.example_interfaces__action__Fibonacci_Goal__Sequence

func Fibonacci_Goal__Sequence_to_Go(goSlice *[]Fibonacci_Goal, cSlice CFibonacci_Goal__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Fibonacci_Goal, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Fibonacci_GoalTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Fibonacci_Goal__Sequence_to_C(cSlice *CFibonacci_Goal__Sequence, goSlice []Fibonacci_Goal) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__action__Fibonacci_Goal)(C.malloc(C.sizeof_struct_example_interfaces__action__Fibonacci_Goal * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Fibonacci_GoalTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Fibonacci_Goal__Array_to_Go(goSlice []Fibonacci_Goal, cSlice []CFibonacci_Goal) {
	for i := 0; i < len(cSlice); i++ {
		Fibonacci_GoalTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Fibonacci_Goal__Array_to_C(cSlice []CFibonacci_Goal, goSlice []Fibonacci_Goal) {
	for i := 0; i < len(goSlice); i++ {
		Fibonacci_GoalTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
