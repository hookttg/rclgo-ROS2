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
	geometry_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/imu.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Imu", ImuTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/Imu", ImuTypeSupport)
}

// Do not create instances of this type directly. Always use NewImu
// function instead.
type Imu struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Orientation geometry_msgs_msg.Quaternion `yaml:"orientation"`
	OrientationCovariance [9]float64 `yaml:"orientation_covariance"`// Row major about x, y, z axes
	AngularVelocity geometry_msgs_msg.Vector3 `yaml:"angular_velocity"`
	AngularVelocityCovariance [9]float64 `yaml:"angular_velocity_covariance"`// Row major about x, y, z axes
	LinearAcceleration geometry_msgs_msg.Vector3 `yaml:"linear_acceleration"`
	LinearAccelerationCovariance [9]float64 `yaml:"linear_acceleration_covariance"`// Row major x, y z
}

// NewImu creates a new Imu with default values.
func NewImu() *Imu {
	self := Imu{}
	self.SetDefaults()
	return &self
}

func (t *Imu) Clone() *Imu {
	c := &Imu{}
	c.Header = *t.Header.Clone()
	c.Orientation = *t.Orientation.Clone()
	c.OrientationCovariance = t.OrientationCovariance
	c.AngularVelocity = *t.AngularVelocity.Clone()
	c.AngularVelocityCovariance = t.AngularVelocityCovariance
	c.LinearAcceleration = *t.LinearAcceleration.Clone()
	c.LinearAccelerationCovariance = t.LinearAccelerationCovariance
	return c
}

func (t *Imu) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Imu) SetDefaults() {
	t.Header.SetDefaults()
	t.Orientation.SetDefaults()
	t.OrientationCovariance = [9]float64{}
	t.AngularVelocity.SetDefaults()
	t.AngularVelocityCovariance = [9]float64{}
	t.LinearAcceleration.SetDefaults()
	t.LinearAccelerationCovariance = [9]float64{}
}

func (t *Imu) GetTypeSupport() types.MessageTypeSupport {
	return ImuTypeSupport
}

// ImuPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ImuPublisher struct {
	*rclgo.Publisher
}

// NewImuPublisher creates and returns a new publisher for the
// Imu
func NewImuPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ImuPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ImuTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ImuPublisher{pub}, nil
}

func (p *ImuPublisher) Publish(msg *Imu) error {
	return p.Publisher.Publish(msg)
}

// ImuSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ImuSubscription struct {
	*rclgo.Subscription
}

// ImuSubscriptionCallback type is used to provide a subscription
// handler function for a ImuSubscription.
type ImuSubscriptionCallback func(msg *Imu, info *rclgo.RmwMessageInfo, err error)

// NewImuSubscription creates and returns a new subscription for the
// Imu
func NewImuSubscription(node *rclgo.Node, topic_name string, subscriptionCallback ImuSubscriptionCallback) (*ImuSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Imu
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ImuTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &ImuSubscription{sub}, nil
}

func (s *ImuSubscription) TakeMessage(out *Imu) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneImuSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneImuSlice(dst, src []Imu) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ImuTypeSupport types.MessageTypeSupport = _ImuTypeSupport{}

type _ImuTypeSupport struct{}

func (t _ImuTypeSupport) New() types.Message {
	return NewImu()
}

func (t _ImuTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Imu
	return (unsafe.Pointer)(C.sensor_msgs__msg__Imu__create())
}

func (t _ImuTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Imu__destroy((*C.sensor_msgs__msg__Imu)(pointer_to_free))
}

func (t _ImuTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Imu)
	mem := (*C.sensor_msgs__msg__Imu)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	geometry_msgs_msg.QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.orientation), &m.Orientation)
	cSlice_orientation_covariance := mem.orientation_covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_orientation_covariance)), m.OrientationCovariance[:])
	geometry_msgs_msg.Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.angular_velocity), &m.AngularVelocity)
	cSlice_angular_velocity_covariance := mem.angular_velocity_covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_angular_velocity_covariance)), m.AngularVelocityCovariance[:])
	geometry_msgs_msg.Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.linear_acceleration), &m.LinearAcceleration)
	cSlice_linear_acceleration_covariance := mem.linear_acceleration_covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_linear_acceleration_covariance)), m.LinearAccelerationCovariance[:])
}

func (t _ImuTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Imu)
	mem := (*C.sensor_msgs__msg__Imu)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	geometry_msgs_msg.QuaternionTypeSupport.AsGoStruct(&m.Orientation, unsafe.Pointer(&mem.orientation))
	cSlice_orientation_covariance := mem.orientation_covariance[:]
	primitives.Float64__Array_to_Go(m.OrientationCovariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_orientation_covariance)))
	geometry_msgs_msg.Vector3TypeSupport.AsGoStruct(&m.AngularVelocity, unsafe.Pointer(&mem.angular_velocity))
	cSlice_angular_velocity_covariance := mem.angular_velocity_covariance[:]
	primitives.Float64__Array_to_Go(m.AngularVelocityCovariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_angular_velocity_covariance)))
	geometry_msgs_msg.Vector3TypeSupport.AsGoStruct(&m.LinearAcceleration, unsafe.Pointer(&mem.linear_acceleration))
	cSlice_linear_acceleration_covariance := mem.linear_acceleration_covariance[:]
	primitives.Float64__Array_to_Go(m.LinearAccelerationCovariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_linear_acceleration_covariance)))
}

func (t _ImuTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Imu())
}

type CImu = C.sensor_msgs__msg__Imu
type CImu__Sequence = C.sensor_msgs__msg__Imu__Sequence

func Imu__Sequence_to_Go(goSlice *[]Imu, cSlice CImu__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Imu, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ImuTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Imu__Sequence_to_C(cSlice *CImu__Sequence, goSlice []Imu) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Imu)(C.malloc(C.sizeof_struct_sensor_msgs__msg__Imu * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ImuTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Imu__Array_to_Go(goSlice []Imu, cSlice []CImu) {
	for i := 0; i < len(cSlice); i++ {
		ImuTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Imu__Array_to_C(cSlice []CImu, goSlice []Imu) {
	for i := 0; i < len(goSlice); i++ {
		ImuTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
