/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/srv/path.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("fog_msgs/Path_Response", &Path_Response{})
}

// Do not create instances of this type directly. Always use NewPath_Response
// function instead.
type Path_Response struct {
	Success bool `yaml:"success"`
	Message rosidl_runtime_c.String `yaml:"message"`
}

// NewPath_Response creates a new Path_Response with default values.
func NewPath_Response() *Path_Response {
	self := Path_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *Path_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Message.SetDefaults("")
	
	return t
}

func (t *Path_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__Path_Response())
}
func (t *Path_Response) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__Path_Response
	return (unsafe.Pointer)(C.fog_msgs__srv__Path_Response__create())
}
func (t *Path_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__Path_Response__destroy((*C.fog_msgs__srv__Path_Response)(pointer_to_free))
}
func (t *Path_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.fog_msgs__srv__Path_Response)(t.PrepareMemory())
	mem.success = C.bool(t.Success)
	mem.message = *(*C.rosidl_runtime_c__String)(t.Message.AsCStruct())
	return unsafe.Pointer(mem)
}
func (t *Path_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.fog_msgs__srv__Path_Response)(ros2_message_buffer)
	t.Success = bool(mem.success)
	t.Message.AsGoStruct(unsafe.Pointer(&mem.message))
}
func (t *Path_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CPath_Response = C.fog_msgs__srv__Path_Response
type CPath_Response__Sequence = C.fog_msgs__srv__Path_Response__Sequence

func Path_Response__Sequence_to_Go(goSlice *[]Path_Response, cSlice CPath_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Path_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__Path_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__Path_Response * uintptr(i)),
		))
		(*goSlice)[i] = Path_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Path_Response__Sequence_to_C(cSlice *CPath_Response__Sequence, goSlice []Path_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__Path_Response)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__Path_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__Path_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__Path_Response * uintptr(i)),
		))
		*cIdx = *(*C.fog_msgs__srv__Path_Response)(v.AsCStruct())
	}
}
func Path_Response__Array_to_Go(goSlice []Path_Response, cSlice []CPath_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Path_Response__Array_to_C(cSlice []CPath_Response, goSlice []Path_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.fog_msgs__srv__Path_Response)(goSlice[i].AsCStruct())
	}
}


