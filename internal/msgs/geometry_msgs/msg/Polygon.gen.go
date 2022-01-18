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

#include <geometry_msgs/msg/polygon.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Polygon", PolygonTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/Polygon", PolygonTypeSupport)
}

// Do not create instances of this type directly. Always use NewPolygon
// function instead.
type Polygon struct {
	Points []Point32 `yaml:"points"`
}

// NewPolygon creates a new Polygon with default values.
func NewPolygon() *Polygon {
	self := Polygon{}
	self.SetDefaults()
	return &self
}

func (t *Polygon) Clone() *Polygon {
	c := &Polygon{}
	if t.Points != nil {
		c.Points = make([]Point32, len(t.Points))
		ClonePoint32Slice(c.Points, t.Points)
	}
	return c
}

func (t *Polygon) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Polygon) SetDefaults() {
	t.Points = nil
}

// PolygonPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PolygonPublisher struct {
	*rclgo.Publisher
}

// NewPolygonPublisher creates and returns a new publisher for the
// Polygon
func NewPolygonPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PolygonPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PolygonTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PolygonPublisher{pub}, nil
}

func (p *PolygonPublisher) Publish(msg *Polygon) error {
	return p.Publisher.Publish(msg)
}

// PolygonSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PolygonSubscription struct {
	*rclgo.Subscription
}

// PolygonSubscriptionCallback type is used to provide a subscription
// handler function for a PolygonSubscription.
type PolygonSubscriptionCallback func(msg *Polygon, info *rclgo.RmwMessageInfo, err error)

// NewPolygonSubscription creates and returns a new subscription for the
// Polygon
func NewPolygonSubscription(node *rclgo.Node, topic_name string, subscriptionCallback PolygonSubscriptionCallback) (*PolygonSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Polygon
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PolygonTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &PolygonSubscription{sub}, nil
}

func (s *PolygonSubscription) TakeMessage(out *Polygon) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePolygonSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePolygonSlice(dst, src []Polygon) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PolygonTypeSupport types.MessageTypeSupport = _PolygonTypeSupport{}

type _PolygonTypeSupport struct{}

func (t _PolygonTypeSupport) New() types.Message {
	return NewPolygon()
}

func (t _PolygonTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Polygon
	return (unsafe.Pointer)(C.geometry_msgs__msg__Polygon__create())
}

func (t _PolygonTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Polygon__destroy((*C.geometry_msgs__msg__Polygon)(pointer_to_free))
}

func (t _PolygonTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Polygon)
	mem := (*C.geometry_msgs__msg__Polygon)(dst)
	Point32__Sequence_to_C(&mem.points, m.Points)
}

func (t _PolygonTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Polygon)
	mem := (*C.geometry_msgs__msg__Polygon)(ros2_message_buffer)
	Point32__Sequence_to_Go(&m.Points, mem.points)
}

func (t _PolygonTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Polygon())
}

type CPolygon = C.geometry_msgs__msg__Polygon
type CPolygon__Sequence = C.geometry_msgs__msg__Polygon__Sequence

func Polygon__Sequence_to_Go(goSlice *[]Polygon, cSlice CPolygon__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Polygon, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PolygonTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Polygon__Sequence_to_C(cSlice *CPolygon__Sequence, goSlice []Polygon) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Polygon)(C.malloc(C.sizeof_struct_geometry_msgs__msg__Polygon * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PolygonTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Polygon__Array_to_Go(goSlice []Polygon, cSlice []CPolygon) {
	for i := 0; i < len(cSlice); i++ {
		PolygonTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Polygon__Array_to_C(cSlice []CPolygon, goSlice []Polygon) {
	for i := 0; i < len(goSlice); i++ {
		PolygonTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
