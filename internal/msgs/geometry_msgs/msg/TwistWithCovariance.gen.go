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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/twist_with_covariance.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/TwistWithCovariance", TwistWithCovarianceTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/TwistWithCovariance", TwistWithCovarianceTypeSupport)
}

// Do not create instances of this type directly. Always use NewTwistWithCovariance
// function instead.
type TwistWithCovariance struct {
	Twist Twist `yaml:"twist"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewTwistWithCovariance creates a new TwistWithCovariance with default values.
func NewTwistWithCovariance() *TwistWithCovariance {
	self := TwistWithCovariance{}
	self.SetDefaults()
	return &self
}

func (t *TwistWithCovariance) Clone() *TwistWithCovariance {
	c := &TwistWithCovariance{}
	c.Twist = *t.Twist.Clone()
	c.Covariance = t.Covariance
	return c
}

func (t *TwistWithCovariance) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TwistWithCovariance) SetDefaults() {
	t.Twist.SetDefaults()
	t.Covariance = [36]float64{}
}

func (t *TwistWithCovariance) GetTypeSupport() types.MessageTypeSupport {
	return TwistWithCovarianceTypeSupport
}

// TwistWithCovariancePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TwistWithCovariancePublisher struct {
	*rclgo.Publisher
}

// NewTwistWithCovariancePublisher creates and returns a new publisher for the
// TwistWithCovariance
func NewTwistWithCovariancePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TwistWithCovariancePublisher, error) {
	pub, err := node.NewPublisher(topic_name, TwistWithCovarianceTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TwistWithCovariancePublisher{pub}, nil
}

func (p *TwistWithCovariancePublisher) Publish(msg *TwistWithCovariance) error {
	return p.Publisher.Publish(msg)
}

// TwistWithCovarianceSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TwistWithCovarianceSubscription struct {
	*rclgo.Subscription
}

// TwistWithCovarianceSubscriptionCallback type is used to provide a subscription
// handler function for a TwistWithCovarianceSubscription.
type TwistWithCovarianceSubscriptionCallback func(msg *TwistWithCovariance, info *rclgo.RmwMessageInfo, err error)

// NewTwistWithCovarianceSubscription creates and returns a new subscription for the
// TwistWithCovariance
func NewTwistWithCovarianceSubscription(node *rclgo.Node, topic_name string, subscriptionCallback TwistWithCovarianceSubscriptionCallback) (*TwistWithCovarianceSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TwistWithCovariance
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TwistWithCovarianceTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &TwistWithCovarianceSubscription{sub}, nil
}

func (s *TwistWithCovarianceSubscription) TakeMessage(out *TwistWithCovariance) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTwistWithCovarianceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTwistWithCovarianceSlice(dst, src []TwistWithCovariance) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TwistWithCovarianceTypeSupport types.MessageTypeSupport = _TwistWithCovarianceTypeSupport{}

type _TwistWithCovarianceTypeSupport struct{}

func (t _TwistWithCovarianceTypeSupport) New() types.Message {
	return NewTwistWithCovariance()
}

func (t _TwistWithCovarianceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__TwistWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__TwistWithCovariance__create())
}

func (t _TwistWithCovarianceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__TwistWithCovariance__destroy((*C.geometry_msgs__msg__TwistWithCovariance)(pointer_to_free))
}

func (t _TwistWithCovarianceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TwistWithCovariance)
	mem := (*C.geometry_msgs__msg__TwistWithCovariance)(dst)
	TwistTypeSupport.AsCStruct(unsafe.Pointer(&mem.twist), &m.Twist)
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)), m.Covariance[:])
}

func (t _TwistWithCovarianceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TwistWithCovariance)
	mem := (*C.geometry_msgs__msg__TwistWithCovariance)(ros2_message_buffer)
	TwistTypeSupport.AsGoStruct(&m.Twist, unsafe.Pointer(&mem.twist))
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_Go(m.Covariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}

func (t _TwistWithCovarianceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__TwistWithCovariance())
}

type CTwistWithCovariance = C.geometry_msgs__msg__TwistWithCovariance
type CTwistWithCovariance__Sequence = C.geometry_msgs__msg__TwistWithCovariance__Sequence

func TwistWithCovariance__Sequence_to_Go(goSlice *[]TwistWithCovariance, cSlice CTwistWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TwistWithCovariance, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TwistWithCovarianceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func TwistWithCovariance__Sequence_to_C(cSlice *CTwistWithCovariance__Sequence, goSlice []TwistWithCovariance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__TwistWithCovariance)(C.malloc(C.sizeof_struct_geometry_msgs__msg__TwistWithCovariance * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TwistWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func TwistWithCovariance__Array_to_Go(goSlice []TwistWithCovariance, cSlice []CTwistWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		TwistWithCovarianceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TwistWithCovariance__Array_to_C(cSlice []CTwistWithCovariance, goSlice []TwistWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		TwistWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
