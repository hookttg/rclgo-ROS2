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
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/region_of_interest.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/RegionOfInterest", RegionOfInterestTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/RegionOfInterest", RegionOfInterestTypeSupport)
}

// Do not create instances of this type directly. Always use NewRegionOfInterest
// function instead.
type RegionOfInterest struct {
	XOffset uint32 `yaml:"x_offset"`// Leftmost pixel of the ROI
	YOffset uint32 `yaml:"y_offset"`// Topmost pixel of the ROI. (0 if the ROI includes the left edge of the image)
	Height uint32 `yaml:"height"`// Height of ROI. (0 if the ROI includes the top edge of the image)
	Width uint32 `yaml:"width"`// Width of ROI
	DoRectify bool `yaml:"do_rectify"`// True if a distinct rectified ROI should be calculated from the "raw"ROI in this message. Typically this should be False if the full imageis captured (ROI not used), and True if a subwindow is captured (ROIused).
}

// NewRegionOfInterest creates a new RegionOfInterest with default values.
func NewRegionOfInterest() *RegionOfInterest {
	self := RegionOfInterest{}
	self.SetDefaults()
	return &self
}

func (t *RegionOfInterest) Clone() *RegionOfInterest {
	c := &RegionOfInterest{}
	c.XOffset = t.XOffset
	c.YOffset = t.YOffset
	c.Height = t.Height
	c.Width = t.Width
	c.DoRectify = t.DoRectify
	return c
}

func (t *RegionOfInterest) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RegionOfInterest) SetDefaults() {
	t.XOffset = 0
	t.YOffset = 0
	t.Height = 0
	t.Width = 0
	t.DoRectify = false
}

func (t *RegionOfInterest) GetTypeSupport() types.MessageTypeSupport {
	return RegionOfInterestTypeSupport
}

// RegionOfInterestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RegionOfInterestPublisher struct {
	*rclgo.Publisher
}

// NewRegionOfInterestPublisher creates and returns a new publisher for the
// RegionOfInterest
func NewRegionOfInterestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RegionOfInterestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RegionOfInterestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RegionOfInterestPublisher{pub}, nil
}

func (p *RegionOfInterestPublisher) Publish(msg *RegionOfInterest) error {
	return p.Publisher.Publish(msg)
}

// RegionOfInterestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RegionOfInterestSubscription struct {
	*rclgo.Subscription
}

// RegionOfInterestSubscriptionCallback type is used to provide a subscription
// handler function for a RegionOfInterestSubscription.
type RegionOfInterestSubscriptionCallback func(msg *RegionOfInterest, info *rclgo.RmwMessageInfo, err error)

// NewRegionOfInterestSubscription creates and returns a new subscription for the
// RegionOfInterest
func NewRegionOfInterestSubscription(node *rclgo.Node, topic_name string, subscriptionCallback RegionOfInterestSubscriptionCallback) (*RegionOfInterestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RegionOfInterest
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RegionOfInterestTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &RegionOfInterestSubscription{sub}, nil
}

func (s *RegionOfInterestSubscription) TakeMessage(out *RegionOfInterest) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRegionOfInterestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRegionOfInterestSlice(dst, src []RegionOfInterest) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RegionOfInterestTypeSupport types.MessageTypeSupport = _RegionOfInterestTypeSupport{}

type _RegionOfInterestTypeSupport struct{}

func (t _RegionOfInterestTypeSupport) New() types.Message {
	return NewRegionOfInterest()
}

func (t _RegionOfInterestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__RegionOfInterest
	return (unsafe.Pointer)(C.sensor_msgs__msg__RegionOfInterest__create())
}

func (t _RegionOfInterestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__RegionOfInterest__destroy((*C.sensor_msgs__msg__RegionOfInterest)(pointer_to_free))
}

func (t _RegionOfInterestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RegionOfInterest)
	mem := (*C.sensor_msgs__msg__RegionOfInterest)(dst)
	mem.x_offset = C.uint32_t(m.XOffset)
	mem.y_offset = C.uint32_t(m.YOffset)
	mem.height = C.uint32_t(m.Height)
	mem.width = C.uint32_t(m.Width)
	mem.do_rectify = C.bool(m.DoRectify)
}

func (t _RegionOfInterestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RegionOfInterest)
	mem := (*C.sensor_msgs__msg__RegionOfInterest)(ros2_message_buffer)
	m.XOffset = uint32(mem.x_offset)
	m.YOffset = uint32(mem.y_offset)
	m.Height = uint32(mem.height)
	m.Width = uint32(mem.width)
	m.DoRectify = bool(mem.do_rectify)
}

func (t _RegionOfInterestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__RegionOfInterest())
}

type CRegionOfInterest = C.sensor_msgs__msg__RegionOfInterest
type CRegionOfInterest__Sequence = C.sensor_msgs__msg__RegionOfInterest__Sequence

func RegionOfInterest__Sequence_to_Go(goSlice *[]RegionOfInterest, cSlice CRegionOfInterest__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RegionOfInterest, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RegionOfInterestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RegionOfInterest__Sequence_to_C(cSlice *CRegionOfInterest__Sequence, goSlice []RegionOfInterest) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__RegionOfInterest)(C.malloc(C.sizeof_struct_sensor_msgs__msg__RegionOfInterest * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RegionOfInterestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RegionOfInterest__Array_to_Go(goSlice []RegionOfInterest, cSlice []CRegionOfInterest) {
	for i := 0; i < len(cSlice); i++ {
		RegionOfInterestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RegionOfInterest__Array_to_C(cSlice []CRegionOfInterest, goSlice []RegionOfInterest) {
	for i := 0; i < len(goSlice); i++ {
		RegionOfInterestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
