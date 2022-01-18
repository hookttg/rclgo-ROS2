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
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/image.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Image", ImageTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/Image", ImageTypeSupport)
}

// Do not create instances of this type directly. Always use NewImage
// function instead.
type Image struct {
	Header std_msgs_msg.Header `yaml:"header"`// Header timestamp should be acquisition time of image
	Height uint32 `yaml:"height"`// image height, that is, number of rows
	Width uint32 `yaml:"width"`// image width, that is, number of columns
	Encoding string `yaml:"encoding"`// Encoding of pixels -- channel meaning, ordering, size
	IsBigendian uint8 `yaml:"is_bigendian"`// is this data bigendian?
	Step uint32 `yaml:"step"`// Full row length in bytes
	Data []uint8 `yaml:"data"`// actual matrix data, size is (step * rows)
}

// NewImage creates a new Image with default values.
func NewImage() *Image {
	self := Image{}
	self.SetDefaults()
	return &self
}

func (t *Image) Clone() *Image {
	c := &Image{}
	c.Header = *t.Header.Clone()
	c.Height = t.Height
	c.Width = t.Width
	c.Encoding = t.Encoding
	c.IsBigendian = t.IsBigendian
	c.Step = t.Step
	if t.Data != nil {
		c.Data = make([]uint8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Image) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Image) SetDefaults() {
	t.Header.SetDefaults()
	t.Height = 0
	t.Width = 0
	t.Encoding = ""
	t.IsBigendian = 0
	t.Step = 0
	t.Data = nil
}

// ImagePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ImagePublisher struct {
	*rclgo.Publisher
}

// NewImagePublisher creates and returns a new publisher for the
// Image
func NewImagePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ImagePublisher, error) {
	pub, err := node.NewPublisher(topic_name, ImageTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ImagePublisher{pub}, nil
}

func (p *ImagePublisher) Publish(msg *Image) error {
	return p.Publisher.Publish(msg)
}

// ImageSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ImageSubscription struct {
	*rclgo.Subscription
}

// ImageSubscriptionCallback type is used to provide a subscription
// handler function for a ImageSubscription.
type ImageSubscriptionCallback func(msg *Image, info *rclgo.RmwMessageInfo, err error)

// NewImageSubscription creates and returns a new subscription for the
// Image
func NewImageSubscription(node *rclgo.Node, topic_name string, subscriptionCallback ImageSubscriptionCallback) (*ImageSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Image
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ImageTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &ImageSubscription{sub}, nil
}

func (s *ImageSubscription) TakeMessage(out *Image) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneImageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneImageSlice(dst, src []Image) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ImageTypeSupport types.MessageTypeSupport = _ImageTypeSupport{}

type _ImageTypeSupport struct{}

func (t _ImageTypeSupport) New() types.Message {
	return NewImage()
}

func (t _ImageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Image
	return (unsafe.Pointer)(C.sensor_msgs__msg__Image__create())
}

func (t _ImageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Image__destroy((*C.sensor_msgs__msg__Image)(pointer_to_free))
}

func (t _ImageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Image)
	mem := (*C.sensor_msgs__msg__Image)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.height = C.uint32_t(m.Height)
	mem.width = C.uint32_t(m.Width)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.encoding), m.Encoding)
	mem.is_bigendian = C.uint8_t(m.IsBigendian)
	mem.step = C.uint32_t(m.Step)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _ImageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Image)
	mem := (*C.sensor_msgs__msg__Image)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Height = uint32(mem.height)
	m.Width = uint32(mem.width)
	primitives.StringAsGoStruct(&m.Encoding, unsafe.Pointer(&mem.encoding))
	m.IsBigendian = uint8(mem.is_bigendian)
	m.Step = uint32(mem.step)
	primitives.Uint8__Sequence_to_Go(&m.Data, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _ImageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Image())
}

type CImage = C.sensor_msgs__msg__Image
type CImage__Sequence = C.sensor_msgs__msg__Image__Sequence

func Image__Sequence_to_Go(goSlice *[]Image, cSlice CImage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Image, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ImageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Image__Sequence_to_C(cSlice *CImage__Sequence, goSlice []Image) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Image)(C.malloc(C.sizeof_struct_sensor_msgs__msg__Image * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ImageTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Image__Array_to_Go(goSlice []Image, cSlice []CImage) {
	for i := 0; i < len(cSlice); i++ {
		ImageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Image__Array_to_C(cSlice []CImage, goSlice []Image) {
	for i := 0; i < len(goSlice); i++ {
		ImageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
